package main

import "fmt"

// slice 一个slice由三部分组成指针、长度和容量
func main() {

	months := [...]string{1: "January", 2: "February",
		3: "March", 4: "Aprill",
		5: "May", 6: "June",
		7: "July", 8: "August",
		9: "September", 10: "October",
		11: "November", 12: "December"}
	summer := months[6:9]
	Q2 := months[4:7]
	fmt.Println(Q2)
	fmt.Println(summer)
	fmt.Println("summer length:", len(summer))
	fmt.Println("summer cap:", cap(summer))
	//fmt.Println(summer[:20])
	endlessSummer := summer[:7]
	fmt.Println(endlessSummer)
	fmt.Println("summer length:", len(endlessSummer))
	fmt.Println("summer cap:", cap(endlessSummer))
	equal(summer, Q2)
}
func equal(s1, s2 []string) {
	for _, ss1 := range s1 {
		for _, ss2 := range s2 {
			if ss1 == ss2 {
				fmt.Printf("%s appears in both\n", ss1)
			}
		}
	}
}
