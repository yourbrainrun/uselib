package main

import (
	"fmt"
	"git.cloudhub.top/server/media-server-agent-go-client/ms_requests"
	"github.com/yourbrainrun/testlib/my_rand"
)

func main() {
	aa := my_rand.GetRand()
	fmt.Println(aa)

	var test ms_requests.BadRequest
	fmt.Println(test)
}
