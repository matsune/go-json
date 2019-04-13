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

	v, err := json.Parse(string(b))
	if err != nil {
		panic(err)
	}
	walk(v.(json.Value), 0)
}

func walk(v json.Value, nest int) {
	switch vv := v.(type) {
	case *json.ObjectValue:
		fmt.Println("{")
		nest++
		for i, kv := range vv.KeyValues {
			indent(nest)
			fmt.Printf("%q: ", kv.Key)
			walk(kv.Value, nest)
			if i < len(vv.KeyValues)-1 {
				fmt.Println(",")
			} else {
				fmt.Println()
			}
		}
		nest--
		indent(nest)
		fmt.Print("}")
	case *json.ArrayValue:
		fmt.Printf("[")
		nest++
		for i, vvv := range vv.Values {
			walk(vvv, nest)
			if i < len(vv.Values)-1 {
				fmt.Print(",")
			}
		}
		nest--
		fmt.Print("]")
	default:
		fmt.Print(vv)
	}
}

func indent(nest int) {
	for i := 0; i < nest; i++ {
		fmt.Print("  ")
	}
}
