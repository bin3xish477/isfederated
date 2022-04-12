package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alexflint/go-arg"
)

type RealmInfo struct {
	Type    string `json:"NameSpaceType"`
	AuthURL string `json:"AuthURL,omitempty"`
}

var args struct {
	Email string `arg:"-e,--email" help:"email to check"`
}

var url = "https://login.microsoftonline.com/getuserrealm.srf?login="

func main() {
	arg.MustParse(&args)

	var r RealmInfo

	resp, err := http.Get(fmt.Sprintf("%s%s", url, args.Email))
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	bdy, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bdy, &r)

	if r.Type == "Managed" {
		fmt.Printf("%s is not federated and is managed by Microsoft\n", args.Email)
	} else if r.Type == "Federated" {
		fmt.Printf("%s is federated and is managed by %s\n", args.Email, r.AuthURL)
	}
}
