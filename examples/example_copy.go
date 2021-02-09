package examples

import (
	"io"
	"os"

	"github.com/serhiy-t/errflow"
)

// CopyPlainGo is an example of copying a file, using plain idiomatic Go.
// Warning: this implementation is broken.
func CopyPlainGo(dstFilename string, srcFilename string) error {
	reader, err := os.Open(srcFilename)
	if err != nil {
		return err
	}
	defer reader.Close()

	writer, err := os.Create(dstFilename)
	if err != nil {
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, reader)

	return err
}

// CopyErrflow is an example of copying a file, using idiomatic errflow.
func CopyErrflow(dstFilename string, srcFilename string) (err error) {
	defer errflow.IfError().ReturnFirst().LogAll().ThenAssignTo(&err)

	reader := errflow.CheckIoReadCloser(os.Open(srcFilename))
	defer errflow.OnlyLog(reader.Close())

	writer := errflow.CheckIoWriteCloser(os.Create(dstFilename))
	defer errflow.Check(writer.Close())

	return errflow.CheckIgnoreValue(io.Copy(writer, reader))
}
