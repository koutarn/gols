package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/jessevdk/go-flags"
)

const (
	appName        = "gols"
	appVersion     = "0.01"
	appUsage       = "[options][file...]"
	appDiscription = "unix ls command for golang"
)

type exitCode int

const (
	exitCodeOK exitCode = iota
	exitCodeErrArgs
	exitCodeErrLs
)

type options struct {
	Version bool   `short:"v" long:"version" description:"Show version"`
	Recerse bool   `short:"r" long:"recurse" description:"recurse"`
	Path    string `short:"p" long:"path" default:"./" description:"path"`
}

func main() {
	code, err := run(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %s\n", err)
	}

	os.Exit(int(code))
}

func run(cliantArgs []string) (exitCode, error) {
	var opts options
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = appName
	parser.Usage = appUsage
	parser.ShortDescription = appDiscription
	parser.LongDescription = appDiscription

	_, err := parser.ParseArgs(cliantArgs)

	if err != nil {
		if flags.WroteHelp(err) {
			return exitCodeOK, nil
		}
		return exitCodeErrArgs, fmt.Errorf("parse error:%w", err)
	}

	if opts.Version {
		fmt.Fprintf(os.Stdout, "%s: v%s\n", appName, appVersion)
		return exitCodeOK, nil
	}

	var dir string = opts.Path
	fmt.Println(dir)
	if dir == "./" {
		dir, err = os.Getwd()
		if err != nil {
			return exitCodeErrLs, errors.New("error ls getwd")
		}
	}

	// ls実行
	// 本当は文字列を吐き出すだけにして後のフィルタリング処理を実行したほうが良いかもしれん
	if opts.Recerse {
		code, err := recurseLs(dir)
		if err != nil {
			return code, err
		}
	} else {
		code, err := ls(dir)
		if err != nil {
			return code, err
		}
	}

	//フィルタリング処理
	//TODO:気が向いたら書く

	return exitCodeOK, nil
}

//再帰表示
func recurseLs(dir string) (exitCode, error) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		printFileInfo(info)

		return nil
	})

	if err != nil {
		return exitCodeErrLs, err
	}

	return exitCodeOK, nil
}

func ls(dir string) (exitCode, error) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return exitCodeErrLs, errors.New("error read dir")
	}

	for _, info := range fileInfos {
		printFileInfo(info)
	}

	return exitCodeOK, nil
}

func printFileInfo(info os.FileInfo) {
	fmt.Printf("%v %s\n",
		color.New(color.FgHiYellow, color.Bold).Sprintf("%v", info.Mode()),
		color.New(color.FgHiWhite, color.Bold).Sprintf("%v", info.Name()))
}
