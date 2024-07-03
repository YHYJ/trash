/*
File: define_filemanager.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-26 11:32:05

Description: 文件管理
*/

package general

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/gookit/color"
)

// GetFsID 获取文件所在文件系统的 ID
//
// 参数：
//   - file: 文件路径
//
// 返回：
//   - 文件系统 ID。如果 ID = 0，表示获取文件系统 ID 失败
func GetFsID(file string) (uint64, error) {
	// 获取文件信息
	fileInfo, err := os.Stat(file)
	if err != nil {
		return uint64(0), err
	}

	// 获取文件所在文件系统的设备 ID
	id := fileInfo.Sys().(*syscall.Stat_t).Dev

	return uint64(id), nil
}

// ReadFileKey 读取文件包含关键字的行
//
// 参数：
//   - file: 文件路径
//   - key: 关键字
//
// 返回：
//   - 包含关键字的行的内容
func ReadFileKey(file, key string) string {
	// 打开文件
	text, err := os.Open(file)
	if err != nil {
		fileName, lineNo := GetCallerInfo()
		color.Printf("%s %s %s\n", DangerText(ErrorInfoFlag), SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
	}
	defer text.Close()

	// 创建一个扫描器对象按行遍历
	scanner := bufio.NewScanner(text)
	// 逐行读取，输出指定行
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), key) {
			return scanner.Text()
		}
	}
	return ""
}

// FileExist 判断文件是否存在
//
// 参数：
//   - filePath: 文件路径
//
// 返回：
//   - 文件存在返回 true，否则返回 false
func FileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

// GetAbsPath 获取指定文件的绝对路径
//
// 参数：
//   - filePath: 文件路径
//
// 返回：
//   - 文件的绝对路径
func GetAbsPath(filePath string) string {
	// 获取绝对路径
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return ""
	} else {
		return absPath
	}
}

// GetFilePureName 获取指定文件的纯粹文件名
//
// 参数：
//   - filePath: 文件路径
//
// 返回：
//   - 纯粹的文件名字
func GetFilePureName(filePath string) string {
	// 获取纯粹文件名
	pureName := filepath.Base(filePath)
	return pureName
}

// FileEmpty 判断文件是否为空
//
//   - 无法判断文件夹
//
// 参数：
//   - file: 文件路径
//
// 返回：
//   - 文件为空返回 true，否则返回 false
func FileEmpty(file string) bool {
	text, err := os.Open(file)
	if err != nil {
		return true
	}
	defer text.Close()

	fi, err := text.Stat()
	if err != nil {
		return true
	}
	return fi.Size() == 0
}

// CreateFile 创建文件，包括其父目录
//
// 参数：
//   - file: 文件路径
//
// 返回：
//   - 错误信息
func CreateFile(file string) error {
	if FileExist(file) {
		return nil
	}
	// 创建父目录
	parentPath := filepath.Dir(file)
	if err := os.MkdirAll(parentPath, os.ModePerm); err != nil {
		return err
	}
	// 创建文件
	if _, err := os.Create(file); err != nil {
		return err
	}

	return nil
}

// CreateDir 创建文件夹
//
// 参数：
//   - dir: 文件夹路径
//
// 返回：
//   - 错误信息
func CreateDir(dir string) error {
	if FileExist(dir) {
		return nil
	}
	return os.MkdirAll(dir, os.ModePerm)
}

// WriteFile 写入内容到文件，文件不存在则创建，不自动换行
//
// 参数：
//   - filePath: 文件路径
//   - content: 内容
//   - mode: 写入模式，追加('a', O_APPEND, 默认)或覆盖('t', O_TRUNC)
//
// 返回：
//   - 错误信息
func WriteFile(filePath, content, mode string) error {
	// 确定写入模式
	writeMode := os.O_WRONLY | os.O_CREATE | os.O_APPEND
	if mode == "t" {
		writeMode = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	}

	// 将内容写入文件
	file, err := os.OpenFile(filePath, writeMode, 0666)
	if err != nil {
		return err
	}
	if _, err = file.WriteString(content); err != nil {
		return err
	}
	return nil
}

// DeleteFile 删除文件，如果目标是文件夹则包括其下所有文件
//
// 参数：
//   - filePath: 文件路径
//
// 返回：
//   - 错误信息
func DeleteFile(filePath string) error {
	if !FileExist(filePath) {
		return nil
	}
	return os.RemoveAll(filePath)
}
