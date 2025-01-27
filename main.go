package main

import (
	"fmt"
	"os"
)

// GitHub ActionsのIssue作成時に実行されることを想定
// Issue番号を取得し表示する
func main() {
	eventPath, exists := os.LookupEnv("GITHUB_EVENT_PATH")
	if !exists {
		fmt.Println("環境変数 GITHUB_EVENT_PATH が見つかりません")
		return
	}
	fmt.Println("GITHUB_EVENT_PATH:", eventPath)

	// eventPathからファイルを読み込む
	file, err := os.Open(eventPath)
	if err != nil {
		fmt.Println("ファイルが開けませんでした")
		return
	}
	defer file.Close()

	// Issue番号を取得
	var issueNumber string
	fmt.Fscanf(file, `{"issue":{"number":%s`, &issueNumber)
	fmt.Println("Issue番号:", issueNumber)
}
