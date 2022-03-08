package fsock

import (
        "io"
        "bufio"
        "encoding/binary"
)

/* Reader is an object that reads entire fsock frames from an io.Reader.
 */
type Reader struct {
        underlying *bufio.Reader
}

/* NewReader constructs a new fsock reader from an io.Reader.
 */
func NewReader (raw io.Reader) (reader *Reader) {
        underlying := bufio.NewReader(raw)
        
        return &Reader {
                underlying,
        }
}

/* Read reads a single frame, and returns the bytes it read. It reports any
 * errors emitted by the underlying *bufio.Reader.
 */
func (reader *Reader) Read () (data []byte, err error) {
        // read frame length
        var frameLen uint32
        err = binary.Read(reader.underlying, binary.BigEndian, frameLen)
        if err != nil { return nil, err }
        
        // read actual data
        data = make([]byte, frameLen)
        _, err = reader.underlying.Read(data)
        if err != nil { return nil, err }
        
        return data, nil
}
