package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parseHex(h string) []int {
	bits := make([]int, 0)
	for i := 0; i < len(h); i++ {
		c := h[i]
		if c == '0' {
			bits = append(bits, []int{0, 0, 0, 0}...)
		} else if c == '1' {
			bits = append(bits, []int{0, 0, 0, 1}...)
		} else if c == '2' {
			bits = append(bits, []int{0, 0, 1, 0}...)
		} else if c == '3' {
			bits = append(bits, []int{0, 0, 1, 1}...)
		} else if c == '4' {
			bits = append(bits, []int{0, 1, 0, 0}...)
		} else if c == '5' {
			bits = append(bits, []int{0, 1, 0, 1}...)
		} else if c == '6' {
			bits = append(bits, []int{0, 1, 1, 0}...)
		} else if c == '7' {
			bits = append(bits, []int{0, 1, 1, 1}...)
		} else if c == '8' {
			bits = append(bits, []int{1, 0, 0, 0}...)
		} else if c == '9' {
			bits = append(bits, []int{1, 0, 0, 1}...)
		} else if c == 'A' {
			bits = append(bits, []int{1, 0, 1, 0}...)
		} else if c == 'B' {
			bits = append(bits, []int{1, 0, 1, 1}...)
		} else if c == 'C' {
			bits = append(bits, []int{1, 1, 0, 0}...)
		} else if c == 'D' {
			bits = append(bits, []int{1, 1, 0, 1}...)
		} else if c == 'E' {
			bits = append(bits, []int{1, 1, 1, 0}...)
		} else if c == 'F' {
			bits = append(bits, []int{1, 1, 1, 1}...)
		} else {
			log.Fatal("got bad input:", h)
		}
	}
	return bits
}

func bitsToDecimal(bits []int) int {
	d := 0
	m := 1
	for i := len(bits) - 1; i >= 0; i-- {
		d += (bits[i] * m)
		m *= 2
	}
	return d
}

/*
	how to restructure this? do we need a pointer to pass around?

	operator packet where the # of bits of the subpackets is specified is
	easy, we know how far to read

	operator packet where the # of subpackets is specified is harder, we
	have to start reading and then pass back control and position once
	we've read the number of packets specified

	maybe rawValueBitsToDecimal isn't a separate method? that way we still
	have the position once it's done

	how to handle the "Read N packets" case?

	can we let the caller control how many times to read?

*/

func readPacket(bits []int, startPos int) (pos int, value int) {
	pos = startPos
	//versionBits := bits[pos : pos+3]
	//version := bitsToDecimal(versionBits)
	typeIDBits := bits[pos+3 : pos+6]
	typeID := bitsToDecimal(typeIDBits)
	//fmt.Println("typeID:", typeID)
	pos += 6

	//fmt.Println("typeID:", typeID)

	if typeID == 4 {
		// literal value
		valueBits := make([]int, 0)
		stillReading := true

		for {
			if !stillReading {
				break
			}

			if bits[pos] == 0 {
				stillReading = false
			}

			valueBits = append(valueBits, bits[pos+1:pos+5]...)
			pos += 5
		}
		value = bitsToDecimal(valueBits)
		//fmt.Println("value:", value)
	} else {
		lengthTypeID := bits[pos]
		pos++
		//fmt.Println("lengthTypeID:", lengthTypeID)

		// operator
		subValues := make([]int, 0)
		var subValue int

		if lengthTypeID == 0 {
			// the next 15 bits are a number that represents the total
			// length in bits of the sub-packets contained by this packet
			lengthBits := bits[pos : pos+15]
			pos += 15
			length := bitsToDecimal(lengthBits)
			//fmt.Println("subpackets length bits:", length)

			endPos := pos + length
			for {
				if pos >= endPos {
					break
				}

				pos, subValue = readPacket(bits, pos)
				subValues = append(subValues, subValue)
			}
		} else {
			// the next 11 bits are a number that represents the number
			// of sub-packets immediately contained by this packet
			lengthBits := bits[pos : pos+11]
			pos += 11
			length := bitsToDecimal(lengthBits)
			//fmt.Println("subpackets count:", length)
			for i := 0; i < length; i++ {
				pos, subValue = readPacket(bits, pos)
				subValues = append(subValues, subValue)
			}
		}

		if typeID == 0 {
			value = 0
			for _, subValue := range subValues {
				value += subValue
			}
		} else if typeID == 1 {
			value = 1
			for _, subValue := range subValues {
				value *= subValue
			}
		} else if typeID == 2 {
			value = subValues[0]
			for i := 1; i < len(subValues); i++ {
				if subValues[i] < value {
					value = subValues[i]
				}
			}
		} else if typeID == 3 {
			value = subValues[0]
			for i := 1; i < len(subValues); i++ {
				if subValues[i] > value {
					value = subValues[i]
				}
			}
		} else if typeID == 5 {
			if len(subValues) != 2 {
				log.Fatal("bad greater than packet")
			}

			if subValues[0] > subValues[1] {
				value = 1
			} else {
				value = 0
			}
		} else if typeID == 6 {
			if len(subValues) != 2 {
				log.Fatal("bad less than packet")
			}

			if subValues[0] < subValues[1] {
				value = 1
			} else {
				value = 0
			}
		} else if typeID == 7 {
			if len(subValues) != 2 {
				log.Fatal("bad equal to packet")
			}

			if subValues[0] == subValues[1] {
				value = 1
			} else {
				value = 0
			}
		}
	}
	return pos, value
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Can't read input")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var line string
	for scanner.Scan() {
		line = scanner.Text()
		bits := parseHex(line)
		fmt.Println(line)
		_, value := readPacket(bits, 0)
		fmt.Println("value:", value)
	}

}
