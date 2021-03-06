// Code generated by go-bindata.
// sources:
// templates/init/rails/Dockerfile
// templates/init/rails/docker-compose.yml
// templates/init/ruby/Dockerfile
// templates/init/ruby/docker-compose.yml
// templates/init/sinatra/Dockerfile
// templates/init/sinatra/docker-compose.yml
// templates/templates.go
// DO NOT EDIT!

package templates

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

var _initRailsDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x50\xbb\x4e\xc4\x30\x10\xec\xfd\x15\x2b\xd1\x5f\x7a\x5a\x24\xa8\xe0\x50\x24\x0a\xba\xf3\x39\xeb\x60\xe2\xd8\x96\x1f\x27\xf2\xf7\x6c\xd6\x09\x21\xa7\xa4\xca\x3c\xec\x99\xf1\x73\x7b\x7e\x05\xe5\xdd\xcd\xff\x34\x51\x1a\x9b\x84\x78\x20\x1c\x26\xf0\xce\x4e\x90\xbf\x10\xb4\xb1\x98\xc0\x21\x76\xd8\x81\xf6\x11\xae\xc5\x75\x16\xc1\xb8\x94\xa5\xb5\xe4\x2f\x4e\xf9\x71\x44\x97\xd9\x7f\x43\xd7\xf9\xd8\x28\xa9\x08\x58\xe3\xc8\xa9\x61\xf2\x05\x2e\xcb\xc1\x20\xd5\x20\x7b\xbc\xcc\x64\x84\x1e\xc7\x24\x9e\xce\xef\x9f\xf0\x82\xe3\x9c\x05\xfc\x35\x32\x84\x66\x61\x76\xf2\xc9\x7a\x35\xec\x64\x66\xa8\x06\xbb\x76\xe9\xec\xfa\xcf\x88\xf6\xe3\xed\xbe\xff\x3a\xf8\xbb\xa4\x7c\x3c\x58\xa6\x84\x39\x3d\x86\x88\xb4\x33\xfc\x15\x6a\xe5\x80\x4b\x61\x0e\x5a\x71\x55\xe9\x51\xb5\xe9\xb7\x2d\x15\x57\x2d\x94\xab\x35\x6a\xd3\x2a\xae\xda\x8c\x6b\x60\xd5\x36\xcc\xe5\x23\x85\x1c\x14\x5a\x47\xcc\xfd\x23\xd2\x10\xaf\xf9\x9f\x4e\xd7\x6b\x4f\x7c\x9b\xf8\x0d\x00\x00\xff\xff\x08\xae\x24\x4e\xf0\x01\x00\x00")

func initRailsDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initRailsDockerfile,
		"init/rails/Dockerfile",
	)
}

func initRailsDockerfile() (*asset, error) {
	bytes, err := initRailsDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/rails/Dockerfile", size: 496, mode: os.FileMode(420), modTime: time.Unix(1463504162, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRailsDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x4f\x4d\xb2\xe2\x52\x50\x48\x2a\xcd\xcc\x49\xb1\x52\xd0\x03\x32\x73\x12\x93\x52\x73\x8a\x41\x82\x0a\x0a\xba\x0a\xc9\xf9\x79\x65\xf9\x15\x7a\x05\xf9\x45\x25\x7a\x26\x26\xc6\x7a\x05\x45\xf9\x25\xf9\xc9\xf9\x39\xb6\x25\x39\xc5\xb8\x95\x54\x54\xda\x96\x14\x95\xa6\x02\x15\x80\x44\xe1\x86\x59\x18\x58\x99\x18\x18\x18\x40\x79\x40\xb5\x20\xae\x21\x17\x20\x00\x00\xff\xff\xc0\xe1\x22\xef\x84\x00\x00\x00")

func initRailsDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initRailsDockerComposeYml,
		"init/rails/docker-compose.yml",
	)
}

