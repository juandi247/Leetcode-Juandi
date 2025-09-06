package main

import (
	"slices"
	"strconv"
	"strings"
	"testing"
)

/*

📌 Exercise 1 – Suspicious Activity From Logs
Problem Statement

You are given a log file represented as a string array.
Each entry records a money transfer in the following format:

"sender_user_id recipient_user_id amount"


sender_user_id: the ID of the user initiating the transfer.

recipient_user_id: the ID of the user receiving the transfer.

amount: the transferred amount.

The three values are numeric (0–9), up to 9 digits long, and cannot start with zero.
The log entries are provided in no particular order.

Write a function to identify suspicious users — users who appear in at least threshold transactions, whether as a sender or as a recipient.

Return an array of user IDs that meet the threshold condition, sorted in ascending numerical order.

Example
logs = ["88 99 200", "88 99 300", "99 32 100", "12 12 15"]
threshold = 2


User 88 appears in 2 transactions.

User 99 appears in 3 transactions.

User 32 appears in 1 transaction.

User 12 appears in 1 transaction (counted once even if sender and recipient in the same log).

Output:

["88", "99"]*/

func Suspicious(miArray []string, threshold int) []int {

	rta := []int{}
	countMap := make(map[string]int)

	for _, v := range miArray {
		tempArray := strings.Split(v, " ")

		countMap[tempArray[0]]++

		// edgeeeee case where the same user sends and receives, this should not count as two
		// but should be here because, if he appeared before, then we need to check that out first before we go to the next ocmparison
		if tempArray[0] == tempArray[1] {
			continue
		}

		countMap[tempArray[1]]++

	}

	for key, val := range countMap {
		integerValue, _ := strconv.Atoi(key)
		if val >= threshold {
			rta = append(rta, integerValue)
		}
	}

	slices.Sort(rta)
	return rta

}

/*
APROACH:
We need to obtain the repeated numbers, or the numbers that appear the most, in order (numerical but thats egal)

1. Use a counter of some type to now how many times a number has appeared (MAP)

2. How we obtain the numbers? -> first we dont need to obtain the number, but the strnig character.
From a 8888 999923 10, we need to obtain the first two "numbers" or substrings. So we can obtain the substrings by
going through all the data in each array (its needed). So its an O(N**2) this part.

	We can also use the split, or something like that but that relly depens on the lenguge, so the easiest, would be
	Store the substrings as we go thourhg it, until we reach an empty space. We should do this twice, becasue the 3rd number is
	the amount, so it doenst matter

	3. For each character we need to save it into the map, and count how many times it has appeared.


	4. Check the Map and take the keys, where the value is more than the thershold or equal to the hterhold

5. save that into an arrray or someting and finally convert it into an array of INTEGER, probalby we would need to ordered it at the end so yeah
*/
func Test_Suspicious(t *testing.T) {
	testingStruct := []struct {
		array     []string
		threshold int
		resultado []int
	}{
		{
			array:     []string{"88 99 200", "88 99 300", "99 32 100", "12 12 15"},
			threshold: 2,
			resultado: []int{88, 99},
		},
		{
			array:     []string{"1 2 100", "2 3 200", "3 1 50", "4 1 75"},
			threshold: 2,
			resultado: []int{1, 2, 3},
		},
		{
			array:     []string{"10 20 500", "20 30 600", "30 40 700"},
			threshold: 1,
			resultado: []int{10, 20, 30, 40},
		},
		{
			array:     []string{"5 5 10", "5 5 20", "6 5 15"},
			threshold: 3,
			resultado: []int{5},
		},
		{
			array:     []string{"7 8 100", "8 9 200", "9 10 300"},
			threshold: 2,
			resultado: []int{8, 9},
		},
		{
			array:     []string{}, // edge case: empty logs
			threshold: 1,
			resultado: []int{},
		},
		{
			array:     []string{"1 2 10"}, // edge case: threshold higher than count
			threshold: 2,
			resultado: []int{},
		},
	}

	for i, v := range testingStruct {
		resultadoTEST := Suspicious(v.array, v.threshold)

		if !slices.Equal(resultadoTEST, v.resultado) {
			t.Errorf("Test case %d failed: got %v, expected %v", i+1, resultadoTEST, v.resultado)
		}
	}
}
