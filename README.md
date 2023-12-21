# mu

Building blocks for life

## Dependencies

- Go toolchain

## Setup

```bash
go install mu.dev/cmd/mu@latest
```

## Usage

Build a binary

```bash
mu build [path/to/source]
```

Run from source

```bash
mu run [path/to/source]
```

List built binaries

```bash
mu list
```

### Examples

Build the binary
```bash
$ mu build ../news
Building news
Built /home/asim/mu/bin/news
```

Check it exists

```bash
$ mu list
news
```

Run it
```
$ mu run news
```

Run from source

```bash
$ mu run .
Building news
Built /home/asim/mu/bin/news
Running news
```

Run any binary
```bash
$ mu run /my/local/news/binary/news
```
