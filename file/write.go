package file

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// create file and append write
func CreateAndAppend(obj interface{}) {
	b, _ := json.MarshalIndent(obj, "", "")
	//创建一个新文件
	filePath := "write.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(string(b))
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}
