package examples

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/serhiy-t/errflow"
)

// CopyPlainGoBroken is an example of copying a file, using plain idiomatic Go.
// Warning: this implementation is broken.
func CopyPlainGoBroken(dstFilename string, srcFilename string) error {
	if len(dstFilename) == 0 {
		return fmt.Errorf("dst file should be specified")
	}

	if len(srcFilename) == 0 {
		return fmt.Errorf("src file should be specified")
	}

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
	if err != nil {
		return err
	}

	return nil
}

// CopyErrflow is an example of copying a file, using idiomatic errflow.
// Unlike CopyPlainGo, this implementation is not broken.
func CopyErrflow(dstFilename string, srcFilename string) (err error) {
	defer errflow.IfError().LogIfSuppressed().ThenAssignTo(&err)

	errflow.TryCondition(len(dstFilename) == 0, "dst file should be specified")
	errflow.TryCondition(len(srcFilename) == 0, "src file should be specified")

	reader := errflow.Io.TryReadCloser(os.Open(srcFilename))
	defer errflow.Log(reader.Close())

	writer := errflow.Io.TryWriteCloser(os.Create(dstFilename))
	defer errflow.TryErr(writer.Close())

	return errflow.TryDiscard(io.Copy(writer, reader))
}

// CopyPlainGoCorrectCopyPaste is an example of copying a file, using plain idiomatic Go.
func CopyPlainGoCorrectCopyPaste(dstFilename string, srcFilename string) error {
	reader, err := os.Open(srcFilename)
	if err != nil {
		return err
	}
	defer reader.Close()

	writer, err := os.Create(dstFilename)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, reader)
	if err != nil {
		writer.Close()
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	return nil
}

// CopyPlainGoCorrectDeferMagic is an example of copying a file, using plain idiomatic Go.
func CopyPlainGoCorrectDeferMagic(dstFilename string, srcFilename string) error {
	reader, err := os.Open(srcFilename)
	if err != nil {
		return err
	}
	defer reader.Close()

	writer, err := os.Create(dstFilename)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := writer.Close()
		if closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	_, err = io.Copy(writer, reader)
	if err != nil {
		return err
	}

	return nil
}

// CopyPlainGoCorrectDeferMagicHelper is an example of copying a file, using plain idiomatic Go.
func CopyPlainGoCorrectDeferMagicHelper(dstFilename string, srcFilename string) (err error) {
	if len(dstFilename) == 0 {
		return fmt.Errorf("dst file should be specified")
	}

	if len(srcFilename) == 0 {
		return fmt.Errorf("src file should be specified")
	}

	reader, err := os.Open(srcFilename)
	if err != nil {
		return err
	}
	defer errflow.Log(reader.Close())

	writer, err := os.Create(dstFilename)
	if err != nil {
		return err
	}
	defer errflow.With(errflow.LogStrategyIfSuppressed).IfErrorAssignTo(writer.Close(), &err)

	_, err = io.Copy(writer, reader)
	if err != nil {
		return err
	}

	return nil
}

// CopyErrflow is an example of copying a file, using idiomatic errflow.
// Unlike CopyPlainGo, this implementation is not broken.
func CopyBufferedErrflow(dstFilename string, srcFilename string) (n int64, err error) {
	defer errflow.IfError().LogIfSuppressed().ThenAssignTo(&err)

	errflow.TryCondition(len(dstFilename) == 0, "dst file should be specified")
	errflow.TryCondition(len(srcFilename) == 0, "src file should be specified")

	reader := errflow.Io.TryReadCloser(os.Open(srcFilename))
	defer errflow.Log(reader.Close())

	writer := errflow.Io.TryWriteCloser(os.Create(dstFilename))
	defer errflow.TryErr(writer.Close())

	bufWriter := bufio.NewWriter(writer)
	defer errflow.TryErr(bufWriter.Flush())

	return errflow.Std.TryInt64Err(io.Copy(bufWriter, bufio.NewReader(reader)))
}

