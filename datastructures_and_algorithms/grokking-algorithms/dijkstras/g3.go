package main

// g3
//              20
//         A________*C
//        * *       / \
//     10/   \     /1  \30
//      /     \1  /     *
// start       \ *       FINISH
//              B
//

var g3 = graph{
	"START": map[string]int{
		"A": 10,
	},
	"A": map[string]int{
		"C": 20,
	},
	"B": map[string]int{
		"A": 1,
	},
	"C": map[string]int{
		"B": 1,
		"FINISH": 30,
	},
	"FINISH": map[string]int{},
}
