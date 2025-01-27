package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// GitHub ActionsのIssue作成時に実行されることを想定
// Issue番号を取得し表示する
func main() {
	// GitHub Actions実行環境でのイベント情報が格納されているファイルパスを取得
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	if eventPath == "" {
		log.Fatal("GITHUB_EVENT_PATH is not set")
	}

	// Issue番号を取得
	issueNumber := getIssueNumber(eventPath)
	fmt.Println("Issue Number:", issueNumber)

	// ghコマンドでIssueに関連付けした、issue-[Issue番号]ブランチを作成
	// gh issue develop [issue番号] -n issue-[issue番号]
	cmd := exec.Command("gh", "issue", "develop", issueNumber, "-n", fmt.Sprintf("issue-%s", issueNumber))
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("コマンド実行に失敗しました。: %v", err)
	}

	fmt.Println(string(output))
}

func getIssueNumber(eventPath string) string {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("cat %s | jq .issue.number", eventPath))
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("コマンド実行に失敗しました。: %v", err)
	}

	return strings.TrimSpace(string(output))
}
