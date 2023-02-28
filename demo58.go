package main

import (
	"flag"
	"fmt"
)

var (
	intv    *int
	boolv   *bool
	stringv *string
)

func init() {
	intv = flag.Int("intv", 0, "int value")
	boolv = flag.Bool("boolv", false, "bool value")
	stringv = flag.String("stringv", "default ", "string value")
}
func main() {
	flag.Parse()

	fmt.Println("int flag:", *intv)
	fmt.Println("bool flag:", *boolv)
	fmt.Println("string flag:", *stringv)

}
