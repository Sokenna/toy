package leetcode

import (
	"fmt"
	"testing"
)

func Test_countBits(t *testing.T) {
	var ans = countBits(5)
	fmt.Println(ans)
}
func Test_generate(t *testing.T) {
	var ans = generate(5)
	fmt.Println(ans)
}

func Test_MinCostClimbingStairs(t *testing.T) {
	var cost = []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	MinCostClimbingStairs(cost)
}

func TestNumWays(t *testing.T) {
	relation := [][]int{{0, 2}, {2, 1}, {3, 4}, {2, 3}, {1, 4}, {2, 0}, {0, 4}}
	r := NumWays(5, relation, 3)
	fmt.Println(r)
}

func Test_getRow(t *testing.T) {
	fmt.Println(1 ^ 1)
	fmt.Println(1 ^ 0)
	fmt.Println(0 ^ 1)
	fmt.Println(0 ^ 0)
	fmt.Println("--------")
	var r = getRow(21)
	fmt.Println(r)
}
func Test_getLongestSubsequence(t *testing.T) {
	words := []string{"a", "b", "c", "d", "e"}
	groups := []int{1, 0, 1, 1, 0}
	getLongestSubsequence(words, groups)
}

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{

		// 测试用例1：两个数组都只有一个元素，且nums1中的元素小于nums2中的元素
		{
			name: "Both arrays have one element and nums1 element is smaller",
			args: args{
				nums1: []int{1, 0},
				m:     1,
				nums2: []int{2},
				n:     1,
			},
		},
		// 测试用例2：nums1为空数组，nums2有多个元素
		{
			name: "nums1 is empty and nums2 has multiple elements",
			args: args{
				nums1: []int{0, 0, 0},
				m:     0,
				nums2: []int{1, 2, 3},
				n:     3,
			},
		},
		// 测试用例3：nums2为空数组，nums1有多个元素
		{
			name: "nums2 is empty and nums1 has multiple elements",
			args: args{
				nums1: []int{1, 2, 3, 0},
				m:     3,
				nums2: []int{},
				n:     0,
			},
		},
		// 测试用例4：两个数组都有多个元素，常规合并情况
		{
			name: "Both arrays have multiple elements for normal merge",
			args: args{
				nums1: []int{1, 3, 5, 0, 0, 0},
				m:     3,
				nums2: []int{2, 4, 6},
				n:     3,
			},
		},
		// 测试用例5：两个数组有重复元素的合并情况
		{
			name: "Both arrays have duplicate elements for merge",
			args: args{
				nums1: []int{1, 3, 3, 0, 0, 0},
				m:     3,
				nums2: []int{3, 4, 5},
				n:     3,
			},
		}, // 测试用例6：nums1数组没有元素
		{
			name: "Both arrays have duplicate elements for merge",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			fmt.Println(tt.args.nums1)
		})
	}
}

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "sd",
			args: args{
				nums: nil,
				val:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
