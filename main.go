package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

const pageUrl = "https://yys.163.com/media/picture.html"
const downloadDirName = "downloads"
const distDirName = "dist"

// const userAgentIOS = "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1"
const userAgentPC = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

func main() {
	fmt.Println("🐣 Pelican Start!")

	prompt := promptui.Select{
		// 選択肢のタイトル
		Label: "何をしますか？",
		// 選択肢の配列
		Items: []string{"Download", "Package"},
	}

	promptIndex, _, err := prompt.Run() //入力を受け取る

	if err != nil {
		fmt.Printf("コマンドが正しくありません: %v\n", err)
		return
	}

	switch promptIndex {
	case 0:
		Downloader()
	case 1:
		Packager()
	default:
		fmt.Println("コマンドが正しくありません")
	}
}
