package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/extintor/advent_of_code/shared/utils"
)

func solve(races [][2]int) int {
  total := 1
	for _, race := range races {
		t := race[0]
		d := race[1]
    r := math.Sqrt(float64((t * t) - (4 * d)))
    a := int((float64(t) + r) / 2)
    b := int((float64(t) - r ) / 2)
    total *= a - b
	}
  return total
}

func parseRace(ts, ds string) ([2]int, error) {
  t, err := strconv.Atoi(ts)
  if err != nil {
    return [2]int{}, err
  }
  d, err := strconv.Atoi(ds)
  if err != nil {
    return [2]int{}, err
  }
  return [2]int{t, d}, nil
}

func solveTwo(timeSlice, distanceSlice []string) (int, error) {
  r, err := parseRace(strings.Join(timeSlice, ""),strings.Join(distanceSlice, ""))
  if err != nil {
    return 0, err
  }
  races := [][2]int{r}
  return solve(races), nil
}

func solveOne(timeSlice, distanceSlice []string) (int, error) {
  races := make([][2]int, 0)
  for i := 0; i < len(timeSlice); i++ {
    r, err := parseRace(timeSlice[i], distanceSlice[i])
    if err != nil {
      return 0, err
    }
    races = append(races, r)
  }
  solve(races)
  return solve(races), nil
}

func main() {
  lines := utils.ReadInput()
  timeSlice := strings.Fields(strings.TrimPrefix(lines[0], "Time:"))
  distanceSlice := strings.Fields(strings.TrimPrefix(lines[1], "Distance:"))
	fmt.Println(solveOne(timeSlice, distanceSlice))
	fmt.Println(solveTwo(timeSlice, distanceSlice))
}
