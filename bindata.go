// Code generated by go-bindata.
// sources:
// files/conf/dhcpcd.conf
// files/conf/wpa_supplicant.conf
// files/init.d/autorun_date
// files/init.d/vncboot
// DO NOT EDIT!

package setup

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

var _filesConfDhcpcdConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x54\x4d\x6f\xe3\x36\x10\xbd\xeb\x57\x0c\xa0\x4b\x8b\x66\x05\xcb\x1f\x41\xf7\x90\x43\x6a\x21\xad\x0f\x59\x04\x71\x82\x1e\x05\x4a\x1c\x45\x84\x29\x52\xe5\x8c\xec\xe6\xdf\x17\x43\x4a\x8e\x17\x05\x74\xe0\xc7\xbc\x99\xc7\x79\x6f\x94\xc3\x23\x90\x1a\x46\x8b\xd0\x7a\xd7\x99\x8f\x29\x28\x36\xde\x41\xe7\x03\xe8\xbe\x1d\x5b\x5d\x64\x39\x1c\x11\x97\x9d\x84\xfd\xb2\xfb\x35\x05\x20\x2b\x63\xa9\xc8\xb2\x1c\x1e\xad\xf5\x17\x98\x08\x03\x81\xef\x80\x7b\x43\xf0\x11\xfc\x34\x02\x7b\x30\x8e\x31\xa8\x96\xe1\x62\xb8\x9f\x33\xc1\xd9\x28\xe0\x3e\x16\xe6\xe0\x2d\x90\x6f\x4f\xc8\x45\x96\xcf\x07\x09\x7d\xe9\x11\xad\x14\x38\xb8\xce\x87\x21\x22\xaa\xbf\xf6\x2f\x40\x18\xce\x18\xa4\x96\x9f\x02\xf4\x9e\xd8\xa9\x01\x23\xaf\xaa\xfa\x71\x2c\xb2\xe5\x48\xc0\xef\x84\x11\xd9\xab\xa0\x2f\x2a\x20\x28\xad\x03\xd2\x4c\x15\x13\xc1\x4e\xb5\x09\x2f\x27\x7b\x6b\xd0\x31\x1c\xaa\x22\x6b\xe3\xd2\xe8\x2c\x07\x1f\x6e\x92\x91\xd4\xab\xde\x0f\x15\xfc\x06\x87\xc7\x43\x05\x8a\x80\x90\xc1\xb8\xc8\xf0\x7c\x9f\xc8\xc8\x72\x3b\xe7\x4b\x41\x23\x06\x78\x7d\xda\x6f\x37\xf7\x65\x91\xe5\x7a\x32\x5a\x38\xbe\x60\x20\x43\x7c\xc3\xe5\x67\x49\x2e\x3d\xba\xa5\x77\xf8\xaf\x61\x2a\xb2\x31\x41\xd0\xb1\x24\x78\x55\xa3\xd1\xd0\xfa\x61\x30\x0c\x34\x8d\xa3\x0f\x1c\xd5\x53\x1d\x8a\x0a\xe8\x54\x63\x11\x9a\x4f\xd0\xd8\xa9\xc9\x32\x34\xd8\xaa\x89\x10\x0c\x43\xc0\x7f\x26\x13\x90\xe2\xcb\x64\x7d\x56\x56\x1a\xe0\xc7\x58\x9c\x90\xe5\xf9\x2e\x3d\x3c\xb5\x9e\x3d\xa8\x96\x27\x65\xed\x27\x5c\x7c\x38\x15\xd9\x1c\x1c\x84\x48\x9d\x88\x44\x6f\x80\x95\x87\x89\x54\x31\x80\x04\x2a\x05\x91\x18\xba\xe0\xff\xa7\xea\x35\x93\xf6\x83\x32\xae\x16\x19\xeb\x74\x45\x77\xb7\x87\xd7\x0d\xa1\x0a\x6d\x7f\x17\x7d\x10\x6f\x96\x0c\xad\x55\x44\x16\x89\x6a\x62\xc5\xa6\xad\x83\x9f\x18\x29\xcb\xe1\xd9\x13\x83\x36\xc4\xc1\x34\x53\xa2\xd5\xab\x33\xc2\x8f\xb7\x97\xaf\xee\xcd\x59\x1c\x8f\x4b\x7d\x69\x34\xd2\x88\x2d\x47\xd6\x0e\x59\xde\x0e\xcf\x6f\xef\xb1\xd7\x7e\xb8\x75\x93\x0e\x46\x30\x10\x50\x8c\x11\x15\x6c\x7b\xe5\x3e\x8c\xfb\x88\xe8\xe7\xb7\x77\x20\x2f\x2c\x44\x1a\x7d\xa3\x4d\x91\xe5\x73\xed\x6b\xb6\x7a\xe0\x29\xb5\xf3\x18\xa9\x1c\x2a\x30\xb4\x08\x17\xb1\xaf\x4f\xfb\x75\xb9\x29\x8b\x6c\x3e\x8c\x76\x99\x89\xd7\x46\x8b\x8b\x3b\x83\x41\x92\xfc\x89\x0e\x83\x62\x84\x23\x47\x57\xbc\x04\x73\x96\xed\x41\x7c\xfb\x98\xa6\x03\x09\x8c\x23\x46\xa5\x45\xbb\xeb\xec\x34\x8a\x50\x83\x77\x48\x19\x59\xa5\x5a\x18\x13\x36\x71\xeb\xbd\x3f\x01\xb5\xc1\x8c\x2c\xf4\xc6\xe0\xcf\x46\xa3\x16\xc9\xad\xf7\x27\xf9\x21\xc8\x20\x2e\xf3\x6a\x3a\x70\x9e\xe3\xdc\x34\x9f\x57\x1f\x64\xf9\xec\x84\x3b\x68\x26\x16\x7f\x52\xef\x27\xab\x63\x6c\x83\x10\x26\xf7\x53\xaf\x9c\x8f\x65\x53\x81\x6f\x5f\x93\xff\xa5\xc4\xc5\x2a\xb7\xca\x88\x8c\x86\xe3\xf1\x50\x7d\xfb\x7b\xb5\xa9\xab\x7d\xf5\xbd\xbc\xdf\x6e\x9e\x56\xdb\x3f\xb2\x64\x0f\x30\x63\x3d\xff\x1b\x1e\xca\xef\xeb\xa2\xbc\xff\xbd\x28\x57\xab\xa2\x5c\x6f\x96\x88\x68\xa0\x40\x0f\xe5\xaa\x90\x6f\x57\xac\x77\xdb\xe5\x8e\xa6\xc6\x21\x0f\x8a\x4e\x0f\xeb\x9d\xdc\xec\x8a\xf5\x76\x55\xac\xfe\x0b\x00\x00\xff\xff\x43\x40\xbe\x3f\x68\x05\x00\x00")