// CopyBufferedPlainGoBroken is an example of copying a file, using plain idiomatic Go.
// Warning: this implementation is broken.
func CopyBufferedPlainGoBroken(dstFilename string, srcFilename string) error {
	if len(dstFilename) == 0 {
		return fmt.Errorf("dst file should be specified")
	}

	if len(srcFilename) == 0 {
		return fmt.Errorf("src file should be specified")
	}

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

	bufWriter := bufio.NewWriter(writer)
	if err != nil {
		return err
	}
	defer bufWriter.Flush()

	_, err = io.Copy(bufWriter, bufio.NewReader(reader))
	if err != nil {
		return err
	}

	return nil
}

func WrapperCaptureStacktrace(ef *errflow.Errflow) *errflow.Errflow {
	return ef.With(errflow.Wrapper(ErrorWithStackTrace))
}

func ErrorWithStackTrace(err error) error {
	// ...
	return nil
}

// CopyPlainGoCorrectDeferMagicHelper is an example of copying a file, using plain idiomatic Go.
func CopyBufferedPlainGoCorrectDeferMagicHelper(dstFilename string, srcFilename string) (n int64, err error) {
	if len(dstFilename) == 0 {
		return 0, fmt.Errorf("dst file should be specified")
	}

	if len(srcFilename) == 0 {
		return 0, fmt.Errorf("src file should be specified")
	}

	reader, err := os.Open(srcFilename)
	if err != nil {
		return 0, fmt.Errorf("error opening file for read: %w", err)
	}
	defer errflow.Log(reader.Close())

	writer, err := os.Create(dstFilename)
	if err != nil {
		return 0, fmt.Errorf("error creating file: %w", err)
	}
	defer errflow.With(errflow.LogStrategyAlways).
		IfErrorAssignTo(writer.Close(), &err)

	bufWriter := bufio.NewWriter(writer)
	defer errflow.With(errflow.LogStrategyAlways).
		IfErrorAssignTo(bufWriter.Flush(), &err)

	n, err = io.Copy(bufWriter, bufio.NewReader(reader))
	if err != nil {
		return 0, err
	}

	return n, nil
}

