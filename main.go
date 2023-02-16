package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"sort"
)

const (
	NumberOfUser  = 74
	DayOfSchedule = 45
)

type period struct {
	p string
}

var planHistory = make(map[int]map[period]int, NumberOfUser)
var planning = make(map[int]map[period][]int, DayOfSchedule)
var userFill []int

var periods = []period{period{"0-2"}, period{"6-8"}, period{"4-6"},
	period{"14-16"}, period{"20-22"}, period{"22-24"}, period{"12-14"},
	period{"2-4"}, period{"16-18"}, period{"18-20"}, period{"8-10"}, period{"10-12"},
}

var exceptionCode = []int{1, 3, 5, 6, 7, 9, 15, 25, 34, 37, 39, 43, 44, 48, 52, 53, 54, 63, 73}

func init() {
	for i := 1; i <= NumberOfUser; i++ {
		planHistory[i] = make(map[period]int)
		for j := 0; j < len(periods); j++ {
			planHistory[i][periods[j]] = 0
		}
	}
	for i := 1; i <= DayOfSchedule; i++ {
		planning[i] = make(map[period][]int)
	}
}

func findMin(a map[period]int) (minKeys []period, min int) {
	minKeys = make([]period, 0, len(a))
	min = a[periods[0]]

	for _, value := range a {
		if value < min {
			min = value
		}
	}

	for key, value := range a {
		if value == min {
			minKeys = append(minKeys, key)
		}
	}

	if len(minKeys) == 0 {
		minKeys = periods
	}

	return minKeys, min
}

func findMinimal(a map[period]int) (minKeys []period, min int) {
	minKeys = make([]period, 0, len(a))
	min = a[periods[0]]

	for _, value := range a {
		if value < min-1 {
			min = value + 1
		}
	}

	for key, value := range a {
		if value == min {
			minKeys = append(minKeys, key)
		}
	}

	if len(minKeys) == 0 {
		minKeys = periods
	}

	return minKeys, min
}

func assert(i, j int, p period) bool {
	if !slices.Contains(exceptionCode, j) {
		minKeys, _ := findMin(planHistory[j])
		minimalKeys, _ := findMinimal(planHistory[j])
		if slices.Contains(minKeys, p) && checkHistoryOfExistSentryTodayAndYesterday(i, j) {
			planning[i][p] = append(planning[i][p], j)
			planHistory[j][p]++
			return true
		} else if slices.Contains(minimalKeys, p) && checkHistoryOfExistSentryTodayAndYesterday(i, j) {
			planning[i][p] = append(planning[i][p], j)
			planHistory[j][p]++
			return true
		} else {
			return false
		}
	}
	return false
}

func checkHistoryOfExistSentryTodayAndYesterday(i, j int) bool {
	for _, p := range periods {
		if slices.Contains(planning[j][p], i) {
			return false
		}
	}
	return true
}

func main() {
	generateHistory()
	j := 1
	for i := 1; i < DayOfSchedule; i++ {
		for _, p := range periods {
			pCounter := 0
			for {
				userFillSuccess := false
				if len(userFill) > 0 {
					for u := 0; u < len(userFill); u++ {
						if assert(i, userFill[u], p) {
							pCounter++
							copy(userFill[u:], userFill[u+1:])
							userFill[len(userFill)-1] = 0
							userFill = userFill[:len(userFill)-1]
							userFillSuccess = true
							if p.p == "0-2" || p.p == "2-4" || p.p == "4-6" || p.p == "22-24" {
								if pCounter > 2 {
									break
								} else {
									continue
								}
							} else {
								if pCounter > 1 {
									break
								} else {
									continue
								}
							}
						}
					}
				}
				if userFillSuccess == true {
					if p.p == "0-2" || p.p == "2-4" || p.p == "4-6" || p.p == "22-24" {
						if pCounter > 2 {
							break
						} else {
							continue
						}
					} else {
						if pCounter > 1 {
							break
						} else {
							continue
						}
					}
				}
				j = (j + 1) % (NumberOfUser + 2)
				if j == 0 {
					j = 2
				}
				if assert(i, j-1, p) {
					pCounter++
					if p.p == "0-2" || p.p == "2-4" || p.p == "4-6" || p.p == "22-24" {
						if pCounter > 2 {
							break
						} else {
							continue
						}
					} else {
						if pCounter > 1 {
							break
						} else {
							continue
						}
					}
				} else {
					if !slices.Contains(exceptionCode, j-1) {
						userFill = append(userFill, j-1)
					}
				}
			}
		}
	}

	keys := make([]int, 0, len(planning))
	for key, _ := range planning {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for key, _ := range keys {
		fmt.Println("day : ", key)
		for k, v := range planning[key] {
			fmt.Print("time : ", k)
			fmt.Println(" codes : ", v)
		}
	}

	keysH := make([]int, 0, len(planning))
	for key, _ := range planHistory {
		keysH = append(keysH, key)
	}
	sort.Ints(keysH)
	for _, key := range keysH {
		fmt.Println("Code History : ", key)
		fmt.Println(planHistory[key])
	}

}

func generateHistory() {

	// Code 1
	//planHistory[1][periods[0]] = 1
	//planHistory[1][periods[1]] = 1
	//planHistory[1][periods[2]] = 0
	//planHistory[1][periods[3]] = 1
	//planHistory[1][periods[4]] = 0
	//planHistory[1][periods[5]] = 0
	//planHistory[1][periods[6]] = 1
	//planHistory[1][periods[7]] = 1
	//planHistory[1][periods[8]] = 0
	//planHistory[1][periods[9]] = 1
	//planHistory[1][periods[10]] = 0
	//planHistory[1][periods[11]] = 1

}
