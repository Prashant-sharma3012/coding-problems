// Run-length encoding is a fast and simple method of encoding strings. The basic idea is to represent repeated successive characters as a single count and character. For example, the string "AAAABBBCCDAA" would be encoded as "4A3B2C1D2A".

// Implement run-length encoding and decoding. You can assume the string to be encoded have no digits and consists solely of alphabetic characters. You can assume the string to be decoded is valid.

package easy

import "strconv"

func RunLengthEncode(input string) string {
	encoded := ""
	count := 1

	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			count++
		} else {
			encoded = encoded + strconv.Itoa(count) + string(input[i])
			count = 1
		}
	}

	// for the last character
	encoded = encoded + strconv.Itoa(count) + string(input[len(input)-1])

	return encoded
}

func RunLengthDecode(input string) string {
	decoded := ""

	for i := 0; i < len(input); i = i + 2 {
		num, _ := strconv.Atoi(string(input[i]))
		for j := 0; j < num; j++ {
			decoded = decoded + string(input[i+1])
		}
	}

	return decoded
}
