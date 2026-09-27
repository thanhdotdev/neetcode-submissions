type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}

	var sizes []string
	for _, str := range strs {
		sizes = append(sizes, strconv.Itoa(len(str)))
	}

	fmt.Print(strings.Join(sizes, ",") + "#" + strings.Join(strs, ""))
	return strings.Join(sizes, ",") + "#" + strings.Join(strs, "")
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}

	part := strings.SplitN(encoded, "#", 2)
	sizes := strings.Split(part[0], ",")

	i := 0
	res := []string{}
	for _, size := range sizes {
		if size == "" {
			continue
		}

		length, _ := strconv.Atoi(size)
		res = append(res, part[1][i:i+length])

		i+=length
	}

	return res
}
