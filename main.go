package main

import (
	"fmt"
	"os"
	"os/exec"
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

	// cat $GITHUB_EVENT_PATH | jq .issue.number
	exec.Command("cat", eventPath).Run()
}
