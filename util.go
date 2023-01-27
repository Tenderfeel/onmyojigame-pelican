package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// ディレクトリ確認（なかったら作る）
func checkDir(path string) {
	// なかったら作る
	if !exists(path) {
		os.MkdirAll(path, 0777)
	}
}

// ファイル存在確認
func exists(path string) bool {

	if _, err := os.Stat(path); err == nil {
		// path exists
		return true

	} else if errors.Is(err, os.ErrNotExist) {
		// path does *not* exist
		return false

	} else {
		// Schrodinger: file may or may not exist. See err for details.

		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence

		panic(err)
	}

}

func parsedImageUrl(href string, contetType string) (string, string, error) {
	arr := strings.Split(href, "/")

	ext := ""

	switch contetType {
	case "image/webp":
		ext = "webp"
	case "image/jpeg":
		ext = "jpg"
	case "image/png":
		ext = "png"
	}

	switch len(arr) {
	case 11:
		sizeAndExt := strings.Split(arr[10], ".")
		imageSize := sizeAndExt[0]
		imageExt := strings.Split(sizeAndExt[1], "?")[0] // remove query
		if ext != "" {
			imageExt = ext
		}
		fileName := fmt.Sprintf("%s-%s.%s", arr[8], arr[9], imageExt)
		dirName := fmt.Sprintf("%s/%s/%s/", getDownloadDir(), arr[8], imageSize)
		imagePath := fmt.Sprintf("%s%s", dirName, fileName)

		return dirName, imagePath, nil
	case 12:
		sizeAndExt := strings.Split(arr[11], ".")
		imageSize := sizeAndExt[0]
		imageExt := strings.Split(sizeAndExt[1], "?")[0]
		if ext != "" {
			imageExt = ext
		}
		fileName := fmt.Sprintf("%s-%s-%s.%s", arr[8], arr[9], arr[10], imageExt)
		dirName := fmt.Sprintf("%s/%s/%s/", getDownloadDir(), arr[8], imageSize)
		imagePath := fmt.Sprintf("%s%s", dirName, fileName)

		return dirName, imagePath, nil

	case 9:
		sizeAndExt := strings.Split(arr[8], ".")
		imageSize := sizeAndExt[0]
		imageExt := strings.Split(sizeAndExt[1], "?")[0]
		if ext != "" {
			imageExt = ext
		}
		fileName := fmt.Sprintf("%s-%s.%s", arr[6], arr[7], imageExt)
		dirName := fmt.Sprintf("%s/%s/%s/", getDownloadDir(), arr[6], imageSize)
		imagePath := fmt.Sprintf("%s%s", dirName, fileName)

		return dirName, imagePath, nil

	case 10:
		sizeAndExt := strings.Split(arr[9], ".")
		imageSize := sizeAndExt[0]
		imageExt := strings.Split(sizeAndExt[1], "?")[0]
		if ext != "" {
			imageExt = ext
		}
		fileName := fmt.Sprintf("%s-%s-%s.%s", arr[6], arr[7], arr[8], imageExt)
		dirName := fmt.Sprintf("%s/%s/%s/", getDownloadDir(), arr[6], imageSize)
		imagePath := fmt.Sprintf("%s%s", dirName, fileName)

		return dirName, imagePath, nil
	}

	return "", "", errors.New("parsedImageUrl: invalid length string")
}

// ダウンロードディレクトリを返す
func getDownloadDir() string {
	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	downloadDir := fmt.Sprintf("%s/%s", path, downloadDirName)

	return downloadDir
}

// 出力用ディレクトリを返す
func getDistDir() string {
	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	return fmt.Sprintf("%s/%s", path, distDirName)

}

func convertToDistPath(beforePath string, resolution string) string {
	path := strings.Replace(beforePath, getDownloadDir(), "", 1)
	arr := strings.Split(path, "/")

	fileName := arr[len(arr)-1]

	return fmt.Sprintf("%s/%s/%s", getDistDir(), resolution, fileName)
}

func CopyFile(in io.Reader, dst string) (err error) {

	// Does file already exist? Skip
	if _, err := os.Stat(dst); err == nil {
		return nil
	}

	err = nil

	out, err := os.Create(dst)
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}

	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()

	if _, err = io.Copy(out, in); err != nil {
		fmt.Println("io.Copy error")
		return
	}
	return
}
