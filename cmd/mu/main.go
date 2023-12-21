package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

var (
	Bin  string
	Home string
)

func init() {
	user, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	home := filepath.Join(user, "mu")
	if err := os.MkdirAll(home, 0700); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// set home
	Home = home

	bin := filepath.Join(home, "bin")
	if err := os.MkdirAll(bin, 0700); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// set bin
	Bin = bin
}

func build(source string) (string, error) {
	path, err := filepath.Abs(source)
	if err != nil {
		return "", err
	}
	name := ""
	if path == "." || path == "/" {
		name = "app"
	} else {
		name = filepath.Base(path)
	}

	cmd := exec.Command("go", "build", "-o", filepath.Join(Bin, name), "./main.go")
	cmd.Dir = source
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		out, _ := cmd.CombinedOutput()
		fmt.Println(string(out))
		return "", err
	}
	return name, nil
}

func run(source string, update, kill chan bool) {
	fmt.Println("Running", source)

	f, err := os.Stat(source)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var name string
	var path string

	// directory assumed to be source
	if f.IsDir() {
		name, err = build(source)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		path = filepath.Join(Bin, name)
	} else {
		// otherwise its a binary
		name = filepath.Base(source)
		path = source
	}

	exit := make(chan bool)

	cmd := exec.Command(path)
	cmd.Dir = Bin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	go func() {
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			out, _ := cmd.CombinedOutput()
			fmt.Println(string(out))
		}

		close(exit)
	}()

	// wait for update or exit
	select {
	case <-exit:
		if cmd.Process != nil {
			cmd.Process.Kill()
			cmd.Process.Wait()
		}
		os.Exit(1)
	case <-kill:
		if cmd.Process != nil {
			cmd.Process.Kill()
			cmd.Process.Wait()
		}
		os.Exit(1)
	case <-update:
		fmt.Println("Received update")

		// do the update
		cmd.Process.Kill()
		cmd.Process.Wait()

		// restart
		run(source, update, kill)
	}
}

func watch(filePath string, update chan bool) error {
	initialStat, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	for {
		stat, err := os.Stat(filePath)
		if err != nil {
			return err
		}

		if stat.Size() != initialStat.Size() || stat.ModTime() != initialStat.ModTime() {
			select {
			case update <- true:
			default:
			}
		}

		time.Sleep(5 * time.Second)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	command := args[0]
	args = args[1:]

	switch command {
	case "build":
		if len(args) == 0 {
			fmt.Println("provide source")
			return
		}
		source := args[0]

		fmt.Println("Building", source)
		name, err := build(source)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Built", filepath.Join(Bin, name))
	case "run":
		if len(args) == 0 {
			fmt.Println("provide source")
			return
		}

		source := args[0]
		update := make(chan bool, 1)
		kill := make(chan bool, 1)

		exit := make(chan os.Signal, 1)
		signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)

		go run(source, update, kill)
		go watch(source, update)

		<-exit

		// kill the process
		close(kill)
	default:
		fmt.Println("unknown command")
	}
}
