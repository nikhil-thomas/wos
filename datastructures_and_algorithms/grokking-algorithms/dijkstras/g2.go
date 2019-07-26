package main

// g2
//              4
//         A_______*C
//        **\       |\
//      5/ | \      | \3
//      /  |  \     |  *
// start  8|   \2   |6  FINISH
//      \  |    \   |   *
//      2\ |     \  |  /
//        *|      * * /1
//         B_______*D/
//               7

var g2 = graph{
	"START": map[string]int{
		"A": 5,
		"B": 2,
	},
	"A": map[string]int{
		"C": 4,
		"D": 2,
	},
	"B": map[string]int{
		"D": 7,
	},
	"C": map[string]int{
		"D": 6,
		"FINISH": 3,
	},
	"D": map[string]int{
		"FINISH": 1,
	},
	"FINISH": map[string]int{},
}
