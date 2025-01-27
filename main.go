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
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	if eventPath == "" {
		log.Fatal("GITHUB_EVENT_PATH is not set")
	}

	cmd := exec.Command("sh", "-c", fmt.Sprintf("cat %s | jq .issue.number", eventPath))
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("コマンド実行に失敗しました。: %v", err)
	}

	issueNumber := strings.TrimSpace(string(output))
	fmt.Printf("Issue number: %s\n", issueNumber)
}
