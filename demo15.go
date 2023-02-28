package main

func main() {
	a := true

	switch a {
	case true:
		println("yes")
		fallthrough
	case false:
		//if a == true {
		//	print("hell")
		//}
		print("no")
	default:
		print("final")
	}
}
