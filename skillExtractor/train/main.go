package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command(
		"/Users/himanshu/fastText/fasttext",
		"skipgram",
		"-input", "../skills_corpus.txt",
		"-output", "skill_model",
		"-dim", "150",
		"-ws", "8",
		"-epoch", "15",
		"-minCount", "1",
		"-neg", "15",
		"-minn", "2",
		"-maxn", "6",
		"-thread", "16",
		"-lr", "0.05",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
