/*
 * A toolkit for Golang development
 * https://www.likexian.com/
 *
 * Copyright 2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package file

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

// FileExists returns path is exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsFile returns path is a file
func IsFile(path string) bool {
	f, err := os.Stat(path)
	return err == nil && f.Mode().IsRegular()
}

// IsDir returns path is a dir
func IsDir(path string) bool {
	f, err := os.Stat(path)
	return err == nil && f.Mode().IsDir()
}

// IsSymlink returns path is a symbolic link
func IsSymlink(path string) bool {
	f, err := os.Lstat(path)
	return err == nil && f.Mode()&os.ModeSymlink != 0
}

// FileSize returns the file size of path
func FileSize(path string) (int64, error) {
	f, err := os.Stat(path)
	if err != nil {
		return -1, err
	}

	return f.Size(), nil
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
	if d != "" && !FileExists(d) {
		err := os.MkdirAll(d, 0755)
		if err != nil {
			return err
		}
	}

	return ioutil.WriteFile(path, []byte(text), 0644)
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
