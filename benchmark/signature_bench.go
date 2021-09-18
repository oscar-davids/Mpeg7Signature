package main

import (	
	"os"
	"fmt"	
    "io/ioutil"
    "log"
    "path/filepath"
	"github.com/livepeer/lpms/ffmpeg"
)

func main() {

	if len(os.Args) <= 2 {
		panic("Usage: <first sign-bin path> <second sign-bin path>")
	}

	path1 := os.Args[1]
	path2 := os.Args[2]

	getfilelist := func(path string) []string {
		flist := make([]string, 0)
		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal(err)
			return flist
		}	
		for _, f := range files {
			if filepath.Ext(f.Name()) == ".bin" {
				fullpath := path + "/" + f.Name()
				flist = append(flist,fullpath)
			}
		}
		return flist
	}

    flist1 := getfilelist(path1)
	flist2 := getfilelist(path2)
	totalcount := len(flist1)
	counttrue := 0
	if len(flist1) == len(flist2) {
		for i, _ := range flist1 {
			fmt.Printf("Compare: %v vs %v ", flist1[i], flist2[i])
			res, err := ffmpeg.CompareSignatureByPath(flist1[i], flist2[i])
			if err == nil && res == true {
				counttrue++
			}
			fmt.Printf("-> result: %v\n", res)
		}
	}
	fmt.Printf("Sign bench: Correct: %v / Total:%v \n", counttrue, totalcount)
}