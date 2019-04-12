package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/matsune/go-json"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	if terminal.IsTerminal(0) {
		return
	}

	// read from pipe
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	if v, err := json.Parse(string(b)); err != nil {
		panic(err)
	} else {
		if vv, ok := v.(json.Value); ok {
			fmt.Println(vv)
		}
	}

}
