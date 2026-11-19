package main

func twoSum(nums []int, target int) []int {
	indexByNum := make(map[int]int)

	for i, num := range nums {
		if complement, exists := indexByNum[target-num]; exists {
			return []int{complement, i}
		}
		indexByNum[num] = i
	}

	return nil
}
