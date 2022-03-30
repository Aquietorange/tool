package tbuff

import (
	"bytes"
	"encoding/hex"
	"io/ioutil"
	"os"
	"strings"
)

type Reader interface { //如果Buffer 有读取完整的bytes，则此结构为多余的，此为学习目地
	Read(b []byte) (n int, err error)
	ReadAt(b []byte, off int64) (n int, err error)
}

type Buffer struct {
	offset int
	bytes  []byte
	file   string
	reader *Reader
	close  func() error
}

func (buf *Buffer) Setoffset(pos int) {
	buf.offset = pos
}

//通过查找指定hex ,并设置当前偏移 为此hex尾部
func (buf *Buffer) SetoffsetHex(h string) bool {
	h = strings.Replace(h, " ", "", -1)
	//
	hbin, _ := hex.DecodeString(h)
	op := bytes.LastIndex(buf.bytes, hbin)
	if op >= 0 {
		buf.offset = op + len(hbin)
		return true
	} else {
		return false
	}
}

func (buf *Buffer) Readbytes(length int) []byte {
	p := make([]byte, length)
	var reader = (*buf.reader)
	reader.ReadAt(p, int64(buf.offset))
	//buf.reader.ReadAt(p, int64(buf.offset))
	//bs := buf.bytes[buf.offset : buf.offset+length]
	buf.offset += length
	return p
}

func (buf *Buffer) ReadInt32() int32 {
	b := buf.Readbytes(4)
	return int32(LittleEndian.Uint32(b))
}
func (buf *Buffer) ReadInt64() int64 {
	b := buf.Readbytes(8)
	return int64(LittleEndian.Uint64(b))
}

func (buf *Buffer) ReadString(length int) string {
	bs := buf.Readbytes(length)
	return string(bs)
}

func (buf *Buffer) ReadHex(length int) string {
	bs := buf.Readbytes(length)
	return hex.EncodeToString(bs)
}

func (buf *Buffer) Close() error {
	if buf.file != "" {
		return buf.close()
	} else {
		return nil
	}

}

func NewBuffer(bins []byte) *Buffer {
	r := bytes.NewReader(bins)
	var v = new(Reader)

	*v = r

	return &Buffer{
		offset: 0,
		bytes:  bins,
		reader: v,
		close:  func() error { return nil },
	}
}

func NewBufferByFile(file string) (*Buffer, error) {
	readerf, err := os.OpenFile(file, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	bins, _ := ioutil.ReadAll(readerf)

	var v = new(Reader)
	*v = readerf

	return &Buffer{
		offset: 0,
		file:   file,
		bytes:  bins,
		reader: v,
		close: func() error {
			//fmt.Println("关闭file")
			return readerf.Close()
		},
	}, nil
}
