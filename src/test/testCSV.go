// csv読込テスト
package test

import (
	"core"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	_ "strings"
)

func TestCSV() {
	file, err := os.Open("./data/sample.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	record, err2 := reader.ReadAll()
	if err2 != nil {
		panic(err2)
	}

	fmt.Printf("in TestCSV:\n")
	fmt.Printf("%#v\n", record)

	for i, v := range record {
		fmt.Printf("%d %v %d\n", i, v, len(v))

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

func TestCSV2() {
	file, err := os.Open("./data/sampleT.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	record, err2 := reader.ReadAll()
	if err2 != nil {
		panic(err2)
	}

	fmt.Printf("in TestCSV2:\n")
	fmt.Printf("%#v\n", record)

	// for i, v := range record {
	// 	fmt.Printf("%d %s %d\n", i, v, len(v))

	// 	for _, num := range v {
	// 		d, _ := strconv.ParseFloat(num, 64)
	// 		fmt.Printf("%s %f\n", num, d)
	// 	}
	// }

	// Tcontainerへ格納
	dd := core.NewTcontainer(record)

	fmt.Printf("%v\n", dd)

	fmt.Printf("\n")

}
