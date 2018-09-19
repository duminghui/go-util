// Package ulog provides ...
package ulog

import (
	"os"
	"path/filepath"
	"time"

	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type Config struct {
	DebugFile string `json:"debugFile"`
	InfoFile  string `json:"infoFile"`
	WarnFile  string `json:"warnFile"`
	ErrorFile string `json:"ErrorFile"`
	FatalFile string `json:"fatalFile"`
	PanicFile string `json:"panicFile"`
}

func New(config *Config) (*logrus.Logger, error) {
	debugWriter, err := getWriter(config.DebugFile)
	if err != nil {
		return nil, err
	}
	infoWriter, err := getWriter(config.InfoFile)
	if err != nil {
		return nil, err
	}
	warnWriter, err := getWriter(config.WarnFile)
	if err != nil {
		return nil, err
	}
	errorWriter, err := getWriter(config.ErrorFile)
	if err != nil {
		return nil, err
	}
	fatalWriter, err := getWriter(config.FatalFile)
	if err != nil {
		return nil, err
	}
	panicWriter, err := getWriter(config.PanicFile)
	if err != nil {
		return nil, err
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: debugWriter,
		logrus.InfoLevel:  infoWriter,
		logrus.WarnLevel:  warnWriter,
		logrus.ErrorLevel: errorWriter,
		logrus.FatalLevel: fatalWriter,
		logrus.PanicLevel: panicWriter,
	}, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   false,
	})
	log := logrus.New()
	log.AddHook(lfHook)
	return log, nil
}

var rotateLogs = make(map[string]*rotatelogs.RotateLogs, 6)

func getWriter(path string) (*rotatelogs.RotateLogs, error) {
	writer, ok := rotateLogs[path]
	if ok {
		return writer, nil
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	os.MkdirAll(filepath.Dir(absPath), os.ModePerm)
	writer, err = rotatelogs.New(
		absPath+".%Y%m%d",
		rotatelogs.WithLinkName(absPath),
		rotatelogs.WithMaxAge(168*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		return nil, err
	}
	rotateLogs[path] = writer
	return writer, nil
}
