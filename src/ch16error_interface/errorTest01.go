package main

import (
	"fmt"
	"os"
)

type error interface {
	Error() string
}

func Open(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_RDONLY, 0)
}

func main() {
	file, err := Open("111.txt")
	if err != nil {
		fmt.Errorf("111打开文件失败，err : %w\n",err)
		fmt.Printf("打开文件失败,err: %#v, file: %#v", err, file)
		return
	}

}
