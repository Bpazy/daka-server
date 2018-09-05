// Code generated by go-bindata.
// sources:
// res/init.sql
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _resInitSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x92\xc1\x4a\xc3\x40\x10\x86\xef\x81\xbc\xc3\x1c\x1b\xe8\xa1\x82\x07\x41\x7a\xd8\x24\xd3\xba\x98\x4c\xc2\xee\x2c\xd2\x53\xb2\x9a\x14\x8b\x34\x91\x24\xd5\xd7\x97\xb6\xc6\x82\xd4\x86\x88\x17\xcf\xff\xf7\xef\xce\x7c\x4c\xa0\x50\x30\x02\x0b\x3f\x42\xc8\x0b\xfb\x62\xb3\x4d\xb5\xae\x73\x98\xb8\x0e\x40\x2e\x69\x91\x64\x32\xcc\xe1\xcd\x36\x4f\xcf\xb6\x99\x5c\xcd\x66\x1e\x50\xc2\x40\x26\x8a\xa6\x07\x86\x44\x8c\x17\x81\x50\x6a\x16\x14\x0c\x40\x82\x31\x87\xc2\x76\xe5\xb7\xe0\x38\x61\xc6\x32\xfe\xcc\xbb\xcd\xb6\x84\x10\x17\xc2\x44\x0c\x81\x51\x0a\x89\x0f\xb1\x66\x11\xa7\xc7\x92\x09\xd3\x31\x25\x48\x08\x4c\xba\x1f\xe1\x87\x07\x53\x25\x63\xa1\x56\x70\x8f\x2b\x98\x7c\x59\xf1\x5c\xc7\x03\xa4\xa5\x24\x9c\xcb\xaa\xaa\x43\xff\xf4\xc5\x9d\x50\x1a\x79\xbe\xeb\xd6\x37\xdb\xc7\xeb\x5b\xd7\x71\x9d\x33\xae\x77\x6d\xd9\xf4\xae\x8d\x46\x35\xe4\x7a\xcf\x0c\xfa\x3e\x6f\xac\x67\x7e\xa7\x6e\xb8\x3d\xd2\x61\xbf\xed\xf4\xb4\x93\x07\x46\x4b\x5a\x82\xcf\x0a\xf1\x4f\xd4\x66\xaf\xb6\x6d\xdf\xeb\xa6\x18\xe3\x38\x15\x5a\x3f\x24\xea\x32\xf4\xaf\x1c\x8f\xba\xd3\x8f\x00\x00\x00\xff\xff\xf9\x84\x21\x6f\x12\x04\x00\x00")

func resInitSqlBytes() ([]byte, error) {
	return bindataRead(
		_resInitSql,
		"res/init.sql",
	)
}

func resInitSql() (*asset, error) {
	bytes, err := resInitSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/init.sql", size: 1042, mode: os.FileMode(438), modTime: time.Unix(1536053525, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"res/init.sql": resInitSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"res": &bintree{nil, map[string]*bintree{
		"init.sql": &bintree{resInitSql, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}