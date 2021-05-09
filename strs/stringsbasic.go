package strs

import (
	"fmt"
	"strings"
	"unicode"
)

func Basicdtring() {
	// Compare : < > == always faster
	result := strings.Compare("Yash", "Yash")
	fmt.Println("Is Values are same ", result)

	// Containes
	isContain := strings.Contains("Yash Kale", "Yash")
	fmt.Println("CheckIf Yash Contains in Yash Kale : ", isContain)

	isContainAny := strings.ContainsAny("sabrina", "sab")
	fmt.Println("CheckIf sab containes any in sabrina :", isContainAny)

	fmt.Println("ContainsRun :", strings.ContainsRune("sabrina", 2))

	fmt.Println("Count Input in string", strings.Count("sabrina carpenter", "e"))

	fmt.Println("CheckIf both input are equal case insensitively :", strings.EqualFold("sabrina", "Sabrina"))

	fmt.Printf("Extract each word from string: %q \n", strings.Fields("sabrina anyln carpenter"))

	f := func(c rune) bool {
		// Not a Letter & Not Letter
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	fmt.Printf("Fields are (Excluding Number and Letter) %q \n", strings.FieldsFunc(" sab;yash;kale;...", f))

	fmt.Printf("Has Prefix : %t \n", strings.HasPrefix("sabrinacarpenter", "sab"))

	fmt.Printf("Has suffix : %t \n", strings.HasSuffix("sabrina", "brina"))

	fmt.Printf("Index b in [ sabrina ] : %v \n", strings.Index("sabrina", "b"))
	fmt.Printf("Index sara in [ saracarpenter ] : %v \n", strings.Index("saracarpenter", "sara"))

	fmt.Printf("IndexAny [sabrina] : %v \n", strings.IndexAny("sabrina", "zs"))

	fmt.Printf("IndexByte [ sabrina ] index of first instance of i %v \n", strings.IndexByte("sabrina", 'i'))

	findIndexIf := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}

	fmt.Println("findIndexIf its in Han : ", strings.IndexFunc("Yash , kale", findIndexIf))

	fmt.Println("Return Index : if 66 (unicode) found in input string ", strings.IndexRune("SABRINA", 66))

	favlang := []string{
		"python",
		"javascript",
		"go-lang",
		"rust",
	}

	fmt.Printf("Join the slice string by | %s \n ", strings.Join(favlang, " |"))

	// SAME AS LASTINDEXANY LASTINDEXBYTE LASTINDEXFUNC
	fmt.Printf("Last Index in Input string of substring %v  \n", strings.LastIndex("carpenter", "e"))

	fmt.Printf("Repeat sabrina again and again for 4 Times %s \n", strings.Repeat("sabrina ", 4))

	fmt.Printf("Replace sabrina with %s \n", strings.Replace("sabrina", "sab", "saby", 1))

	fmt.Printf("[sabrina sabrina sabrina ] with sab replace all %v \n", strings.ReplaceAll("sabrina sabrina sabrina", "sabrina", "sab"))

	fmt.Printf("seprate it out each word : %q \n", strings.Split("sabrina", ""))

	list := []string{"sabrina", "sara", "arina", "liza", "joey"}

	toLower := []string{}
	for index, item := range list {
		capitalized := strings.Title(item)
		toLower = append(toLower, capitalized)
		fmt.Printf("[%v] : From %s => %s \n", index, item, strings.Title(capitalized))
	}
	fmt.Println("Len of slice :", len(toLower))

	for index, item := range toLower {
		fmt.Printf("[%v] : From %s => %s \n", index, item, strings.ToLower(item))

	}

	for index, item := range toLower {
		fmt.Printf("[%v] : All Caps : %v \n", index, strings.ToTitle(item))
	}

	fmt.Println("Trimed by !", strings.Trim("sabybaby!!", "!"))

	tobeTrimmed := " sabrina    "
	fmt.Println("Trimmed", strings.TrimSpace(tobeTrimmed))
}
