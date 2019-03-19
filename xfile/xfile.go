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
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Version returns package version
func Version() string {
	return "0.5.0"
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
func Exists(fpath string) bool {
	_, err := os.Stat(fpath)
	return !os.IsNotExist(err)
}

// IsFile returns path is a file, symbolic link will check the target
func IsFile(fpath string) bool {
	f, err := os.Stat(fpath)
	return err == nil && f.Mode().IsRegular()
}

// IsDir returns path is a dir, symbolic link will check the target
func IsDir(fpath string) bool {
	f, err := os.Stat(fpath)
	return err == nil && f.Mode().IsDir()
}

// IsSymlink returns path is a symbolic link
func IsSymlink(fpath string) bool {
	f, err := os.Lstat(fpath)
	return err == nil && f.Mode()&os.ModeSymlink != 0
}

// Size returns the file size of path, symbolic link will check the target
func Size(fpath string) (int64, error) {
	f, err := os.Stat(fpath)
	if err != nil {
		return -1, err
	}

	return f.Size(), nil
}

// MTime returns the file mtime of path, symbolic link will check the target
func MTime(fpath string) (int64, error) {
	f, err := os.Stat(fpath)
	if err != nil {
		return -1, err
	}

	return f.ModTime().Unix(), nil
}

// New open a file and return fd
func New(fpath string) (*os.File, error) {
	dir, _ := filepath.Split(fpath)
	if dir != "" && !IsDir(dir) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return nil, err
		}
	}

	return os.OpenFile(fpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
}

// Write write bytes data to file
func Write(fpath string, data []byte) (err error) {
	fd, err := New(fpath)
	if err != nil {
		return
	}

	n, err := fd.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}

	if e := fd.Close(); err == nil {
		err = e
	}

	return
}

// WriteText write text data to file
func WriteText(fpath, text string) (err error) {
	return Write(fpath, []byte(text))
}

// Read returns bytes of file
func Read(fpath string) ([]byte, error) {
	return ioutil.ReadFile(fpath)
}

// ReadText returns text of file
func ReadText(fpath string) (string, error) {
	b, err := Read(fpath)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// ReadLines returns N lines of file
func ReadLines(fpath string, n int) (lines []string, err error) {
	fd, err := os.Open(fpath)
	if err != nil {
		return
	}

	defer fd.Close()
	nRead := 0
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		nRead += 1
		if n > 0 && nRead >= n {
			break
		}
	}

	err = scanner.Err()

	return
}

// ListDir list dir and children, filter by type, returns up to n
func ListDir(fpath, ftype string, n int) (ls [][]string, err error) {
	if fpath == "" {
		fpath = "."
	}

	if !strings.HasSuffix(fpath, "/") {
		fpath += "/"
	}

	fd, err := os.Open(fpath)
	if err != nil {
		return
	}

	defer fd.Close()
	fs, err := fd.Readdir(-1)
	if err != nil {
		return
	}

	for _, f := range fs {
		tpath := fpath + f.Name()
		if f.IsDir() {
			if ftype == "" || ftype == "dir" {
				ls = append(ls, []string{"dir", tpath})
				if n > 0 && len(ls) >= n {
					return
				}
			}
			tls, err := ListDir(tpath, ftype, n-len(ls))
			if err != nil {
				return ls, err
			}
			ls = append(ls, tls...)
			if n > 0 && len(ls) >= n {
				return ls, nil
			}
		} else {
			if ftype == "" || ftype == "file" {
				ls = append(ls, []string{"file", tpath})
				if n > 0 && len(ls) >= n {
					return
				}
			}
		}
	}

	return
}

// ChmodAll chmod to path and children, returns the first error it encounters
func ChmodAll(root string, mode os.FileMode) error {
	return filepath.Walk(root, func(fpath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return os.Chmod(fpath, mode)
	})
}

// ChownAll chown to path and children, returns the first error it encounters
func ChownAll(root string, uid, gid int) error {
	return filepath.Walk(root, func(fpath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return os.Chown(fpath, uid, gid)
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
