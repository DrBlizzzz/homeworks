package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	var progress int64
	var step int64
	fileInfo, _ := os.Stat(fromPath)
	fileSize := fileInfo.Size()
	// если файл бесконечный то результат будет 0
	if fileSize == 0 {
		return ErrUnsupportedFile
	}
	if fileSize < offset {
		return ErrOffsetExceedsFileSize
	}
	if limit == 0 || limit > fileSize {
		step = fileSize
		progress = 1
	} else {
		step = 10
		if offset+limit > fileSize {
			progress = ((fileSize - offset) / step) + 1
		} else {
			progress = limit / step
		}
	}
	reader, _ := os.Open(fromPath)
	writer, _ := os.Create(toPath)
	defer reader.Close()
	defer writer.Close()
	bar := pb.StartNew(int(progress))
	defer bar.Finish()
	reader.Seek(offset, io.SeekStart)
	for i := 0; i < int(progress); i++ {
		bar.Increment()
		_, err := io.CopyN(writer, reader, step)
		if errors.Is(err, io.EOF) {
			break
		}
	}
	return nil
}
