package main

import "fmt"

func main(){
	s := "akhil"

	fmt.Println(reverse("razi"))
	fmt.Println(reverse2(s))
	fmt.Println(palindrome("malayalam"))
	c := Count{}
	fmt.Println(c.count("malayalam"))
}

func reverse(s string) string{
	arr := []rune(s)
	l := len(arr)-1

	for i:=0; i < len(arr)/2; i++{
		arr[i],arr[l] = arr[l], arr[i]
		l--
	}

	return string(arr)
}

func reverse2(s string) string{
	arr := []rune(s)

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1{
		arr[i], arr[j] = arr[j], arr[i]
	}

	return string(arr)
}

func palindrome(s string) bool{
	return s == reverse2(s)
}

type Count struct{
	vowels int
	consonants int
}

// vowels & consonants
func (c Count) count(s string)Count{
	vowels := []rune("aeiouAEIOU")

	for _, v := range s{

		isVowel := false

		for _, d := range vowels{
			if v == d{
				c.vowels++
				isVowel = true
				break
			}
		}
		
		if !isVowel{
			c.consonants++
		}
	}

	return c
}