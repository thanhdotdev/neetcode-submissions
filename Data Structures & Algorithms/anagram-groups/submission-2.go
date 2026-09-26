func groupAnagrams(strs []string) [][]string {
	res := make(map[[26]int][]string)

	for _, s := range strs {

		var count [26]int
		for _, ch := range s {
			count[ch-'a']++
		}

		res[count] = append(res[count], s)
	}

	var result [][]string
	for _,r := range res {
		result = append(result, r)
	}

	return result
}
