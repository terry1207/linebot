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
