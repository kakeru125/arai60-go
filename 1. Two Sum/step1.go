package main

func twoSum(nums []int, target int) []int {
	indexByNum := make(map[int]int)

	for i, num := range nums {
		if pair, ok := indexByNum[target-num]; ok {
			return []int{pair, i}
		}
		indexByNum[num] = i
	}

	return []int{}
}
