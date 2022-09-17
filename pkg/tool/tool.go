package tool

import (
	"fmt"
	"os"
)

func IsExist_in_Arr(ele interface{}, arr interface{}) bool {
	switch t := arr.(type) {
	case []int:
		v := arr.([]int)

		for _, com := range v {
			if ele == com {
				return true
			}

		}

	case []string:

		v := arr.([]string)

		for _, com := range v {
			if ele == com {
				return true
			}

		}
	// case pq.StringArray:

	// 	v := arr.(pq.StringArray)

	// 	for _, com := range v {
	// 		if ele == com {
	// 			return true
	// 		}

	// 	}
	default:
		fmt.Printf("type is not match %v\n", t)
		os.Exit(1)
	}

	return false
}

func Find_Min_and_Max(arr []int) (min, Max int) {
	min = arr[0]
	Max = arr[0]
	for _, e := range arr {
		if e < min {
			min = e
		}
		if e > Max {
			Max = e
		}
	}
	return min, Max
}
