package examples

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/serhiy-t/errf"
)

// GzipFileErrflow compresses file srcFilename into dstFilename.
func GzipFileErrflow(dstFilename string, srcFilename string) (err error) {
	// defer IfError()... creates and configures
	// ErrorFlow error handler for this function.
	// When any of Check* functions encounters non-nil error
	// it immediately sends error to this handler
	// unwinding all stacked defers.
	errWrapper := errf.WrapperFmtErrorw("error compressing file")
	defer errf.IfError().ReturnFirst().LogIfSuppressed().Apply(errWrapper).ThenAssignTo(&err)

	errf.CheckCondition(len(dstFilename) == 0, "dst file should be specified")
	errf.CheckCondition(len(srcFilename) == 0, "src file should be specified")

	reader := errf.Io.CheckReadCloser(os.Open(srcFilename))
	defer errf.With(errWrapper).Log(reader.Close())

	writer := errf.Io.CheckWriteCloser(os.Create(dstFilename))
	defer errf.CheckErr(writer.Close())

	gzipWriter := gzip.NewWriter(writer)
	defer errf.CheckErr(gzipWriter.Close())

	return errf.CheckDiscard(io.Copy(gzipWriter, reader))
}

// GzipFilePlainGo compresses file srcFilename into dstFilename.
func GzipFilePlainGo(dstFilename string, srcFilename string) (err error) {
	if len(dstFilename) == 0 {
		return fmt.Errorf("error compressing file: dst file should be specified")
	}
	if len(srcFilename) == 0 {
		return fmt.Errorf("error compressing file: src file should be specified")
	}

	reader, err := os.Open(srcFilename)
	if err != nil {
		return fmt.Errorf("error compressing file: %w", err)
	}
	defer func() {
		closeErr := reader.Close()
		if closeErr != nil {
			log.Println(closeErr)
		}
	}()

	writer, err := os.Create(dstFilename)
	if err != nil {
		return fmt.Errorf("error compressing file: %w", err)
	}
	defer func() {
		closeErr := writer.Close()
		if closeErr != nil {
			if err == nil {
				err = fmt.Errorf("error compressing file: %w", closeErr)
			} else {
				log.Println(fmt.Errorf("[suppressed] error compressing file: %w", closeErr))
			}
		}
	}()

	gzipWriter := gzip.NewWriter(writer)
	defer func() {
		closeErr := gzipWriter.Close()
		if closeErr != nil {
			if err == nil {
				err = fmt.Errorf("error compressing file: %w", closeErr)
			} else {
				log.Println(fmt.Errorf("[suppressed] error compressing file: %w", closeErr))
			}
		}
	}()

	_, err = io.Copy(gzipWriter, reader)
	if err != nil {
		return fmt.Errorf("error compressing file: %w", err)
	}

	return nil
}

// GzipFileErrflowLite compresses file srcFilename into dstFilename.
func GzipFileErrflowLite(dstFilename string, srcFilename string) (err error) {
	errflow := errf.With(
		errf.LogStrategyIfSuppressed,
		errf.WrapperFmtErrorw("error compressing file"),
	)

	if len(dstFilename) == 0 {
		return fmt.Errorf("error compressing file: dst file should be specified")
	}
	if len(srcFilename) == 0 {
		return fmt.Errorf("error compressing file: src file should be specified")
	}

	reader, err := os.Open(srcFilename)
	if err != nil {
		return fmt.Errorf("error compressing file: %w", err)
	}
	defer errflow.Log(reader.Close())

	writer, err := os.Create(dstFilename)
	if err != nil {
		return fmt.Errorf("error compressing file: %w", err)
	}
	defer errflow.IfErrorAssignTo(&err, writer.Close())

	gzipWriter := gzip.NewWriter(writer)
	defer errflow.IfErrorAssignTo(&err, gzipWriter.Close())

	_, err = io.Copy(gzipWriter, reader)
	if err != nil {
		return fmt.Errorf("error compressing file: %w", err)
	}

	return nil
}