func filesConfDhcpcdConfBytes() ([]byte, error) {
	return bindataRead(
		_filesConfDhcpcdConf,
		"files/conf/dhcpcd.conf",
	)
}

func filesConfDhcpcdConf() (*asset, error) {
	bytes, err := filesConfDhcpcdConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/conf/dhcpcd.conf", size: 1384, mode: os.FileMode(420), modTime: time.Unix(1503552024, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesConfWpa_supplicantConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x4b\x2d\x29\xcf\x2f\xca\xb6\xad\xe6\x52\x50\x50\x50\x28\x2e\xce\x4c\xb1\x55\x0a\x37\x30\x8e\x77\x71\x76\xb1\x34\x34\x33\x31\x76\x33\x30\x71\x52\x02\xcb\x15\x14\x67\xdb\x2a\xe5\xa5\xa6\x64\x9a\x1b\x66\x14\xe5\x5a\x98\xe4\x19\xa7\x15\x42\x64\xb2\x53\x2b\xe3\x73\xd3\x73\x4b\x6c\xc3\x03\x1c\x75\x03\x82\xbd\xb9\x6a\x01\x01\x00\x00\xff\xff\xd0\x4e\xdc\x77\x56\x00\x00\x00")

func filesConfWpa_supplicantConfBytes() ([]byte, error) {
	return bindataRead(
		_filesConfWpa_supplicantConf,
		"files/conf/wpa_supplicant.conf",
	)
}

