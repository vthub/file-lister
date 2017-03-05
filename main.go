package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Collector struct {
	formatString string
	formatter    Formatter
}

func (c Collector) collect(path string, f os.FileInfo, err error) error {
	fmt.Println(c.formatter.format(c.formatString, FileInfoProvider{info: f}))
	return nil
}

func main() {
	var conf, err = Parse()
	if (err != nil) {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Sprintln("Listing...")
	var c = Collector{formatString: "%p %s", formatter: DefaultFormatter{}}
	filepath.Walk(conf.Path, c.collect)
}
