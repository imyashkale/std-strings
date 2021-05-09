package timego

import (
	"fmt"
	"time"
)

func TimeGO() {
	t1 := time.Now()
	time.Sleep(1 * time.Second)
	t2 := time.Until(t1)
	fmt.Println("Total Time ", t2)

	hrs, _ := time.ParseDuration("1h")
	hrmin, _ := time.ParseDuration("1h30m")

	fmt.Printf("Convert to Hours : %f \n", hrs.Hours())
	fmt.Printf("Convert to Minute : %f \n", hrs.Minutes())
	fmt.Printf("Convert to Seconds : %f \n", hrs.Seconds())

	fmt.Printf("H : %d : M : %d \n", hrs, hrmin)
	fmt.Println("since t1 :", time.Since(t1))

	y, m, d := time.Now().Date()
	fmt.Printf("Y:%d - M:%d - D:%d \n", y, m, d)
}
