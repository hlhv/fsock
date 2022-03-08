package fsock

import (
        "io"
        "encoding/binary"
)

type Writer struct {
        underlying io.Writer
        buffer []byte
}

/* NewWriter constructs a new fsock writer from an io.Writer.
 */
func NewWriter (underlying io.Writer) (reader *Writer) {
        return &Writer {
                underlying: underlying,
                buffer:     make([]byte, 0),
        }
}

/* Flush writes the currently active frame, and clears the internal buffer. It
 * reports any errors emitted by the underlying io.Writer.
 */
func (writer *Writer) Flush () (err error) {
        _, err = writer.WriteFrame(writer.buffer)
        writer.Reset()
        return err
}

/* Reset discards any unflushed buffered data.
 */
func (writer *Writer) Reset () {
        writer.buffer = make([]byte, 0)
}

/* Size returns the size of the underlying buffer in bytes.
 */
func (writer *Writer) Size () (size int) {
        return len(writer.buffer)
}

/* WriteFrame writes the contents of data, as a single frame. It returns the
 * number of bytes written, excluding the frame header. It reports any errors
 * emitted by the underlying io.Writer.
 */
func (writer *Writer) WriteFrame (data []byte) (nn int, err error) {
        // write frame length
        var frameLen uint32 = uint32(len(data))
        err = binary.Write(writer.underlying, binary.BigEndian, frameLen)
        if err != nil { return 0, err }

        // write actual data
        return writer.underlying.Write(data)
}

/* Write writes into the currently active frame. It does not actually write
 * anything to the underlying io.Writer. Flush() must be called to actually
 * write the data.
 */
func (writer *Writer) Write (data []byte) (nn int, err error) {
        writer.buffer = append(writer.buffer, data...)
        return len(data), nil
}

/* WriteByte writes a single byte.
 */
func (writer *Writer) WriteByte (data byte) (err error) {
        writer.buffer = append(writer.buffer, data)
        return nil
}

/* TODO: implement WriteRune and WriteString
 * https://pkg.go.dev/bufio#Writer.Write
 */
