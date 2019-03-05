/*
 * A toolkit for Golang development
 * https://www.likexian.com/
 *
 * Copyright 2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package xfile

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

// Version returns package version
func Version() string {
	return "0.3.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Apache License, Version 2.0"
}

// Exists returns path is exists, symbolic link will check the target
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsFile returns path is a file, symbolic link will check the target
func IsFile(path string) bool {
	f, err := os.Stat(path)
	return err == nil && f.Mode().IsRegular()
}

// IsDir returns path is a dir, symbolic link will check the target
func IsDir(path string) bool {
	f, err := os.Stat(path)
	return err == nil && f.Mode().IsDir()
}

// IsSymlink returns path is a symbolic link
func IsSymlink(path string) bool {
	f, err := os.Lstat(path)
	return err == nil && f.Mode()&os.ModeSymlink != 0
}

// Size returns the file size of path, symbolic link will check the target
func Size(path string) (int64, error) {
	f, err := os.Stat(path)
	if err != nil {
		return -1, err
	}

	return f.Size(), nil
}

// MTime returns the file mtime of path, symbolic link will check the target
func MTime(path string) (int64, error) {
	f, err := os.Stat(path)
	if err != nil {
		return -1, err
	}

	return f.ModTime().Unix(), nil
}

// ReadLines returns N lines of file
func ReadLines(path string, n int) (lines []string, err error) {
	fd, err := os.Open(path)
	if err != nil {
		return
	}

	defer fd.Close()
	line_read := 0
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		line_read += 1
		if n > 0 && line_read >= n {
			break
		}
	}

	err = scanner.Err()

	return
}

// ReadText returns text of file
func ReadText(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// WriteText write text to file
func WriteText(path, text string) error {
	d := filepath.Dir(path)
	if d != "" && !Exists(d) {
		err := os.MkdirAll(d, 0755)
		if err != nil {
			return err
		}
	}

	return ioutil.WriteFile(path, []byte(text), 0644)
}

// ChmodAll chmod to path and children, returns the first error it encounters
func ChmodAll(root string, mode os.FileMode) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return os.Chmod(path, mode)
	})
}

// ChownAll chown to path and children, returns the first error it encounters
func ChownAll(root string, uid, gid int) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return os.Chown(path, uid, gid)
	})
}

// GetPwd returns the abs dir of current path
func GetPwd() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

// GetProcPwd returns the abs dir of current execution file
func GetProcPwd() string {
	file, _ := exec.LookPath(os.Args[0])
	dir, _ := filepath.Abs(filepath.Dir(file))
	return dir
}
