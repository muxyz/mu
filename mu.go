package mu

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"time"
)

// user home dir
var Home string

// the file cache
var Cache string

// the secret key
var Key string

func init() {
	user, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}

	home := filepath.Join(user, "mu")
	if err := os.MkdirAll(home, 0700); err != nil {
		panic(err.Error())
	}

	// set home
	Home = home

	cache := filepath.Join(home, "cache")
	if err := os.MkdirAll(cache, 0700); err != nil {
		panic(err.Error())
	}

	// set cache
	Cache = cache

	path := filepath.Join(home, "key")
	b, _ := os.ReadFile(path)

	if len(b) == 0 {
		// generate a new key
		bytes := make([]byte, 32) //generate a random 32 byte key for AES-256
		if _, err := rand.Read(bytes); err != nil {
			panic(err.Error())
		}

		key := hex.EncodeToString(bytes) //encode key in bytes to string and keep as secret, put in a vault
		fmt.Println("generating new key", path)

		// write the file
		if err := os.WriteFile(path, []byte(key), 0600); err != nil {
			panic(err.Error())
		}

		Key = key
	} else {
		fmt.Println("loading key", path)
		Key = string(b)
	}
}

func encrypt(stringToEncrypt string, keyString string) (encryptedString string) {
	//Since the key is in string, we need to convert decode it to bytes
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

func decrypt(encryptedString string, keyString string) (decryptedString string) {
	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}

func save(data interface{}, file, key string) error {
	cache := filepath.Join(Cache, file)

	// encode data
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// encrypt it
	e := encrypt(string(b), Key)

	// write the data
	return os.WriteFile(cache, []byte(e), 0644)
}

func load(data interface{}, file, key string) error {
	cache := filepath.Join(Cache, file)

	_, err := os.Stat(cache)
	if err != nil {
		return err
	}

	// file exists
	b, err := os.ReadFile(cache)
	if err != nil {
		return err
	}

	if len(b) == 0 {
		return nil
	}

	d := decrypt(string(b), Key)

	return json.Unmarshal([]byte(d), data)
}

// Backoff is for exponential backoff
func Backoff(attempts int) time.Duration {
	if attempts > 13 {
		return time.Hour
	}
	return time.Duration(math.Pow(float64(attempts), math.E)) * time.Millisecond * 100
}

// Encrypt text using AES-256 and secret key
func Encrypt(text string) string {
	return encrypt(text, Key)
}

// Decrypt text using AES-256 and secret key
func Decrypt(text string) string {
	return decrypt(text, Key)
}

// Save data to the cache
func Save(data interface{}, file string) error {
	return save(data, file, Key)
}

// Load data from cache
func Load(data interface{}, file string) error {
	return load(data, file, Key)
}
