// csv読込テスト
package test

import (
	"encoding/csv"
	"fmt"
	"os"
	_ "strings"
	"strconv"
	"core"
)

func TestCSV() {
	file, err := os.Open("./sample.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	record,err2 := reader.ReadAll()
	if err2 != nil {
		panic(err2)
	}

	fmt.Printf("in TestCSV:\n")
	fmt.Printf("%#v\n", record)

	for i,v := range record {
		fmt.Printf("%d %v\n", i, v)

		for _, num := range v {
			d, _ := strconv.ParseFloat(num, 64)
			// d, _ := strconv.Atof(num)
			fmt.Printf("%s %f\n", num, d)
		}
		// slice := strings.Split(v, " ")
		// for _, str := range slice {
		// 	fmt.Printf("%s\n", str)
		// }
	}

	// Dcontainerへ格納
	dd := core.NewDcontainer(record)

	fmt.Printf("\n")
	fmt.Printf(" hokan = %f\n", dd.GetValue(2.5))

	// 下限チェック
	fmt.Printf(" hokan = %f\n", dd.GetValue(0.99))
	// 上限チェック
	fmt.Printf(" hokan = %f\n", dd.GetValue(4.001))

}