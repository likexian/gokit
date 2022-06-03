/*
 * Copyright 2012-2022 Li Kexian
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

package xtar

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/likexian/gokit/xfile"
)

// Version returns package version
func Version() string {
	return "0.2.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// Create compress a list of files
func Create(tarFile string, files ...string) (err error) {
	if xfile.Exists(tarFile) {
		err = fmt.Errorf("xtar: file name %s is exists", tarFile)
		return
	}

	if len(files) == 0 {
		err = fmt.Errorf("xtar: no input file specify")
		return
	}

	fd, err := xfile.New(tarFile)
	if err != nil {
		return
	}
	defer fd.Close()

	var tw *tar.Writer
	if IsGzName(tarFile) {
		gw := gzip.NewWriter(fd)
		defer gw.Close()

		tw = tar.NewWriter(gw)
		defer tw.Close()
	} else {
		tw = tar.NewWriter(fd)
		defer tw.Close()
	}

	for _, f := range files {
		err = addFile(tw, f, "")
		if err != nil {
			return
		}
	}

	return
}

// addFile do compress a file
func addFile(tw *tar.Writer, file string, prefix string) error {
	if prefix != "" && !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}

	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	f, err := os.Lstat(file)
	if err != nil {
		return err
	}

	fl := ""
	if f.Mode()&os.ModeSymlink != 0 {
		fl, err = os.Readlink(file)
		if err != nil {
			return err
		}
	}

	h, err := tar.FileInfoHeader(f, fl)
	if err != nil {
		return err
	}

	h.Name = prefix + h.Name
	err = tw.WriteHeader(h)
	if err != nil {
		return err
	}

	switch mode := f.Mode(); {
	case mode.IsRegular():
		_, err = io.Copy(tw, fd)
		if err != nil {
			return err
		}
	case mode&os.ModeSymlink != 0:
	case mode.IsDir():
		prefix += f.Name()
		fs, err := fd.Readdir(0)
		if err != nil {
			return err
		}
		for _, ff := range fs {
			err = addFile(tw, fd.Name()+"/"+ff.Name(), prefix)
			if err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("xtar: unsupport file mode: %v", mode)
	}

	return nil
}

// Extract decompress a tar file to folder
func Extract(tarFile, dstFolder string) (err error) {
	if dstFolder != "" && !strings.HasSuffix(dstFolder, "/") {
		dstFolder += "/"
	}

	fd, err := os.Open(tarFile)
	if err != nil {
		return
	}
	defer fd.Close()

	var tr *tar.Reader
	if IsGzName(tarFile) {
		gr, err := gzip.NewReader(fd)
		if err != nil {
			return err
		}
		defer gr.Close()

		tr = tar.NewReader(gr)
	} else {
		tr = tar.NewReader(fd)
	}

	for {
		h, err := tr.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		fname := strings.TrimPrefix(h.Name, "/")
		if fname == "" {
			continue
		}
		dstFile := dstFolder + fname
		switch h.Typeflag {
		case tar.TypeReg:
			var ffd *os.File
			ffd, err = xfile.New(dstFile)
			if err != nil {
				return err
			}
			defer ffd.Close()
			_, err = io.Copy(ffd, tr)
		case tar.TypeLink:
			err = os.Link(h.Linkname, dstFile)
		case tar.TypeSymlink:
			err = os.Symlink(h.Linkname, dstFile)
		case tar.TypeDir:
			if !xfile.Exists(dstFile) {
				err = os.MkdirAll(dstFile, 0755)
			}
		default:
			err = fmt.Errorf("xtar: unsupport file type: %v", h.Typeflag)
		}
		if err != nil {
			return err
		}
		_ = os.Chtimes(dstFile, h.AccessTime, h.ModTime)
		_ = os.Chmod(dstFile, os.FileMode(h.Mode))
		_ = os.Chown(dstFile, h.Uid, h.Gid)
	}

	return
}

// IsGzName returns is a tar.gz file name
func IsGzName(name string) bool {
	name = strings.Trim(name, ".")
	if strings.HasSuffix(name, ".tgz") {
		return true
	}

	if strings.HasSuffix(name, ".tar.gz") {
		return true
	}

	return false
}
