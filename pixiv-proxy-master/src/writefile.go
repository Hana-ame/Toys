package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeFile(path string, content string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(content)
	write.Flush()
}
