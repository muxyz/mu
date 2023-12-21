# mu

Building blocks for life

## Dependencies

- Go toolchain

## Setup

```
go install mu.dev/cmd/mu@latest
```

## Usage

Build a binary

```
mu build [path/to/source]
```

Run from source

```
mu run [path/to/source]
```

### Examples

Build the binary
```
# build it
$ mu build github.com/muxyz/news
Building github.com/muxyz/news
Built /home/asim/mu/bin/news
# run it
$ /home/asim/mu/bin/news
```

Run from source
```
$ mu run github.com/muxyz/news
Running github.com/muxyz/news
```

