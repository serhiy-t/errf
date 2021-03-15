# **err**or**f**low
Declarative error handling for Go.

## Motivational example

```go
func CopyFile(dstFilename string, srcFilename string) error {
	reader, _ := os.Open(srcFilename)
	defer reader.Close()

	writer, _ := os.Create(dstFilename)
	defer writer.Close()

	bufWriter := bufio.NewWriter(writer)
	defer bufWriter.Flush()

	_, _ = io.Copy(bufWriter, bufio.NewReader(reader))

	return nil
}
```

```go
func CopyFile(dstFilename string, srcFilename string) (err error) {
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
```

```go
func CopyFile(dstFilename string, srcFilename string) (err error) {
	defer errf.IfError().LogIfSuppressed().Apply(
		errf.WrapperFmtErrorw("error copying file"),
	).ThenAssignTo(&err)

	errf.CheckCondition(len(dstFilename) == 0, "dst file should be specified")
	errf.CheckCondition(len(srcFilename) == 0, "src file should be specified")

	reader := errf.Io.CheckReadCloser(
		os.Open(srcFilename))
	defer errf.Log(
		reader.Close())

	writer := errf.Io.CheckWriteCloser(
		os.Create(dstFilename))
	defer errf.CheckErr(
		writer.Close())

	bufWriter := bufio.NewWriter(writer)
	defer errf.CheckErr(
		bufWriter.Flush())

	return errf.CheckDiscard(
		io.Copy(bufWriter, bufio.NewReader(reader)))
}
```

```go
func CopyFile(dstFilename string, srcFilename string) (err error) {
	errflow := errf.With(
		errf.LogStrategyIfSuppressed,
		errf.WrapperFmtErrorw("error copying file"),
	)

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
	defer errflow.Log(
		reader.Close())

	writer, err := os.Create(dstFilename)
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}
	defer errflow.IfErrorAssignTo(&err,
		writer.Close())

	bufWriter := bufio.NewWriter(writer)
	defer errflow.IfErrorAssignTo(&err,
		bufWriter.Flush())

	_, err = io.Copy(bufWriter, bufio.NewReader(reader))
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}

	return nil
}
```