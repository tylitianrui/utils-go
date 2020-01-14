package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// current path
func CurrentPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}

// 指定的路径是否存在
// false  不存在;true 存在
func Exist(path string) (bool, error) {
	var err error

	if _, err = os.Stat(path); err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// 获取目录下的文件，不包括目录
func SubFileName(path string) ([]string, error) {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	list := make([]string, len(files))
	// 获取文件，并输出它们的名字
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		list = append(list, file.Name())
	}
	return list, nil
}
