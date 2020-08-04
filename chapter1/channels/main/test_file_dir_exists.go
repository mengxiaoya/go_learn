package main

import (
	"log"
	"os"
)

func main() {

	var file = "D:\\tmp\\1.txt"
	dirIsExists(file)

}

func fileIsExist(fileAddr string) bool {
	// 读取文件信息，判断文件是否存在
	_, err := os.Stat(fileAddr)
	if err != nil {
		if os.IsExist(err) { // 根据错误类型进行判断
			return true
		}
		return false
	}
	return true
}

func dirIsExists(dirAddr string) {
	_, err := os.Stat(dirAddr)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(dirAddr, os.ModePerm)
			if err != nil {
				log.Printf("mkdir[%s] failed![%v]", dirAddr, err)
			} else {
				log.Printf("create task root path[%s] success!", dirAddr)
			}
		} else {
			log.Printf("task root dir[%s] already exists!", dirAddr)
		}
	} else {
		log.Printf("task root dir[%s] already exists!", dirAddr)
	}

}
