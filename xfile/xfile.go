/*
 * Copyright 2012-2023 Li Kexian
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * A toolkit for Golang development
 * https://www.likexian.com/
 */

package xfile

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	// TypeAll list dir and file
	TypeAll int = iota
	// TypeDir list only dir
	TypeDir
	// TypeFile list only file
	TypeFile
)

var (
	// ErrNotExists file is exists error
	ErrNotExists = errors.New("xfile: file is not exists")
	// ErrHasExists file is exists error
	ErrHasExists = errors.New("xfile: the file is exists")
)

// LsFile is list file info
type LsFile struct {
	Type int
	Path string
	Name string
}

// Version returns package version
func Version() string {
	return "0.15.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// Exists returns path is exists, symbolic link will check the target
func Exists(fpath string) bool {
	_, err := os.Stat(fpath)
	return !os.IsNotExist(err)
}

// Lexists returns path is exists, symbolic link will not follow
func Lexists(fpath string) bool {
	_, err := os.Lstat(fpath)
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

// Copy copys file and folder from src to dst
func Copy(src, dst string) error {
	if src == "" {
		src = "."
	}

	if dst == "" {
		dst = "."
	}

	if strings.TrimRight(src, "/") == strings.TrimRight(dst, "/") {
		return ErrHasExists
	}

	if Exists(dst) {
		return ErrHasExists
	}

	f, err := os.Lstat(src)
	if err != nil {
		return err
	}

	if f.Mode()&os.ModeSymlink != 0 {
		target, err := os.Readlink(src)
		if err != nil {
			return err
		}
		return os.Symlink(target, dst)
	} else if f.Mode().IsDir() {
		if !strings.HasSuffix(src, "/") {
			src += "/"
		}
		if !strings.HasSuffix(dst, "/") {
			dst += "/"
		}
		err = os.MkdirAll(dst, 0755)
		if err != nil {
			return err
		}
		ls, err := ListDir(src, TypeAll, -1)
		if err != nil {
			return err
		}
		for _, v := range ls {
			err = Copy(src+v.Name, dst+v.Name)
			if err != nil {
				return err
			}
		}
		if err = os.Chtimes(dst, f.ModTime(), f.ModTime()); err != nil {
			return err
		}
		return os.Chmod(dst, f.Mode())
	} else {
		fd, err := os.Open(src)
		if err != nil {
			return err
		}
		defer fd.Close()
		td, err := New(dst)
		if err != nil {
			return err
		}
		defer td.Close()
		if _, err = io.Copy(td, fd); err != nil {
			return err
		}
	}

	if err = os.Chtimes(dst, f.ModTime(), f.ModTime()); err != nil {
		return err
	}

	return os.Chmod(dst, f.Mode())
}

// New opens a file for new and return fd
func New(fpath string) (*os.File, error) {
	return NewFile(fpath, false)
}

// NewFile opens a file and return fd
func NewFile(fpath string, isAppend bool) (*os.File, error) {
	dir, _ := filepath.Split(fpath)
	if dir != "" && !IsDir(dir) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return nil, err
		}
	}

	if isAppend {
		return os.OpenFile(fpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	}

	return os.OpenFile(fpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
}

// WriteText writes string data to file
func WriteText(fpath, text string) (err error) {
	return Write(fpath, []byte(text))
}

// AppendText appends string data to file
func AppendText(fpath, text string) (err error) {
	return Append(fpath, []byte(text))
}

// Write writes bytes data to file
func Write(fpath string, data []byte) (err error) {
	return writeFile(fpath, data, false)
}

// Append appends bytes data to file
func Append(fpath string, data []byte) (err error) {
	return writeFile(fpath, data, true)
}

// writeFile writes bytes data to file
func writeFile(fpath string, data []byte, isAppend bool) (err error) {
	fd, err := NewFile(fpath, isAppend)
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

// Read returns bytes of file
func Read(fpath string) ([]byte, error) {
	return os.ReadFile(fpath)
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
		nRead++
		if n > 0 && nRead >= n {
			break
		}
	}

	err = scanner.Err()

	return
}

// ReadFirstLine returns first NOT empty line
func ReadFirstLine(fpath string) (line string, err error) {
	fd, err := os.Open(fpath)
	if err != nil {
		return
	}

	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if line != "" {
			return
		}
	}

	err = scanner.Err()

	return
}

// ReadLastLine returns last NOT empty line
func ReadLastLine(fpath string) (line string, err error) {
	fd, err := os.Open(fpath)
	if err != nil {
		return
	}

	defer fd.Close()

	stat, err := fd.Stat()
	if err != nil {
		return
	}

	size := stat.Size()
	if size == 0 {
		return
	}

	var cursor int64
	data := make([]byte, 0)

	for {
		cursor--
		_, err = fd.Seek(cursor, io.SeekEnd)
		if err != nil {
			return
		}

		buf := make([]byte, 1)
		_, err = fd.Read(buf)
		if err != nil {
			return
		}

		if buf[0] != '\r' && buf[0] != '\n' {
			data = append([]byte{buf[0]}, data...)
		} else {
			if cursor != -1 && strings.TrimSpace(string(data)) != "" {
				break
			}
			data = make([]byte, 0)
		}

		if cursor == -size {
			break
		}
	}

	return string(data), nil
}

// ListDir lists dir without recursion
func ListDir(fpath string, ftype, n int) (ls []LsFile, err error) {
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
			if ftype == TypeAll || ftype == TypeDir {
				ls = append(ls, LsFile{TypeDir, tpath, f.Name()})
				if n > 0 && len(ls) >= n {
					return
				}
			}
		} else {
			if ftype == TypeAll || ftype == TypeFile {
				ls = append(ls, LsFile{TypeFile, tpath, f.Name()})
				if n > 0 && len(ls) >= n {
					return
				}
			}
		}
	}

	return
}

// ListDirAll lists dir and children, filter by type, returns up to n
func ListDirAll(fpath string, ftype, n int) (ls []LsFile, err error) {
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
			if ftype == TypeAll || ftype == TypeDir {
				ls = append(ls, LsFile{TypeDir, tpath, f.Name()})
				if n > 0 && len(ls) >= n {
					return
				}
			}
			tls, err := ListDirAll(tpath, ftype, n-len(ls))
			if err != nil {
				return ls, err
			}
			ls = append(ls, tls...)
			if n > 0 && len(ls) >= n {
				return ls, nil
			}
		} else {
			if ftype == TypeAll || ftype == TypeFile {
				ls = append(ls, LsFile{TypeFile, tpath, f.Name()})
				if n > 0 && len(ls) >= n {
					return
				}
			}
		}
	}

	return
}

// Chmod chmods to path without recursion
func Chmod(fpath string, mode os.FileMode) error {
	return os.Chmod(fpath, mode)
}

// ChmodAll chmods to path and children, returns the first error it encounters
func ChmodAll(root string, mode os.FileMode) error {
	return filepath.Walk(root, func(fpath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return Chmod(fpath, mode)
	})
}

// Chown chowns to path without recursion
func Chown(fpath string, uid, gid int) error {
	return os.Chown(fpath, uid, gid)
}

// ChownAll chowns to path and children, returns the first error it encounters
func ChownAll(root string, uid, gid int) error {
	return filepath.Walk(root, func(fpath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return Chown(fpath, uid, gid)
	})
}
