package main

import "fmt"

func removeLeftRight(arr []any, position string) []any {
	//yout code here!
	var last = []interface{}{}

	if position == "left" {
		last = append(last, arr[1:]...)
	} else if position == "right" {
		last = append(last, arr[:len(arr)-1]...)
	} else {
		//
	}
	return last
}

func main() {
	var arr = []interface{}{"Edo", "Budi", "Joko", "Tono"}
	fmt.Println(removeLeftRight(arr, "left"))
}
