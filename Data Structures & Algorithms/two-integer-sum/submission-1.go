func twoSum(nums []int, target int) []int {
     
	numMap := make(map[int]int)
	for i, v := range nums {
		completed := target - v
		if _, ok := numMap[completed];  ok {
			return []int{numMap[completed], i}
		}

		numMap[v] = i
	}

	return nil
}
