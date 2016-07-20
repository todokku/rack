// Code generated by go-bindata.
// sources:
// templates/init/django/.dockerignore
// templates/init/django/Dockerfile
// templates/init/django/docker-compose.yml
// templates/init/rails/.dockerignore
// templates/init/rails/Dockerfile
// templates/init/rails/docker-compose.yml
// templates/init/ruby/.dockerignore
// templates/init/ruby/Dockerfile
// templates/init/ruby/docker-compose.yml
// templates/init/sinatra/.dockerignore
// templates/init/sinatra/Dockerfile
// templates/init/sinatra/docker-compose.yml
// templates/init/unknown/.dockerignore
// templates/init/unknown/Dockerfile
// templates/init/unknown/docker-compose.yml
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

var _initDjangoDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\x4b\xce\xcf\x2b\xcb\xaf\xe0\xd2\x4b\xcd\x2b\xe3\xd2\x4b\xcf\x2c\x01\x13\x99\xe9\x79\xf9\x45\xa9\x5c\x80\x00\x00\x00\xff\xff\x57\x31\x5f\xce\x1d\x00\x00\x00")

func initDjangoDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initDjangoDockerignore,
		"init/django/.dockerignore",
	)
}

func initDjangoDockerignore() (*asset, error) {
	bytes, err := initDjangoDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/django/.dockerignore", size: 29, mode: os.FileMode(420), modTime: time.Unix(1469004867, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initDjangoDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x91\xcd\xae\xd3\x30\x10\x85\xf7\x79\x8a\x91\x60\xdb\x66\xd1\x27\x40\x25\x2c\x40\xb4\x51\x28\x48\x5d\x21\xe3\x4c\x52\x17\xc7\x63\xfc\x03\x8d\x50\xdf\x9d\xb1\xd3\xd0\xe6\xde\xbb\xb8\xbb\xcc\xf1\xcc\x99\x6f\x4e\x3e\x34\xfb\xcf\x20\xc9\xfc\xa6\x4b\xd9\x9e\x85\xe9\xa9\x28\xde\x80\x43\xab\x85\x44\xc0\x8b\x18\xac\x46\x61\x2d\x08\xd3\xce\xa5\x75\x74\x46\x19\x20\x10\x04\xa1\x34\x39\x08\x27\x04\x35\x88\x1e\x93\x36\x52\x74\x70\xeb\x61\xaf\xba\xd9\x7f\xac\xb6\x07\x50\x1e\x84\xf6\x04\xd1\x63\x0b\x3f\x46\xe8\xa3\x51\x92\x9c\x01\x65\xf2\xfc\x02\x02\xde\x93\xfc\x89\xae\x53\x1a\x8b\x6a\xf7\x0d\xde\xd5\xf5\x03\x4c\x96\x66\xdf\x25\x54\xa2\x17\x06\x70\xb0\x61\x84\x2f\xd5\xb6\xa9\x0e\xdf\x3f\x55\x47\x68\xa3\x53\xa6\x87\x41\x18\xa6\x5c\xdb\x91\xd7\x0d\x5c\xb4\x1e\xfe\x28\xad\xf9\x60\x1f\x75\x48\x28\x69\xd8\x39\x72\x79\xc7\x83\x41\x47\x39\x19\x49\x3c\x4b\x46\x8f\x99\x39\xf1\x79\x30\x88\x2d\xdf\xd4\x71\x10\x56\x59\x36\xf1\x41\x68\x5d\x6c\xf7\xf5\x91\x8d\x7f\x45\xe5\x70\x40\x13\xfc\x3a\x5c\x02\x94\xcc\x5f\x3e\x55\x8b\xe6\xeb\x2e\xcd\x6e\xe6\x61\x58\xad\xa2\xed\x9d\x68\x31\xc9\x2f\x3c\xbb\x67\xce\xaf\xa0\x93\xa4\x35\x67\xc4\x16\x41\xc9\x89\xef\xed\x5f\x8e\xf6\x5a\x4e\xd2\x04\xb7\x90\xe6\xae\x5b\xda\xd7\xb9\x65\xae\xa7\xf7\x7b\xac\xf9\xf9\x7f\x39\x81\x8f\xe1\x44\x66\xb3\xc8\xfe\x81\x83\x4f\x35\xa4\x8c\x8d\xf7\x0b\x12\x3c\xff\x91\x00\xd4\xe5\xef\xf4\xcb\xf3\x9e\x75\xf6\x2f\xfe\x05\x00\x00\xff\xff\x22\xcb\xe6\x65\xb5\x02\x00\x00")

func initDjangoDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initDjangoDockerfile,
		"init/django/Dockerfile",
	)
}

func initDjangoDockerfile() (*asset, error) {
	bytes, err := initDjangoDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/django/Dockerfile", size: 693, mode: os.FileMode(420), modTime: time.Unix(1469004867, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initDjangoDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x8d\xc1\xaa\xc2\x30\x10\x45\xf7\xfd\x8a\xfc\xc0\xcb\x8b\x36\x82\x04\xba\x92\xae\xdc\xa9\x1b\x57\x92\xb4\x43\x09\xa6\x99\x92\x4c\x6b\xfd\x7b\x13\x68\xbb\x10\xdc\xcd\xbd\xf7\x70\xe6\x05\x46\x15\x8c\x99\xd1\xba\x56\x31\x9e\x4e\xf0\x93\x0d\xe8\x7b\xf0\x94\x17\xc6\xfe\xd8\xb5\x3e\x5d\xea\xdb\xe3\x5c\xdf\x53\xe1\xb4\x01\x17\xd7\xa9\x41\x3f\xe1\xcc\x07\x0c\xc4\xa5\x2c\xf9\x10\x90\xb0\x41\x57\x91\x8b\xbf\x91\xf9\x5d\x51\x18\x21\xdb\xac\x7f\x6e\xb2\x56\x93\x36\x3a\xe6\x3e\xd3\x5b\x7f\x14\x4a\x0a\x21\x96\x94\x1c\x39\xee\x8a\x15\xcf\x98\xed\x75\x07\x6a\xf9\xf5\x3f\x60\xa4\x2e\x40\xfc\x16\x1d\x64\xb9\x2f\x3e\x01\x00\x00\xff\xff\x25\x21\x30\xfe\xf3\x00\x00\x00")

func initDjangoDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initDjangoDockerComposeYml,
		"init/django/docker-compose.yml",
	)
}

func initDjangoDockerComposeYml() (*asset, error) {
	bytes, err := initDjangoDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/django/docker-compose.yml", size: 243, mode: os.FileMode(420), modTime: time.Unix(1469004867, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRailsDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x4b\x2a\xcd\x4b\xc9\x49\xe5\xd2\xd7\x4b\xce\xcf\x2b\xcb\xaf\x00\x32\x52\xf3\xca\x80\x64\x7a\x66\x09\x97\x7e\x4a\x92\xbe\x96\x5e\x71\x61\x4e\x66\x49\xaa\x31\x2a\x4f\x37\x2b\xbf\xb4\x28\x2f\x31\x87\x4b\x3f\x27\x3f\x5d\x5f\x8b\x4b\xbf\x24\xb7\x80\x0b\x10\x00\x00\xff\xff\xa0\x04\x95\x56\x4e\x00\x00\x00")

func initRailsDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initRailsDockerignore,
		"init/rails/.dockerignore",
	)
}

func initRailsDockerignore() (*asset, error) {
	bytes, err := initRailsDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/rails/.dockerignore", size: 78, mode: os.FileMode(420), modTime: time.Unix(1464029683, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
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

	info := bindataFileInfo{name: "init/rails/Dockerfile", size: 496, mode: os.FileMode(420), modTime: time.Unix(1463624980, 0)}
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

	info := bindataFileInfo{name: "init/rails/docker-compose.yml", size: 132, mode: os.FileMode(420), modTime: time.Unix(1463624980, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRubyDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x4b\x2a\xcd\x4b\xc9\x49\xe5\xd2\xd7\x4b\xce\xcf\x2b\xcb\xaf\x00\x32\x52\xf3\xca\x80\x64\x7a\x66\x09\x17\x20\x00\x00\xff\xff\xc9\x68\x92\x70\x1e\x00\x00\x00")

func initRubyDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initRubyDockerignore,
		"init/ruby/.dockerignore",
	)
}

func initRubyDockerignore() (*asset, error) {
	bytes, err := initRubyDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/ruby/.dockerignore", size: 30, mode: os.FileMode(420), modTime: time.Unix(1464029683, 0)}
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

	info := bindataFileInfo{name: "init/ruby/Dockerfile", size: 298, mode: os.FileMode(420), modTime: time.Unix(1463624980, 0)}
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

	info := bindataFileInfo{name: "init/ruby/docker-compose.yml", size: 200, mode: os.FileMode(420), modTime: time.Unix(1463624980, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initSinatraDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x4b\x2a\xcd\x4b\xc9\x49\xe5\xd2\xd7\x4b\xce\xcf\x2b\xcb\xaf\x00\x32\x52\xf3\xca\x80\x64\x7a\x66\x09\x17\x20\x00\x00\xff\xff\xc9\x68\x92\x70\x1e\x00\x00\x00")

func initSinatraDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initSinatraDockerignore,
		"init/sinatra/.dockerignore",
	)
}

func initSinatraDockerignore() (*asset, error) {
	bytes, err := initSinatraDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/sinatra/.dockerignore", size: 30, mode: os.FileMode(420), modTime: time.Unix(1464029683, 0)}
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

	info := bindataFileInfo{name: "init/sinatra/Dockerfile", size: 301, mode: os.FileMode(420), modTime: time.Unix(1463624980, 0)}
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

	info := bindataFileInfo{name: "init/sinatra/docker-compose.yml", size: 132, mode: os.FileMode(420), modTime: time.Unix(1463624980, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initUnknownDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\x4b\xcd\x2b\xe3\xd2\x4b\xcf\x2c\xe1\x02\x04\x00\x00\xff\xff\x9c\x10\x28\x7b\x0a\x00\x00\x00")

func initUnknownDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initUnknownDockerignore,
		"init/unknown/.dockerignore",
	)
}

func initUnknownDockerignore() (*asset, error) {
	bytes, err := initUnknownDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/unknown/.dockerignore", size: 10, mode: os.FileMode(420), modTime: time.Unix(1464952501, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initUnknownDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0b\xf2\xf7\x55\x28\x4d\x2a\xcd\x2b\x29\xb5\x32\x34\xd3\x33\x30\xe1\xe2\x72\xf6\x0f\x88\x54\xd0\x53\xd0\x4f\x2c\x28\xe0\x02\x04\x00\x00\xff\xff\xf1\xa3\x65\xfc\x1f\x00\x00\x00")

func initUnknownDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initUnknownDockerfile,
		"init/unknown/Dockerfile",
	)
}

func initUnknownDockerfile() (*asset, error) {
	bytes, err := initUnknownDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/unknown/Dockerfile", size: 31, mode: os.FileMode(420), modTime: time.Unix(1464952501, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initUnknownDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xca\x4d\xcc\xcc\xb3\xe2\x52\x50\x48\x2a\xcd\xcc\x49\xb1\x52\xd0\xe3\x02\x04\x00\x00\xff\xff\xa6\xe1\xc1\x85\x11\x00\x00\x00")

func initUnknownDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initUnknownDockerComposeYml,
		"init/unknown/docker-compose.yml",
	)
}