// CopyPlainGoCorrectDeferMagic is an example of copying a file, using plain idiomatic Go.
func CopyBufferedPlainGoCorrectDeferMagic(dstFilename string, srcFilename string) error {
	if len(dstFilename) == 0 {
		return fmt.Errorf("dst file should be specified")
	}

	if len(srcFilename) == 0 {
		return fmt.Errorf("src file should be specified")
	}

	reader, err := os.Open(srcFilename)
	if err != nil {
		return err
	}
	defer reader.Close()

	writer, err := os.Create(dstFilename)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := writer.Close()
		if closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	bufWriter := bufio.NewWriter(writer)
	defer func() {
		closeErr := bufWriter.Flush()
		if closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	_, err = io.Copy(bufWriter, bufio.NewReader(reader))
	if err != nil {
		return err
	}

	return nil
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

func copy2Errflow(dstFilename string, srcFilename string) (n int64, err error) {
	defer errflow.IfError().ReturnFirst().LogIfSuppressed().ThenAssignTo(&err)

	errflow.TryCondition(len(dstFilename) == 0, "dst file should be specified")
	errflow.TryCondition(len(srcFilename) == 0, "src file should be specified")

	reader := errflow.Io.TryReadCloser(os.Open(srcFilename))
	defer errflow.Log(reader.Close())

	writer := errflow.Io.TryWriteCloser(os.Create(dstFilename))
	defer errflow.With(errflow.WrapperFmtErrorW("error closing file")).
		TryErr(writer.Close())

	bufWriter := bufio.NewWriter(writer)
	defer errflow.TryErr(bufWriter.Flush())

	return errflow.Std.TryInt64Err(io.Copy(bufWriter, bufio.NewReader(reader)))
}

func copy2PlainGo(dstFilename string, srcFilename string) error {
	if len(dstFilename) == 0 {
		return fmt.Errorf("dst file should be specified")
	}

	if len(srcFilename) == 0 {
		return fmt.Errorf("src file should be specified")
	}

	reader, err := os.Open(srcFilename)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer reader.Close()

	writer, err := os.Create(dstFilename)
	if err != nil {
		return err
	}
	defer writer.Close()

	bufWriter := bufio.NewWriter(writer)
	if err != nil {
		return err
	}
	defer bufWriter.Flush()

	_, err = io.Copy(bufWriter, bufio.NewReader(reader))
	if err != nil {
		return err
	}

	return nil
}

//
//
//
//

func copyPlainGo(dstFilename string, srcFilename string) (err error) {
	if len(dstFilename) == 0 {
		return fmt.Errorf("dst file should be specified")
	}

	if len(srcFilename) == 0 {
		return fmt.Errorf("src file should be specified")
	}

	reader, err := os.Open(srcFilename)
	if err != nil {
		return err
	}
	defer reader.Close()

	writer, err := os.Create(dstFilename)
	if err != nil {
		return err
	}
	defer func() {
		err2 := writer.Close()
		if err == nil {
			err = err2
		}
	}()

	bufWriter := bufio.NewWriter(writer)
	if err != nil {
		return err
	}
	defer func() {
		err2 := bufWriter.Flush()
		if err == nil {
			err = err2
		}
	}()

	_, err = io.Copy(bufWriter, bufio.NewReader(reader))
	if err != nil {
		return err
	}

	return nil
}

func copyErrflow(dstFilename string, srcFilename string) (err error) {
	defer errflow.IfError().LogIfSuppressed().ThenAssignTo(&err)

	errflow.TryCondition(len(dstFilename) == 0, "dst file should be specified")
	errflow.TryCondition(len(srcFilename) == 0, "src file should be specified")

	reader := errflow.Io.TryReadCloser(
		os.Open(srcFilename))
	defer errflow.Log(
		reader.Close())

	writer := errflow.Io.TryWriteCloser(
		os.Create(dstFilename))
	defer errflow.TryErr(
		writer.Close())

	bufWriter := bufio.NewWriter(writer)
	defer errflow.TryErr(
		bufWriter.Flush())

	return errflow.TryDiscard(
		io.Copy(bufWriter, bufio.NewReader(reader)))
}

// func copyErrflowFantasy(dstFilename string, srcFilename string) (err error) {
// 	defer errflow.IfError().LogIfSuppressed().ThenAssignTo(&err)

// 	errflow.TryCondition(len(dstFilename) == 0, "dst file should be specified")
// 	errflow.TryCondition(len(srcFilename) == 0, "src file should be specified")

// 	reader := os.Open(srcFilename) |> errflow.Io.TryReadCloser
// 	defer reader.Close() |> errflow.Log

// 	writer := os.Create(dstFilename) |> errflow.Io.TryWriteCloser
// 	defer writer.Close() |> errflow.Try

// 	bufWriter := bufio.NewWriter(writer)
// 	defer bufWriter.Flush() |> errflow.Try

// 	return io.Copy(bufWriter, bufio.NewReader(reader)) |> errflow.TryDiscard
// }

// func copyErrflowFantasy2(dstFilename string, srcFilename string) (err error) {
// 	defer errflow.IfError().LogIfSuppressed().ThenAssignTo(&err)

// 	errflow.TryCondition(len(dstFilename) == 0, "dst file should be specified")
// 	errflow.TryCondition(len(srcFilename) == 0, "src file should be specified")

// 	reader := os.Open(srcFilename) |>
// 		errflow.With(errflow.WrapperFmtErrorW("error reading file")).Try1
// 	defer reader.Close() |> errflow.Log

// 	writer := os.Create(dstFilename) |> errflow.Try1
// 	defer writer.Close() |> errflow.Try

// 	bufWriter := bufio.NewWriter(writer)
// 	defer bufWriter.Flush() |> errflow.Try

// 	return io.Copy(bufWriter, bufio.NewReader(reader)) |> errflow.Try1
// }
