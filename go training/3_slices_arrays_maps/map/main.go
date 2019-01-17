package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	m["todd"] = 33

	fmt.Println(m)

	delete(m, "James")
	fmt.Println(m)

	delete(m, "Ian") //you can delete things that don't exist
	fmt.Println(m)

	fmt.Println(m["Miss Moneypenny"])
	fmt.Println(m["Ian"])

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss Moneypenny")
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println(m)
}
