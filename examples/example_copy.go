package examples

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	errf "github.com/serhiy-t/errflow"
)

// CopyFileErrflow copies file srcFilename into dstFilename.
func CopyFileErrflow(dstFilename string, srcFilename string) (err error) {
	defer errf.IfError().LogIfSuppressed().Apply(
		errf.WrapperFmtErrorw("error copying file"),
	).ThenAssignTo(&err)

	errf.TryCondition(len(dstFilename) == 0, "dst file should be specified")
	errf.TryCondition(len(srcFilename) == 0, "src file should be specified")

	reader := errf.Io.TryReadCloser(
		os.Open(srcFilename))
	defer errf.Log(
		reader.Close())

	writer := errf.Io.TryWriteCloser(
		os.Create(dstFilename))
	defer errf.TryErr(
		writer.Close())

	bufWriter := bufio.NewWriter(writer)
	defer errf.TryErr(
		bufWriter.Flush())

	return errf.TryDiscard(
		io.Copy(bufWriter, bufio.NewReader(reader)))
}

// CopyFilePlainGo copies file srcFilename into dstFilename.
func CopyFilePlainGo(dstFilename string, srcFilename string) (err error) {
	if len(dstFilename) == 0 {
		return fmt.Errorf("error copying file: dst file should be specified")
	}
	if len(srcFilename) == 0 {
		return fmt.Errorf("error copying file: src file should be specified")
	}

	reader, err := os.Open(srcFilename)
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}
	defer func() {
		closeErr := reader.Close()
		if closeErr != nil {
			log.Println(closeErr)
		}
	}()

	writer, err := os.Create(dstFilename)
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}
	defer func() {
		closeErr := writer.Close()
		if closeErr != nil {
			if err == nil {
				err = fmt.Errorf("error copying file: %w", closeErr)
			} else {
				log.Println(fmt.Errorf("[suppressed] error copying file: %w", closeErr))
			}
		}
	}()

	bufWriter := bufio.NewWriter(writer)
	defer func() {
		closeErr := bufWriter.Flush()
		if closeErr != nil {
			if err == nil {
				err = fmt.Errorf("error copying file: %w", closeErr)
			} else {
				log.Println(fmt.Errorf("[suppressed] error copying file: %w", closeErr))
			}
		}
	}()

	_, err = io.Copy(bufWriter, bufio.NewReader(reader))
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}

	return nil
}

// CopyFileErrflowNoPanics copies file srcFilename into dstFilename.
func CopyFileErrflowNoPanics(dstFilename string, srcFilename string) (err error) {
	if len(dstFilename) == 0 {
		return fmt.Errorf("error copying file: dst file should be specified")
	}
	if len(srcFilename) == 0 {
		return fmt.Errorf("error copying file: src file should be specified")
	}

	reader, err := os.Open(srcFilename)
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}
	defer errf.Log(
		reader.Close())

	writer, err := os.Create(dstFilename)
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}
	defer errf.With(
		errf.LogStrategyIfSuppressed,
		errf.WrapperFmtErrorw("error copying file"),
	).IfErrorAssignTo(&err,
		writer.Close())

	bufWriter := bufio.NewWriter(writer)
	defer errf.With(
		errf.LogStrategyIfSuppressed,
		errf.WrapperFmtErrorw("error copying file"),
	).IfErrorAssignTo(&err,
		bufWriter.Flush())

	_, err = io.Copy(bufWriter, bufio.NewReader(reader))
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}

	return nil
}
