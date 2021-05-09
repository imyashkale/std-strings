package switchcontrolflow

import (
	"fmt"
)

//export Fallthrough
func Fallthrough(brk int) {
	switch brk {
	case brk + 1:
		fmt.Println("First Case Matched")
	case brk + 0:
		fmt.Println("5th Matched")
		// fallthrough
	case 6:
		fmt.Println("6th also matched")
	}
}

func SwitchGo() {
	swt := 2
	// If case matched
	switch swt {
	case 1:
		fmt.Println("1st Matched")
	case 2:
		fmt.Println("2nd matched")
	default:
		fmt.Println("Deafualt")
	}

	dft := "A"
	switch dft {
	case "N":
		fmt.Println("N Case Matched")
	default:
		fmt.Println("Default Matched")
	}

}
