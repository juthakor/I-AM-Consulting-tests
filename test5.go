package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
)

func findFee(t string) {
    hour := 0
    min := 0
    var total float64
    if strings.Index(t, "hours") == -1 {
        min, _ = strconv.Atoi(strings.Split(t, " minutes")[0])
    } else{
	hour, _ = strconv.Atoi(strings.Split(t, " hours")[0])
        min, _ = strconv.Atoi(strings.Split(strings.Split(t, " hours ")[1], " minutes")[0])
    }
    total = float64(60*hour + min)
    if total < 30 {
	fmt.Println("Free")
    } else if total <= 180 {
	tmp := strconv.Itoa(int(math.Ceil(total/60)*10))
	fmt.Println(tmp + " bath")
    } else if total <= 360 {
	tmp := strconv.Itoa(int(math.Ceil(total/60)*20))
	fmt.Println(tmp + " bath")
    } else if total > 360 {
	tmp := strconv.Itoa(int(math.Ceil(total/1440)*200))
	fmt.Println(tmp + " bath")
    }
}

func main() {
	findFee("4 hours 15 minutes")
}
