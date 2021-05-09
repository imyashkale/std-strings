package mapgo

import (
	"fmt"
)

func Generator() func() int {
	i := 0
	return func() int {
		i++
		return i

	}
}

func GeneratorGo() {

	generator := Generator()

	for i := 0; i <= 10; i++ {
		fmt.Printf("Called %d  %d \n", i, generator())
	}

	const KEY = "KSHFJHG"
	fmt.Println("Config Key : ", KEY)

}
