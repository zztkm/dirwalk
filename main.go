package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// DirWalk root以下のファイルのパスをスライスで返却する
func DirWalk(root string) []string {

	var paths []string
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				// 特定のディレクトリ以下を無視する場合は
				// return filepath.SkipDir
				return nil
			}

			// windowsの場合スラッシュをバックスラッシュで返却する
			if runtime.GOOS == "windows" {
				path = filepath.FromSlash(path)
			}
			paths = append(paths, path)
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

	return paths
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("パスくれ！")
		os.Exit(1)
	}

	root := flag.Arg(0)

	DirWalk(root)
}
