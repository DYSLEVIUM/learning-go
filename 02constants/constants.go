package main

import "fmt"

const (
	iot = iota + 5
	iot2
	iot3
)

func main() {
	println(iot)
	println(iot2)
	println(iot3)

	const m int = 20
	fmt.Println(m)

	const (
		_  = iota // ignore first value by assigning to black identifier
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	filesize := 4e9
	fmt.Printf("%.2fGB\n", filesize/GB)

	const (
		isAdmin = 1 << iota
		isHeadquarters
		canSeeFinancials

		canSeeAfrica
		canSeeAsia
		canSeeEurope
		canSeeNorthAmerica
		canSeeSouthAmerica
	)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope

	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is Headquarters? %v", isHeadquarters&roles == isHeadquarters)
}
