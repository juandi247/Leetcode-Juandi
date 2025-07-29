package arrays

import (
	"testing"
)

func GroupAnagrams(miArray []string) [][]string {

	/*
	   1. Crear un mapa de tipo
	   Array de integers (simulando las 26 letras), y Array de strings

	   2. Evaluamos cada palabra y le ponemos sus datos en el propio mapa, si ya existe una con un mismo clave, le appendeamos
	   al array de sitnrgs, y sino la creamos en el map

	   3. al final lo que haremos es recorerr el map y poner todo en un array final y listo
	*/

	miMap := make(map[[26]int][]string)

	for _, v := range miArray {
		runeArray := [26]int{}
		for _, theRune := range v {
			runeValue := theRune - 'a'

			runeArray[runeValue] = runeArray[runeValue] + 1
			// this is the same as
			// runeArray[runeValue]++

		}

		_, exists := miMap[runeArray]
		if exists {
			miMap[runeArray] = append(miMap[runeArray], v)
		} else {
			miMap[runeArray] = []string{v}
		}
	}

	result := [][]string{}
	for _, value := range miMap {
		result = append(result, value)
	}

	return result

}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{
			name:  "caso 1",
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			name:  "caso 2: vacío",
			input: []string{},
			want:  [][]string{},
		},
		{
			name:  "caso 3: sin anagramas",
			input: []string{"abc", "def", "ghi"},
			want: [][]string{
				{"abc"},
				{"def"},
				{"ghi"},
			},
		},
		{
			name:  "caso 4: todo iguales",
			input: []string{"a", "a", "a"},
			want: [][]string{
				{"a", "a", "a"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupAnagrams(tt.input)

			if !equalUnordered(got, tt.want) {
				t.Errorf("GroupAnagrams(%v) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

func equalUnordered(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	used := make([]bool, len(b))

	for _, groupA := range a {
		found := false
		for j, groupB := range b {
			if used[j] {
				continue
			}
			if sameGroup(groupA, groupB) {
				used[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func sameGroup(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	count := make(map[string]int)
	for _, s := range a {
		count[s]++
	}
	for _, s := range b {
		count[s]--
		if count[s] < 0 {
			return false
		}
	}
	return true
}
