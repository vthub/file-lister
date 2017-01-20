package main

import (
	"fmt"
	"flag"
	"os"
	"io/ioutil"
)

func main() {
	var path = flag.String("dir", ".", "directory to be scanned")
	flag.Parse()
	fmt.Println("Directory to be scanned", *path)
	if pathStat, err := os.Stat(*path); err != nil || !pathStat.IsDir() {
		fmt.Println("Sorry, cannot list", *path)
	}
	infos, err := ioutil.ReadDir(*path);
	if err != nil {
		fmt.Println("Error during path listing");
	}
	for _, f := range infos {
		fmt.Println(f.Name())
	}
}
