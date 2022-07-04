package fsock

import (
	"bytes"
	"testing"
)

var (
	data0 []byte = []byte{34, 10, 24, 123}
	data1 []byte = []byte{0}
	data2 []byte = []byte{
		23, 32, 100, 230, 210, 50, 1, 2, 3, 3, 3, 3, 3, 3, 3, 34, 231,
		30, 102, 32,
	}
)

func TestWrite(t *testing.T) {
	var err error

	var buffer bytes.Buffer
	writer := NewWriter(&buffer)

	t.Log("Writing to buffer")
	_, err = writer.Write(data0)
	if err != nil {
		t.Error("Failed! writer.Write returned error")
		t.Error(err)
		return
	}
	err = writer.Flush()
	if err != nil {
		t.Error("Failed! writer.Flush returned error")
		t.Error(err)
		return
	}
	_, err = writer.Write(data1)
	if err != nil {
		t.Error("Failed! writer.Write returned error")
		t.Error(err)
		return
	}
	err = writer.Flush()
	if err != nil {
		t.Error("Failed! writer.Flush returned error")
		t.Error(err)
		return
	}
	_, err = writer.Write(data2)
	if err != nil {
		t.Error("Failed! writer.Write returned error")
		t.Error(err)
		return
	}
	err = writer.Flush()
	if err != nil {
		t.Error("Failed! writer.Flush returned error")
		t.Error(err)
		return
	}
	t.Log("Wrote successfully")

	t.Log(buffer.Bytes())

	if readData(buffer, t) {
		t.Log("Read successfully")
	}
}

func TestWriteFrame(t *testing.T) {
	var err error

	var buffer bytes.Buffer
	writer := NewWriter(&buffer)

	t.Log("Writing to buffer")
	_, err = writer.WriteFrame(data0)
	if err != nil {
		t.Error("Failed! writer.WriteFrame returned error")
		t.Error(err)
		return
	}
	_, err = writer.WriteFrame(data1)
	if err != nil {
		t.Error("Failed! writer.WriteFrame returned error")
		t.Error(err)
		return
	}
	_, err = writer.WriteFrame(data2)
	if err != nil {
		t.Error("Failed! writer.WriteFrame returned error")
		t.Error(err)
		return
	}
	t.Log("Wrote successfully")

	t.Log(buffer.Bytes())

	if readData(buffer, t) {
		t.Log("Read successfully")
	}
}

func readData(buffer bytes.Buffer, t *testing.T) bool {
	byteReader := bytes.NewReader(buffer.Bytes())
	reader := NewReader(byteReader)

	t.Logf("Reading from buffer")
	read0, err := reader.Read()
	if err != nil {
		t.Error("Failed! reader.Read returned error")
		t.Error(err)
		return false
	}
	if !bytes.Equal(read0, data0) {
		t.Error("Failed! read0 != data0: read0 is")
		t.Error(read0)
		return false
	}

	read1, err := reader.Read()
	if err != nil {
		t.Error("Failed! reader.Read returned error")
		t.Error(err)
		return false
	}
	if !bytes.Equal(read1, data1) {
		t.Error("Failed! read0 != data0: read0 is")
		t.Error(read1)
		return false
	}

	read2, err := reader.Read()
	if err != nil {
		t.Error("Failed! reader.Read returned error")
		t.Error(err)
		return false
	}
	if !bytes.Equal(read2, data2) {
		t.Error("Failed! read0 != data0: read0 is")
		t.Error(read2)
		return false
	}

	return true
}
