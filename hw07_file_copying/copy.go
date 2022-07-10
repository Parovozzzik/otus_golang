package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	flag.Parse()

	var file *os.File
	file, err := os.OpenFile(fromPath, os.O_RDWR, 0o644)
	if err != nil && os.IsNotExist(err) {
		return ErrUnsupportedFile
	}
	defer file.Close()

	fi, _ := file.Stat()
	fileSize := fi.Size()
	if fileSize < offset {
		return ErrOffsetExceedsFileSize
	}

	if !fi.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	start := int64(0)
	if offset != 0 {
		start = offset
	}

	if (limit+offset) > fileSize || limit == 0 {
		limit = fileSize - offset
	}

	data := make([]byte, limit)
	for start < fileSize {
		read, err := file.ReadAt(data, start)
		start += int64(read)

		if err == io.EOF || start >= limit {
			break
		}
	}

	emptyFile, _ := os.Create(toPath)
	defer emptyFile.Close()

	bar := pb.StartNew(int(limit))
	defer bar.Finish()

	for i := 0; i < int(limit); i++ {
		bar.Increment()
		emptyFile.Write(data[i : i+1])
	}

	return nil
}
