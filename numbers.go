package main

type Numbers []int

func (n Numbers) Len() int {
	return len(n)
}

func (n Numbers) Less(i, j int) bool {
	return n[i] < n[j]
}

func (n Numbers) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
