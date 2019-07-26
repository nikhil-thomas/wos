package main

// g1
//         A
//        **\
//      5/ | \1
//      /  |  *
// start  3|   finish
//      \  |  *
//      2\ | /5
//        *|/
//         B


var g1 = graph{
	"START": map[string]int{
		"A": 6,
		"B": 2,
	},
	"A": map[string]int{
		"FINISH": 1,
	},
	"B": map[string]int{
		"A": 3,
		"FINISH": 5,
	},
	"FINISH": map[string]int{},
}
