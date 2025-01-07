package copier

import (
	"io"
)

func CopySourceToDest(source io.Reader, dest io.Writer) error {
	_, err := io.Copy(dest, source)
	return err
}
