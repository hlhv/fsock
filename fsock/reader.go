package fsock

import (
        "io"
        "encoding/binary"
)

/* Reader is an object that reads entire fsock frames from an io.Reader.
 */
type Reader struct {
        underlying io.Reader
}

/* NewReader constructs a new fsock reader from an io.Reader.
 */
func NewReader (underlying io.Reader) (reader *Reader) {
        return &Reader {
                underlying: underlying,
        }
}

/* Read reads a single frame, and returns the bytes it read. It reports any
 * errors emitted by the underlying *bufio.Reader.
 */
func (reader *Reader) Read () (data []byte, err error) {
        // read frame length
        frameBytes := make([]byte, 4)
        _, err = io.ReadFull(reader.underlying, frameBytes)
        if err != nil { return nil, err }
        
        frameLen := binary.BigEndian.Uint32(frameBytes)
        
        // read actual data
        data = make([]byte, frameLen)
        _, err = io.ReadFull(reader.underlying, data)
        if err != nil { return nil, err }
        
        return data, nil
}
