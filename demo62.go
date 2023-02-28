package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jessevdk/go-flags"
)

type MathCommand struct {
	Op     string `long:"op" description:"operation to execute"`
	Args   []string
	Result int64
}

func (this *MathCommand) Execute(args []string) error {
	if this.Op != "+" && this.Op != "-" && this.Op != "x" && this.Op != "/" {
		return errors.New("invalid op")
	}

	for _, arg := range args {
		num, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			return err
		}

		this.Result += num
	}

	this.Args = args
	return nil
}

type Option struct {
	Math MathCommand `command:"math"`
}

func main() {
	var opt Option
	_, err := flags.Parse(&opt)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The result of %s is %d", strings.Join(opt.Math.Args, opt.Math.Op), opt.Math.Result)
}
