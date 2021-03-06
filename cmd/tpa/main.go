package main

import (
	"fmt"
	"os"
	"togglepx/lib"
	"togglepx/lib/cli"
)

func main() {
	p := lib.PathIPConfig{}
	//ファイルを読み込んで構造体に適用
	err := p.ReadJsonTransfer(lib.JsonPath)
	if err != nil {
		fmt.Print("ファイルの読み込みに失敗")
		os.Exit(1)
	}
	stream := cli.Stream{os.Stdout, os.Stderr}
	//cli起動、cliに構造体渡す
	os.Exit(stream.Run(os.Args, &p))
}
