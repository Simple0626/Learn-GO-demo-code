package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
)

type Option struct {
	Basic GroupBasicOption `description:"basic type" group:"basic"`
	Slice GroupSliceOption `description:"slice of basic type" group:"slice"`
	//	分了两个组 basic和slice
}

// basic组的结构体
type GroupBasicOption struct {
	IntFlag    int     `short:"i" long:"intflag" description:"int flag"`
	BoolFlag   bool    `short:"b" long:"boolflag" description:"bool flag"`
	FloatFlag  float64 `short:"f" long:"floatflag" description:"float flag"`
	StringFlag string  `short:"s" long:"stringflag" description:"string flag"`
}

// slice的结构体
type GroupSliceOption struct {
	IntSlice    int     `long:"intslice" description:"int slice"`
	BoolSlice   bool    `long:"boolslice" description:"bool slice"`
	FloatSlice  float64 `long:"floatslice" description:"float slice"`
	StringSlice string  `long:"stringslice" description:"string slice"`
}

func main() {
	var opt Option
	p := flags.NewParser(&opt, flags.Default)
	_, err := p.Parse()
	if err != nil {
		log.Fatal("Parse error:", err)
	}

	//拿到basic组
	basicGroup := p.Command.Group.Find("basic")
	for _, option := range basicGroup.Options() {
		fmt.Printf("name:%s value:%v\n", option.LongNameWithNamespace(), option.Value())
	}
	//拿到slice组
	sliceGroup := p.Command.Group.Find("slice")
	for _, option := range sliceGroup.Options() {
		fmt.Printf("name:%s value:%v\n", option.LongNameWithNamespace(), option.Value())
	}

}
