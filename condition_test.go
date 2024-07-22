package corekit

import (
	"testing"
)

func TestTernary(t *testing.T) {
	// 定义测试用例
	tests := []struct {
		name       string
		condition  bool
		ifOutput   int
		elseOutput int
		want       int
	}{
		{"condition is true", true, 1, 2, 1},
		{"condition is false", false, 1, 2, 2},
	}

	// 运行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Ternary(tt.condition, tt.ifOutput, tt.elseOutput)
			if got != tt.want {
				t.Errorf("Ternary() = %v, want %v", got, tt.want)
			}
		})
	}
}
