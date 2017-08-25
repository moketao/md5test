package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var which = flag.Bool("which", true, "")
var path = flag.String("path", "", "")
var cnt = flag.Int("cnt", 100, "")

func aaa() {
	f, err := os.Open(*path)
	if err != nil {
		fmt.Println("Open", err)
		return
	}

	defer f.Close()

	body, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("ReadAll", err)
		return
	}

	md5.Sum(body)
	fmt.Printf("%x\n", md5.Sum(body))
}

func bbb() {
	f, err := os.Open(*path)
	if err != nil {
		fmt.Println("Open", err)
		return
	}

	defer f.Close()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		fmt.Println("Copy", err)
		return
	}

	md5hash.Sum(nil)
	fmt.Printf("%x\n", md5hash.Sum(nil))
}
func ccc(theName string) {
	f, err := os.Open(theName)
	if err != nil {
		fmt.Println("Open", err)
		return
	}

	defer f.Close()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		fmt.Println("Copy", err)
		return
	}

	//md5hash.Sum(nil)
	fmt.Printf("%x\n", md5hash.Sum(nil))
}
func main() {
	flag.Parse()
	theName := os.Args[1]
	ccc(theName)
	// for i := 0; i < *cnt; i++ {
	// 	if *which {
	// 		aaa()
	// 	} else {
	// 		bbb()
	// 	}
	// }
}
