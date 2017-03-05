package main

import (
	"os"
	"strings"
	"strconv"
)

type InfoProvider interface {
	Path() string
	Size() int64
}

type FileInfoProvider struct {
	info os.FileInfo
}

func (p FileInfoProvider) Path() string {
	return p.info.Name()
}

func (p FileInfoProvider) Size() int64 {
	return p.info.Size()
}

type FormatFunc func(info InfoProvider) string

type Flag struct {
	flag        string
	function    FormatFunc
	description string
}

var defaultFlags = []Flag{
	{flag: "p", function: func(info InfoProvider) string {
		return info.Path()
	}, description: "Full file path", },
	{flag: "s", function: func(info InfoProvider) string {
		return strconv.FormatInt(info.Size(), 10)
	}, description: "Size in bytes", },
}

type Formatter interface {
	format(formatString string, info InfoProvider) string
}

type DefaultFormatter struct {
}

func (f DefaultFormatter) format(formatString string, info InfoProvider) string {
	var tmp = formatString
	for _, flag := range defaultFlags {
		tmp = strings.Replace(tmp, "%" + flag.flag, flag.function(info), -1)
	}
	return tmp
}
