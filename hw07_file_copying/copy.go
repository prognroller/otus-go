package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrNegativeParams        = errors.New("negative value of offset or limit")
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if offset < 0 || limit < 0 {
		return ErrNegativeParams
	}

	copyFile, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer copyFile.Close()

	copyFileStat, err := copyFile.Stat()
	if err != nil {
		return err
	}
	if copyFileStat.Size() < offset {
		return ErrOffsetExceedsFileSize
	}

	newFile, err := ioutil.TempFile("", "fedya")
	if err != nil {
		return err
	}

	_, err = copyFile.Seek(offset, 0)
	if err != nil {
		return err
	}

	if (limit == 0) || (limit > copyFileStat.Size()-offset) {
		limit = copyFileStat.Size() - offset
	}

	reader := io.LimitReader(copyFile, limit)

	bar := pb.Full.Start64(limit)

	barReader := bar.NewProxyReader(reader)

	_, err = io.Copy(newFile, barReader)
	if err != nil {
		return err
	}

	bar.Finish()

	dirPath, err := os.Getwd()
	if err != nil {
		return err
	}
	newPath := fmt.Sprintf("%v/%v", dirPath, toPath)
	err = os.Rename(newFile.Name(), newPath)
	if err != nil {
		return err
	}

	return nil
}