func filesConfWpa_supplicantConf() (*asset, error) {
	bytes, err := filesConfWpa_supplicantConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/conf/wpa_supplicant.conf", size: 86, mode: os.FileMode(420), modTime: time.Unix(1503551847, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesInitDAutorun_date = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xce\xbf\x4a\xc3\x50\x14\xc7\xf1\xfd\x3e\xc5\x91\xe3\x1a\xe2\xff\xa1\x9b\x92\x2a\x59\xa2\xd8\xee\x72\x4d\x4e\xc9\xa5\x69\x6e\xbd\xf7\x44\xd0\xad\xb7\x20\x05\x07\xc5\x45\x44\x07\x17\x27\xdd\x45\x7c\x9c\x43\xf3\x1c\xd2\x41\x68\xc6\x1f\x1f\x7e\xf0\xc5\x8d\xf8\xd2\xd4\xb1\x2f\x15\x22\xc2\x51\xff\x24\xcd\x20\xcd\xd2\x21\xa4\xd9\xf1\xa9\x42\x38\x73\xf6\xda\x14\xe4\x7b\x30\xd1\x5c\x8e\xc9\xd5\x54\x29\x84\x73\xba\x6a\x8c\xa3\x22\x1a\xb0\x76\xdc\x83\xcd\xca\xe6\xba\xba\x18\xf9\xae\xd9\x69\x97\x12\x1a\xe9\xa6\xe2\xff\xd7\x0e\xec\xc2\x1e\xec\x77\x60\x75\xd9\x82\xed\x03\x85\x30\x28\xad\xe3\x28\x21\x9f\x3b\x33\x65\x63\xeb\x6e\x04\x22\xf4\xb3\x64\xad\x56\xa1\x84\x6f\x09\x3f\x32\xbf\x93\xf9\x7b\xfb\x12\x96\x8b\x5f\x09\x4f\xed\xf3\x47\xfb\xf6\x25\xb3\xcf\xe5\xe3\x42\xc2\x83\xcc\x5e\x25\xdc\x2b\xdf\x14\x16\xd8\x4c\xa8\xd0\x4c\x39\x57\xe0\x89\xa3\xd5\xbe\xb5\x35\xc1\xa1\x37\x3a\x1e\xda\xf1\x8d\xfd\x0b\x00\x00\xff\xff\x7d\x20\xb3\x9d\x21\x01\x00\x00")

func filesInitDAutorun_dateBytes() ([]byte, error) {
	return bindataRead(
		_filesInitDAutorun_date,
		"files/init.d/autorun_date",
	)
}

