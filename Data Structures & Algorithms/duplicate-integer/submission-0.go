func hasDuplicate(nums []int) bool {
    hashMap := make(map[int]int)

	for i,num := range nums {
		if _, ok := hashMap[num]; ok {
			return true
		}

		hashMap[num] = i
	}

	return false;
}
