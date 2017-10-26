package main

// #include <stdlib.h>
// #include <pwd.h>
import "C"

import (
	"bytes"
	"encoding/binary"
	"log"
	"os"
	"time"
	"unsafe"
)

type Passwd struct {
	Name    string
	Passwd  string
	Uid     uint32
	Gid     uint32
	Comment string
	Home    string
	Shell   string
}

func passwdC2Go(passwdC *C.struct_passwd) *Passwd {
	return &Passwd{
		Name:    C.GoString(passwdC.pw_name),
		Passwd:  C.GoString(passwdC.pw_passwd),
		Uid:     uint32(passwdC.pw_uid),
		Gid:     uint32(passwdC.pw_gid),
		Comment: C.GoString(passwdC.pw_gecos),
		Home:    C.GoString(passwdC.pw_dir),
		Shell:   C.GoString(passwdC.pw_shell),
	}
}

type record struct {
	time int32
	line [32]byte
	host [256]byte
}

type UserInfo struct {
	Name string
	Line string
	Host string
	Last string
}

var rsize = unsafe.Sizeof(record{})

func main() {
	f, err := os.Open("/var/log/lastlog")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	stats, err := f.Stat()
	if err != nil {
		panic(err)
	}
	size := stats.Size()

	passwds := make([]*Passwd, 0)
	C.setpwent()
	for passwdC, err := C.getpwent(); passwdC != nil && err == nil; passwdC, err = C.getpwent() {
		passwd := passwdC2Go(passwdC)
		passwds = append(passwds, passwd)
	}
	C.endpwent()

	for _, p := range passwds {
		last, line, host, err := getLogByUID(int64(p.Uid), f, size)
		if err != nil {
			panic(err)
		}

		var lastlog string
		if last == time.Unix(0, 0) {
			lastlog = "**Never logged in**"
		} else {
			lastlog = last.String()
		}

		var info = &UserInfo{
			Name: p.Name,
			Line: line,
			Host: host,
			Last: lastlog,
		}
		log.Printf("%#v", info)
	}
}

func getLogByUID(uid int64, lastLog *os.File, lastLogSize int64) (time.Time, string, string, error) {
	offset := uid * int64(rsize)
	if offset+int64(rsize) <= lastLogSize {
		_, err := lastLog.Seek(offset, 0)
		if err != nil {
			return time.Unix(0, 0), "", "", err
		}
		rawRecord := make([]byte, rsize)
		_, err = lastLog.Read(rawRecord)
		if err != nil {
			return time.Unix(0, 0), "", "", err
		}
		return bytes2time(rawRecord[:4]), string(bytes.Trim(rawRecord[4:36], "\x00")), string(bytes.Trim(rawRecord[36:], "\x00")), nil
	}
	return time.Unix(0, 0), "", "", nil
}

func bytes2time(b []byte) time.Time {
	return time.Unix(int64(binary.LittleEndian.Uint32(b)), 0)
}
