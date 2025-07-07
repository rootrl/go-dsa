package bubble

import (
	"reflect"
	"testing"
)

func TestBubbleStrict(t *testing.T) {
	testCases := []struct {
		name     string // 测试用例的描述
		input    []int  // 输入数据
		expected []int  // 期望的输出
	}{
		// 2. 设计一系列覆盖不同场景的测试用例
		{
			name:     "通用随机数组",
			input:    []int{6, 1, 3, 2, 5, 4},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "已排序数组",
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "逆序数组 (最坏情况)",
			input:    []int{6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "包含重复元素的数组",
			input:    []int{5, 8, 5, 2, 9},
			expected: []int{2, 5, 5, 8, 9},
		},
		{
			name:     "包含负数的数组",
			input:    []int{-1, -5, 3, 0},
			expected: []int{-5, -1, 0, 3},
		},
		{
			name:     "空数组",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "单元素数组",
			input:    []int{42},
			expected: []int{42},
		},
	}

	// 3. 遍历所有测试用例并执行测试
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 创建输入的副本，防止排序修改原始测试用例数据
			inputCopy := make([]int, len(tc.input))
			copy(inputCopy, tc.input)

			bubble(inputCopy)

			if !reflect.DeepEqual(inputCopy, tc.expected) {
				t.Errorf("Bubble failed for case '%s'.\nInput:    %v\nGot:      %v\nExpected: %v", 
					tc.name, tc.input, inputCopy, tc.expected)
			}
		})
	}
}
