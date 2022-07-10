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
	size := fileSize
	if offset != 0 {
		start = offset
		size = fileSize - offset
	}

	barCount := 1
	extraBarCount := 0
	if limit != 0 {
		barCount = int(size) / int(limit)
		if size%limit > 0 {
			extraBarCount = 1
		}
	} else {
		limit = fileSize
	}

	bar := pb.StartNew(barCount + extraBarCount)
	defer bar.Finish()

	data := make([]byte, size)
	for start < fileSize {
		bar.Increment()
		dataCurrent := make([]byte, limit)
		read, err := file.ReadAt(dataCurrent, start)
		start += int64(read)

		if err == io.EOF || (start-limit) >= size {
			copy(data[size-(size%limit):], dataCurrent[:(size%limit)])
			break
		} else {
			copy(data[(start-limit-offset):(start-offset)], dataCurrent[:limit])
		}
	}

	emptyFile, _ := os.Create(toPath)
	defer emptyFile.Close()
	emptyFile.Write(data)

	return nil
}
