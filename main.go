package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

const (
	NumberOfUser  = 74
	DayOfSchedule = 18
)

type period struct {
	p string
}

var planHistory = make(map[int]map[period]int, NumberOfUser)
var planning = make(map[int]map[period][]int, DayOfSchedule)
var userFill []int

var periods = []period{period{"0-2"}, period{"2-4"}, period{"4-6"},
	period{"6-8"}, period{"8-10"}, period{"10-12"}, period{"12-14"},
	period{"14-16"}, period{"16-18"}, period{"18-20"}, period{"20-22"}, period{"22-24"},
}

var exceptionCode = []int{1, 3, 5, 6, 7, 15, 25, 34, 37, 39, 43, 44, 48, 52, 53, 54, 63}

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

func assert(i, j int, p period) bool {
	if !slices.Contains(exceptionCode, j) {
		minKeys, _ := findMin(planHistory[j])
		if slices.Contains(minKeys, p) && checkHistoryOfExistSentryTodayAndYesterday(i, j) {
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
							if p.p == "0-2" || p.p == "2-4" || p.p == "4-6" {
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
					if p.p == "0-2" || p.p == "2-4" || p.p == "4-6" {
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
					if p.p == "0-2" || p.p == "2-4" || p.p == "4-6" {
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

	for key, value := range planning {
		fmt.Println("day : ", key)
		for k, v := range value {
			fmt.Print("time : ", k)
			fmt.Println(" codes : ", v)
		}
	}

	for key, value := range planHistory {
		fmt.Println(key)
		fmt.Println(value)
	}

}

func generateHistory() {

	// Code 1
	planHistory[1][periods[0]] = 1
	planHistory[1][periods[1]] = 1
	planHistory[1][periods[2]] = 0
	planHistory[1][periods[3]] = 1
	planHistory[1][periods[4]] = 0
	planHistory[1][periods[5]] = 0
	planHistory[1][periods[6]] = 1
	planHistory[1][periods[7]] = 1
	planHistory[1][periods[8]] = 0
	planHistory[1][periods[9]] = 1
	planHistory[1][periods[10]] = 0
	planHistory[1][periods[11]] = 1

	// Code 2
	planHistory[2][periods[0]] = 2
	planHistory[2][periods[1]] = 1
	planHistory[2][periods[2]] = 1
	planHistory[2][periods[3]] = 1
	planHistory[2][periods[4]] = 1
	planHistory[2][periods[5]] = 0
	planHistory[2][periods[6]] = 1
	planHistory[2][periods[7]] = 2
	planHistory[2][periods[8]] = 0
	planHistory[2][periods[9]] = 0
	planHistory[2][periods[10]] = 2
	planHistory[2][periods[11]] = 1

	// Code 3
	planHistory[3][periods[0]] = 0
	planHistory[3][periods[1]] = 0
	planHistory[3][periods[2]] = 0
	planHistory[3][periods[3]] = 0
	planHistory[3][periods[4]] = 0
	planHistory[3][periods[5]] = 0
	planHistory[3][periods[6]] = 0
	planHistory[3][periods[7]] = 0
	planHistory[3][periods[8]] = 0
	planHistory[3][periods[9]] = 0
	planHistory[3][periods[10]] = 0
	planHistory[3][periods[11]] = 0

	// Code 4
	planHistory[4][periods[0]] = 1
	planHistory[4][periods[1]] = 1
	planHistory[4][periods[2]] = 1
	planHistory[4][periods[3]] = 1
	planHistory[4][periods[4]] = 1
	planHistory[4][periods[5]] = 0
	planHistory[4][periods[6]] = 0
	planHistory[4][periods[7]] = 2
	planHistory[4][periods[8]] = 1
	planHistory[4][periods[9]] = 1
	planHistory[4][periods[10]] = 0
	planHistory[4][periods[11]] = 2

	// Code 5
	planHistory[5][periods[0]] = 0
	planHistory[5][periods[1]] = 0
	planHistory[5][periods[2]] = 0
	planHistory[5][periods[3]] = 0
	planHistory[5][periods[4]] = 0
	planHistory[5][periods[5]] = 0
	planHistory[5][periods[6]] = 0
	planHistory[5][periods[7]] = 0
	planHistory[5][periods[8]] = 0
	planHistory[5][periods[9]] = 0
	planHistory[5][periods[10]] = 0
	planHistory[5][periods[11]] = 0

	// Code 6
	planHistory[6][periods[0]] = 0
	planHistory[6][periods[1]] = 0
	planHistory[6][periods[2]] = 0
	planHistory[6][periods[3]] = 0
	planHistory[6][periods[4]] = 0
	planHistory[6][periods[5]] = 1
	planHistory[6][periods[6]] = 0
	planHistory[6][periods[7]] = 0
	planHistory[6][periods[8]] = 0
	planHistory[6][periods[9]] = 0
	planHistory[6][periods[10]] = 0
	planHistory[6][periods[11]] = 0

	// Code 7
	planHistory[7][periods[0]] = 1
	planHistory[7][periods[1]] = 1
	planHistory[7][periods[2]] = 1
	planHistory[7][periods[3]] = 0
	planHistory[7][periods[4]] = 0
	planHistory[7][periods[5]] = 1
	planHistory[7][periods[6]] = 0
	planHistory[7][periods[7]] = 0
	planHistory[7][periods[8]] = 3
	planHistory[7][periods[9]] = 0
	planHistory[7][periods[10]] = 0
	planHistory[7][periods[11]] = 0

	// Code 8
	planHistory[8][periods[0]] = 1
	planHistory[8][periods[1]] = 1
	planHistory[8][periods[2]] = 1
	planHistory[8][periods[3]] = 0
	planHistory[8][periods[4]] = 1
	planHistory[8][periods[5]] = 0
	planHistory[8][periods[6]] = 1
	planHistory[8][periods[7]] = 0
	planHistory[8][periods[8]] = 1
	planHistory[8][periods[9]] = 2
	planHistory[8][periods[10]] = 1
	planHistory[8][periods[11]] = 1

	// Code 9
	planHistory[9][periods[0]] = 1
	planHistory[9][periods[1]] = 1
	planHistory[9][periods[2]] = 2
	planHistory[9][periods[3]] = 0
	planHistory[9][periods[4]] = 0
	planHistory[9][periods[5]] = 0
	planHistory[9][periods[6]] = 1
	planHistory[9][periods[7]] = 0
	planHistory[9][periods[8]] = 0
	planHistory[9][periods[9]] = 2
	planHistory[9][periods[10]] = 2
	planHistory[9][periods[11]] = 1

	// Code 10
	planHistory[10][periods[0]] = 1
	planHistory[10][periods[1]] = 1
	planHistory[10][periods[2]] = 2
	planHistory[10][periods[3]] = 1
	planHistory[10][periods[4]] = 0
	planHistory[10][periods[5]] = 0
	planHistory[10][periods[6]] = 0
	planHistory[10][periods[7]] = 0
	planHistory[10][periods[8]] = 0
	planHistory[10][periods[9]] = 1
	planHistory[10][periods[10]] = 2
	planHistory[10][periods[11]] = 2

	// Code 11
	planHistory[11][periods[0]] = 1
	planHistory[11][periods[1]] = 1
	planHistory[11][periods[2]] = 1
	planHistory[11][periods[3]] = 1
	planHistory[11][periods[4]] = 1
	planHistory[11][periods[5]] = 0
	planHistory[11][periods[6]] = 1
	planHistory[11][periods[7]] = 0
	planHistory[11][periods[8]] = 0
	planHistory[11][periods[9]] = 0
	planHistory[11][periods[10]] = 1
	planHistory[11][periods[11]] = 3

	// Code 12
	planHistory[12][periods[0]] = 0
	planHistory[12][periods[1]] = 1
	planHistory[12][periods[2]] = 1
	planHistory[12][periods[3]] = 1
	planHistory[12][periods[4]] = 2
	planHistory[12][periods[5]] = 0
	planHistory[12][periods[6]] = 0
	planHistory[12][periods[7]] = 0
	planHistory[12][periods[8]] = 1
	planHistory[12][periods[9]] = 1
	planHistory[12][periods[10]] = 1
	planHistory[12][periods[11]] = 3

	// Code 13
	planHistory[13][periods[0]] = 3
	planHistory[13][periods[1]] = 2
	planHistory[13][periods[2]] = 0
	planHistory[13][periods[3]] = 1
	planHistory[13][periods[4]] = 0
	planHistory[13][periods[5]] = 0
	planHistory[13][periods[6]] = 0
	planHistory[13][periods[7]] = 1
	planHistory[13][periods[8]] = 0
	planHistory[13][periods[9]] = 1
	planHistory[13][periods[10]] = 1
	planHistory[13][periods[11]] = 2

	// Code 14
	planHistory[14][periods[0]] = 2
	planHistory[14][periods[1]] = 1
	planHistory[14][periods[2]] = 1
	planHistory[14][periods[3]] = 1
	planHistory[14][periods[4]] = 1
	planHistory[14][periods[5]] = 0
	planHistory[14][periods[6]] = 0
	planHistory[14][periods[7]] = 1
	planHistory[14][periods[8]] = 0
	planHistory[14][periods[9]] = 1
	planHistory[14][periods[10]] = 1
	planHistory[14][periods[11]] = 1

	// Code 15
	planHistory[15][periods[0]] = 0
	planHistory[15][periods[1]] = 0
	planHistory[15][periods[2]] = 0
	planHistory[15][periods[3]] = 0
	planHistory[15][periods[4]] = 0
	planHistory[15][periods[5]] = 0
	planHistory[15][periods[6]] = 0
	planHistory[15][periods[7]] = 0
	planHistory[15][periods[8]] = 0
	planHistory[15][periods[9]] = 0
	planHistory[15][periods[10]] = 0
	planHistory[15][periods[11]] = 0

	// Code 16
	planHistory[16][periods[0]] = 2
	planHistory[16][periods[1]] = 1
	planHistory[16][periods[2]] = 1
	planHistory[16][periods[3]] = 2
	planHistory[16][periods[4]] = 2
	planHistory[16][periods[5]] = 0
	planHistory[16][periods[6]] = 0
	planHistory[16][periods[7]] = 0
	planHistory[16][periods[8]] = 1
	planHistory[16][periods[9]] = 1
	planHistory[16][periods[10]] = 1
	planHistory[16][periods[11]] = 1

	// Code 17
	planHistory[17][periods[0]] = 1
	planHistory[17][periods[1]] = 1
	planHistory[17][periods[2]] = 1
	planHistory[17][periods[3]] = 1
	planHistory[17][periods[4]] = 0
	planHistory[17][periods[5]] = 1
	planHistory[17][periods[6]] = 1
	planHistory[17][periods[7]] = 1
	planHistory[17][periods[8]] = 1
	planHistory[17][periods[9]] = 1
	planHistory[17][periods[10]] = 1
	planHistory[17][periods[11]] = 1

	// Code 18
	planHistory[18][periods[0]] = 1
	planHistory[18][periods[1]] = 2
	planHistory[18][periods[2]] = 1
	planHistory[18][periods[3]] = 1
	planHistory[18][periods[4]] = 1
	planHistory[18][periods[5]] = 2
	planHistory[18][periods[6]] = 0
	planHistory[18][periods[7]] = 0
	planHistory[18][periods[8]] = 1
	planHistory[18][periods[9]] = 1
	planHistory[18][periods[10]] = 1
	planHistory[18][periods[11]] = 1

	// Code 19
	planHistory[19][periods[0]] = 0
	planHistory[19][periods[1]] = 3
	planHistory[19][periods[2]] = 0
	planHistory[19][periods[3]] = 2
	planHistory[19][periods[4]] = 1
	planHistory[19][periods[5]] = 1
	planHistory[19][periods[6]] = 0
	planHistory[19][periods[7]] = 0
	planHistory[19][periods[8]] = 0
	planHistory[19][periods[9]] = 1
	planHistory[19][periods[10]] = 2
	planHistory[19][periods[11]] = 0

	// Code 20
	planHistory[20][periods[0]] = 1
	planHistory[20][periods[1]] = 1
	planHistory[20][periods[2]] = 2
	planHistory[20][periods[3]] = 1
	planHistory[20][periods[4]] = 1
	planHistory[20][periods[5]] = 1
	planHistory[20][periods[6]] = 3
	planHistory[20][periods[7]] = 0
	planHistory[20][periods[8]] = 0
	planHistory[20][periods[9]] = 0
	planHistory[20][periods[10]] = 1
	planHistory[20][periods[11]] = 1

	// Code 21
	planHistory[21][periods[0]] = 1
	planHistory[21][periods[1]] = 1
	planHistory[21][periods[2]] = 3
	planHistory[21][periods[3]] = 1
	planHistory[21][periods[4]] = 1
	planHistory[21][periods[5]] = 1
	planHistory[21][periods[6]] = 3
	planHistory[21][periods[7]] = 0
	planHistory[21][periods[8]] = 1
	planHistory[21][periods[9]] = 0
	planHistory[21][periods[10]] = 0
	planHistory[21][periods[11]] = 0

	// Code 22
	planHistory[22][periods[0]] = 0
	planHistory[22][periods[1]] = 0
	planHistory[22][periods[2]] = 2
	planHistory[22][periods[3]] = 1
	planHistory[22][periods[4]] = 2
	planHistory[22][periods[5]] = 1
	planHistory[22][periods[6]] = 1
	planHistory[22][periods[7]] = 2
	planHistory[22][periods[8]] = 0
	planHistory[22][periods[9]] = 1
	planHistory[22][periods[10]] = 0
	planHistory[22][periods[11]] = 1

	// Code 23
	planHistory[23][periods[0]] = 1
	planHistory[23][periods[1]] = 1
	planHistory[23][periods[2]] = 1
	planHistory[23][periods[3]] = 1
	planHistory[23][periods[4]] = 1
	planHistory[23][periods[5]] = 1
	planHistory[23][periods[6]] = 1
	planHistory[23][periods[7]] = 2
	planHistory[23][periods[8]] = 0
	planHistory[23][periods[9]] = 0
	planHistory[23][periods[10]] = 1
	planHistory[23][periods[11]] = 2

	// Code 24
	planHistory[24][periods[0]] = 2
	planHistory[24][periods[1]] = 1
	planHistory[24][periods[2]] = 1
	planHistory[24][periods[3]] = 1
	planHistory[24][periods[4]] = 0
	planHistory[24][periods[5]] = 2
	planHistory[24][periods[6]] = 1
	planHistory[24][periods[7]] = 1
	planHistory[24][periods[8]] = 1
	planHistory[24][periods[9]] = 0
	planHistory[24][periods[10]] = 0
	planHistory[24][periods[11]] = 1

	// Code 25
	planHistory[25][periods[0]] = 0
	planHistory[25][periods[1]] = 0
	planHistory[25][periods[2]] = 0
	planHistory[25][periods[3]] = 0
	planHistory[25][periods[4]] = 0
	planHistory[25][periods[5]] = 0
	planHistory[25][periods[6]] = 0
	planHistory[25][periods[7]] = 0
	planHistory[25][periods[8]] = 0
	planHistory[25][periods[9]] = 0
	planHistory[25][periods[10]] = 0
	planHistory[25][periods[11]] = 0

	// Code 26
	planHistory[26][periods[0]] = 2
	planHistory[26][periods[1]] = 0
	planHistory[26][periods[2]] = 1
	planHistory[26][periods[3]] = 1
	planHistory[26][periods[4]] = 1
	planHistory[26][periods[5]] = 1
	planHistory[26][periods[6]] = 1
	planHistory[26][periods[7]] = 1
	planHistory[26][periods[8]] = 1
	planHistory[26][periods[9]] = 0
	planHistory[26][periods[10]] = 1
	planHistory[26][periods[11]] = 1

	// Code 27
	planHistory[27][periods[0]] = 2
	planHistory[27][periods[1]] = 1
	planHistory[27][periods[2]] = 0
	planHistory[27][periods[3]] = 0
	planHistory[27][periods[4]] = 1
	planHistory[27][periods[5]] = 1
	planHistory[27][periods[6]] = 2
	planHistory[27][periods[7]] = 0
	planHistory[27][periods[8]] = 1
	planHistory[27][periods[9]] = 1
	planHistory[27][periods[10]] = 0
	planHistory[27][periods[11]] = 1

	// Code 28
	planHistory[28][periods[0]] = 1
	planHistory[28][periods[1]] = 2
	planHistory[28][periods[2]] = 0
	planHistory[28][periods[3]] = 0
	planHistory[28][periods[4]] = 1
	planHistory[28][periods[5]] = 1
	planHistory[28][periods[6]] = 1
	planHistory[28][periods[7]] = 1
	planHistory[28][periods[8]] = 1
	planHistory[28][periods[9]] = 1
	planHistory[28][periods[10]] = 1
	planHistory[28][periods[11]] = 1

	// Code 29
	planHistory[29][periods[0]] = 1
	planHistory[29][periods[1]] = 2
	planHistory[29][periods[2]] = 1
	planHistory[29][periods[3]] = 0
	planHistory[29][periods[4]] = 1
	planHistory[29][periods[5]] = 0
	planHistory[29][periods[6]] = 1
	planHistory[29][periods[7]] = 1
	planHistory[29][periods[8]] = 0
	planHistory[29][periods[9]] = 1
	planHistory[29][periods[10]] = 1
	planHistory[29][periods[11]] = 1

	// Code 30
	planHistory[30][periods[0]] = 1
	planHistory[30][periods[1]] = 2
	planHistory[30][periods[2]] = 1
	planHistory[30][periods[3]] = 1
	planHistory[30][periods[4]] = 0
	planHistory[30][periods[5]] = 1
	planHistory[30][periods[6]] = 0
	planHistory[30][periods[7]] = 2
	planHistory[30][periods[8]] = 0
	planHistory[30][periods[9]] = 1
	planHistory[30][periods[10]] = 1
	planHistory[30][periods[11]] = 0

	// Code 31
	planHistory[31][periods[0]] = 1
	planHistory[31][periods[1]] = 2
	planHistory[31][periods[2]] = 1
	planHistory[31][periods[3]] = 0
	planHistory[31][periods[4]] = 0
	planHistory[31][periods[5]] = 1
	planHistory[31][periods[6]] = 0
	planHistory[31][periods[7]] = 1
	planHistory[31][periods[8]] = 2
	planHistory[31][periods[9]] = 0
	planHistory[31][periods[10]] = 1
	planHistory[31][periods[11]] = 1

	// Code 32
	planHistory[32][periods[0]] = 0
	planHistory[32][periods[1]] = 1
	planHistory[32][periods[2]] = 2
	planHistory[32][periods[3]] = 0
	planHistory[32][periods[4]] = 0
	planHistory[32][periods[5]] = 1
	planHistory[32][periods[6]] = 1
	planHistory[32][periods[7]] = 1
	planHistory[32][periods[8]] = 2
	planHistory[32][periods[9]] = 0
	planHistory[32][periods[10]] = 1
	planHistory[32][periods[11]] = 1

	// Code 33
	planHistory[33][periods[0]] = 0
	planHistory[33][periods[1]] = 1
	planHistory[33][periods[2]] = 2
	planHistory[33][periods[3]] = 1
	planHistory[33][periods[4]] = 1
	planHistory[33][periods[5]] = 1
	planHistory[33][periods[6]] = 1
	planHistory[33][periods[7]] = 0
	planHistory[33][periods[8]] = 0
	planHistory[33][periods[9]] = 2
	planHistory[33][periods[10]] = 0
	planHistory[33][periods[11]] = 1

	// Code 34
	planHistory[34][periods[0]] = 0
	planHistory[34][periods[1]] = 0
	planHistory[34][periods[2]] = 0
	planHistory[34][periods[3]] = 0
	planHistory[34][periods[4]] = 0
	planHistory[34][periods[5]] = 0
	planHistory[34][periods[6]] = 0
	planHistory[34][periods[7]] = 0
	planHistory[34][periods[8]] = 0
	planHistory[34][periods[9]] = 0
	planHistory[34][periods[10]] = 0
	planHistory[34][periods[11]] = 0

	// Code 35
	planHistory[35][periods[0]] = 1
	planHistory[35][periods[1]] = 1
	planHistory[35][periods[2]] = 2
	planHistory[35][periods[3]] = 1
	planHistory[35][periods[4]] = 1
	planHistory[35][periods[5]] = 0
	planHistory[35][periods[6]] = 0
	planHistory[35][periods[7]] = 1
	planHistory[35][periods[8]] = 1
	planHistory[35][periods[9]] = 0
	planHistory[35][periods[10]] = 1
	planHistory[35][periods[11]] = 1

	// Code 36
	planHistory[36][periods[0]] = 1
	planHistory[36][periods[1]] = 1
	planHistory[36][periods[2]] = 1
	planHistory[36][periods[3]] = 1
	planHistory[36][periods[4]] = 1
	planHistory[36][periods[5]] = 1
	planHistory[36][periods[6]] = 0
	planHistory[36][periods[7]] = 1
	planHistory[36][periods[8]] = 2
	planHistory[36][periods[9]] = 0
	planHistory[36][periods[10]] = 1
	planHistory[36][periods[11]] = 1

	// Code 37
	planHistory[37][periods[0]] = 0
	planHistory[37][periods[1]] = 0
	planHistory[37][periods[2]] = 0
	planHistory[37][periods[3]] = 0
	planHistory[37][periods[4]] = 0
	planHistory[37][periods[5]] = 0
	planHistory[37][periods[6]] = 0
	planHistory[37][periods[7]] = 0
	planHistory[37][periods[8]] = 0
	planHistory[37][periods[9]] = 0
	planHistory[37][periods[10]] = 0
	planHistory[37][periods[11]] = 0

	// Code 38
	planHistory[38][periods[0]] = 1
	planHistory[38][periods[1]] = 1
	planHistory[38][periods[2]] = 1
	planHistory[38][periods[3]] = 1
	planHistory[38][periods[4]] = 1
	planHistory[38][periods[5]] = 1
	planHistory[38][periods[6]] = 0
	planHistory[38][periods[7]] = 0
	planHistory[38][periods[8]] = 2
	planHistory[38][periods[9]] = 0
	planHistory[38][periods[10]] = 0
	planHistory[38][periods[11]] = 2

	// Code 39
	planHistory[39][periods[0]] = 0
	planHistory[39][periods[1]] = 0
	planHistory[39][periods[2]] = 0
	planHistory[39][periods[3]] = 0
	planHistory[39][periods[4]] = 0
	planHistory[39][periods[5]] = 0
	planHistory[39][periods[6]] = 0
	planHistory[39][periods[7]] = 0
	planHistory[39][periods[8]] = 0
	planHistory[39][periods[9]] = 0
	planHistory[39][periods[10]] = 0
	planHistory[39][periods[11]] = 0

	// Code 40
	planHistory[40][periods[0]] = 1
	planHistory[40][periods[1]] = 1
	planHistory[40][periods[2]] = 1
	planHistory[40][periods[3]] = 1
	planHistory[40][periods[4]] = 1
	planHistory[40][periods[5]] = 0
	planHistory[40][periods[6]] = 1
	planHistory[40][periods[7]] = 0
	planHistory[40][periods[8]] = 1
	planHistory[40][periods[9]] = 0
	planHistory[40][periods[10]] = 1
	planHistory[40][periods[11]] = 2

	// Code 41
	planHistory[41][periods[0]] = 1
	planHistory[41][periods[1]] = 1
	planHistory[41][periods[2]] = 1
	planHistory[41][periods[3]] = 1
	planHistory[41][periods[4]] = 1
	planHistory[41][periods[5]] = 1
	planHistory[41][periods[6]] = 0
	planHistory[41][periods[7]] = 0
	planHistory[41][periods[8]] = 0
	planHistory[41][periods[9]] = 1
	planHistory[41][periods[10]] = 0
	planHistory[41][periods[11]] = 1

	// Code 42
	planHistory[42][periods[0]] = 1
	planHistory[42][periods[1]] = 2
	planHistory[42][periods[2]] = 0
	planHistory[42][periods[3]] = 0
	planHistory[42][periods[4]] = 0
	planHistory[42][periods[5]] = 0
	planHistory[42][periods[6]] = 0
	planHistory[42][periods[7]] = 0
	planHistory[42][periods[8]] = 0
	planHistory[42][periods[9]] = 1
	planHistory[42][periods[10]] = 0
	planHistory[42][periods[11]] = 1

	// Code 43
	planHistory[43][periods[0]] = 0
	planHistory[43][periods[1]] = 0
	planHistory[43][periods[2]] = 0
	planHistory[43][periods[3]] = 0
	planHistory[43][periods[4]] = 0
	planHistory[43][periods[5]] = 0
	planHistory[43][periods[6]] = 0
	planHistory[43][periods[7]] = 0
	planHistory[43][periods[8]] = 0
	planHistory[43][periods[9]] = 0
	planHistory[43][periods[10]] = 0
	planHistory[43][periods[11]] = 0

	// Code 44
	planHistory[44][periods[0]] = 0
	planHistory[44][periods[1]] = 0
	planHistory[44][periods[2]] = 0
	planHistory[44][periods[3]] = 0
	planHistory[44][periods[4]] = 0
	planHistory[44][periods[5]] = 0
	planHistory[44][periods[6]] = 0
	planHistory[44][periods[7]] = 0
	planHistory[44][periods[8]] = 0
	planHistory[44][periods[9]] = 0
	planHistory[44][periods[10]] = 0
	planHistory[44][periods[11]] = 0

	// Code 45
	planHistory[45][periods[0]] = 1
	planHistory[45][periods[1]] = 1
	planHistory[45][periods[2]] = 1
	planHistory[45][periods[3]] = 1
	planHistory[45][periods[4]] = 1
	planHistory[45][periods[5]] = 1
	planHistory[45][periods[6]] = 1
	planHistory[45][periods[7]] = 0
	planHistory[45][periods[8]] = 0
	planHistory[45][periods[9]] = 1
	planHistory[45][periods[10]] = 0
	planHistory[45][periods[11]] = 2

	// Code 46
	planHistory[46][periods[0]] = 1
	planHistory[46][periods[1]] = 0
	planHistory[46][periods[2]] = 1
	planHistory[46][periods[3]] = 1
	planHistory[46][periods[4]] = 1
	planHistory[46][periods[5]] = 1
	planHistory[46][periods[6]] = 1
	planHistory[46][periods[7]] = 0
	planHistory[46][periods[8]] = 0
	planHistory[46][periods[9]] = 0
	planHistory[46][periods[10]] = 1
	planHistory[46][periods[11]] = 1

	// Code 47
	planHistory[47][periods[0]] = 1
	planHistory[47][periods[1]] = 0
	planHistory[47][periods[2]] = 1
	planHistory[47][periods[3]] = 1
	planHistory[47][periods[4]] = 1
	planHistory[47][periods[5]] = 1
	planHistory[47][periods[6]] = 1
	planHistory[47][periods[7]] = 0
	planHistory[47][periods[8]] = 0
	planHistory[47][periods[9]] = 0
	planHistory[47][periods[10]] = 1
	planHistory[47][periods[11]] = 2

	// Code 48
	planHistory[48][periods[0]] = 0
	planHistory[48][periods[1]] = 0
	planHistory[48][periods[2]] = 0
	planHistory[48][periods[3]] = 0
	planHistory[48][periods[4]] = 0
	planHistory[48][periods[5]] = 0
	planHistory[48][periods[6]] = 0
	planHistory[48][periods[7]] = 0
	planHistory[48][periods[8]] = 0
	planHistory[48][periods[9]] = 0
	planHistory[48][periods[10]] = 0
	planHistory[48][periods[11]] = 0

	// Code 49
	planHistory[49][periods[0]] = 1
	planHistory[49][periods[1]] = 1
	planHistory[49][periods[2]] = 1
	planHistory[49][periods[3]] = 0
	planHistory[49][periods[4]] = 1
	planHistory[49][periods[5]] = 1
	planHistory[49][periods[6]] = 2
	planHistory[49][periods[7]] = 1
	planHistory[49][periods[8]] = 0
	planHistory[49][periods[9]] = 1
	planHistory[49][periods[10]] = 0
	planHistory[49][periods[11]] = 2

	// Code 50
	planHistory[50][periods[0]] = 1
	planHistory[50][periods[1]] = 1
	planHistory[50][periods[2]] = 0
	planHistory[50][periods[3]] = 0
	planHistory[50][periods[4]] = 1
	planHistory[50][periods[5]] = 0
	planHistory[50][periods[6]] = 1
	planHistory[50][periods[7]] = 2
	planHistory[50][periods[8]] = 1
	planHistory[50][periods[9]] = 1
	planHistory[50][periods[10]] = 0
	planHistory[50][periods[11]] = 2

	// Code 51
	planHistory[51][periods[0]] = 1
	planHistory[51][periods[1]] = 1
	planHistory[51][periods[2]] = 0
	planHistory[51][periods[3]] = 0
	planHistory[51][periods[4]] = 0
	planHistory[51][periods[5]] = 1
	planHistory[51][periods[6]] = 0
	planHistory[51][periods[7]] = 2
	planHistory[51][periods[8]] = 1
	planHistory[51][periods[9]] = 0
	planHistory[51][periods[10]] = 1
	planHistory[51][periods[11]] = 2

	// Code 52
	planHistory[52][periods[0]] = 1
	planHistory[52][periods[1]] = 0
	planHistory[52][periods[2]] = 1
	planHistory[52][periods[3]] = 0
	planHistory[52][periods[4]] = 1
	planHistory[52][periods[5]] = 0
	planHistory[52][periods[6]] = 0
	planHistory[52][periods[7]] = 1
	planHistory[52][periods[8]] = 2
	planHistory[52][periods[9]] = 0
	planHistory[52][periods[10]] = 1
	planHistory[52][periods[11]] = 1

	// Code 53
	planHistory[53][periods[0]] = 0
	planHistory[53][periods[1]] = 0
	planHistory[53][periods[2]] = 0
	planHistory[53][periods[3]] = 0
	planHistory[53][periods[4]] = 0
	planHistory[53][periods[5]] = 0
	planHistory[53][periods[6]] = 0
	planHistory[53][periods[7]] = 1
	planHistory[53][periods[8]] = 0
	planHistory[53][periods[9]] = 0
	planHistory[53][periods[10]] = 0
	planHistory[53][periods[11]] = 0

	// Code 54
	planHistory[54][periods[0]] = 1
	planHistory[54][periods[1]] = 2
	planHistory[54][periods[2]] = 0
	planHistory[54][periods[3]] = 0
	planHistory[54][periods[4]] = 1
	planHistory[54][periods[5]] = 0
	planHistory[54][periods[6]] = 0
	planHistory[54][periods[7]] = 0
	planHistory[54][periods[8]] = 1
	planHistory[54][periods[9]] = 0
	planHistory[54][periods[10]] = 0
	planHistory[54][periods[11]] = 0

	// Code 55
	planHistory[55][periods[0]] = 1
	planHistory[55][periods[1]] = 2
	planHistory[55][periods[2]] = 1
	planHistory[55][periods[3]] = 0
	planHistory[55][periods[4]] = 0
	planHistory[55][periods[5]] = 1
	planHistory[55][periods[6]] = 1
	planHistory[55][periods[7]] = 0
	planHistory[55][periods[8]] = 1
	planHistory[55][periods[9]] = 2
	planHistory[55][periods[10]] = 1
	planHistory[55][periods[11]] = 0

	// Code 56
	planHistory[56][periods[0]] = 0
	planHistory[56][periods[1]] = 2
	planHistory[56][periods[2]] = 1
	planHistory[56][periods[3]] = 0
	planHistory[56][periods[4]] = 0
	planHistory[56][periods[5]] = 1
	planHistory[56][periods[6]] = 1
	planHistory[56][periods[7]] = 0
	planHistory[56][periods[8]] = 1
	planHistory[56][periods[9]] = 2
	planHistory[56][periods[10]] = 1
	planHistory[56][periods[11]] = 0

	// Code 57
	planHistory[57][periods[0]] = 0
	planHistory[57][periods[1]] = 1
	planHistory[57][periods[2]] = 2
	planHistory[57][periods[3]] = 1
	planHistory[57][periods[4]] = 0
	planHistory[57][periods[5]] = 1
	planHistory[57][periods[6]] = 1
	planHistory[57][periods[7]] = 0
	planHistory[57][periods[8]] = 1
	planHistory[57][periods[9]] = 0
	planHistory[57][periods[10]] = 2
	planHistory[57][periods[11]] = 1

	// Code 58
	planHistory[58][periods[0]] = 0
	planHistory[58][periods[1]] = 1
	planHistory[58][periods[2]] = 2
	planHistory[58][periods[3]] = 1
	planHistory[58][periods[4]] = 0
	planHistory[58][periods[5]] = 0
	planHistory[58][periods[6]] = 1
	planHistory[58][periods[7]] = 0
	planHistory[58][periods[8]] = 1
	planHistory[58][periods[9]] = 0
	planHistory[58][periods[10]] = 2
	planHistory[58][periods[11]] = 1

	// Code 59
	planHistory[59][periods[0]] = 0
	planHistory[59][periods[1]] = 0
	planHistory[59][periods[2]] = 3
	planHistory[59][periods[3]] = 2
	planHistory[59][periods[4]] = 1
	planHistory[59][periods[5]] = 0
	planHistory[59][periods[6]] = 0
	planHistory[59][periods[7]] = 1
	planHistory[59][periods[8]] = 1
	planHistory[59][periods[9]] = 1
	planHistory[59][periods[10]] = 0
	planHistory[59][periods[11]] = 1

	// Code 60
	planHistory[60][periods[0]] = 1
	planHistory[60][periods[1]] = 1
	planHistory[60][periods[2]] = 1
	planHistory[60][periods[3]] = 2
	planHistory[60][periods[4]] = 1
	planHistory[60][periods[5]] = 0
	planHistory[60][periods[6]] = 0
	planHistory[60][periods[7]] = 1
	planHistory[60][periods[8]] = 0
	planHistory[60][periods[9]] = 1
	planHistory[60][periods[10]] = 0
	planHistory[60][periods[11]] = 3

	// Code 61
	planHistory[61][periods[0]] = 2
	planHistory[61][periods[1]] = 0
	planHistory[61][periods[2]] = 1
	planHistory[61][periods[3]] = 1
	planHistory[61][periods[4]] = 1
	planHistory[61][periods[5]] = 0
	planHistory[61][periods[6]] = 0
	planHistory[61][periods[7]] = 1
	planHistory[61][periods[8]] = 1
	planHistory[61][periods[9]] = 1
	planHistory[61][periods[10]] = 2
	planHistory[61][periods[11]] = 2

	// Code 62
	planHistory[62][periods[0]] = 2
	planHistory[62][periods[1]] = 0
	planHistory[62][periods[2]] = 0
	planHistory[62][periods[3]] = 1
	planHistory[62][periods[4]] = 1
	planHistory[62][periods[5]] = 1
	planHistory[62][periods[6]] = 0
	planHistory[62][periods[7]] = 0
	planHistory[62][periods[8]] = 1
	planHistory[62][periods[9]] = 1
	planHistory[62][periods[10]] = 2
	planHistory[62][periods[11]] = 1

	// Code 63
	planHistory[63][periods[0]] = 0
	planHistory[63][periods[1]] = 0
	planHistory[63][periods[2]] = 0
	planHistory[63][periods[3]] = 0
	planHistory[63][periods[4]] = 0
	planHistory[63][periods[5]] = 0
	planHistory[63][periods[6]] = 0
	planHistory[63][periods[7]] = 0
	planHistory[63][periods[8]] = 0
	planHistory[63][periods[9]] = 0
	planHistory[63][periods[10]] = 0
	planHistory[63][periods[11]] = 0

	// Code 64
	planHistory[64][periods[0]] = 4
	planHistory[64][periods[1]] = 1
	planHistory[64][periods[2]] = 0
	planHistory[64][periods[3]] = 2
	planHistory[64][periods[4]] = 1
	planHistory[64][periods[5]] = 3
	planHistory[64][periods[6]] = 0
	planHistory[64][periods[7]] = 0
	planHistory[64][periods[8]] = 0
	planHistory[64][periods[9]] = 1
	planHistory[64][periods[10]] = 0
	planHistory[64][periods[11]] = 0

	// Code 65
	planHistory[65][periods[0]] = 3
	planHistory[65][periods[1]] = 2
	planHistory[65][periods[2]] = 0
	planHistory[65][periods[3]] = 1
	planHistory[65][periods[4]] = 1
	planHistory[65][periods[5]] = 1
	planHistory[65][periods[6]] = 0
	planHistory[65][periods[7]] = 2
	planHistory[65][periods[8]] = 0
	planHistory[65][periods[9]] = 0
	planHistory[65][periods[10]] = 0
	planHistory[65][periods[11]] = 0

	// Code 66
	planHistory[66][periods[0]] = 2
	planHistory[66][periods[1]] = 3
	planHistory[66][periods[2]] = 0
	planHistory[66][periods[3]] = 0
	planHistory[66][periods[4]] = 2
	planHistory[66][periods[5]] = 1
	planHistory[66][periods[6]] = 1
	planHistory[66][periods[7]] = 1
	planHistory[66][periods[8]] = 0
	planHistory[66][periods[9]] = 0
	planHistory[6][periods[10]] = 0
	planHistory[66][periods[11]] = 0

	// Code 67
	planHistory[67][periods[0]] = 0
	planHistory[67][periods[1]] = 4
	planHistory[67][periods[2]] = 1
	planHistory[67][periods[3]] = 0
	planHistory[67][periods[4]] = 1
	planHistory[67][periods[5]] = 1
	planHistory[67][periods[6]] = 1
	planHistory[67][periods[7]] = 1
	planHistory[67][periods[8]] = 1
	planHistory[67][periods[9]] = 0
	planHistory[67][periods[10]] = 0
	planHistory[67][periods[11]] = 0

	// Code 68
	planHistory[68][periods[0]] = 0
	planHistory[68][periods[1]] = 2
	planHistory[68][periods[2]] = 3
	planHistory[68][periods[3]] = 0
	planHistory[68][periods[4]] = 1
	planHistory[68][periods[5]] = 2
	planHistory[68][periods[6]] = 0
	planHistory[68][periods[7]] = 1
	planHistory[68][periods[8]] = 1
	planHistory[68][periods[9]] = 0
	planHistory[68][periods[10]] = 0
	planHistory[68][periods[11]] = 0

	// Code 69
	planHistory[69][periods[0]] = 0
	planHistory[69][periods[1]] = 2
	planHistory[69][periods[2]] = 2
	planHistory[69][periods[3]] = 0
	planHistory[69][periods[4]] = 0
	planHistory[69][periods[5]] = 2
	planHistory[69][periods[6]] = 1
	planHistory[69][periods[7]] = 1
	planHistory[69][periods[8]] = 1
	planHistory[69][periods[9]] = 0
	planHistory[69][periods[10]] = 0
	planHistory[69][periods[11]] = 0

	// Code 70
	planHistory[70][periods[0]] = 1
	planHistory[70][periods[1]] = 0
	planHistory[70][periods[2]] = 4
	planHistory[70][periods[3]] = 0
	planHistory[70][periods[4]] = 0
	planHistory[70][periods[5]] = 1
	planHistory[70][periods[6]] = 2
	planHistory[70][periods[7]] = 0
	planHistory[70][periods[8]] = 1
	planHistory[70][periods[9]] = 1
	planHistory[70][periods[10]] = 1
	planHistory[70][periods[11]] = 1

	// Code 71
	planHistory[71][periods[0]] = 1
	planHistory[71][periods[1]] = 0
	planHistory[71][periods[2]] = 2
	planHistory[71][periods[3]] = 1
	planHistory[71][periods[4]] = 0
	planHistory[71][periods[5]] = 0
	planHistory[71][periods[6]] = 2
	planHistory[71][periods[7]] = 1
	planHistory[71][periods[8]] = 1
	planHistory[71][periods[9]] = 1
	planHistory[71][periods[10]] = 1
	planHistory[71][periods[11]] = 1

	// Code 72
	planHistory[72][periods[0]] = 1
	planHistory[72][periods[1]] = 1
	planHistory[72][periods[2]] = 2
	planHistory[72][periods[3]] = 2
	planHistory[72][periods[4]] = 0
	planHistory[72][periods[5]] = 0
	planHistory[72][periods[6]] = 1
	planHistory[72][periods[7]] = 2
	planHistory[72][periods[8]] = 0
	planHistory[72][periods[9]] = 1
	planHistory[72][periods[10]] = 1
	planHistory[72][periods[11]] = 1

	// Code 73
	planHistory[73][periods[0]] = 2
	planHistory[73][periods[1]] = 0
	planHistory[73][periods[2]] = 0
	planHistory[73][periods[3]] = 2
	planHistory[73][periods[4]] = 1
	planHistory[73][periods[5]] = 1
	planHistory[73][periods[6]] = 1
	planHistory[73][periods[7]] = 1
	planHistory[73][periods[8]] = 1
	planHistory[73][periods[9]] = 1
	planHistory[73][periods[10]] = 1
	planHistory[73][periods[11]] = 1

	// Code 74
	planHistory[74][periods[0]] = 1
	planHistory[74][periods[1]] = 0
	planHistory[74][periods[2]] = 0
	planHistory[74][periods[3]] = 1
	planHistory[74][periods[4]] = 0
	planHistory[74][periods[5]] = 1
	planHistory[74][periods[6]] = 1
	planHistory[74][periods[7]] = 0
	planHistory[74][periods[8]] = 2
	planHistory[74][periods[9]] = 2
	planHistory[74][periods[10]] = 0
	planHistory[74][periods[11]] = 1
}
