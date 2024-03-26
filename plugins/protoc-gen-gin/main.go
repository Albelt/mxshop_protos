package main

import (
	"flag"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
)

const version = "v0.1"

var (
	showVersion = flag.Bool("version", false, "protoc-gen-gin version")
)

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-gin version:%s\n", version)
		return
	}

	var flags flag.FlagSet
	gg := GenGin{}

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if err := gg.GenerateFile(gen, f); err != nil {
				return err
			}
		}
		return nil
	})
}
