package main

import (
	"fmt"
	"sort"
	"strings"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func wordGrouping(words string) []string {
	wordsArray := strings.Split(words, ",")
	var resultArr []string

	for _, item1 := range wordsArray {
		var tmp []string
		var tmpStr string
		for _, item2 := range wordsArray {
			tmpItem1 := strings.Split(item1, "")
			sort.Strings(tmpItem1)
			tmpItem2 := strings.Split(item2, "")
			sort.Strings(tmpItem2)
			if strings.Join(tmpItem1, "") == strings.Join(tmpItem2, "") {
				tmp = append(tmp, item2)
			}

		}

		tmpStr = strings.Join(tmp, "-")
		if len(resultArr) == 0 {
			resultArr = append(resultArr, tmpStr)
		}
		if !contains(resultArr, tmpStr) {
			resultArr = append(resultArr, tmpStr)
		}
	}
	return resultArr
}

func main() {
	fmt.Println(wordGrouping("VMRCO,VORCM,MCRTOX,ZXTAC,XZATC,XMTCOR,OXVS,AZTXC,VXOS,RZAT,MRCOTX,SVXO,TRAZ,ZTAR,MVOCR"))
}
