<div align="center">

# <img src="https://cdn.discordapp.com/emojis/1268324750879621243.webp" width="25"> osaka-codec
<sub>The codec that will replace everything 🗣️‼️</sub>

<img src="https://cdn.ananas.moe/osaka-codec.png">
</div>

## Installation

### Install via `go install`

```bash
go install github.com/osaka-lab/osaka-codec@latest
```

### Download the Release

You can download the [latest release](https://github.com/osaka-lab/osaka-codec/releases/latest) and move the binary to your `~/.local/bin` directory.

## Usage
```
osaka-codec [command] {file/text}

Available Commands:
  decode      Decode data with osaka-codec
  encode      Encode data with osaka-codec

Flags:
  -o, --output string   Specify an output file to save the result. (Only on file encoding/decoding)
  -s, --string          Specify that the input arguments are a string rather than a file path.
```

> Encoding a file

> ```bash
> osaka-codec encode {file/text}
> ```

> Decoding a file
> ```
> osaka-codec decode {file/text}
> ```

> Add `-s`/`--string` if it's a string.

## Inspiration
[uwu-codec](https://github.com/THEGOLDENPRO/uwu-codec)
