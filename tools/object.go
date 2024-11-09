/*
*@author  chengkenli
*@project pjstack
*@package tools
*@file    object
*@date    2024/11/9 17:19
 */
package tools

import (
	"bufio"
	"fmt"
	"os"
	"pjstack/util"
)

type SrAvgs struct {
	Host string
	Port int
	User string
	Pass string
}

// WriteFile 文件落地
func WriteFile(fname, msg string) {
	fileHandle, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fileHandle.Close()
	// NewWriter 默认缓冲区大小是 4096
	// 需要使用自定义缓冲区的writer 使用 NewWriterSize()方法
	buf := bufio.NewWriterSize(fileHandle, len(msg))

	buf.WriteString(msg)

	err = buf.Flush()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

// InitCommon 初始化变量是否填写
func InitCommon() bool {
	pstack := util.Config.GetString("common.Pstack")
	if len(pstack) == 0 {
		return false
	}
	jstack := util.Config.GetString("common.Jstack")
	if len(jstack) == 0 {
		return false
	}
	user := util.Config.GetString("common.User")
	if len(user) == 0 {
		return false
	}
	private := util.Config.GetString("common.Private")
	if len(private) == 0 {
		return false
	}
	logpath := util.Config.GetString("common.LogPath")
	if len(logpath) == 0 {
		return false
	}
	return true
}

// StrInSlice 检查切片中是否存在某个元素
func StrInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
