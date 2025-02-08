# b64

A shorthand command for encoding/decoding text to/from base 64.

## Installation

To install `b64`, you need to have Go installed on your machine. If you don't have Go installed, you can download it from the [official website](https://golang.org/dl/). Once you have Go installed, you can install `b64` by running the following command:

```bash
go install github.com/pawelkowalak/b64@latest
```

## Usage

### Encoding

```
$ b64 text
dGV4dA==
```

### Decoding

```
$ b64 -d dGV4dA==
text
```