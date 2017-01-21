package main

import (
	"fmt"
	"io/ioutil"
	"github.com/vthub/file-lister/manager"
	"os"
)

func main() {
	var conf, err = manager.Parse()
	if(err != nil) {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	infos, err := ioutil.ReadDir(conf.Path);
	if err != nil {
		fmt.Println("Error during path listing");
	}
	for _, f := range infos {
		fmt.Println(f.Name())
	}
}
