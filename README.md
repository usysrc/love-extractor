# love-extractor

A simple tool written in Go that extracts an appended `.love` file from a LÖVE game executable.

## Prerequisites

- Go (tested with Go 1.20)

## Compilation

1. Clone this repository or download the `love-extractor.go` source file.
2. Navigate to the directory containing `love-extractor.go`.
3. Compile the program:
```bash
go build
```

## Usage

After compiling the program, you can extract the `.love` file from a LÖVE game executable using:

```bash
./love-extractor path/to/your/exe/file.exe output.love
```

This will search for the appended `.love` file within the given executable and, if found, will save it as `output.love`.

## Important Notes

- This tool assumes that the `.love` file is simply appended to the LÖVE executable.
- If the `.exe` uses a different method of bundling or has any form of obfuscation, this tool might not work as intended.
- Always ensure you have the rights to extract and utilize the contents of any software you use with this tool. Unauthorized extraction or decompilation may violate copyright laws or terms of service.

## License

MIT (or whatever license you'd like to use)

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.


