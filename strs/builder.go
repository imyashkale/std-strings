package strs

import (
	"fmt"
	"strings"
)
//Builder : Various strings method Implementation
func Builder() {

	var builder strings.Builder
	builder.WriteString("Yash Kale")

	fmt.Println("Builder string :", builder.String())
	fmt.Println("Builder string :", builder.Cap())
	builder.WriteString("SABRINA CARPENTER")
	builder.Grow(64)
	fmt.Println("Builder string :", builder.Cap())

	fmt.Println("len of string", builder.Len())
	builder.Reset()
	fmt.Println("Builder string :", builder.String())

	builder.WriteString("SA")
	builder.WriteRune(66)
	builder.WriteString("Y")

	fmt.Println("Builder String", builder.String())
}
