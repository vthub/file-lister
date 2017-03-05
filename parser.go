package main

import (
	"os"
	"flag"
	"fmt"
	"errors"
)

type Configuration struct {
	Path string
}

func Validate(conf Configuration) error {
	if pathStat, err := os.Stat(conf.Path); err != nil {
		return errors.New("Cannot list " + conf.Path + ". Cannot open directory")
	} else if !pathStat.IsDir() {
		return errors.New("Cannot list " + conf.Path + ". Provided path is not a directory")
	}
	return nil
}

func Parse() (*Configuration, error) {
	var path = flag.String("dir", ".", "directory to be scanned")
	flag.Parse()
	fmt.Println("Directory to be scanned", *path)

	conf := &Configuration{*path};
	return conf, Validate(*conf)
}
