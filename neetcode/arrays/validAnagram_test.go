package arrays

import (
	"fmt"
	"strconv"
	"testing"
)

/*
Valid Anagram
Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.

An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.
Example 1:

Input: s = "racecar", t = "carrace"

Output: true
Example 2:

Input: s = "jar", t = "jam"

Output: false

*/

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		fmt.Println("retornamos false porque no tienenn la misma length")
		return false
	}

	miMap := make(map[rune]int)

	for i, value := range s {
		_, exists := miMap[value]
		fmt.Println("value", value)
		if !exists {
			miMap[value] = 1
		} else {
			valor:= miMap[value]
			miMap[value]= valor-1
		}

		_, exists= miMap[rune(t[i])]

		if !exists {
			miMap[value] = 1
		} else {
			valor:= miMap[value]
			miMap[value]= valor-1
		}


	}



	for key, value:= range miMap{
		if value!=0{
			fmt.Println("flase",key, value )
			return false
		}
	}


return true

}



func TestAnagram(t *testing.T){

miStruct:= []struct{
	palabra string
	palabra2 string	
	result bool
}{

	{palabra: "jar",
	palabra2: "jar",	
	result: true,
	},
	// {palabra: "miiiimo",
	// palabra2: "miiiimo",	
	// result: true,
	// },


	// {palabra: "",
	// palabra2: "",	
	// result: true,
	// },
	// {palabra: "sdsajksajk",
	// palabra2: "weuiwejd",	
	// result: false,
	// },

	// {palabra: "a",
	// palabra2: "",	
	// result: false,
	// },
	// {palabra: "sdsajksajk",
	// palabra2: "wwewewjejkdsksdsjdjksdkjskjkjdseuiwejd",	
	// result: false,
	// },

}


for i,v:= range miStruct{


t.Run("test numero "+ strconv.Itoa(i), func(t *testing.T) {
	isA:= isAnagram(v.palabra, v.palabra2)
	if isA != v.result{
		t.Error("error on test")
	}
})

}

}