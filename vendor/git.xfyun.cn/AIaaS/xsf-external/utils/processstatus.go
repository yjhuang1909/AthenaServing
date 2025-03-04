/*
 @brief		this is a limiter for qps
 @file		qpslimiter
 @author	sqjain
 @version	1.0
 @date		2017.11.17
*/

// this is a status set for of process.
package utils

import (
	"os"
	"os/user"
	"time"
	"runtime"
	"strings"
	"strconv"
)

// Pid（进程ID）
// User（进程所有者）
// Uptime（启动的时间）
// Goroutines（协程数）
// GoroutineID（协程标识）
var basetime int64

type ps struct {
}

func init() {
	basetime = time.Now().UnixNano()
}

// get process id
func (p *ps) GetPid() int {
	return os.Getpid()
}

// get process user
func (p *ps) GetUser() (string, error) {
	u, e := user.Current()
	if nil != e {
		return "", e
	}
	return u.Username, nil
}

// indicates seconds
func (p *ps) GetUptime() float64 {
	return float64(time.Now().UnixNano()-basetime) / 1000000000
}

// get goroutines
func (p *ps) GetGoroutines() int {
	return runtime.NumGoroutine()
}

//get gorouotine id
// -1 stand for err
func (p *ps) GetGoroutineID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if nil != err {
		return -1
	}
	return id
}