func initRailsDockerComposeYml() (*asset, error) {
	bytes, err := initRailsDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/rails/docker-compose.yml", size: 132, mode: os.FileMode(420), modTime: time.Unix(1463504162, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRubyDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8f\xcd\x4e\xc0\x20\x10\x84\xef\x3c\xc5\x26\xde\xcb\x43\x98\xe8\x49\x6b\x9a\x78\xf0\x56\x0a\x4b\x6d\x0a\xbb\x84\x9f\x46\xde\x5e\x8a\x35\xb1\x72\x62\x67\xbf\x61\x86\xa7\x69\x7c\x01\xcd\x74\xf0\x97\x8c\x65\xa9\x42\x3c\xb4\x31\x54\x60\x72\x15\xf2\x27\x82\xdd\x1c\x26\x20\x44\x83\x06\x2c\x47\x58\x0a\x19\x87\xb0\x51\xca\xca\xb9\xc6\x17\xd2\xec\x3d\x52\xee\xfc\x81\x64\x38\x4a\xad\x74\x1b\xdc\x46\x8d\xb4\x50\xb9\xc0\x7c\x19\x83\xd2\xbb\x5a\x71\x3e\xc5\x08\x2b\xfa\x24\x1e\xc7\xb7\x0f\x78\x46\x7f\x66\x41\x3f\x52\x85\x20\x2f\xe5\xb6\x1e\x1c\xeb\xfd\xb6\xee\x4a\xab\xd1\xa9\x5b\x7a\xa7\xfe\x2a\x62\x7a\x7f\xfd\xdf\xff\xf7\xc3\x67\xf7\x88\x29\x03\xdb\x7e\x6f\xde\x9f\xe0\xa1\xbf\x23\xbe\x03\x00\x00\xff\xff\x8b\xae\xa0\xae\x2a\x01\x00\x00")

func initRubyDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initRubyDockerfile,
		"init/ruby/Dockerfile",
	)
}

func initRubyDockerfile() (*asset, error) {
	bytes, err := initRubyDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/ruby/Dockerfile", size: 298, mode: os.FileMode(420), modTime: time.Unix(1463504162, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRubyDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x8c\xc1\x0a\x83\x30\x0c\x86\xef\x3e\x45\xf0\x6e\xe9\x98\x87\x21\xf8\x30\x5a\x03\xca\xa2\x29\x69\x3a\xed\xdb\x2f\x85\x6d\xb7\xdd\xfa\xf5\xff\xf2\x9d\x38\x0f\x0d\xc0\x9c\x37\x5a\x06\x70\xf6\x0c\xbc\xef\xd3\x61\x80\x61\x65\x68\x71\xd9\x14\x16\x0e\x4f\x94\xce\xa6\xc8\x09\x5d\xd9\x09\xce\x4d\x57\x28\x9c\x05\x92\x4e\xa2\x39\x7e\x0f\x5b\x6b\xd0\x34\x23\xa5\x1a\x06\xe8\x6c\x38\x5e\x7c\xb9\xc8\xa2\xae\xef\xef\x2e\x0a\x2b\x07\xa6\x51\x29\xfd\x57\xae\x32\xaa\x64\x34\xa1\xfe\xfe\x62\x0f\x3f\xf4\xde\xfb\x0f\x99\x5b\xf1\xd6\xbc\x03\x00\x00\xff\xff\xae\x01\x4e\xf5\xc8\x00\x00\x00")

func initRubyDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initRubyDockerComposeYml,
		"init/ruby/docker-compose.yml",
	)
}

func initRubyDockerComposeYml() (*asset, error) {
	bytes, err := initRubyDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/ruby/docker-compose.yml", size: 200, mode: os.FileMode(420), modTime: time.Unix(1463504162, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initSinatraDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8f\xc1\x6e\x84\x20\x10\x86\xef\x3c\xc5\x24\xbd\xcb\x43\x34\x69\x4f\xad\x8d\x49\x0f\xbd\x49\x61\x70\x89\x30\x43\x00\xcd\xfa\xf6\x8b\xac\x9b\xac\xeb\x49\xbe\xf9\x86\xff\xe7\x63\xe8\xbf\x40\x33\xad\x7c\x95\xd9\x91\x2a\x49\x09\xf1\x56\x49\xdc\x80\xc9\x6f\x50\x2e\x08\xd6\x79\xcc\x40\x88\x06\x0d\x58\x4e\xf0\xbf\x90\xf1\x08\x8e\x72\x51\xde\x57\x7f\x21\xcd\x21\x20\x95\xe6\xaf\x48\x86\x93\xd4\x4a\xd7\x83\x77\x54\x4d\x0b\x1b\x2f\x30\x1e\x8b\x51\xe9\x59\x4d\x38\xee\x30\xc1\x84\x21\x8b\xf7\xfe\xe7\x0f\x3e\x31\xec\x59\xd0\x3e\xa9\x62\x94\x07\x39\x8d\x3b\xcf\x7a\x3e\x8d\x1b\xa9\x35\x9a\x75\x4a\x6f\xd6\x33\x11\xc3\xef\xf7\x6b\xff\xc7\x83\xf7\xee\x09\x73\x01\xb6\xed\xbf\xee\xde\x83\xbb\x76\x8f\xb8\x05\x00\x00\xff\xff\x9c\x51\x49\xbe\x2d\x01\x00\x00")

func initSinatraDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initSinatraDockerfile,
		"init/sinatra/Dockerfile",
	)
}

func initSinatraDockerfile() (*asset, error) {
	bytes, err := initSinatraDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/sinatra/Dockerfile", size: 301, mode: os.FileMode(420), modTime: time.Unix(1463504162, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initSinatraDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x4f\x4d\xb2\xe2\x52\x50\x48\x2a\xcd\xcc\x49\xb1\x52\xd0\x03\x32\x73\x12\x93\x52\x73\x8a\x41\x82\x0a\x0a\xba\x0a\xc9\xf9\x79\x65\xf9\x15\x7a\x05\xf9\x45\x25\x7a\x26\x26\xc6\x7a\x05\x45\xf9\x25\xf9\xc9\xf9\x39\xb6\x25\x39\xc5\xb8\x95\x54\x54\xda\x96\x14\x95\xa6\x02\x15\x80\x44\xe1\x86\x59\x18\x58\x99\x18\x18\x18\x40\x79\x40\xb5\x20\xae\x21\x17\x20\x00\x00\xff\xff\xc0\xe1\x22\xef\x84\x00\x00\x00")

func initSinatraDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initSinatraDockerComposeYml,
		"init/sinatra/docker-compose.yml",
	)
}

func initSinatraDockerComposeYml() (*asset, error) {
	bytes, err := initSinatraDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/sinatra/docker-compose.yml", size: 132, mode: os.FileMode(420), modTime: time.Unix(1463504162, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x99\x5b\x6f\x1b\xc9\xd1\x86\xaf\xc5\x5f\x31\x2b\x60\x17\xe4\x07\x2d\x35\xe7\x83\x00\xdf\xac\xed\x0f\xf0\x45\xbc\xc0\xae\x73\x11\x44\xc1\x62\x0e\xdd\xca\x60\x29\x52\x21\xa9\x8d\x64\xc3\xff\x3d\xf5\x74\xf5\x90\x94\x44\x4a\x8e\xac\x60\x17\x41\x0c\x8c\x39\xd3\x87\xea\xaa\xee\xaa\xb7\xde\x6a\x9d\x9e\x06\xaf\x17\x9d\x09\x2e\xcc\xdc\x2c\xeb\xb5\xe9\x82\xe6\x36\xb8\x58\x7c\xdf\xf4\xf3\xae\x5e\xd7\xd3\x91\x0c\x58\x2d\xae\x97\xad\x59\x9d\xf1\xbe\x36\x97\x57\x33\x19\xb7\x3a\xed\xe7\xfd\xfa\x74\x59\xf7\xb3\xd5\xe9\x9b\x45\xfb\xab\x59\xda\x7e\x66\x0e\x0e\xe9\xdc\x90\xef\xdb\xc5\xe5\xd5\x62\x65\xa6\xb7\x97\xb3\x7d\x43\xaf\x9b\xdb\x27\x84\x31\xe2\x8b\x64\xad\xfa\x79\xbd\x5e\xd6\x8f\x8b\x1b\x06\x3d\x25\x71\xf3\x36\xbd\x58\xd0\xf3\xe6\xc7\xe0\xfd\x8f\x1f\x82\xb7\x6f\xde\x7d\xf8\x66\x34\xba\xaa\xdb\x5f\xeb\x0b\xb3\x1d\x3f\x1a\xf5\x22\x68\xb9\x0e\xc6\xa3\xa3\xe3\xe6\x56\x5a\x8e\xe5\x05\xe9\x4b\xb3\x5a\x9d\x5e\x7c\xec\xaf\x68\xb0\x97\x6b\x7e\xfa\x85\xfe\x7f\xda\x2f\xae\xd7\xfd\x8c\x8f\x85\x9b\x70\x55\xaf\xff\x7e\x8a\xe6\xbc\xd0\xb0\x5a\x2f\xfb\xf9\x85\xeb\x5b\xf7\x97\xe6\x78\x34\x19\x8d\xec\xf5\xbc\x0d\xfc\x69\xfd\x64\xea\x6e\xcc\x4b\xf0\xd7\xbf\xb1\xec\x49\x30\xaf\x2f\x4d\xa0\xd3\x26\xc1\x78\x68\x35\xcb\xe5\x62\x39\x09\x3e\x8d\x8e\x2e\x3e\xba\xaf\xe0\xec\x55\x80\x56\xd3\xf7\xe6\x9f\x08\x31\xcb\xb1\x53\x9b\xef\x1f\xae\xad\x95\x6f\xc4\x4e\x26\xa3\xa3\xde\xba\x09\xdf\xbc\x0a\xe6\xfd\x0c\x11\x47\x4b\xb3\xbe\x5e\xce\xf9\x3c\x09\xc4\xa4\xe9\x5b\xa4\xdb\xf1\x31\x82\x82\x6f\xff\x71\x16\x7c\xfb\xdb\xb1\x6a\xe2\xd6\x12\x19\x9f\x47\xa3\xa3\xdf\xea\x65\xd0\x5c\xdb\x40\xd7\xd1\x45\x46\x47\xbf\xa8\x3a\xaf\x82\x7e\x31\x7d\xbd\xb8\xba\x1d\x7f\x27\x63\x4e\x44\x37\x99\xd5\xce\xde\x0e\x9a\x4e\x5f\xcf\xe4\x9c\xc6\x62\xfe\x0b\xe9\x83\x18\x95\x7f\x40\x90\x0c\x54\xbd\x7d\xa3\xa8\x35\xfd\x01\xd5\xc7\x93\x13\x46\x8c\xa4\x6f\x7d\x7b\x65\x82\x7a\xb5\x32\x6b\xb6\xfc\xba\x5d\x23\xc5\xd9\xe7\xcf\x43\x96\x99\xdb\x45\x10\x2c\x56\xd3\xff\x97\x63\x7d\x27\x1f\x9b\x79\xfe\x08\x87\xf6\x1d\x09\xee\x0c\xe5\x9f\x1e\xe3\xe8\x68\xd5\x7f\x74\xdf\xfd\x7c\x9d\xa7\xa3\xa3\x4b\xc2\x37\xd8\x08\xfd\x93\x7c\xba\xc6\x0f\xe2\x21\x01\x6e\x32\xe5\x8d\x75\x9c\xab\x8c\x6d\x7f\x7f\xad\x49\xf0\x5e\x96\x18\x4f\xfc\x0a\xac\xe9\xad\xb4\xfd\x94\xd5\x65\xf2\xe1\xb9\x3f\x8b\x3a\x32\xd7\x69\x73\x77\x2a\x8a\x3e\x3a\x15\x5d\x65\xea\x8e\xe6\x77\x05\x60\xda\x53\x02\x30\x4e\x64\x6c\x0c\x7d\x20\xc1\x5b\x7f\x58\xc8\xbb\xd5\x9b\x7e\x29\x22\x9a\xc5\x62\xb6\x3b\xbb\x9e\xad\x9e\xb0\xfc\x76\xa5\x86\x0b\xbe\xd4\xad\xf9\xf4\x79\x67\xb6\x77\x09\xbc\xfc\x17\xa0\xe6\x27\x50\x70\x8b\x45\xe2\xe1\xea\x12\xe3\xe3\xf3\x9b\xc8\x9e\xdf\x94\xcd\xf9\x4d\x58\xca\x13\xfa\xa7\x3a\xbf\xc9\x8d\xb4\xfb\x36\x2b\x63\xf2\xf6\xfc\x26\x93\xf7\x46\xc6\xa6\xd2\xd7\xa6\xe7\x37\x89\x7c\x47\xf2\x18\xe9\xb3\x9d\xbc\x67\xe7\x37\xb1\xf4\x77\x91\x8c\x95\x39\x45\x2d\xbf\xf2\xc4\x32\xb6\x16\x59\x26\x54\x19\x7c\x87\xd2\xde\xc8\x63\x13\x91\x23\xeb\x19\x99\x97\x23\x2b\x96\xf9\x32\xb6\xca\x55\xb7\xb8\x90\x31\xd2\x66\x0b\xd5\xa1\xcb\x55\xbf\x58\xd6\xa8\xa5\xad\x16\x59\xad\xc8\x49\x5a\xd5\xa3\x92\x3e\x2b\x7d\x85\xc8\x2d\x44\x66\x21\xba\x86\xa2\x97\x91\xa7\x13\x1d\xdb\x4e\xed\x49\x64\x5e\x26\xe3\x22\x99\x5b\xc9\xb8\x52\xbe\x0b\x59\x37\xc6\x26\x91\x13\xcb\x3a\x56\xde\x5b\xa3\xf6\x56\xd8\x6e\xd5\xde\x46\xc6\x36\x32\xb7\x92\xf1\x6d\xa8\xba\x14\xb9\xea\x5d\xca\xbb\x45\x77\xf4\x63\x9f\x64\xdd\x42\x9e\x48\xda\x5a\x69\x6b\xb0\x8d\xfd\x90\xb6\x06\xbd\xe4\x3d\xb6\xba\x86\x15\xdd\x0b\x19\x93\x4a\x5b\x25\xb6\x75\x7c\xd3\x27\xb6\x74\x85\xf6\xb3\x46\x2c\x4f\x5d\xe9\x99\x65\xec\xab\xf4\xb7\xbe\x2d\x8f\x74\xbf\xb0\x39\x66\xfd\x46\xd7\xc5\xae\x2e\xd3\x5f\xf6\xa5\x11\x1b\x5b\x79\xf2\x54\x6d\x8f\x18\x5b\xe8\xd9\x54\x46\xed\x36\xb2\x76\x65\x75\x1f\xca\x58\xd7\xa9\x5a\x95\x6d\xe5\x37\x91\xdf\x24\xd6\xf9\x79\xae\x6b\xb3\x0f\xac\xdf\x56\x7a\xfe\x8c\x31\x7e\x1d\xc6\xe0\x07\x61\xae\x7e\xc4\x58\x53\x79\xff\x91\xb1\xad\x55\x9f\x63\xff\x58\xcb\xc8\x59\x35\xc6\x9f\x17\xbe\x52\xea\x3c\xce\xbc\x68\xf5\xcc\xb1\xbf\xce\x55\x07\xfc\x28\x91\x39\x65\xae\x72\x38\xa7\x24\x51\x5d\xf1\xcd\xbc\x56\x3f\x28\xc5\x9e\xb4\x51\x7d\x43\xa3\x3e\xca\xbe\x67\x99\xea\x83\x5f\x34\xfe\xbd\x61\x4f\xbc\x6f\x67\xfe\xdd\x9d\x5f\xa3\xb6\x21\xb3\x93\x35\xe2\x5a\xcf\xbb\x93\xdf\x12\x7b\xbd\x3f\xb3\xe7\x49\xae\x7b\xcd\xda\xb1\x8c\x2d\x33\xf5\xb1\x28\xd5\xb5\xd3\x42\xfb\x89\x21\xfa\xbb\x58\xfd\xac\xb6\x7a\xe6\x95\xf7\x0f\x7c\x00\xb9\xa9\x55\xfb\xf1\x5b\xf6\x24\xec\xb6\xb1\xcb\x5e\xf1\xe0\x1b\xe8\xc3\x79\x3a\xdf\xa2\x3f\x1a\xc6\x1d\x0f\xc9\x7b\x0f\x3c\xf8\xcc\xb2\x2f\x63\x0f\xf9\x67\x27\xe3\x4b\xaa\xda\x87\x31\x27\xd2\x7e\xbc\x97\xa5\x1d\x4b\xd7\x64\x93\x11\xf6\x4c\x65\xe9\xff\x73\x99\x6c\x77\x69\x97\xca\x36\x7c\xe1\xb0\xda\x4f\x25\xe5\x4d\x2e\x75\xd9\x50\x44\xdd\x43\xd6\x4f\xe4\x9c\xb3\xe0\x90\xee\x01\x79\xe5\x2c\x48\xab\xfc\x24\x20\x43\x9c\xed\x26\x90\x71\x1a\x87\x13\xd7\x0e\xee\x9f\x69\x5e\xf8\xf3\xbc\xbf\x19\x47\x69\x9e\x64\x61\x1a\xe5\xf1\x49\x10\x4e\x24\xe5\xd7\x2c\xfd\x9d\xb3\xf2\x93\x33\xed\x2c\xf0\x16\xa2\xd7\x99\xfb\xff\xf3\x66\xc3\xeb\x93\xc7\x30\xfd\xb5\x32\xc7\xbf\x5c\xce\x9e\x85\xec\x78\x2e\x1e\x95\x12\x6d\xb1\x22\x70\x16\x2b\x52\xa7\xa5\xf6\x83\x9a\x78\x68\x5a\x29\xf2\xd1\xdf\x21\x2f\x51\x04\x00\x6d\x23\xf9\xad\x12\xed\xe3\xbb\x44\x6e\xa4\xc8\x01\xda\x0f\x88\xcf\x2f\xe8\xe0\xd0\x0e\x7d\x32\x7d\x27\xc2\x40\x8c\xd0\x7f\xa7\x44\x8b\x6f\x03\x8d\x79\xda\x7c\x3b\x26\xf5\xe3\x62\xff\x3b\xc8\x04\x05\x9a\x5c\xdb\x79\x07\x81\x1d\xea\x12\xc1\xa9\x3e\x44\xa9\xcb\x30\xa9\xb6\x83\x20\x61\xac\x3a\x38\x74\x90\x76\x13\x29\x92\x64\xe8\xe6\x91\xb6\xf2\xef\xc3\x93\x86\x6a\x03\xbf\x8d\x47\x57\x17\x75\x64\x96\x42\xbf\xef\x47\x26\x19\x03\xd9\x71\xac\x08\x0b\x72\x6d\xce\xe7\x60\x64\x6e\x0f\xf9\xab\xe3\x73\x2b\xea\x7e\x94\x3e\x2c\x45\x1e\x8d\xd6\xad\xa0\x67\xc4\xec\x03\x83\xfe\x03\x91\xbb\xcf\x1e\x1f\xc1\x51\x12\xff\xde\x11\x2c\xe5\xe4\x57\x92\xb2\x4c\x42\xb2\xb4\x1a\x9e\x8e\x94\x85\x9e\xc0\x84\xea\x56\xb8\x17\x49\x92\x10\x20\x7c\x3a\x4f\x08\x20\x14\x90\x17\x53\x6a\x48\x93\x54\xaa\x5a\x09\x10\x89\x22\xcb\x35\x4c\x5d\x92\xec\x7c\xe8\x7a\x92\xe4\x92\x91\x97\x95\x19\x0d\x73\x12\x21\xb0\x50\x88\x4b\xe7\x3c\x85\x92\xa5\xdc\x87\x10\x44\x2d\xaf\x34\x69\x91\x84\xd0\xb7\x48\x75\x2d\x88\x8e\x4b\x98\x99\x12\x98\x34\x56\x9d\x5d\xf2\x8e\x34\x5c\x49\xa4\x85\x0f\x4f\xc8\x05\xa4\x10\x58\x81\xd0\x0d\x64\x2d\xf6\x61\x5b\x26\x4a\x30\xe2\x56\x13\x2b\x61\x8b\xfe\x91\xc8\x2e\xd1\x2b\xd4\x04\x0f\x69\xe4\x69\x2a\x85\x15\xc2\x95\xc4\x0b\x19\x81\x44\x00\x63\xc6\x28\x79\x80\xd4\x41\x60\x12\xf6\x8b\x44\xef\xf5\x8f\x20\xbb\xad\xee\x47\xd9\xa9\xce\x8e\x24\x57\x7a\x16\xd8\xeb\xd6\xf5\x49\x9d\x3d\x1c\xc8\x46\xe2\xc9\x21\x67\x0c\x69\xb6\x9e\x28\x0f\x44\x0c\xe2\x11\x5a\x3d\x13\xe4\x43\x04\x21\x50\x40\x6a\x62\x55\xbf\xd2\x43\x0e\x04\x12\x72\x01\xec\x31\x9f\x3d\x81\x50\x43\x28\x68\xc3\x0e\xa7\xb7\x27\xe7\xec\x73\xd6\x28\x94\x72\x36\xd6\x28\xcc\x73\x76\xb4\xb1\x16\x73\x3b\x0f\x59\x90\xef\x36\xd1\x73\xe5\x1d\x1f\x8c\x2b\x85\xff\xce\x93\xec\xdc\xaa\x4f\xe0\x1f\x10\xfd\x3a\x52\x1f\x80\xd0\x40\xe2\x18\x7b\x1f\x0a\xf1\x71\xe0\xb2\x1e\x60\xb3\x3e\x4c\x52\xee\x44\xcb\x33\x31\xf0\x8e\x8c\x1d\xf0\xbb\x7b\xf5\xb3\x07\xf3\xee\x4c\xfc\x62\xb0\xdb\xa7\xf2\xcb\xa2\xdc\x7d\xc5\x3d\xb8\xc5\x55\xf9\x87\x01\xb7\xaf\x64\x27\x84\x19\xf0\x40\x0d\x43\x18\x13\x4a\xd4\x9d\x61\xab\xd0\xe2\x20\xce\x78\x36\x10\xea\x7c\x78\x3b\xf0\x46\xb8\x93\x89\xe1\xc8\xcc\x21\xf4\x70\x43\xc2\xbe\x8e\xd5\x85\x81\x25\x6a\x0b\x6a\x0d\x5c\x99\x50\x22\xac\x80\x3d\x42\x10\x88\x21\x34\x6d\xe6\xdd\x16\x96\xd3\x29\x14\x10\x9e\x70\x6f\x42\x9d\xba\x28\x29\xfc\x1a\x84\x6b\xa8\x75\x0c\x7a\x52\x6b\xa1\x27\x75\x82\x83\xc4\x50\x7f\x81\xbc\xbc\x54\x38\xa0\xd6\x83\x8d\x50\x27\x52\x93\xc0\xc8\xa8\x33\xa8\xc3\x60\x27\xd4\x77\xec\x4f\xd6\xe9\x58\xde\x5d\x6d\xda\xf9\xba\xa4\xdc\xd6\x66\x55\xac\xe9\x00\x1b\x61\x40\xae\x06\xb6\x1a\xf2\x40\x3d\xcc\x8d\xfa\x97\xd0\xa4\x0e\x8d\xbc\xce\xa4\x03\x6a\x6c\x6c\x03\xda\x81\x2f\x07\x65\xa5\xca\xaa\x7d\x4d\xc8\x43\x6d\xc9\x79\x00\x5d\x61\xa1\x3a\x02\xab\xec\xa9\xf5\x3a\x31\x1e\x76\x58\xd7\x1e\x3e\x53\x85\x05\xe0\x86\x07\xb8\x41\x2f\xe0\xcc\xa6\x0a\x1f\xb6\xd1\xb6\xca\x43\x14\xb5\x3c\x35\x3f\x7b\xb8\x0f\x42\x58\x03\xc8\x70\x35\x4e\xe6\xf7\xe9\x00\x9b\xda\xe3\x93\x5f\x0b\x24\x7b\xb9\xd4\xfe\x7b\xe2\xc7\x60\xe5\x19\x4c\xea\xb0\x31\x2f\x0f\x31\x8f\xf0\xa8\x38\x0c\x7f\x67\xa8\xf9\x59\xef\xd1\x5f\x8a\x4a\x45\xbe\x2f\xdd\xa1\x52\xf9\x3d\x2a\xc5\xfd\x4f\xb7\xa5\x52\xf8\x35\x38\x42\xcc\x42\x0d\xa0\x02\xd0\x2a\xfc\xb8\xf1\xef\xee\x6e\x04\x59\x95\x62\x51\xea\x7d\xb9\x1d\xf0\x25\xf7\xe9\xd0\xd7\xf5\xb5\x4f\xe1\xae\xe2\xf2\x77\x43\xe8\x81\xae\x06\x3a\x95\x68\xbc\x92\x62\xa9\x3e\xb8\xe7\x70\xb4\xa4\xd5\x8a\x06\x8c\xa8\x22\x5f\xdd\x79\xbc\x20\x96\xa0\x76\x7c\x43\x5b\xc0\x21\xaa\x26\x52\xb7\xbb\x97\x33\xba\x4f\xc4\x1b\x55\x8d\xab\xf8\x42\xbf\x4f\xb9\x62\x1d\x74\x6a\xb8\x4f\x60\x6d\xee\xc5\x90\xcb\xbc\xd2\x28\x4d\x03\x07\x1c\x3d\x2c\x94\x4a\xc4\x9e\xf2\xc5\xfe\x1e\x0a\xfd\x4c\xae\xf7\x1b\x54\x98\xd8\x45\xe5\xd8\xa5\x4a\x53\xb8\xe7\x00\xe7\x8b\x42\xe9\x1f\x18\x17\x36\x4a\xb7\xc0\x67\x77\xef\x68\x74\xbf\x58\xcb\xdd\xd9\x54\x3a\x0e\xec\x82\xee\x30\x86\xb3\x74\xe7\x18\xe9\x58\xeb\xa9\x8e\xf1\xd4\x8a\xbc\x01\xc5\x84\xf2\x31\xdf\xe1\xa2\xf1\xf7\x68\xb1\xe2\x2f\xd8\x06\x86\x72\xa6\x49\xa3\x77\x54\xee\x5e\xcb\x68\x1b\x3a\x31\x1e\x0a\xec\x28\x74\xa6\x36\x38\x3a\x94\xeb\xb9\x70\x17\x08\x7d\x32\x9e\x4a\x81\xc1\xae\x12\xf4\x76\xd2\x86\xce\x9c\x11\x95\x33\x78\xe6\x68\x5b\xa7\x7b\x4c\x1f\xfb\xc9\xd8\xe1\xde\x0c\x5f\xa5\xaa\x05\xef\xef\x63\x22\x79\x80\x7d\x19\x7c\x87\x3b\xab\x03\xb4\xea\x41\xf0\x3c\x0f\x10\x1f\x88\xd9\xa2\xe1\xc3\xbf\x85\x3d\x04\xc2\x07\xd3\xbf\x14\x05\x0f\xa9\xff\xa2\x10\xb8\xcf\x02\x8f\x7e\x49\x18\xfd\x91\xd0\xef\x7f\x37\x41\xff\xad\x37\x41\x07\x8e\xf9\x05\xa2\x75\x1f\x83\x39\xfc\xa7\xe9\x27\x62\xf7\xdf\xe3\x31\xff\x0a\x00\x00\xff\xff\x07\x7a\xfa\xf1\x00\x20\x00\x00")

