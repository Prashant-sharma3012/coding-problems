package easy

import "fmt"

var numOfOperation = 2
var inputStr = "10101110011"

// O(n*m) where n is number of operations and m is length of string
func RemoveLongestOccurence() {
	var currStart, currEnd, removeFrom, removeTill int64

	for i := 0; i < numOfOperation; i++ {
		if len(inputStr) == 0 {
			break
		}

		for indx, c := range inputStr {
			// reset pointers
			if indx == 0 {
				currStart = int64(indx)
				currEnd = int64(indx)
				removeFrom = int64(indx)
				removeTill = int64(indx)
			} else {
				if string(c) == string(inputStr[indx-1]) {
					currEnd = int64(indx)
					// reached end
					if len(inputStr) == indx+1 && (removeTill-removeFrom) < (currEnd-currStart) {
						removeFrom = currStart
						removeTill = currEnd
					}
				} else {
					if (removeTill - removeFrom) < (currEnd - currStart) {
						removeFrom = currStart
						removeTill = currEnd
					}
					currStart = int64(indx)
				}
			}
		}

		inputStr = inputStr[:removeFrom] + inputStr[removeTill+1:]
	}

	if len(inputStr) == 0 {
		fmt.Println("KHALI")
	} else {
		fmt.Println(inputStr)
	}

}
