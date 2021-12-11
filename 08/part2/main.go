package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg

1, 4, 7, 8 have unique # of letters

 abcdefg
0abc efg
1  c  f
2a cde g
3a cd fg
4 bcd f
5ab d fg
6ab defg
7a c  f
8abcdefg
9abcd fg

known
8abcdefg
1  c  f
4 bcd f
7a c  f  (a)
5ab d fg (a,f)
0abc efg (a,f)
2a cde g


1. find digit with 2 segments, this is 1 and the segments are c and f
2. find digit with 4 segments, this is 4 and the segments are bcdf
   the 2 segments not in 1 are bd
3. find digit with 3 segments, this is 7 and the segments are acf
   the 1 segment not in 1 is *a* (all other numbers include a though)
4. find the digit with 7 (all) segments, this is 8
5. 2, 3, 5 have 5 segments
   	7. only 5 includes b, d (which we know from above)
      the 1 segment other than a that is in both 5 and 7 is *f*
	9. only 2 does not include f
	10. the remaining 5 segment digit is 3
6. 0, 6, 9 have 6 segments
   	8. only 0 doesn't include both b,d (which we know from above)
	11. 9 includes all the digits in 3, 6 does not
	12. remaining digit is 6
*/

func subtract(one string, two string) string {
	// get all characters in one not in two

	mtwo := make(map[byte]bool)
	for i := 0; i < len(two); i++ {
		b := two[i]
		mtwo[b] = true
	}

	s := make([]byte, 0)
	for i := 0; i < len(one); i++ {
		b := one[i]
		if _, isPresent := mtwo[b]; !isPresent {
			s = append(s, b)
		}
	}
	return string(s)
}

func overlap(one string, two string) string {
	mtwo := make(map[byte]bool)
	for i := 0; i < len(two); i++ {
		b := two[i]
		mtwo[b] = true
	}

	s := make([]byte, 0)
	for i := 0; i < len(one); i++ {
		b := one[i]
		if _, isPresent := mtwo[b]; isPresent {
			s = append(s, b)
		}
	}
	return string(s)

}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func getOutputValue(input []string, output []string) int {
	key := make(map[string]int)
	var one, four, seven string
	fives := make([]string, 0)
	sixes := make([]string, 0)
	for _, digit := range input {
		if len(digit) == 2 {
			one = digit
			key[digit] = 1
		} else if len(digit) == 3 {
			seven = digit
			key[digit] = 7
		} else if len(digit) == 4 {
			four = digit
			key[digit] = 4
		} else if len(digit) == 7 {
			key[digit] = 8
		} else if len(digit) == 5 {
			fives = append(fives, digit)
		} else if len(digit) == 6 {
			sixes = append(sixes, digit)
		} else {
			fmt.Println("what is this?", digit)
			return 0
		}
	}

	// find 5:
	var five string
	bd := subtract(four, one)
	for _, digit := range fives {
		missing_bd := subtract(bd, digit)
		if len(missing_bd) == 0 {
			five = digit
			key[digit] = 5
			break
		}
	}
	// find 2 and 3
	var three string
	af := overlap(five, seven)
	a := subtract(seven, one)
	f := subtract(af, a)

	for _, digit := range fives {
		if digit == five {
			continue
		}
		missing_f := subtract(f, digit)
		if len(missing_f) == 0 {
			// 3 includes f
			three = digit
			key[digit] = 3
		} else if digit != five {
			key[digit] = 2
		}
	}

	// find 0
	var zero string
	for _, digit := range sixes {
		missing_bd := subtract(bd, digit)
		if len(missing_bd) != 0 {
			// 0 doesn't include both b,d
			zero = digit
			key[digit] = 0
		}
	}
	// find 9 and 6
	for _, digit := range sixes {
		if digit == zero {
			continue
		}

		if len(subtract(three, digit)) == 0 {
			// 9 includes all the digits in 3, 6 does not
			key[digit] = 9
		} else {
			key[digit] = 6
		}
	}

	sortedKey := make(map[string]int)
	for k, v := range key {
		sk := SortString(k)
		sortedKey[sk] = v
	}
	s := 0
	m := 1
	fmt.Println(key)
	for i := len(output) - 1; i >= 0; i-- {
		so := SortString(output[i])
		fmt.Println(output[i], "is", sortedKey[so])
		s += (m * sortedKey[so])
		m *= 10
	}
	return s
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	s := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " | ")
		input := strings.Split(parts[0], " ")
		output := strings.Split(parts[1], " ")
		os := getOutputValue(input, output)
		fmt.Println("output", os)
		s += os
	}
	fmt.Println("total", s)
}
