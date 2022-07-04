# fsock

[![A+ Golang report card.](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat)](https://goreportcard.com/report/github.com/hlhv/fsock)

Simple socket framing module for HLHV.

## How to use

First, you need to install the module,

`go get github.com/hlhv/fsock`

And require it.

```
// *.go
import "github.com/hlhv/fsock/fsock"
```

`fsock` has a very similar set of methods to `bufio`. It is possible to create
readers and writers, and the constructors for both take in an io.Reader and an
io.Writer respectively.

## Frame structure

An `fsock` frame is composed of a block of data, preceded by an unsigned big
endian 32 bit integer with the size of the data in bytes. Because the size of
each block of data is known, they can be sent sequentially.

## Methods

### type Reader

```
type Reader struct {
        // contains filtered or unexported fields
}
```

Reader is an object that reads fsock frames from an io.Reader.

### func NewReader

`NewReader (raw io.Reader) (reader *Reader)`

NewReader constructs a new fsock reader from an io.Reader.

### func (reader *Reader) Read

`func (reader *Reader) Read () (data []byte, err error)`

Read reads a single frame, and returns the bytes it read. It reports any errors
emitted by the underlying *bufio.Reader.

### type Writer

```
type Writer struct {
        // contains filtered or unexported fields
}
```

### func NewWriter

`NewWriter (raw io.Writer) (writer *Writer)`

NewWriter constructs a new fsock writer from an io.Writer.

### func (writer *Writer) Flush 

`func (writer *Writer) Flush (data []byte) (err error)`

Flush writes the currently active frame, and clears the internal buffer. It
reports any errors emitted by the underlying io.Writer.

### func (writer *Writer) Reset

`func (writer *Writer) Reset ()`

Reset discards any unflushed buffered data.

### func (writer *Writer) Size

`func (writer *Writer) Size () (size int)`

Size returns the size of the underlying buffer in bytes.

### func (writer *Writer) WriteFrame

`func (writer *Writer) WriteFrame (data []byte) (nn int, err error)`

WriteFrame writes the contents of data, as a single frame. It returns the
number of bytes written, excluding the frame header. It reports any errors
emitted by the underlying io.Writer.

### func (writer *Writer) Write

`func (writer *Writer) Write (data []byte) (nn int, err error)`

Write writes into the currently active frame. It does not actually write
anything to the underlying io.Writer. Flush() must be called to actually
write the data.

### func (writer *Writer) WriteByte

`func (writer *Writer) WriteByte (data byte) (err error)`

WriteByte writes a single byte.
