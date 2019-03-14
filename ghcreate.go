package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"

	"bitbucket.org/mischief/libauth"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	description = flag.String("d", "", "Short description of repository")
	homepage = flag.String("p", "", "Home page for repository")
	orgname = flag.String("o", "", "Name of github organization")
)

func main() {
	flag.Parse()
	if flag.Lookup("h") != nil {
		flag.Usage()
		os.Exit(1)
	}
	usr, err := user.Current()
	token, err := libauth.Getuserpasswd("proto=pass service=github role=client user=%s", usr.Username)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token.Password,
	})
	Reponame := flag.Arg(0)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	repo := &github.Repository{
		Name: &Reponame,
		Description: description,
		Homepage: homepage,
	}
	newrepo, _, err := client.Repositories.Create(ctx, *orgname, repo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newrepo.HTMLURL)
}

