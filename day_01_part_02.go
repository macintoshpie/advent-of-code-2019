// https://play.golang.org/p/BrMtnvTvLOK
package main

import (
	"fmt"
	"strings"
	"strconv"
	"errors"
	"os"
)

// toInts converts a slice of strings into a slice of ints
func toInts(vals []string) ([]int, error) {
	intVals := make([]int, len(vals))
	for i, val := range vals {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return []int{}, errors.New(fmt.Sprintf("Error parsing int for index %v: %v", i, err.Error()))
		}
		intVals[i] = intVal
	}
	return intVals, nil
}

// fuelRequired calculates the amount of fuel required for a given mass
func fuelRequired(mass int) int {
	fuel := mass / 3 - 2	
	if fuel < 0 {
		return 0	
	}
	
	return fuel
}

// moduleFuelRequired calculates the total fuel required for a module, fuel for that fuel, etc
func moduleFuelRequired(mass int) int {
	totalFuelRequired := 0
	lastAmountRequired := fuelRequired(mass)
	for lastAmountRequired > 0 {
		totalFuelRequired += lastAmountRequired		
		lastAmountRequired = fuelRequired(lastAmountRequired)
	}	

	return totalFuelRequired
}

// main calculates the total fuel required for the given input
func main() {
	vals := strings.Split(input, "\n")
	
	masses, err := toInts(vals)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	
	total := 0
	for _, mass := range masses {
		total += moduleFuelRequired(mass)
	}
	
	fmt.Println(total)
}

const input = `140403
114038
56226
132100
62000
111395
91372
126850
145044
79273
91088
84429
58971
107626
149678
85268
105251
115850
115947
74982
75008
97761
121022
148319
125244
138640
86968
144443
137218
139558
128776
53593
133805
64245
113120
63085
59209
51671
63956
139163
119501
77432
51040
137313
58973
64708
76505
108041
101124
133219
95907
57933
117791
76209
102960
90848
141969
91297
146254
84585
103447
83172
76648
111340
118543
52957
86004
131965
90898
90909
52217
144674
97058
72387
57962
147792
114025
100193
77582
146708
54283
143979
99582
149890
73229
56045
63240
124091
103324
125187
74027
120344
105333
100939
131454
109570
149398
140535
57379
138385`
