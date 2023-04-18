package excel

import (
	"os"
	"encoding/csv"
	"bufio"
	"io"
	"log"
	"strconv"
	"fmt"
)

var ExcelHandle = &excelHandle{}

type excelHandle struct{}

func (e *excelHandle) Do() {
	csvFile, err := os.Open("excel/files/example.csv")
	if err != nil {
		fmt.Printf("file err:%v", err)
		return
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var ret = make(map[uint32]int)
	i := 0
	j := 0
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if len(line) >= 2 {
			i++
			uid, _ := strconv.Atoi(line[0])
			score, _ := strconv.Atoi(line[1])
			if uid > 0 && score > 0 {
				level := -1
				if score >= 5000 {
					level = 3
				} else if score >= 1000 {
					level = 2
				} else if score >= 500 {
					level = 1
				} else {
					level = 0
				}
				if level >= 0 {
					ret[uint32(uid)] = level
				}
			} else {
				j++
			}
		}
	}
	fmt.Printf("line num:%d ,left:%d \r\n", i, i-j)
	file, err := os.Create("test.txt")
	defer file.Close()
	os.Chmod("test.txt", 0777)
	write := bufio.NewWriter(file)
	write.WriteString("map[uint32]int{\r\n")
	for k, v := range ret {
		_, err = write.WriteString(fmt.Sprintf("     %d:%d,\r\n", k, v))
		if err != nil {
			fmt.Printf("write error: %v\n", err)
			write.WriteString(fmt.Sprintf("write error:%v", err))
		}
	}
	write.WriteString("}")
	write.Flush()
	fmt.Println("Done")
}