func templatesGoBytes() ([]byte, error) {
	return bindataRead(
		_templatesGo,
		"templates.go",
	)
}

func templatesGo() (*asset, error) {
	bytes, err := templatesGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates.go", size: 16384, mode: os.FileMode(420), modTime: time.Unix(1463504165, 0)}
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
	"init/rails/Dockerfile": initRailsDockerfile,
	"init/rails/docker-compose.yml": initRailsDockerComposeYml,
	"init/ruby/Dockerfile": initRubyDockerfile,
	"init/ruby/docker-compose.yml": initRubyDockerComposeYml,
	"init/sinatra/Dockerfile": initSinatraDockerfile,
	"init/sinatra/docker-compose.yml": initSinatraDockerComposeYml,
	"templates.go": templatesGo,
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
	"init": &bintree{nil, map[string]*bintree{
		"rails": &bintree{nil, map[string]*bintree{
			"Dockerfile": &bintree{initRailsDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initRailsDockerComposeYml, map[string]*bintree{}},
		}},
		"ruby": &bintree{nil, map[string]*bintree{
			"Dockerfile": &bintree{initRubyDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initRubyDockerComposeYml, map[string]*bintree{}},
		}},
		"sinatra": &bintree{nil, map[string]*bintree{
			"Dockerfile": &bintree{initSinatraDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initSinatraDockerComposeYml, map[string]*bintree{}},
		}},
	}},
	"templates.go": &bintree{templatesGo, map[string]*bintree{}},
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

