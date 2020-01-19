package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
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

func CopyFile(src, dst, rename string) error {

	var (
		b   bool
		err error
	)

	if b, err = Exist(src); err != nil || !b {
		return err
	}
	rename = strings.ReplaceAll(rename, " ", "")
	if rename != "" {
		dst = path.Join(dst, rename)

	}

	if b, err = Exist(dst); b {

		return errors.New(fmt.Sprintf("%s已经存在", dst))
	}

	return copy(src, dst)

}

func copy(src, dst string) error {
	var (
		rf, wf *os.File
		err    error
	)

	if rf, err = os.Open(src); err != nil {
		return err
	}
	defer rf.Close()
	reader := bufio.NewReader(rf)

	wf, err = os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer wf.Close()
	writer := bufio.NewWriter(wf)

	_, err = io.Copy(writer, reader)
	return err
}
