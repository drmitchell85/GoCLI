package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type Article struct {
	Id      string `json:"article_id"`
	Title   string `json:"articleTitle"`
	Desc    string `json:"articleDesc"`
	Content string `json:"articleContent"`
}

type ArticleResponse struct {
	Status   string    `json:"Status"`
	Code     int       `json:"Code"`
	Count    int       `json:"Count"`
	Articles []Article `json:"Articles"`
}

func main() {
	cliFlags()
}

func cliFlags() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'API' command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "getAllArticles":
		fmt.Println("getting all articles...")
		fmt.Println("curl -v http://localhost:10000/articles/list")
		excCmd()
	default:
		fmt.Println("API does not exist...")
	}
}

func excCmd() {
	cmd, err := exec.Command("bash", "-c", "curl -v http://localhost:10000/articles/list").Output()
	if err != nil {
		fmt.Println(err)
	}

	var articleResponse ArticleResponse
	err = json.Unmarshal(cmd, &articleResponse)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("articleResponse: ", articleResponse)
}
