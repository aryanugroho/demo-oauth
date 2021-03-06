// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../svc/partner/views/hi.html
// ../../svc/partner/views/homepage.html

package views


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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataSvcPartnerViewsHiHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x51\x74\xf1\x77\x0e\x89\x0c\x70\x55\xc8\x28\xc9\xcd\xb1\xe3\xb2\x01\x51\x0a\x39\x89\x79\xe9\xb6\x4a\xa9\x79\x4a\x20\x81\xd4\xc4\x14\x3b\x2e\x05\x05\x05\x05\x9b\xdc\xd4\x92\x44\x85\xe4\x8c\xc4\xa2\xe2\xd4\x12\x5b\xa5\xd0\x10\x37\x5d\x0b\x25\xa8\x54\x49\x66\x49\x4e\xaa\x9d\x47\x7e\x6e\x6a\x41\x62\x7a\xaa\x8d\x3e\x84\xcf\x65\xa3\x0f\xd1\x6e\x93\x94\x9f\x52\x09\x51\xea\x93\x9f\x9e\x99\xa7\x50\x5c\x9a\x9c\x9c\x5a\x5c\x9c\x56\x9a\x93\x53\xc9\x65\xa3\x0f\x91\xb6\xd1\x07\x3b\x02\x10\x00\x00\xff\xff\x8f\x3b\x7c\xff\x94\x00\x00\x00")

func bindataSvcPartnerViewsHiHtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataSvcPartnerViewsHiHtml,
		"../../svc/partner/views/hi.html",
	)
}



func bindataSvcPartnerViewsHiHtml() (*asset, error) {
	bytes, err := bindataSvcPartnerViewsHiHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "../../svc/partner/views/hi.html",
		size: 148,
		md5checksum: "",
		mode: os.FileMode(436),
		modTime: time.Unix(1636103855, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataSvcPartnerViewsHomepageHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8e\xbd\x0e\x82\x40\x10\x84\x7b\x9e\x62\xdd\x5e\xaf\xb5\x58\x2e\x31\xfe\xc4\x82\x44\x63\xa0\xb0\x5c\x65\xbd\x23\x81\x3b\x83\x6b\x61\x08\xef\x6e\xe0\xa8\xf6\x67\xe6\x9b\x0c\xad\x0e\x97\x7d\x79\xbf\x1e\xc1\x6b\xd7\xda\x8c\xa6\x01\x2d\x07\x97\xa3\x04\x9c\x1e\xc2\xb5\xcd\x00\x00\xa8\x13\x65\x78\x7a\xee\x3f\xa2\x39\x56\xe5\x69\xbd\xc5\x45\xd2\x46\x5b\xb1\xe7\xd8\xc9\x9b\x9d\x90\x49\x77\x46\x26\xe1\xf4\x88\xf5\x6f\xb1\x32\xf8\x5e\x5e\x39\x0e\xc3\xa6\x88\xae\x09\xd5\xad\x18\x47\xb4\xf3\x0e\xb5\x04\xc7\x01\xe2\xee\xab\x9e\x0c\x4f\x09\x09\x25\x33\x17\xfc\x07\x00\x00\xff\xff\x4b\xcf\xec\xdc\xb0\x00\x00\x00")

func bindataSvcPartnerViewsHomepageHtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataSvcPartnerViewsHomepageHtml,
		"../../svc/partner/views/homepage.html",
	)
}



func bindataSvcPartnerViewsHomepageHtml() (*asset, error) {
	bytes, err := bindataSvcPartnerViewsHomepageHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "../../svc/partner/views/homepage.html",
		size: 176,
		md5checksum: "",
		mode: os.FileMode(436),
		modTime: time.Unix(1636103462, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"../../svc/partner/views/hi.html":       bindataSvcPartnerViewsHiHtml,
	"../../svc/partner/views/homepage.html": bindataSvcPartnerViewsHomepageHtml,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"..": {Func: nil, Children: map[string]*bintree{
		"..": {Func: nil, Children: map[string]*bintree{
			"svc": {Func: nil, Children: map[string]*bintree{
				"partner": {Func: nil, Children: map[string]*bintree{
					"views": {Func: nil, Children: map[string]*bintree{
						"hi.html": {Func: bindataSvcPartnerViewsHiHtml, Children: map[string]*bintree{}},
						"homepage.html": {Func: bindataSvcPartnerViewsHomepageHtml, Children: map[string]*bintree{}},
					}},
				}},
			}},
		}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