func initUnknownDockerComposeYml() (*asset, error) {
	bytes, err := initUnknownDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/unknown/docker-compose.yml", size: 17, mode: os.FileMode(420), modTime: time.Unix(1464952501, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9b\x59\x6f\x23\xc7\x11\xc7\x9f\xc5\x4f\x31\x16\x60\x43\x0a\x64\x69\x6e\xce\x08\xf0\x8b\x77\x1d\xc0\x0f\xb1\x01\x1f\x0f\x41\x36\x30\xe6\xe8\x51\x18\x4b\xe4\x86\xa4\xec\x95\x17\xfe\xee\xe9\x5f\x55\x8d\xc8\x15\x87\x94\x44\x49\xf0\xe6\x58\x60\x56\x64\x4f\x77\x75\x55\x77\xd5\xbf\x8e\x6e\x9e\x9d\x05\xaf\x66\xad\x0b\x2e\xdc\xd4\xcd\xab\xa5\x6b\x83\xfa\x26\xb8\x98\x7d\x5e\x4f\xa6\x6d\xb5\xac\x4e\x47\xbe\xc3\x62\x76\x3d\x6f\xdc\xe2\x9c\xcf\x4b\x77\xf5\xf6\xd2\xf7\x5b\x9c\x4d\xa6\x93\xe5\x59\xfb\xcf\x6a\x7a\x31\x3b\x3b\x6d\x67\xcd\xcf\x6e\x3e\xb9\x98\xce\xe6\x6e\x7b\xb7\xd7\xd2\xab\x9b\x5c\xee\xe8\xa3\x94\x3e\x6f\x66\x57\x6f\x67\x0b\x77\x7a\x73\x75\x39\xd0\x77\x5e\x4d\x2e\x17\xf7\xce\xaa\xbd\x76\x4e\xaa\x5d\x1e\x36\xe7\x75\x7d\x73\xff\x94\x74\xda\x3d\x23\x3d\x1e\x34\xe1\x62\x32\xad\x96\xf3\xea\xde\x39\xfb\x7e\x3b\xa7\xed\x3b\x3d\x68\xe6\xeb\xe9\xcf\xd3\xd9\xaf\xd3\x7b\x67\xee\xfb\xed\x9c\xb9\xef\x74\xdf\xcc\xb7\x9f\x4e\x2f\x66\xbc\x79\xfd\x6d\xf0\xcd\xb7\x3f\x04\x5f\xbd\xfe\xfa\x87\x4f\x46\xa3\xb7\x55\xf3\x73\x75\xe1\x56\xfd\x47\xa3\x89\x27\x34\x5f\x06\x47\xa3\x83\xc3\xfa\xc6\xb7\x1c\xfa\x0f\x50\x9f\xbb\xc5\xe2\xec\xe2\xb7\xc9\x5b\x1a\xba\xab\x25\x7f\x26\x33\xfd\xff\x6c\x32\xbb\x5e\x4e\x2e\xf9\x32\x93\x01\x6f\xab\xe5\x3f\xce\xe0\x9c\x0f\x34\x2c\x96\xf3\xc9\xf4\x42\xde\x2d\x27\x57\xee\x70\x74\x3c\x1a\x75\xd7\xd3\x26\x30\x8b\xf8\xce\x55\xed\x11\x1f\x82\xbf\xfd\x9d\x69\x4f\x82\x69\x75\xe5\x02\x1d\x76\x1c\x1c\xf5\xad\x6e\x3e\x9f\xcd\x8f\x83\xf7\xa3\x83\x8b\xdf\xe4\x5b\x70\xfe\x45\x00\x57\xa7\xdf\xb8\x5f\x21\xe2\xe6\x47\xc2\x36\xdf\xbf\xbc\xee\x3a\xff\x1d\xb2\xc7\xc7\xa3\x83\x49\x27\x03\x3e\xf9\x22\x98\x4e\x2e\x21\x71\x30\x77\xcb\xeb\xf9\x94\xaf\x27\x81\x17\xe9\xf4\x2b\xa8\x77\x47\x87\x10\x0a\x3e\xfd\xd7\x79\xf0\xe9\x2f\x87\xca\x89\xcc\xe5\x69\xfc\x3e\x1a\x1d\xfc\x52\xcd\x83\xfa\xba\x0b\x74\x1e\x9d\x64\x74\xf0\x93\xb2\xf3\x45\x30\x99\x9d\xbe\x9a\xbd\xbd\x39\xfa\xcc\xf7\x39\xf1\xbc\xf9\x51\xcd\xe5\x57\x3d\xa7\xa7\xaf\x2e\xfd\x3e\x1d\x79\xf1\x9f\x89\x1f\xc8\x28\xfd\x2d\x84\x7c\x47\xe5\xdb\x1a\x3d\x5b\xa7\x5f\xc2\xfa\xd1\xf1\x09\x3d\x46\xfe\xdd\xf2\xe6\xad\x0b\xaa\xc5\xc2\x2d\x59\xf2\xeb\x66\x09\x15\x91\xcf\xf6\xc3\x4f\x33\xed\x66\x41\x30\x5b\x9c\xfe\xd9\x6f\xeb\xd7\xfe\xcb\xed\x38\xdb\xc2\xbe\x7d\x8d\x82\xec\xa1\xff\xa7\xdb\x38\x3a\x58\x4c\x7e\x93\xef\x93\xe9\x32\x4f\x47\x07\x57\x40\x64\x70\x4b\xf4\x2f\xfe\xab\x34\xfe\xe0\x35\x24\x40\x4d\x4e\xf9\xc4\x3c\xa2\x2a\x47\xdd\xe4\xee\x5c\xc7\xc1\x37\x7e\x8a\xa3\x63\x9b\x81\x39\x4d\xca\x6e\x72\xca\xec\x7e\xf0\xf6\xb1\xdf\x7b\x76\xfc\x58\xe1\xe6\xc3\xa1\x30\xba\x73\x28\xbc\xfa\xa1\x6b\x9c\x7f\x48\x00\xd1\xee\x23\x80\x70\x9e\xc6\xad\xa0\x1b\x14\x4c\xfa\xed\x44\xbe\x5e\xbc\x9e\xcc\x3d\x89\x7a\x36\xbb\x5c\x1f\x5d\x5d\x2e\xee\x91\xfc\x66\xa1\x82\x7b\x7c\xa9\x1a\xf7\xfe\xf7\xb5\xd1\xa6\x12\x68\xf9\x4f\x40\xcd\x6b\xf1\x20\xaf\xd7\x30\xcb\x2b\xb9\x6a\xc5\xd1\xe1\x9b\x77\x51\xf7\xe6\x5d\x51\xbf\x79\x17\x16\xfe\x09\xed\x29\xdf\xbc\xcb\x9d\x6f\xb7\xb6\xce\xf7\x69\xe3\x37\xef\x52\xdf\xaf\xf1\xed\x8d\xff\x1e\xf3\xd9\x3f\x95\xff\xec\xc2\xb5\xf7\xad\xbe\x73\xc9\x5a\x1b\xfd\x1b\x4f\x2b\xf2\xf3\xf9\xf6\xd2\xd3\x77\xfe\x19\xfb\xa7\xf3\x4f\x9a\x79\x3a\xfe\x6f\xe6\xfb\x14\xe1\x1a\x1f\x36\x37\x4f\x36\x7e\xf3\x2e\xf1\xe3\xb3\x4e\x79\x88\xda\xf5\x7e\x87\x3d\x1e\x0d\x4b\x6c\xf6\x32\x84\x43\xbd\x55\xad\xe1\x98\x37\xc0\x2d\x2b\x77\xe2\x5f\x1d\x6e\x75\xf1\x87\xfe\xf5\xf1\xad\xba\x0f\x53\x80\x89\x3f\x89\xa5\xae\x33\x21\xa6\x7a\x8b\x87\x3b\x65\xb8\x0f\x77\x6e\xe1\x42\x0c\xde\x53\xbb\xa3\x3c\xef\x31\xab\xf3\x60\x97\x14\x01\xe6\x73\x1e\xc4\xe5\x49\x80\x1d\x9c\xaf\x9b\xc9\x51\x1a\x87\xc7\xd2\x8e\x76\x9f\xab\xf6\xff\x38\x9d\xbc\x3b\x8a\xd2\xbc\x0c\xc3\xb4\xc8\xc7\x27\x41\x78\xec\x81\xad\x62\xf6\xcf\x44\xd6\xf7\x22\xe0\x79\x60\x72\xc2\xda\xb9\xfc\xff\xfb\xed\x06\x54\x27\x3b\x35\x17\x67\xb4\x97\xde\x16\x5e\xa7\xca\x48\xf5\xb2\xf2\xef\x5a\xaf\x7f\x89\x7f\x17\xf9\xa7\xf0\x7a\xd7\x8d\x55\x0f\x8b\x4a\xfb\xe5\xe8\xb2\xa7\x9b\xe7\xfe\xaf\xff\x1e\xfb\xf7\xa9\x6f\x8b\x33\xd5\x61\x3e\xd7\xa9\xd7\x43\xde\xf9\x79\x52\xff\x64\xe8\x7c\xa4\x3a\x9f\xfa\x3e\x99\xd7\xfb\xc8\x8f\x6b\xfc\x93\xfb\xb6\x0e\xdd\xf7\x7f\x0b\xdf\x2f\x83\xbe\xe7\xab\xf4\x9f\xeb\x48\xf9\x69\x7d\x9b\x63\x3e\xcf\x5f\xed\xe7\xae\x0b\xfd\xdb\xf8\x71\x5d\xa4\x7f\xb1\x99\xdc\x8f\x4b\x7d\x9f\x84\xc7\xf3\xd0\xf5\xb6\xe5\xc7\x37\xa5\xce\x53\xe5\x6a\x73\xad\xff\x5e\x3a\x95\x11\x5b\xc3\xbe\xe0\x97\x39\xb0\xb1\xd4\xcf\x5b\xd5\xfa\x3e\xf5\xb4\x9a\x50\xd7\x33\xf2\x7d\x2a\xe4\xf4\x74\x72\x64\x6c\x75\x8d\xe1\x13\xbb\xab\x7c\xff\x31\x4f\xaa\x7d\xa2\x52\xe7\x67\x3d\x43\xdf\x56\x45\xca\x5b\x52\xea\x38\xd6\x8f\xf6\x24\xd3\x7d\x89\x3c\x8d\x92\x3d\xc8\x75\x9d\xa0\x33\x46\xfe\x5a\xe7\x03\x4f\xea\x4a\xf9\x1f\x77\xca\x4b\xed\xfb\x86\x63\x5d\x3b\xc6\x17\xc8\x9e\x2b\x5d\xf6\x88\x35\x0e\xfd\xf8\xa4\x53\x9e\x1c\x32\x24\xba\x47\xa5\x9f\xa3\x34\xec\xc9\xd9\xef\xd8\xf6\x23\xd6\xa7\x35\x7e\x68\x2b\x4a\xd5\x91\xcc\x7f\x8f\x2a\x5d\x8f\xbc\x52\x1d\x09\x5b\xed\xdb\x42\x23\xd3\xfd\x64\xaf\xcb\xdc\x74\xa5\x53\x1d\xc9\x58\x03\xdb\xff\x30\x57\xd9\xea\x50\x65\x83\xef\xb8\x53\x1a\xc8\xc4\x9e\x84\x4e\xc7\xc2\x7b\xc6\x5e\xa0\x33\x3d\xff\x89\xee\x67\x81\x0e\x46\xb6\x37\xb9\xe2\x24\x3a\x8a\xbe\xb6\xc6\x1b\x6d\xe8\x25\xeb\xd3\x39\xdd\xeb\xaa\x55\x7c\x45\xa7\xb1\x17\xf6\x0d\x7d\xe5\x5d\xee\xdb\xdb\x42\xf7\x69\x1c\xab\x0d\xa0\xaf\x45\xa2\x73\xc1\x07\xef\xd8\xdf\xd4\x3f\x49\xa3\x7a\xc5\xfa\x16\x9d\xea\x23\xef\xd1\x4f\xc6\x62\x53\xec\x2f\xfa\x82\x3c\x2d\xfb\x1a\xa9\x5e\x64\xf0\x5c\xea\x9e\xd3\x1f\xfa\xb9\xd9\x4d\xde\xe8\xfa\xb2\xa6\xc8\x83\x8d\xb0\xef\xf8\x04\x97\xe9\xfa\x61\x73\x91\xed\x51\x52\xa9\xac\xec\x5d\x99\xaa\x6d\xe0\x13\xb0\x09\xd6\x8f\x3d\xc3\x96\xd0\xa7\xd8\xa9\xdd\x83\x09\xce\xf4\x39\xb3\x75\x61\x8f\x5c\xab\x76\x08\x2f\xf8\x16\x6c\x88\xfd\x41\x56\xec\x2f\x1f\x9b\xce\xa3\x87\xa1\xea\x49\x65\xba\x2c\xef\x58\xef\x5c\xe5\x61\x2c\xfa\xe3\x3a\xa5\x0b\x4f\x85\x53\x3d\xcd\x2a\xb5\x5b\xfc\x21\x3a\x5b\xfb\xb1\xa5\xd9\xbc\xe8\x1b\xf6\x5a\xe9\x5e\xd6\xa5\xea\x29\xed\xd5\xd8\xf0\xa9\x56\x3b\xe8\xcc\x5f\xb2\x3e\xac\x7d\x11\xe9\x5e\xb8\x48\x6d\x18\x3d\xac\xb1\xd3\x42\x75\x40\xde\xb3\x9f\x9d\xf2\x0c\xef\xe0\x21\x6b\x2c\x3a\x8d\xbd\xc7\x2a\x2f\x58\xc9\xfa\x83\x9b\xec\x1d\x6b\x8f\x2c\x5d\xaa\x7e\xbe\x4b\x14\x4f\xd0\x21\x64\x62\x9d\x98\x23\xcc\x36\x7d\x75\x1c\xeb\x18\x59\x73\x74\x3d\x33\x7b\xdb\xed\xab\xc1\xf8\xa7\x7b\x6a\xa8\x6c\xf8\xe9\xd5\xab\xdd\x4e\x9a\x1e\xfb\xb8\xe8\x35\xd6\x5f\xc2\x41\xaf\xb3\x6f\xde\x39\x2f\x93\x8f\xc8\x3d\xbf\xd2\xfc\xf5\xaf\x57\x97\x7b\x39\x69\x9c\x00\x4a\xd9\xe0\x00\xbc\x21\x34\xf1\xca\x49\xa7\xe6\xa4\xbb\x56\x9d\x34\x20\x80\xb3\x42\xc1\xa0\x0d\xa8\x14\xbd\x61\x55\x0a\xf8\xe2\xe8\x1b\x05\xd8\xa8\xd6\xe0\x91\x76\x00\x12\xc7\x07\x0f\x00\x29\x20\x46\x3b\x40\x9e\xd7\x3a\x07\xc6\x06\xd8\xe4\xe6\x84\xe1\x01\x5a\x00\x49\x6d\x86\x33\x36\xe3\x45\xf9\xc5\x09\x66\x16\x68\x94\xea\x90\xe0\x03\xd0\x03\x54\x30\x1a\x8c\xbf\x33\x00\x01\xac\x71\x50\xcc\x23\x6d\xa9\x05\x0b\xb9\x1a\x14\xa0\x8c\xc1\x08\xa0\xd1\xb7\x52\xb0\x27\xb8\x10\xe0\xef\xd4\x31\x60\xf4\xc8\x23\x41\xf5\x58\xc1\x03\x79\x01\xa3\xc4\x40\x01\x70\xc4\x71\x86\x8d\x82\x55\x65\x41\x0c\x20\x82\x5c\xa5\x39\x27\xc6\xc8\x1a\x45\xba\xa6\xb5\x81\x01\xfd\xe0\x21\x33\xe7\x43\x90\xd3\x1a\x18\x01\x42\xec\x63\x1d\xab\xac\xbd\x53\x07\xa0\x59\x9b\xc4\x82\x2d\xc0\x8d\xbe\x11\x6b\xef\xdf\x85\x95\xd2\xc1\x71\x8a\xec\x8d\x82\x99\x73\xba\xbf\xac\x25\x41\x4d\x59\x28\x90\x02\x38\xc8\x20\x8e\xb8\x54\x47\x81\x7c\x38\x35\x00\x0d\x90\xc7\x41\xa0\x17\x80\x31\x89\x41\x9e\x2a\x90\xc6\xe6\x18\xc2\x68\x00\xa4\x32\xe5\x03\x3d\x63\xdd\x01\xb8\x07\x24\x14\x2b\x4d\x7f\x3a\x54\xad\x68\x6d\x00\xd6\x66\x5d\x68\x37\x70\xad\x48\xed\x03\x5f\x1b\x42\xbd\x04\x88\x0d\x89\xd4\xa7\x1a\xe9\x1f\x0d\x66\xdf\x51\xf2\x7c\x96\x24\x19\x3b\x24\x50\x8a\x2b\xc5\x0c\x49\x80\x2d\x88\xc2\xa9\xae\xf7\x19\x4a\xa6\xa1\x95\xc4\x1a\x1c\xa3\x93\x4d\xa5\x01\x3b\x3a\x3d\xae\xd4\xc9\x33\x77\x09\x0e\x39\xb5\x29\xc1\x36\xa7\x36\x95\x39\x0d\xbc\x08\x76\xc0\x1d\xfa\x33\x37\xb8\x0a\x1e\xc1\x97\x04\x04\x63\x9d\x17\x5b\x07\x47\x08\x0a\xc5\x5e\x22\x0b\x38\x2d\xa0\x26\x80\x97\xe0\xb4\x0f\x6a\x6a\x7d\xc7\xb8\xd8\x82\x19\x49\xda\x0d\x2b\xef\xda\x59\x65\x89\x41\x99\x29\x4e\xc0\xd3\x16\x3b\xdb\xd8\x84\xfd\x4c\x6c\x83\xcc\xca\xba\x06\x4a\xe4\x9b\x76\xb5\x31\xfe\xa1\x26\xb5\x8d\xff\x67\xb5\xa6\x41\x11\xcc\x8e\xc6\xc5\x63\xcd\x28\x0d\xe3\x32\x2f\x92\x17\x30\xa3\xbd\x33\x76\x32\x02\x32\x3d\x9c\x2f\xca\x82\xe3\xea\x83\x01\xd7\xa8\x93\xc6\x49\xa0\xbc\xad\x55\x84\x30\x0c\xa2\x6c\x14\xb2\x2a\xb4\x22\x05\x0d\xbe\xe3\x70\x70\xc8\x18\x93\x38\x9a\x5a\x9d\x87\x8b\xd5\xd9\x62\x34\x91\x29\xba\x64\x87\x63\xe5\xa1\x35\x43\xc3\x49\x10\x89\xe3\x6c\x30\x46\xa2\x65\x67\xce\x5e\xb2\x9f\x44\x9d\xf9\xd8\xa2\x61\x31\xf2\x56\x8d\x1f\x79\x12\xab\x16\x90\x51\x96\xb5\x66\x81\xe3\xc2\xb2\x66\x4f\x27\xce\x35\x10\x68\x2c\x50\x10\xa7\xdc\xa9\xbc\x18\x26\x55\x02\x9c\x21\x41\x0f\xbc\x8c\x73\xe5\x9b\xe8\x9e\x48\x3c\x8a\x34\xe0\x69\x2c\x33\xc7\x09\x12\x48\x51\x45\x28\xcd\xe9\x13\x50\x38\xcb\x3e\x25\x18\x29\x15\x24\xc8\x1a\x00\x92\x3e\xdb\x26\x38\x02\x9c\x3a\xcb\x0e\x25\x8b\xb4\x4c\x9e\xcc\xc9\x59\x76\x41\x1b\x00\xc3\x7a\x21\xb3\x64\x5a\xb5\xce\x8b\x5c\x64\xb8\xfc\x65\x5d\xea\x46\x03\x11\x00\xac\xb0\x0c\x95\x4c\x93\xbd\x21\x38\x42\x6e\x32\x8d\xb2\xd3\x75\x20\x98\x63\x1e\xc9\xda\x33\xcb\xfa\x32\x05\xc5\xc2\xb2\x9c\xdc\xd6\x81\xf9\x01\x57\xf6\x9f\x3e\xce\xe6\xa1\x0f\x7a\x40\x06\x8e\x1e\xd1\x97\x2a\x88\xe8\x4f\xa5\xa0\x8b\xce\xb1\x7e\xcc\x45\xa0\x01\x80\xca\x7e\xa1\x2b\x16\xfc\xb1\xe7\xe3\x46\xf7\x3c\xb3\xec\xbb\xb0\xea\x46\x42\x40\x9a\x2b\x1d\xf6\x29\xb1\x0c\x15\xdd\x24\x58\x41\x0f\xa4\xd2\x62\x95\x0d\x32\x7d\x74\x94\x75\xcf\xac\x2a\x83\x5e\xd4\xf6\x19\x20\x8d\x4d\xb7\x33\xfb\x2c\xfb\x57\x5b\xc6\x59\x69\x40\x05\x80\x4b\x80\x5b\x69\xd0\x99\x9b\x3e\xb3\xe6\x04\x60\xac\x35\x73\xc7\x89\x56\x77\xd0\x31\x82\x2f\xc9\x42\xc7\x56\x55\x6a\xf5\x3d\xce\x08\x3d\xc3\xf1\xb0\xe7\xa5\xe9\x07\x3a\x00\x5d\x1c\x05\xf2\xa3\xb7\xac\x49\xd8\x6e\x02\x3c\xba\x01\x3f\xec\x67\x1f\xe8\xae\x02\xae\x6d\x00\xbf\x7f\xb2\x77\x87\xc8\x5d\x70\xdf\x95\xea\xdd\x19\xba\x07\xae\xbf\x54\xa2\xb7\xc9\xbb\x41\x7a\x5a\xe6\x8f\xc5\xf4\x24\x8f\xd3\xb2\x08\x5f\x00\xd3\x9f\x98\xe6\xf5\xa1\x47\xda\x6a\xda\xe0\x2c\xcc\x01\xa9\xa5\x2e\x65\x21\x13\x1a\x4a\xc8\x02\xf2\xf1\x9e\xfa\x28\xf5\x36\x10\x00\xb4\xa5\x66\x41\xba\x91\xd9\x77\xd2\x06\x52\x18\x49\x03\xab\x15\xe2\xf3\xb7\xe9\x6b\x26\x56\x57\x91\xd4\x27\x53\xc4\x08\xb3\xd5\xf9\x43\x6c\x6d\xa0\x31\x0f\xe9\x58\xdf\x27\xb5\x7e\xb1\xfd\xed\x69\x82\x02\xd4\xf1\x68\xe7\x33\x08\x2c\xa8\x6b\x75\x40\x1e\xac\x54\x3c\x8c\x85\x40\x20\x08\xb5\x90\xa8\x47\x87\x54\x53\xb5\xc2\x6a\x64\x91\x21\x6d\x69\x9f\xfb\x87\x14\x07\x19\xa4\xfe\x6c\xe8\x2a\x56\x67\xe9\x5f\x3c\x10\x7a\xe1\x31\xa0\x4d\x3d\x06\x84\x05\xb9\xee\x0f\xbd\x9e\x9a\xe1\x0c\x92\xba\x6b\xa5\x0f\xc9\x6f\x06\x09\xed\x61\xb3\x2f\x9b\xdd\x6c\x97\xc7\x2c\x38\x4a\xe2\x3f\xda\x82\xaf\xeb\x9b\xff\xa8\xdc\x66\xab\x42\x97\x5a\x97\x21\xdf\x19\xdb\x61\xc0\x36\x85\xbe\x23\xf3\x9e\xba\x7c\x87\xca\x9a\x1a\x6f\x5c\x7c\x19\x50\xe0\x3b\xa3\x1f\xac\xbb\xc3\xbc\x3f\xaf\xda\x0e\xf0\x6f\x0a\x9b\x84\x7f\x74\x16\x71\x2b\xff\xde\x49\x84\x1c\x25\x77\xa6\xa5\xce\x82\xe7\xfe\xd8\x2f\x55\x38\x24\xa8\x03\xb2\x63\x3b\x7a\x43\x4b\x09\x80\xa5\xf2\x54\xa8\x56\x13\x04\x95\x95\x06\xec\x9d\x1d\x65\xe0\x56\x24\xa8\x6b\xcd\xd5\x58\x50\x5f\xda\xf1\x13\xb4\x32\xab\x66\x11\xb8\xe1\xc6\x38\x62\xc9\x63\x3d\x36\x20\xb8\xcf\x0d\xf2\x49\x2c\x38\x8a\x19\xdb\x11\x15\xfc\x52\x0d\x95\xca\xe0\xd8\x02\x3c\x3b\x1e\x4f\xad\x54\x2f\xc1\x66\xa4\xee\x85\xc0\x6f\x6c\xee\x84\x60\x98\x24\x06\x37\x48\x02\xd2\x27\x17\xb1\xb9\x19\x8e\x02\x08\x88\xe5\x18\x73\xac\x6e\x06\xfe\x39\xe2\x21\xe3\xa7\x12\x48\x40\x8a\x55\xf2\x50\x41\xc3\x0d\x4a\x75\x31\xd6\xe0\x99\xa0\x17\xeb\xa5\x5a\x27\x47\x66\x91\x5a\x6f\x52\x68\x25\xa2\x32\xfe\xa9\xba\xb6\x76\x8c\x48\x65\x57\x8e\x4c\x43\xa3\x19\xaa\xbc\x32\xaf\x05\xa1\xac\x61\x1f\x1c\x27\x96\xcc\xc8\xf1\x61\xad\x01\x7b\xbc\x96\x38\x10\x28\x87\x9d\xee\x09\xf4\xa5\x12\x99\xaf\x2a\x9f\xf0\x57\x98\x8b\x24\xe1\x71\x76\xdc\xc6\x78\xd6\xc4\x59\xa5\x97\x36\x67\x15\x94\xca\x92\x49\xb9\x86\x50\x2b\x1a\xc9\xf1\x8b\x53\xb4\x63\xef\x68\x63\xae\xce\x8e\xb8\x04\x95\x08\xa8\x13\xdd\x57\x3e\xcb\x11\x58\xa9\xe1\x4a\x6b\x49\x21\xc7\xb6\x72\xcc\x68\x57\x25\x38\x22\x95\xea\x49\xa2\x49\x47\x98\x6c\x22\x1d\x3a\x2e\x47\x40\xbd\x9b\xaf\xb6\x07\xd5\x1f\x58\xcb\x53\x71\xee\x4e\x48\xfd\xe1\xcd\xbd\x5d\x10\xf7\xa8\x80\x7a\x88\xe5\xe7\x87\xb7\x81\x70\x3a\x2e\x1f\x5d\x22\x79\x31\x67\xfc\x1c\x87\x26\x8d\xe6\xdc\x98\x71\x61\x37\x1b\x28\xf6\x03\x2d\xce\x8a\xdf\x12\xbd\x86\x76\x0a\x5a\x2a\xbc\x61\xee\x44\x8e\xe4\x74\x8c\xc1\xf4\x42\x73\xc8\x9c\x8c\xa2\xc2\xc0\x92\x14\xf3\x5b\x55\x65\x4c\x09\xb3\x02\xf6\x30\x41\x20\x06\xd3\xa4\x20\x2f\x6a\x1b\xeb\x69\x2a\x50\x80\x79\x92\x2b\x62\xea\xe4\xf1\x14\x1b\x65\x0e\x3b\x11\x97\xd3\xfb\x46\x6b\x03\xf0\x49\x5e\x2b\x90\x68\x27\xe5\x40\x1e\x8e\x1e\x38\xa0\x36\x11\xd9\x49\x2b\x39\x34\x19\x84\x9c\xfc\x3a\x8d\xa6\xa9\x47\x84\x76\x50\x41\x5f\x3e\x4b\x2d\xa5\xb5\x3c\xba\x58\xd5\x12\xe4\xb0\xc7\xa9\x8c\x72\xf2\xed\x94\x57\x4c\x1e\xa8\x27\xd3\xa0\x5e\x83\x69\x52\x37\x89\x8c\x67\xdc\x01\x35\x21\x64\x93\x02\x6b\x63\x50\x56\x28\xad\xca\x6a\x18\x3c\xd4\x42\xe4\x90\xa3\xd6\x83\x0b\x39\xb1\x8f\x74\x4d\x3b\xe3\x89\xfe\x04\x42\x14\x65\x05\x3e\xed\xe6\x44\x67\xa7\xf7\xc0\x4d\x68\xb7\x02\x38\x81\x05\x3e\xb8\xed\x11\xda\xc9\x3b\xfc\x52\x7b\xa2\x46\x55\x37\xc3\x10\x52\xd9\xc1\x87\xe4\xe4\x99\xad\xd3\x7d\xc1\xd2\x93\x83\xff\x01\x4a\x77\xe0\xe4\x41\xa1\xff\x00\x99\xc7\x83\xcb\x0b\x07\xfe\xdb\x84\xe9\xa1\x26\x7c\x74\x1c\xf5\xcc\x50\xf3\xbd\x5e\x6f\xfe\x5f\x0b\xfd\x07\xc4\xde\x4f\x99\x07\x08\xad\x74\x79\xf0\x22\xfa\xa6\x26\x0f\xd0\x78\xa8\x22\x6f\x97\xe3\x59\xf5\x78\x8b\x20\x1f\x4b\x32\xf0\xc1\x2a\x3c\x3d\x1f\x88\xec\x5d\xba\x96\x0f\xe4\x77\xf2\x81\x54\x4f\xfa\xfb\x7c\x00\x70\xc6\x19\xe2\x78\x88\x6f\x89\x67\x51\x7b\xc0\xb8\xb6\xcf\x52\x90\x0e\xf5\x0a\x10\x0e\x35\x35\x40\x6e\x7a\x27\x99\x5b\x4c\x67\xc5\xd4\xca\xe2\x50\x29\x73\x59\x41\x1e\x3e\xe0\xd5\xd9\xb5\x41\x9c\x0e\x71\x22\x25\x1f\x8a\xcb\x12\x5b\x37\x5a\x46\x92\xeb\x7d\xfd\x69\x9e\x39\x3d\x1c\x42\x66\xa7\x7d\xad\x5d\xa5\x6d\xec\xda\xa0\x1c\x86\x38\x5d\x27\x9c\x86\x5c\xf7\xb2\xeb\x8c\x85\x15\xca\xe5\x1a\x59\xb1\x2a\xe2\xca\x2d\x80\x50\xe9\x32\x8e\xeb\x4c\xe4\x1a\x38\x33\xc9\x71\xc6\x1a\x0f\xc7\x96\xb7\xc4\x56\xfc\x87\x3f\xb9\xda\xd4\x69\x59\xaf\xb0\xdb\x01\xad\xdd\x00\x68\xed\x9a\xd9\x78\xac\x39\x0c\x8e\x5a\x4e\x0f\x2d\xc8\x48\xcc\xbc\x59\x2f\xe6\x92\x42\x79\xa9\xfd\x70\xc0\xc4\xec\xf4\x69\xfa\x2b\x75\x76\x6b\xa1\xb3\x78\xdd\x59\x7e\x20\xd7\x24\x23\xcd\x5b\x18\x2f\xce\xdd\xd9\xe1\x85\x5d\x1f\x0b\xed\x9a\x1b\x7b\x9a\xd8\x75\x4a\x39\x4c\x70\xda\x06\x4f\xf4\x27\x8f\x93\x3c\x30\x53\x19\x24\xa6\xcf\x75\x5f\x38\x80\x21\x07\x70\x96\x0f\xc8\x4d\x92\x6e\x75\x9d\x8e\x36\x78\x66\x8f\x28\x57\xe2\x94\x25\xf7\x68\x75\x8d\x79\x27\xd7\x32\x93\xd5\x61\x05\xba\x4a\x29\x71\xe8\x7a\x15\xc1\x0c\xeb\xd2\xeb\x8e\x5c\xd1\x1c\xce\x0d\x36\x8c\xe7\x19\x80\xf0\xc3\x0c\x61\xf3\x77\x36\xf7\x60\xe0\x63\xf2\x84\x6d\xec\xbf\x08\xfe\x0d\x64\x0b\x49\x18\x7d\x4c\x2e\xfc\xff\xe5\xf7\xff\xd6\xf2\xfb\x96\x6d\x7e\x06\x6b\x1d\x0a\xc3\xb7\xff\xec\xed\x1e\xdb\x7d\x7c\x30\xbe\x5b\xb0\x17\xb1\xe3\x8f\xba\x14\xff\xa3\xfe\xee\xef\xf9\x7e\x8e\xb3\xe3\xe7\x36\xe8\x5e\x68\x37\x23\x87\x7c\x08\xb1\x4f\x6c\x37\x11\xc5\x6e\x87\x75\x73\x80\xe5\xfd\xf4\x72\x80\xd0\x4a\x27\x07\x7f\x5d\xb9\xa9\x8e\x03\x34\x1e\xaa\x8a\xdb\xe5\x78\x56\x35\xdc\x22\x48\xaf\x81\x8f\x0f\xa7\xcb\x2c\xce\xf0\x43\x2f\xa0\x80\x7b\x87\xd3\x84\x81\x84\x6c\xfd\x7d\x19\xee\x33\xc8\xaf\x61\xda\x95\x1b\x41\x2d\xa9\x79\xd4\x76\x77\x83\x30\x5a\x7e\x7d\x63\x75\x2a\xd4\x13\x17\x04\x2d\xa9\x07\xd9\x9d\x0c\x81\xed\x50\x2f\xb9\xf2\x37\x35\x75\x8e\xed\x5e\xcf\x36\x95\x26\x4c\xe5\x57\x13\xb9\xdd\x19\x89\xba\x87\xa9\xf4\xfe\x61\xd1\x06\x99\x4d\x75\xde\x15\x16\x6d\x0c\xdf\x4b\x93\x5f\x2a\x2c\x1a\x92\xa0\x0f\x8b\x1e\x1d\x15\xbd\xa4\x12\x3f\x31\x2a\xa2\x10\x91\x5a\xd4\xc3\xc3\x0f\x28\xfe\x1d\x00\x00\xff\xff\xb4\x02\x7d\xab\x00\x40\x00\x00")

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

	info := bindataFileInfo{name: "templates.go", size: 32768, mode: os.FileMode(420), modTime: time.Unix(1469005271, 0)}
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
	"init/django/.dockerignore": initDjangoDockerignore,
	"init/django/Dockerfile": initDjangoDockerfile,
	"init/django/docker-compose.yml": initDjangoDockerComposeYml,
	"init/rails/.dockerignore": initRailsDockerignore,
	"init/rails/Dockerfile": initRailsDockerfile,
	"init/rails/docker-compose.yml": initRailsDockerComposeYml,
	"init/ruby/.dockerignore": initRubyDockerignore,
	"init/ruby/Dockerfile": initRubyDockerfile,
	"init/ruby/docker-compose.yml": initRubyDockerComposeYml,
	"init/sinatra/.dockerignore": initSinatraDockerignore,
	"init/sinatra/Dockerfile": initSinatraDockerfile,
	"init/sinatra/docker-compose.yml": initSinatraDockerComposeYml,
	"init/unknown/.dockerignore": initUnknownDockerignore,
	"init/unknown/Dockerfile": initUnknownDockerfile,
	"init/unknown/docker-compose.yml": initUnknownDockerComposeYml,
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
		"django": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initDjangoDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initDjangoDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initDjangoDockerComposeYml, map[string]*bintree{}},
		}},
		"rails": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initRailsDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initRailsDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initRailsDockerComposeYml, map[string]*bintree{}},
		}},
		"ruby": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initRubyDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initRubyDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initRubyDockerComposeYml, map[string]*bintree{}},
		}},
		"sinatra": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initSinatraDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initSinatraDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initSinatraDockerComposeYml, map[string]*bintree{}},
		}},
		"unknown": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initUnknownDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initUnknownDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initUnknownDockerComposeYml, map[string]*bintree{}},
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

