package twoSum

import (
	"testing"
)

func twoSum(numbers []int, target int) []int {
	left := 0
	rigth := len(numbers) - 1
	for left != rigth {
		if numbers[left]+numbers[rigth] == target {
			return []int{left + 1, rigth + 1}
		}
		if numbers[left]+numbers[rigth] > target {
			rigth--
		} else {
			left++
		}
	}
	return []int{}
}

func TestTwoSum(t *testing.T) {

}
