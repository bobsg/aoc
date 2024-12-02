package main

func Similarity(list1, list2 []int64) int64 {
	lookup := make(map[int64]int64)
	for i := 0; i < len(list2); i++ {
		lookup[list2[i]] += 1
	}

	var similarity int64
	for i := 0; i < len(list1); i++ {
		similarity += list1[i] * lookup[list1[i]]
	}

	return similarity
}
