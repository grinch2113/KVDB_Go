package KVDBHandler

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const path = "./Database/"

var dbName = "MyDB.kvdb"

func init() {
	os.MkdirAll(path, os.ModePerm)
	if _, err := os.Stat(path + dbName); os.IsNotExist(err) {
		os.Create(path + dbName)
	}
}

// Set 写入一条记录/**
func Set(key, value string) {
	file, err := os.OpenFile(path+dbName, os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开错误：" + err.Error())
		return
	}
	_, err = file.Write([]byte(key + " " + value + "\n"))
	if err != nil {
		fmt.Println("文件写入错误：" + err.Error())
		return
	}
}

// Get 读取一条记录/**
func Get(key string) string {
	file, err := os.Open(path + dbName)
	if err != nil {
		fmt.Println("打开文件失败：" + err.Error())
		return ""
	}
	defer file.Close()
	value := ""
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("读取文件错误：" + err.Error())
			return ""
		}
		kvGroup := strings.Split(line, " ")
		if kvGroup[0] == key {
			value = kvGroup[1]
		}
	}
	return value
}

func Remove() {

}

func purge() {

}
