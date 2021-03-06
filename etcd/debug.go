package etcd

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	stdlog "github.com/joefitzgerald/standardlog"
)

var logger *etcdLogger

func SetLogger(l stdlog.Logger) {
	logger = &etcdLogger{l}
}

func GetLogger() stdlog.Logger {
	return logger.log
}

type etcdLogger struct {
	log		stdlog.Logger
}

func (p *etcdLogger) Debug(args ...interface{}) {
	msg := "DEBUG: " + fmt.Sprint(args...)
	p.log.Println(msg)
}

func (p *etcdLogger) Debugf(f string, args ...interface{}) {
	msg := "DEBUG: " + fmt.Sprintf(f, args...)
	// Append newline if necessary
	if !strings.HasSuffix(msg, "\n") {
		msg = msg + "\n"
	}
	p.log.Print(msg)
}

func (p *etcdLogger) Warning(args ...interface{}) {
	msg := "WARNING: " + fmt.Sprint(args...)
	p.log.Println(msg)
}

func (p *etcdLogger) Warningf(f string, args ...interface{}) {
	msg := "WARNING: " + fmt.Sprintf(f, args...)
	// Append newline if necessary
	if !strings.HasSuffix(msg, "\n") {
		msg = msg + "\n"
	}
	p.log.Print(msg)
}

func init() {
	// Default logger uses the go default log.
	SetLogger(log.New(ioutil.Discard, "go-etcd", log.LstdFlags))
}
