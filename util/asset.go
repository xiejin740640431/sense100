// Code generated by go-bindata.
// sources:
// asset/config.json
// DO NOT EDIT!

package util

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

var _assetConfigJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xc1\x4a\xc3\x40\x10\x86\xef\x85\xbe\x43\x99\x43\x51\x90\xac\x69\xd2\xd0\x2c\x04\xbd\x2a\x11\xc4\xf4\x2e\xdb\x74\xdc\x44\xb2\x9d\xba\x33\x89\x88\xf8\xee\xb2\x71\x0f\x9e\xc4\xe3\xff\xcd\xff\x0d\xc3\x7c\x2e\x17\xab\x15\x34\xe8\x27\xf4\xa0\xa1\x13\x39\x6b\xa5\x06\x6a\xcd\xd0\x11\x8b\x2e\xf3\x32\x87\xab\xb9\x74\xe7\x8c\xc5\xbf\x9b\x8a\xc5\x48\xdf\xaa\x3e\x54\x39\x7a\xcd\xcc\xfe\x25\x46\xe3\x69\x3c\x49\xef\xf0\x81\x8e\x38\x80\x86\x23\x4e\x71\x50\x93\xad\x71\x8a\xf0\x30\xda\x88\xdd\x07\xbf\x05\xe6\x89\x44\xa7\x9b\x2c\xdf\x16\xb7\x17\x69\xb9\x49\xd2\x62\x97\x64\x49\x5a\x6c\x75\x96\x5d\x17\x97\xca\xd2\xb3\x20\xcb\x4d\xdb\x19\xcf\x28\xd5\x28\x2f\x3b\x77\xc8\xd7\xe7\x10\xf7\xbd\xc3\x6a\xef\x47\x5c\x0f\xd4\x56\x75\xb8\x2f\xee\x7f\x24\x2f\xa0\xe1\xd7\x2f\x9a\x77\x63\x2d\xfa\x7b\xa6\x13\x68\xe0\x9f\x94\xbc\x86\xb8\x5c\x7c\x7d\x07\x00\x00\xff\xff\x1e\x74\x7d\xe7\x55\x01\x00\x00")

func assetConfigJsonBytes() ([]byte, error) {
	return bindataRead(
		_assetConfigJson,
		"asset/config.json",
	)
}

func assetConfigJson() (*asset, error) {
	bytes, err := assetConfigJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "asset/config.json", size: 341, mode: os.FileMode(438), modTime: time.Unix(1520584385, 0)}
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
	"asset/config.json": assetConfigJson,
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
	"asset": &bintree{nil, map[string]*bintree{
		"config.json": &bintree{assetConfigJson, map[string]*bintree{}},
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

