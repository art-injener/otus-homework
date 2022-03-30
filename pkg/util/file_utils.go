package util

import (
	"fmt"
	"os"
)

func WriteFile(pathFile string, fileBytes []byte, fileName string) error {
	err := os.MkdirAll(pathFile, os.ModePerm)
	if err != nil {
		return err
	}

	fileOnDisk, err := os.Create(fmt.Sprintf("%s/%s", pathFile, fileName))
	if err != nil {
		return err
	}
	defer fileOnDisk.Close()

	_, err = fileOnDisk.Write(fileBytes)
	if err != nil {
		return err
	}
	return nil
}
