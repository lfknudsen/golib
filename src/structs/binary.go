package structs

import (
	"fmt"
	"io"
)

type Binary struct {
	length uint32
	data   []byte
}

func (b *Binary) Write(w io.Writer) error {
	bytesWritten, err := w.Write(b.data)
	if err != nil {
		return err
	}
	if uint32(bytesWritten) != b.length {
		return fmt.Errorf("expected to write %d, got %d", b.length, bytesWritten)
	}
	return nil
}

func (b *Binary) Read(r io.Reader) error {
	bytesRead, err := r.Read(b.data)
	if err != nil {
		return err
	}
	b.length = uint32(bytesRead)
	return err
}
