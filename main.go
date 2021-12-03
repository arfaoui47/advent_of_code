package main

func main() {
	v := "0101010111"
	m := make(map[int]int)
	count := 0
	for index, i := range v {

		if i == '1' {
			count++
		}
		m[index] = count
		println(index, i, count)
	}
	// println(count)
	for g, k := range m {

		println(g, k)
	}

}
