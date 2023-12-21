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

### Examples

Build the binary
```bash
# build it
$ mu build github.com/muxyz/news
Building github.com/muxyz/news
Built /home/asim/mu/bin/news

# run it
$ /home/asim/mu/bin/news
```

Run from source
```bash
$ mu run github.com/muxyz/news
Running github.com/muxyz/news
```

Run any binary
```bash
$ mu run /home/asim/mu/bin/news
```
