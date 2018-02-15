package die

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

// Roll takes in a string such as "1d10" validates it. Will then roll the die
// and return the total of the number of rolls requested.
func Roll(s string) (int, error) {
	matched, err := regexp.MatchString("^[1-9]d[1-9]\\d*", s)
	if matched == false {
		return -1, err
	}
	split := strings.Split(s, "d")
	total := 0
	times, _ := strconv.Atoi(split[0])
	max, _ := strconv.Atoi(split[1])

	for i := 0; i < times; i++ {
		total += (rand.Intn(max) + 1)
	}

	return total, nil
}
