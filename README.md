# tinygo-logger

Logger for [TinyGo](https://tinygo.org/). It builds log lines in a fixed, pre-allocated byte buffer (avoiding the allocations and `fmt` overhead of the standard library) and writes them to `os.Stdout` with a `HH:MM:SS.mmm` timestamp and a level header.

**Note:** This repository is archived and read-only.

## Installation

```bash
go get github.com/ralvarezdev/tinygo-logger
```

Depends on `tinygo-buffers` and `tinygo-errors` (same author).

## Usage

```go
logger := tinygologger.NewDefaultLogger(256) // buffer size in bytes

logger.AddMessageWithUint32([]byte("Value:"), 42, true, true, false)
logger.Debug()

logger.InfoMessage([]byte("started"))
logger.ErrorMessageWithErrorCode([]byte("failed:"), errCode, true)
```

Fields accumulate in the buffer through the `Add*` methods and are flushed by `Debug`, `Info`, `Warning` or `Error` (or the `*Message` shortcuts). A full buffer is flushed with a `FULL_BUFFER` header.

## API

- **`Logger`** — interface in `interfaces.go` with the `Add*`, `Debug*`, `Info*`, `Warning*` and `Error*` methods.
- **`NewDefaultLogger(bufferSize uint64)`** — returns the `DefaultLogger` (`types.go`).
- **`DebugMemory(logger Logger)`** — logs current and cumulative allocated memory in KB from `runtime.MemStats` (`utils.go`).
- **Headers** — `DebugHeader`, `WarningHeader`, `ErrorHeader`, `InfoHeader`, `FullBufferHeader` in `constants.go`.

## License

GNU General Public License v3.0. See [LICENSE](LICENSE).
