package main

import (
	"encoding/json"
	"flag"
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

	cmdGetArtByID := flag.NewFlagSet("getArticleByID", flag.ExitOnError)
	article_id := cmdGetArtByID.String("article_id", "", "id of the article to be searched")
	flag.Parse()

	switch os.Args[1] {
	case "getAllArticles":
		fmt.Println("getting all articles...")
		path := fmt.Sprintln("curl -v http://localhost:10000/articles/list")
		fmt.Println(path)
		excCmd(path)
	case "getArticleByID":
		cmdGetArtByID.Parse(os.Args[2:])
		fmt.Println("getting article by id: ", *article_id)
		path := fmt.Sprintf("curl -v http://localhost:10000/articles/fetch?article_id=%s", *article_id)
		fmt.Println(path)
		excCmd(path)
	default:
		fmt.Println("that API does not exist...")
	}
}

func excCmd(path string) {
	cmd, err := exec.Command("bash", "-c", path).Output()
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
