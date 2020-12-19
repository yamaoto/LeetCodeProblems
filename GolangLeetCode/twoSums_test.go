package twosums

import (
    "fmt"
    "testing"
)

func twoSum(nums []int, target int) []int {
	return twoSumMap(nums, target)
}

func twoSumMap(nums []int, target int) []int {
	m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        part := target - nums[i];
        if j, ok := m[part]; ok && i != j {
            return []int {i, m[part]}
        }
        m[nums[i]] = i;
    }
    return []int {}
}

// TestTwoSumMap is unit test for twoSumMap
func TestTwoSumMap(t *testing.T) {
	var tests = []struct {
        nums []int
		target int
		want []int
    }{
        {[]int {3, 2, 4}, 6, []int{1,2}},
        {[]int {2, 7, 11, 15}, 9, []int{0,1}},
        {[]int {3, 3}, 6, []int{0,1}},
	}
	
	for _, tt := range tests {
        testname := fmt.Sprintf("%p,%d", tt.nums, tt.target)
        t.Run(testname, func(t *testing.T) {
            ans := twoSumMap(tt.nums, tt.target)
            if (ans[0] != tt.want[0] || ans[1] != tt.want[1]) && (ans[1] != tt.want[0] || ans[0] != tt.want[1]) {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}

func twoSumSimple(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// TestTwoSimple is unit test for twoSumSimple
func TestTwoSimple(t *testing.T) {
	var tests = []struct {
        nums []int
		target int
		want []int
    }{
        {[]int {3, 2, 4}, 6, []int{1,2}},
        {[]int {2, 7, 11, 15}, 9, []int{0,1}},
        {[]int {3, 3}, 6, []int{0,1}},
	}
	
	for _, tt := range tests {
        testname := fmt.Sprintf("%p,%d", tt.nums, tt.target)
        t.Run(testname, func(t *testing.T) {
            ans := twoSumSimple(tt.nums, tt.target)
            if ans[0] != tt.want[0] || ans[1] != tt.want[1] {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
