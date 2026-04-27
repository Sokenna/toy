package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestEcho1(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "无参数",
			args:    []string{},
			wantErr: false,
		},
		{
			name:    "单个参数",
			args:    []string{"hello"},
			wantErr: false,
		},
		{
			name:    "多个参数",
			args:    []string{"hello", "world", "test"},
			wantErr: false,
		},
		{
			name:    "包含空字符串参数",
			args:    []string{"hello", "", "world"},
			wantErr: false,
		},
		{
			name:    "包含特殊字符参数",
			args:    []string{"hello", "world!@#$%", "test"},
			wantErr: false,
		},
		{
			name:    "包含空格参数",
			args:    []string{"hello world", "test"},
			wantErr: false,
		},
		{
			name:    "长字符串参数",
			args:    []string{"a very long string with many words to test performance"},
			wantErr: false,
		},
		{
			name:    "大量参数",
			args:    []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 保存原始 os.Args
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			// 设置测试参数
			os.Args = append([]string{"cmd"}, tt.args...)

			// 捕获标准输出
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// 执行被测函数
			Echo1()

			// 恢复标准输出
			w.Close()
			os.Stdout = oldStdout

			// 读取输出
			var buf bytes.Buffer
			buf.ReadFrom(r)
			output := buf.String()

			// 验证输出包含时间（毫秒数）
			if output == "" {
				t.Errorf("Echo1() 没有输出")
			}

			// 验证输出是数字（毫秒）
			var ms int64
			_, err := fmt.Sscanf(output, "%d", &ms)
			if err != nil {
				t.Errorf("Echo1() 输出格式错误: %v, 输出: %s", err, output)
			}

			// 验证毫秒数是非负数
			if ms < 0 {
				t.Errorf("Echo1() 输出毫秒数为负数: %d", ms)
			}
		})
	}
}

// BenchmarkEcho1-10    	 1000000	    121625 ns/op
func BenchmarkEcho1(b *testing.B) {
	// 设置测试参数
	oldArgs := os.Args
	os.Args = []string{"cmd", "hello", "world", "test", "benchmark"}
	defer func() { os.Args = oldArgs }()

	// 捕获标准输出
	oldStdout := os.Stdout
	_, w, _ := os.Pipe()
	os.Stdout = w
	defer func() {
		w.Close()
		os.Stdout = oldStdout
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Echo1()
	}
}

// BenchmarkEcho2-10    	  756393	    160025 ns/op
func BenchmarkEcho2(b *testing.B) {
	// 设置测试参数
	oldArgs := os.Args
	os.Args = []string{"cmd", "hello", "world", "test", "benchmark"}
	defer func() { os.Args = oldArgs }()

	// 捕获标准输出
	oldStdout := os.Stdout
	_, w, _ := os.Pipe()
	os.Stdout = w
	defer func() {
		w.Close()
		os.Stdout = oldStdout
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Echo2()
	}
}
