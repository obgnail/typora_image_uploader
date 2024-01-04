package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"io/ioutil"
	"path"
	"strconv"
	"time"
)

var (
	token  string
	owner  string
	repo   string
	images []string
)

func initVar() {
	flag.StringVar(&token, "token", "", "GitHub personal access token")
	flag.StringVar(&owner, "owner", "", "GitHub repository owner")
	flag.StringVar(&repo, "repo", "", "GitHub repository name")
	flag.Parse()

	images = flag.Args()

	if token == "" || owner == "" || repo == "" {
		panic("token/owner/repo is empty")
	}
}

func initClient() (context.Context, *github.Client) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return ctx, client
}

func uploadImages(ctx context.Context, client *github.Client, images []string) []string {
	urls := make([]string, len(images))
	copy(urls, images)
	for idx, image := range images {
		data, err := ioutil.ReadFile(image)
		if err != nil {
			fmt.Printf("Error reading file %s: %v", image, err)
			continue
		}

		suffix := path.Ext(image)
		timestamp := time.Now().Unix()
		_path := fmt.Sprintf("image/%d_%d%s", timestamp, idx, suffix)
		message := strconv.FormatInt(timestamp, 10)
		options := &github.RepositoryContentFileOptions{Message: &message, Content: data}
		_, _, err = client.Repositories.CreateFile(ctx, owner, repo, _path, options)
		if err != nil {
			fmt.Printf("Error committing file %s: %v", image, err)
			continue
		}
		urls[idx] = fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/master/%s", owner, repo, _path)
	}
	return urls
}

func main() {
	initVar()
	ctx, client := initClient()
	urls := uploadImages(ctx, client, images)

	fmt.Println("Upload Success:")
	for _, url := range urls {
		fmt.Println(url)
	}
}
