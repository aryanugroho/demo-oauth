// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../svc/user/views/consent.html
// ../../svc/user/views/login.html

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

var _bindataSvcUserViewsConsentHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\x6d\x6f\xdb\x36\x10\xfe\x9e\x5f\xc1\x72\xd8\xa7\x95\x96\x9d\xb7\x76\xaa\x64\x20\xcd\xd6\x02\xeb\x5a\xb4\xcb\xd6\x6d\x18\x86\x82\x22\x4f\x12\x6b\x8a\xd4\xc8\x93\x5f\xe1\xff\x3e\x50\xb2\x5c\x5b\x71\xda\x38\x40\x24\xde\x3d\x3c\xde\xf3\xf0\x8e\x54\xf2\x44\x5a\x81\xab\x1a\x48\x89\x95\x9e\x9e\x25\xe1\x41\x34\x37\x45\x4a\xc1\xd0\x60\x00\x2e\xa7\x67\x84\x10\x92\x54\x80\x9c\x88\x92\x3b\x0f\x98\xd2\x06\x73\xf6\x9c\x1e\xba\x0c\xaf\x20\xa5\x73\x05\x8b\xda\x3a\xa4\x44\x58\x83\x60\x30\xa5\x0b\x25\xb1\x4c\x25\xcc\x95\x00\xd6\x0e\x9e\x12\x65\x14\x2a\xae\x99\x17\x5c\x43\x3a\x79\x4a\x7c\xe9\x94\x99\x31\xb4\x2c\x57\x98\x1a\x7b\x22\xb4\x04\x2f\x9c\xaa\x51\x59\x73\x10\xfd\x04\x90\x37\x58\x5a\x77\x80\xf9\xbb\xf1\x4d\xde\x03\x51\xa1\x86\xe9\x4d\x8b\x51\x6b\x20\x37\x86\xdc\xd4\x75\x12\x75\xf6\xb3\xb3\x0e\xf5\x84\x31\xf2\xd2\x5a\xf4\xe8\x78\x4d\x84\x75\x40\x6e\xef\xee\x08\x63\xbb\x28\x5a\x99\x19\x71\xa0\x53\xea\x71\xa5\xc1\x97\x00\x48\x49\xe9\x20\x4f\x69\x89\x58\xfb\x38\x8a\x84\x34\x9f\xfd\x48\x68\xdb\xc8\x5c\x73\x07\x23\x61\xab\x88\x7f\xe6\xcb\x48\xab\xcc\x47\xb8\x50\x88\xe0\x58\xd6\x2f\x13\x5d\x8e\xae\x47\xe3\x48\x78\x1f\xed\x6d\xa3\x4a\x99\x91\xf0\x9e\x12\x65\x10\x0a\xa7\x70\x95\x52\x5f\xf2\xab\xc9\x39\x7b\x7f\xf5\xb6\x78\x6b\x26\x9f\x5f\xbe\x1b\x4f\xb8\x7f\x59\xfc\x31\x5e\x5f\x8f\x3f\xcc\x2e\x3f\x2c\xd5\x5f\xf6\xf9\xf5\x0f\x0b\xfd\x8a\x97\x6f\x9c\xff\x90\x5f\x3c\x13\x3f\x0a\x67\xaf\x26\xcf\xfe\x5c\x7f\xbc\x7b\xff\xfe\xe3\x04\x7f\x7a\xb3\x2e\x67\xcd\xb9\xfa\xe5\xfc\xd5\xc7\xe2\xd7\xeb\x67\x8b\xd7\xe3\x8b\xbb\xd7\xe6\xdd\x4d\x9a\x52\x22\x9c\xf5\xde\x3a\x55\x28\x93\x52\x6e\xac\x59\x55\xb6\xf1\x94\x44\x3b\xf6\x2d\xe7\xee\x3d\xfc\x42\xd9\x3c\xdd\x8f\x32\x2b\x57\x64\xb3\x1f\xb6\x00\x50\x45\x89\x31\x99\x8c\xc7\xdf\xbf\xd8\x7b\xb6\x67\x5f\x9b\x23\x95\xaf\x35\x5f\xc5\x84\x55\x9e\xe5\x1a\x96\x99\x5d\xbe\x38\x8d\x08\xde\x63\x57\x3f\x87\x71\xad\x0a\x13\x13\x01\x06\xc1\x1d\x63\x5a\x17\x53\x08\x95\x3f\x0d\xa8\xb9\x94\xca\x14\x0c\x6d\x1d\x93\xcb\x71\xbd\x3c\xed\xce\x2c\xa2\xad\x4e\x21\x32\x2e\x66\x85\xb3\x8d\x91\x4c\x58\x6d\x5d\x4c\xbe\xcb\xaf\xc2\xdf\x49\x0d\x46\xb9\x75\x15\x13\xd6\x78\x30\x38\x10\xa3\x6d\x9b\xa1\x7e\xe1\x57\xf1\x25\xdb\x39\x2f\x2e\x1e\x4a\x31\x26\x93\xab\xa1\xab\xe2\xae\x50\x26\x26\xbc\x41\x7b\x98\xce\xe9\x6c\x46\xa2\x04\x31\xcb\xec\x72\x90\x57\x6e\x0d\xb2\xc5\x6e\x77\x2f\xc7\xe3\x47\x44\xea\x47\xe8\xac\x1e\x44\xab\xad\x57\xa1\xbd\xe3\xd0\x59\x1c\xd5\x1c\x06\x7a\xda\x25\xf3\x6a\xdd\x32\xca\xac\x93\x6d\xf3\x0c\x78\xf5\xa5\x76\xcc\xeb\x58\x8c\x7b\x3a\xb5\x34\xbc\x5a\x43\x4c\x26\xd7\x87\xce\x47\xb1\x88\x73\x2b\x1a\x3f\xe0\xb2\x66\xca\x48\x58\xc6\xe4\xfc\xdb\xd1\x94\xa9\x1b\xfc\x27\x1c\xc4\x29\x85\x8a\x2b\x4d\xff\x1d\x44\xeb\x76\x6b\x5f\x6a\x6c\x72\xaf\xd4\x7a\x39\x02\x80\xb9\xa0\x01\x73\x5c\xaa\xc6\xc7\x64\xfc\x35\xa8\x86\xfc\x24\xf2\x31\xb9\xd6\xdc\xfb\x85\x75\xf2\x5b\xe9\xde\x17\x7c\x97\x03\xda\xfa\xc1\x04\x06\xb0\x07\x29\xed\x3a\x28\x89\x76\x87\xd2\x59\x12\x75\xb7\x56\xd2\x1e\x2a\x42\x73\xef\x53\x8a\xb0\x44\xd6\x35\x39\x0d\x98\xc0\xa9\xf7\x1d\xf2\xa3\xa4\x02\x2c\xad\x4c\x69\x6d\x3d\x52\xc2\x45\x28\xc8\x94\x46\xe1\x52\x89\x7a\x50\x77\xf4\x6d\x36\x2a\x27\xc6\x22\x19\xdd\x76\xf6\xdb\x92\x6b\x0d\xa6\x80\xed\x17\xf5\x12\xa9\xe6\x27\x92\x20\x5c\x83\xc3\xee\x3f\x93\xdc\x14\xe0\x28\x71\x56\x87\xeb\x2b\xd8\xe8\xf4\x48\x88\x24\x9b\x6e\x36\x64\xf4\xb3\x73\xd6\xfd\x1e\x2e\x2a\xb2\xdd\x26\x51\x36\x04\xb9\x63\xc3\x7e\xca\x6d\x77\x13\x92\xc3\xbc\x22\xa9\xe6\x3d\x0f\x30\x72\xbb\xd3\xb1\x25\xf5\x10\xa1\xa4\x9c\xf4\x5c\xca\x0b\x52\x65\xec\xe2\xf0\x04\x60\xc6\xba\x8a\x6b\xba\xbf\x5c\x79\xd0\x2e\x89\xca\xc9\xee\xda\xa8\xa7\x49\x54\xf7\x6b\xba\x40\x9a\x8c\x7e\x83\xff\x1a\xf0\x08\xf2\x4e\xd8\x1a\xfc\x03\xca\x75\x5b\x14\xce\xa0\xa1\x30\x6d\x3d\xde\x87\xb1\xd6\x4e\x49\x57\xa8\xfd\xe9\x45\x77\x5f\x08\x85\xe3\x06\x3f\xf9\xb0\x22\x25\x73\xae\x1b\x48\xe9\x66\x33\xda\x6e\x29\x51\x72\xff\xda\xce\x02\x39\x58\x50\xf3\x0c\xf4\x89\x05\x5b\x3b\x25\xb9\x75\x7d\x80\x69\xfb\x48\xa2\xd6\x33\xfd\x9a\xf2\x07\x4c\xba\x84\x4b\x25\x25\x98\x3e\xdd\x5d\xd9\x7d\x12\xfd\x76\x1c\x26\x7d\x7f\xab\xfa\x2f\x9d\xac\x41\xb4\xa6\x4f\x35\x43\x43\x32\x34\x4c\x17\xed\xa3\x76\xaa\xe2\x6e\xd5\xbe\x67\xda\x8a\x59\xaf\x95\x6f\xb2\x4a\xe1\x97\x4d\x84\x24\xea\x02\x1d\xa7\x9c\x44\x81\x7b\xdb\x6e\xa1\xcf\xa6\xa1\xed\xda\x8f\xc8\xff\x03\x00\x00\xff\xff\xf5\x24\x39\x82\x55\x0a\x00\x00")

func bindataSvcUserViewsConsentHtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataSvcUserViewsConsentHtml,
		"../../svc/user/views/consent.html",
	)
}



func bindataSvcUserViewsConsentHtml() (*asset, error) {
	bytes, err := bindataSvcUserViewsConsentHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "../../svc/user/views/consent.html",
		size: 2645,
		md5checksum: "",
		mode: os.FileMode(436),
		modTime: time.Unix(1636103151, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataSvcUserViewsLoginHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x57\xfb\x6f\xdb\xb6\x13\xff\x3d\x7f\x05\xcb\x2f\xbe\xc3\x86\x95\x96\x9d\x67\xa7\x4a\xc6\xd6\xae\x2d\xb6\x3e\x90\x2e\x5b\xb7\x61\x18\x0a\x8a\x3c\x4b\xac\xf9\xd0\x48\x2a\xb6\x13\xf8\x7f\x1f\xa8\x87\x23\xcb\x4a\xd3\x1f\xe6\x00\xb1\x74\x77\x3c\x7e\xee\xc3\xbb\xe3\x39\x79\xc4\x0d\xf3\x9b\x12\x50\xe1\x95\x9c\x1f\x25\xe1\x0b\x49\xaa\xf3\x14\x83\xc6\x41\x00\x94\xcf\x8f\x10\x42\x28\x51\xe0\x29\x62\x05\xb5\x0e\x7c\x8a\x2b\xbf\x20\x4f\x70\x5f\xa5\xa9\x82\x14\x5f\x0b\x58\x95\xc6\x7a\x8c\x98\xd1\x1e\xb4\x4f\xf1\x4a\x70\x5f\xa4\x1c\xae\x05\x03\x52\xbf\x3c\x46\x42\x0b\x2f\xa8\x24\x8e\x51\x09\xe9\xec\x31\x72\x85\x15\x7a\x49\xbc\x21\x0b\xe1\x53\x6d\x46\x5c\x73\x70\xcc\x8a\xd2\x0b\xa3\x7b\xde\x47\x0c\x69\xe5\x0b\x63\x7b\x36\x7f\x56\xae\x5a\x74\x86\x5e\x78\x09\xf3\x2b\x91\x6b\xf4\x93\x4e\xa2\xe6\xf5\xe8\xa8\x51\x3e\x22\x04\x3d\x33\xc6\x3b\x6f\x69\x89\x98\xb1\x80\x9e\x5f\x5d\x21\x42\xda\xc5\x52\xe8\x25\xb2\x20\x53\xec\xfc\x46\x82\x2b\x00\x3c\x46\x85\x85\x45\x8a\x0b\xef\x4b\x17\x47\x11\xe3\xfa\x93\x9b\x30\x69\x2a\xbe\x90\xd4\xc2\x84\x19\x15\xd1\x4f\x74\x1d\x49\x91\xb9\xc8\xaf\x84\xf7\x60\x49\xd6\x6d\x13\x9d\x4e\xce\x27\xd3\x88\x39\x17\xed\x64\x13\x25\xf4\x84\x39\x87\x91\xd0\x1e\x72\x2b\xfc\x26\xc5\xae\xa0\x67\xb3\x63\x72\x79\xf6\x36\x7f\xab\x67\x9f\x9e\xbd\x9b\xce\xa8\x7b\x96\xff\x36\xbd\x39\x9f\xbe\x5f\x9e\xbe\x5f\x8b\x3f\xcc\x93\xf3\x6f\x57\xf2\x25\x2d\x5e\x5b\xf7\x7e\x71\x72\xc1\xbe\x63\xd6\x9c\xcd\x2e\x7e\xbf\xf9\x70\x75\x79\xf9\x61\xe6\x7f\x7c\x7d\x53\x2c\xab\x63\xf1\xf3\xf1\xcb\x0f\xf9\x9b\xf3\x8b\xd5\xab\xe9\xc9\xd5\x2b\xfd\xee\x87\x34\xc5\x88\x59\xe3\x9c\xb1\x22\x17\x3a\xc5\x54\x1b\xbd\x51\xa6\x72\x18\x45\x6d\xf4\x75\xcc\xcd\x73\xf8\x4c\x32\x4e\x4a\x49\x19\x14\x46\x72\xb0\x44\xa8\x1c\xdd\xee\xb4\xe1\xb3\x30\xda\x13\x27\x6e\x20\x46\xb3\xc9\xec\xf8\xcc\x82\x7a\xba\x67\xe0\x61\xed\x09\xd5\xac\x30\x36\x46\x4a\x70\x2e\x61\xdf\x80\xac\x20\x5b\x0a\x4f\x2a\x07\x96\x38\x90\xc0\x7c\x8c\xb4\xd1\x43\x33\x65\x6e\x1e\xb6\x71\x0f\x99\x7c\x46\xbd\x3d\xda\x3d\x7e\xaf\x80\x0b\x8a\xbe\x56\x42\x37\xf9\x1c\xa3\x8b\xf3\x27\xe5\xfa\x9b\x41\xf4\x23\xfc\x10\x39\xa4\x68\x40\xd3\xc9\xe4\x90\xa4\xed\x18\x88\x50\xaa\x8f\x77\x6f\x99\xe1\x9b\x81\xe3\x02\x44\x5e\xf8\x18\xcd\xa6\xd3\xff\x8f\x86\x31\xb2\x86\x0b\x57\x4a\xba\x89\x6b\xae\x16\x12\xd6\x99\x59\x3f\x1d\xb7\x08\xda\x43\x7e\x83\x94\x50\x29\x72\x1d\x23\x06\xda\x83\xdd\xb7\xa9\x55\x44\x78\x50\x6e\xdc\xa0\xa4\x9c\x0b\x9d\x13\x6f\xca\x18\x9d\x4e\xcb\xf5\xb8\x3a\x33\xde\x1b\x35\x66\x91\x51\xb6\xcc\xad\xa9\x34\x27\xcc\xc8\x90\x56\xff\x5b\x9c\x85\xbf\x51\x0e\x26\x0b\x63\x15\x71\x22\xd7\x42\x0f\xb8\x68\x4f\x76\x9f\xbe\xf0\x51\x74\xdd\x1d\xfb\xc9\xc9\x7d\x08\x63\x34\x3b\x1b\xaa\x14\xb5\xb9\xd0\x31\xa2\x95\x37\x7d\x34\xa3\x60\x26\xac\x00\xb6\xcc\xcc\x7a\xac\xa4\x56\xed\xd9\x9e\x4e\xa7\x0f\x3b\xaa\x5f\x42\x2f\xb4\x46\x0e\x9c\x95\xc6\x89\xd0\x4f\xe3\xd0\xd3\xa8\x17\xd7\x83\x7a\xc8\xcc\x3a\xe4\x65\x1d\x4f\x66\x2c\xaf\xdb\xd6\x20\xaa\x2e\xcf\xf6\xa3\xda\xa7\xe2\x80\xa5\x7e\x63\x38\xef\x2b\xbf\x24\x88\x78\x61\x58\xe5\x06\xa1\xdc\x10\xa1\x39\xac\x63\x74\xfc\xa0\x33\xa1\xcb\xca\xff\x15\xee\xbd\x14\x83\xa2\x42\xe2\xbf\x07\xce\x9a\x93\xda\x65\x19\x99\x1d\x64\x59\x47\x46\x30\x20\x36\x30\x40\x2c\xe5\xa2\x72\x31\x9a\x7e\xce\x54\xc2\x62\xd4\xf2\x0b\xa0\x96\xd4\xb9\x95\xb1\xfc\x21\xb4\x87\x6c\xb7\x10\xbc\x29\xef\xdd\x7f\x60\x76\x6f\x44\x6d\xed\x24\x51\x7b\x17\x1c\x25\x51\x33\x23\x24\x75\x3b\x61\x92\x3a\x97\xe2\xba\xad\x37\xe5\x8d\x83\x4d\x08\xa9\xd3\xf5\xc2\xc3\x48\x81\x2f\x0c\x4f\x71\x69\x9c\xc7\x88\xb2\x90\x8c\x29\x8e\xc2\x0d\x1e\x49\x93\x0b\xdd\xde\xda\xb7\xb7\x62\x81\xb4\xf1\x68\xf2\x26\x48\x9f\x17\x54\x4a\xd0\x39\x6c\xef\x78\x4b\xb8\xb8\x1e\xd9\x1f\x51\x09\xd6\x37\xff\x09\xa7\x3a\x07\x8b\x91\x35\x32\xcc\x09\x41\x86\xe7\x7b\x1c\x24\xd9\xfc\xf6\x16\x4d\x5e\x58\x6b\xec\xaf\x61\x34\x40\xdb\x6d\x12\x65\x43\x23\xbb\x2f\xd8\x2d\x79\xde\x8c\x1c\xa8\x8f\x2b\xe2\xe2\xba\x8b\x02\x34\xdf\x6e\xdb\x61\xa3\x8e\x69\x3c\x9e\x46\x55\x7b\x74\xff\x79\x88\x3b\xb0\xee\x01\x98\xb5\xb8\x98\x75\x3b\x16\x27\x48\x65\xe4\xa4\xdf\x81\x88\x36\x56\x51\x89\xe7\x97\x12\xa8\x03\x14\x8e\x15\x09\x9d\x44\xc5\xac\x1d\x19\xea\x04\x46\x4d\x02\x17\x82\x73\xd0\xb8\x9d\xd2\xea\xe3\xfd\xc8\xba\xc8\x31\xba\xa6\xb2\x82\x14\xdf\xde\x1e\x70\xd2\x8d\x6e\x92\x66\x20\xd1\xc2\xd8\x14\xd7\x7e\x5f\xd4\xc5\xdb\xe1\x73\x96\x18\x2d\x37\x78\x5e\x8b\x11\xe5\xdc\x82\x73\x49\x54\xaf\x1a\x81\xd3\x94\x3e\x12\x7c\xd4\x5b\xbf\xe1\x74\x90\xdb\x15\xbd\x3b\x3d\xc5\x7b\x9b\x61\x64\xe1\x9f\x4a\x58\xe0\x75\x3b\xac\xfb\xd4\x3d\xd8\x2f\xbb\x6a\x3e\x80\xdf\x69\xee\x47\xbe\xeb\x04\x77\xe0\x0f\xdc\x8d\xe1\xbf\x5b\xb7\x17\xc2\xdd\xda\x0e\x7d\xbb\x67\x2f\xdd\x76\xb7\x51\x48\x81\x5e\x42\x25\x3d\x8c\x3b\x59\x1f\x6b\xb7\xb0\xc3\x60\x41\x81\xca\xc0\x7e\x54\x77\x47\xee\x6d\x05\x78\x8e\x7e\x69\x55\x48\x41\x2f\x2d\xfb\x24\xdc\xe5\x68\x92\x55\xde\x1b\xdd\xc1\xcb\xbc\x46\x99\xd7\x61\xbc\x0a\x5f\xa5\x15\x8a\xda\x4d\xfd\x9c\x49\xc3\x96\xb8\x45\xe3\xaa\x4c\x09\x8f\x9b\xf1\x3f\x24\x6a\xe3\xa6\xf5\x59\x76\xee\x94\x27\x67\x4d\xb2\xd7\x65\xa6\x2a\x0f\x1c\xcf\xbf\x62\xa6\xdc\x3c\x45\xc7\xd3\xe3\x59\x12\x95\xfb\xb5\x92\x44\x81\xf0\xf0\x3b\xe2\x28\x89\x42\x2b\x9c\x87\xce\x58\xff\xaa\xfa\x37\x00\x00\xff\xff\x0f\xbd\x47\xf5\x66\x0d\x00\x00")

func bindataSvcUserViewsLoginHtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataSvcUserViewsLoginHtml,
		"../../svc/user/views/login.html",
	)
}



func bindataSvcUserViewsLoginHtml() (*asset, error) {
	bytes, err := bindataSvcUserViewsLoginHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "../../svc/user/views/login.html",
		size: 3430,
		md5checksum: "",
		mode: os.FileMode(436),
		modTime: time.Unix(1636104752, 0),
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
	"../../svc/user/views/consent.html": bindataSvcUserViewsConsentHtml,
	"../../svc/user/views/login.html":   bindataSvcUserViewsLoginHtml,
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
				"user": {Func: nil, Children: map[string]*bintree{
					"views": {Func: nil, Children: map[string]*bintree{
						"consent.html": {Func: bindataSvcUserViewsConsentHtml, Children: map[string]*bintree{}},
						"login.html": {Func: bindataSvcUserViewsLoginHtml, Children: map[string]*bintree{}},
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
