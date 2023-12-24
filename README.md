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
mu build [source]
```

List binaries

```bash
mu list
```

Run an app

```bash
mu run [app]
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