func filesInitDAutorun_date() (*asset, error) {
	bytes, err := filesInitDAutorun_dateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/init.d/autorun_date", size: 289, mode: os.FileMode(420), modTime: time.Unix(1503553766, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesInitDVncboot = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x41\x6f\xd3\x40\x10\x85\xef\xfb\x2b\x1e\x76\xa4\x0a\x24\xd7\x76\x09\x48\xb8\xea\x05\x12\xc0\x07\x5c\xd4\x50\xae\xc8\xb5\x27\xf1\x8a\x64\x67\xd9\x19\x5b\x89\x80\xff\x8e\xec\xa8\xd0\x20\x10\x1c\xad\x79\xef\x7d\x6f\xc6\x1b\x3f\x42\x7a\x67\x5d\x2a\x9d\x89\x91\x92\x36\xa9\x75\x56\xcf\xdb\x74\x70\xcd\x1d\xb3\x9a\x38\x8e\xf1\x72\xf9\xa6\xac\x50\x56\xe5\x07\x94\xd5\xeb\x6b\x13\xe3\x7d\xe0\xc1\xb6\x24\x05\x7e\xea\x70\x43\x5f\x7a\x1b\xa8\x4d\x56\x5a\x07\x2d\x30\x0b\xb4\x63\xa5\x4f\x6b\xc1\x4c\x0e\xb2\xe5\xcd\xa9\x88\xfd\x5f\x34\x0b\x5a\xd7\xfd\x56\xef\x73\x2e\xf0\x14\x73\x3c\x3b\x19\x8c\xde\x0c\x39\x9e\x9b\x18\xab\x8e\x83\x26\x0b\x92\x26\x58\xaf\x96\x5d\x81\xc9\x89\x8f\xd5\x2b\xac\x28\x0c\x14\x50\x2b\xc6\x96\x50\xbb\xa3\x29\xe8\x7f\xc5\xe7\xd3\x01\x96\xd5\xe2\xc1\xfa\xb7\xab\xe5\xcd\x95\xb7\xe6\xed\xf5\xbb\xe5\x55\xda\xf1\x8e\x52\x6f\x0d\xed\x3d\x07\xc5\x38\xc4\x38\x31\xa6\xa9\x85\x10\xcd\xf2\x08\xd6\x19\x19\x29\x8f\x0d\x35\x1d\x23\x9a\x90\xd6\x6d\x1e\x50\x23\x13\x97\x4e\x28\x28\x0e\xdc\x07\xac\xeb\x81\xfb\x40\x2d\x84\x74\x54\x0a\xd6\x1c\x50\x4f\x06\x21\x11\xcb\xce\x48\x8f\xd9\x84\x4b\x1a\x9c\xa5\xbd\x84\xe9\x4f\x0e\xae\x91\xe3\x1e\x45\x8e\x64\x43\xbc\x23\x0d\x07\xe4\xf3\x79\xb6\x7f\x91\x65\x48\x5a\xf2\xda\xe1\x62\x7e\x66\x2e\x2f\x8d\x11\x65\xff\xab\x16\x7b\xff\x7b\xad\x7f\x51\x92\xcf\x76\xbb\x45\x91\x1f\xe3\x9e\xdc\x67\xdd\x4a\xbd\xa1\xe2\x4f\x4f\x0a\x5f\xa7\x5b\x7c\x1b\xc9\xdf\x23\x43\x7b\xab\xc8\x47\x33\x49\xdd\x98\xe3\x77\xf6\x23\x00\x00\xff\xff\xad\x97\x07\x85\x98\x02\x00\x00")

func filesInitDVncbootBytes() ([]byte, error) {
	return bindataRead(
		_filesInitDVncboot,
		"files/init.d/vncboot",
	)
}

func filesInitDVncboot() (*asset, error) {
	bytes, err := filesInitDVncbootBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/init.d/vncboot", size: 664, mode: os.FileMode(420), modTime: time.Unix(1503557815, 0)}
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
	"files/conf/dhcpcd.conf": filesConfDhcpcdConf,
	"files/conf/wpa_supplicant.conf": filesConfWpa_supplicantConf,
	"files/init.d/autorun_date": filesInitDAutorun_date,
	"files/init.d/vncboot": filesInitDVncboot,
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
	"files": &bintree{nil, map[string]*bintree{
		"conf": &bintree{nil, map[string]*bintree{
			"dhcpcd.conf": &bintree{filesConfDhcpcdConf, map[string]*bintree{}},
			"wpa_supplicant.conf": &bintree{filesConfWpa_supplicantConf, map[string]*bintree{}},
		}},
		"init.d": &bintree{nil, map[string]*bintree{
			"autorun_date": &bintree{filesInitDAutorun_date, map[string]*bintree{}},
			"vncboot": &bintree{filesInitDVncboot, map[string]*bintree{}},
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

