package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
)

var resolutions = []string{
	"640x960",
	"640x1136",
	"720x1280",
	"750x1334",
	"1080x1920",
	"1125x2436",
	"1366x768",
	"1440x900",
	"1920x1080",
	"2048x1536",
	"2208x1242",
	"2436x1125",
	"2732x2048",
}

// 指定した解像度の画像をパッケージングする
func Packager() {

	prompt := promptui.Select{
		// 選択肢のタイトル
		Label: "解像度を選択してください",
		// 選択肢の配列
		Items: resolutions,
	}

	_, resolutionInput, err := prompt.Run()

	if err != nil {
		fmt.Printf("コマンドが正しくありません: %v\n", err)
		return
	}

	checkDir(fmt.Sprintf("%s/%s", getDistDir(), resolutionInput))

	filepath.WalkDir(getDownloadDir(), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return errors.Wrap(err, "failed filepath.WalkDir")
		}

		if d.IsDir() {
			return nil
		}

		if strings.Contains(path, ".DS_Store") {
			return nil
		}

		if strings.Contains(path, resolutionInput) {

			newPath := convertToDistPath(path, resolutionInput)

			if reader, err := os.Open(path); err == nil {
				defer reader.Close()

				if err != nil {
					return errors.Wrap(err, "reader.Stat")
				}

				// Directory exists and is writable
				err = CopyFile(reader, newPath)

				if err != nil {
					return errors.Wrap(err, "CopyFile")
				}
				fmt.Println(path + " -> " + newPath)

			} else {
				return errors.Wrap(err, "Impossible to open the file")
			}

		}

		return err
	})
}
