package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	if eventPath == "" {
		log.Fatal("GITHUB_EVENT_PATH is not set")
	}

	issueNumber := getIssueNumber(eventPath)
	fmt.Println("Issue Number:", issueNumber)

	// ここで呼び出し元リポジトリを取得
	// (action.yamlの inputs.target-repo を env:"TARGET_REPO" で受け取る)
	repo := os.Getenv("TARGET_REPO")
	if repo == "" {
		log.Fatal("TARGET_REPO is not set")
	}

	// --repo でリポジトリを明示的に指定
	cmd := exec.Command(
		"gh", "issue", "develop", issueNumber,
		"--repo", repo,
		"-n", fmt.Sprintf("issue-%s", issueNumber),
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("コマンド実行に失敗しました。: %v\n出力: %s", err, string(output))
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
