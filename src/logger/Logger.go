package logger

import (
	"errors"
	"fmt"
	"os"
)

type LogWriter interface {
	Write(data interface{}) error
}
type consoleWriter struct {
}
type fileWriter struct {
	file *os.File
}

func (f *fileWriter) SetFile(fileName string) (err error) {
	if nil != f.file {
		f.file.Close()
	}
	f.file, err = os.Create(fileName)
	return err
}

/**
 *@description:TODO
 *@author:txl
 *@date:2023/1/17 23:29
 *
 */
func (f *fileWriter) Write(data interface{}) error {
	if f.file == nil {
		return errors.New("文件不存在！")
	}
	str := fmt.Sprintf("%v\n", data)
	_, err := f.file.Write([]byte(str))
	return err
}

func newConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}
func newFileWriter() *fileWriter {
	return &fileWriter{}
}
func (c *consoleWriter) Write(data interface{}) error {
	str := fmt.Sprintf("%v\n", data)
	_, err := os.Stdout.Write([]byte(str))
	return err
}

type LoggerTool struct {
	writerList []LogWriter
}

func (l *LoggerTool) RegisterWriter(Writer LogWriter) {
	l.writerList = append(l.writerList, Writer)
}
func (l *LoggerTool) Log(data interface{}) {
	for _, writer := range l.writerList {
		writer.Write(data)
	}
}

func NewLogger() *LoggerTool {
	return &LoggerTool{}
}
func CreateLoggerTool() *LoggerTool {
	l := NewLogger()
	cw := newConsoleWriter()
	l.RegisterWriter(cw)
	fw := newFileWriter()
	os.ReadFile("")
	if err := fw.SetFile("学习地址路径大全.log"); err != nil {
		fmt.Println(err)
	}
	l.RegisterWriter(fw)
	return l
}
