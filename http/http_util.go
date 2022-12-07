package http

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"syscall"
	"time"
)

type Common struct {
}

/**
   获取今天
 */
func (this *Common) GetToDay() string {
	return time.Now().Format("20060102") // yyyymmdd
}

func (this *Common) GetToNow() string {
	//return time.Now().Format("2006-01-02 18:18:18")
	//return time.Now().Format("2006-01-02 15:06:05")
	return time.Now().Format("2006-01-02 15:04:05") // yyyy-mm-dd HH:mm:ss
}

/**
   获取昨天日期
 */
func (this *Common) GetYesterDay() string {
	return this.GetDayFromTimeStamp(time.Now().AddDate(0, 0, -1).Unix())
}

func (this *Common) GetDayFromTimeStamp(timeStamp int64) string {
	return time.Unix(timeStamp, 0).Format("20060102")
}


func (this *Common) MD5(str string) string {
	md := md5.New()
	md.Write([]byte(str))
	return fmt.Sprintf("%x", md.Sum(nil))
}

/**
  c.util.Contains(peer, fileInfo.Peers)
 */
func (this *Common) Contains(obj interface{}, arrayobj interface{}) bool {
	targetValue := reflect.ValueOf(arrayobj)
	switch reflect.TypeOf(arrayobj).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj { //interface  判断类型是否相同， 值是否相等
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return false
}


func (this *Common) GetClientIp(r *http.Request) string {
	client_ip := ""
	headers := []string{"X_Forwarded_For", "X-Forwarded-For", "X-Real-Ip",
		"X_Real_Ip", "Remote_Addr", "Remote-Addr"}
	for _, v := range headers {
		if _v, ok := r.Header[v]; ok {
			if len(_v) > 0 {
				client_ip = _v[0]
				break
			}
		}
	}
	if client_ip == "" {
		clients := strings.Split(r.RemoteAddr, ":")
		client_ip = clients[0]
	}
	return client_ip
}

func (this *Common) FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}
func (this *Common) WriteFile(path string, data string) bool {
	if err := os.WriteFile(path, []byte(data), 0775); err == nil {
		return true
	} else {
		return false
	}
}

/**
	自己实现
 */
func (this *Common) WriteAppendFile(path string, data string) bool {
	if err := WriteAppendFile(path, []byte(data), 0775); err == nil {
		return true
	} else {
		return false
	}
}

func WriteAppendFile(name string, data []byte, perm fs.FileMode) error {

	f, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, perm)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}


func (this *Common) Exec(cmd []string, timeout int) (string, int) {
	var out bytes.Buffer
	duration := time.Duration(timeout) * time.Second
	ctx, _ := context.WithTimeout(context.Background(), duration)
	var command *exec.Cmd
	command = exec.CommandContext(ctx, cmd[0], cmd[1:]...)
	command.Stdin = os.Stdin
	command.Stdout = &out
	command.Stderr = &out
	err := command.Run()
	if err != nil {
		log.Println(err, cmd)
		return  err.Error(), -1
	}
	status := command.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	return out.String(), status
}