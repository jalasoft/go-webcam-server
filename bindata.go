// Code generated by go-bindata.
// sources:
// assets/index.html
// assets/js/logic.js
// DO NOT EDIT!

package webcamserver

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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x51\x74\xf1\x77\x0e\x89\x0c\x70\x55\xc8\x28\xc9\xcd\xb1\xe3\xe2\xb2\x81\xd0\x0a\x0a\x0a\x0a\x36\x19\xa9\x89\x29\x10\x26\x98\x5b\x9c\x5c\x94\x59\x50\xa2\x90\x92\x9a\x96\x5a\xa4\x50\x5c\x94\x6c\xab\xa4\x9f\x55\xac\x9f\x93\x9f\x9e\x99\xac\x97\x55\xac\x64\x67\xa3\x0f\x51\x01\xd5\xad\x8f\xd0\x6e\x93\x94\x9f\x52\x89\x64\x52\x66\x6e\x3a\x4c\x11\x44\xc6\x46\x1f\x6c\x2d\x20\x00\x00\xff\xff\x07\x8b\xb6\x2a\x8e\x00\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 142, mode: os.FileMode(420), modTime: time.Unix(1594313335, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsLogicJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x54\x4d\x6f\xd4\x30\x10\xbd\xfb\x57\x0c\x06\x69\xe3\xb6\x4a\x38\xb7\x0a\x12\x2a\x1c\x2a\x28\x1c\x16\xc4\x91\xba\xf6\xec\xd6\xaa\xe3\x09\xb6\xb3\x11\xac\xf6\xbf\xa3\x38\x1f\x4a\xdb\xcd\x72\xe0\x16\x8f\x9f\xdf\x3c\x3f\xbf\x09\x53\xe4\x42\x04\x4b\x0a\x4a\x68\x8d\xd3\xd4\xe6\x96\x94\x8c\x86\xdc\xb0\x17\x48\x3d\x62\x84\x12\x1c\xb6\xf0\x03\xef\xd7\x69\x9d\xdd\xb5\xe1\xb2\x28\xde\xec\x2d\xa9\xfc\x81\x42\x3c\x14\x4a\x56\xe8\x65\xb1\x33\x1a\xe9\x6d\x11\xa2\x47\x59\xdd\x09\xc6\x7a\x82\x5c\x6a\xfd\x71\x87\x2e\x7e\x36\x21\xa2\x43\x9f\xad\xa8\x46\xb7\xba\x00\xec\xaa\x50\xbe\x83\x3d\x03\x00\xe8\xba\x92\xc5\xdc\xd2\x36\xe3\xd7\xe4\x1c\xaa\x4e\x0d\x60\x88\xf2\xde\x9a\xf0\x80\x9a\x8b\x84\x1c\x88\x03\x3a\x9d\xf1\x6f\x37\xd7\x9f\xb8\x60\x07\xb1\xdc\x4f\x59\x0a\x38\x6f\xb8\xd4\x2a\x01\x35\x17\x27\xb8\x2a\x0c\x41\x6e\x97\xd9\x6e\xfb\x7d\xf0\xa8\xd0\xec\xfe\x41\x86\xde\x93\x5f\xa2\x5a\xbd\x77\x90\x00\x40\x4a\x35\xde\xa3\x5e\x09\xc1\x18\x2b\x8a\xe1\xbd\x5e\xf0\xf1\x0f\x5f\x6f\xaf\xc9\xc5\xae\x46\x52\xa3\xe6\x17\xb0\x69\x5c\xba\x5b\x26\x60\xcf\x8a\xb3\xc9\xe9\x08\xfd\xb3\x41\x09\x9a\x54\x53\xa1\x8b\xf9\xaf\x06\xfd\xef\x35\x5a\x54\x91\x7c\xc6\x5f\xf7\x88\x9f\x4e\x56\xc8\x45\xbe\x93\xb6\xc1\x74\x7e\x46\x72\x34\x40\x57\x33\x40\x1b\xbe\x7b\x03\x25\xf0\x36\x74\xb9\xe1\x70\x0e\x63\x72\xe0\x1c\xf8\x18\x9e\xae\x3e\x08\xea\xaa\x7d\x88\xf8\xd3\x6e\x93\xc7\x37\xce\x44\x23\xad\xf9\x93\xba\xc1\x86\x3c\x3c\x3d\x0f\x8a\xaa\xda\x62\xec\xec\xbf\x62\x89\x60\x27\xfd\x90\x9b\x7e\xbd\x78\xe9\x10\xa5\x8f\x5c\x1c\x71\x57\x59\xa3\x1e\xe7\x96\xa2\x18\xa2\x3b\xc9\x5c\xcc\xb1\xdb\x42\xa4\xa4\x32\x19\x22\x26\xf4\xc2\x94\x0d\xa8\x67\xb0\x9c\xdc\x90\x3f\x28\x17\x64\xbc\x50\xf0\x05\x15\x41\xed\x4d\xb0\x74\x99\x04\x60\xae\x65\x94\xe2\x6a\x3a\x72\x48\x5f\x87\xd1\xa9\x13\xce\x50\xfd\x3f\xc6\x98\x0d\x64\xaf\xfa\x8b\x9c\x56\x4c\xa3\x2b\xdd\x8f\xa2\x7f\xc3\x39\xd8\x63\x6c\xbc\x9b\xeb\x3f\xee\xfd\x3a\x52\x5d\x77\xce\xab\x69\xc2\xf3\x91\xeb\xb9\xb1\x69\xf2\x33\x31\x3a\x71\x56\xb0\xa2\x38\x08\xf6\x37\x00\x00\xff\xff\xb4\xd9\xfc\x30\x27\x05\x00\x00")

func jsLogicJsBytes() ([]byte, error) {
	return bindataRead(
		_jsLogicJs,
		"js/logic.js",
	)
}

func jsLogicJs() (*asset, error) {
	bytes, err := jsLogicJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/logic.js", size: 1319, mode: os.FileMode(420), modTime: time.Unix(1594315731, 0)}
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
	"index.html": indexHtml,
	"js/logic.js": jsLogicJs,
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
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
	"js": &bintree{nil, map[string]*bintree{
		"logic.js": &bintree{jsLogicJs, map[string]*bintree{}},
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

