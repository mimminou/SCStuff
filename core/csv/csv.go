package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func WriteCSV(fileName string, fieldNames []string, data [][]string) error {
	cwd, cwdErr := os.Getwd()
	if cwdErr != nil {
		fmt.Println(cwdErr)
		return cwdErr
	}
	dirErr := os.Mkdir("csv", 0755)
	if dirErr != nil {
		if os.IsExist(dirErr) {
			fmt.Println("dir exists")
		} else {
			fmt.Println(dirErr)
		}
	}

	file, fileErr := os.Create(cwd + "/csv/" + fileName + ".csv")
	if fileErr != nil {
		if os.IsExist(fileErr) {
			fmt.Println("file exists, overwriting")
		} else {
			fmt.Println(fileErr)
			return fileErr
		}
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	//write header
	if csvErr := csvWriter.Write(fieldNames); csvErr != nil {
		fmt.Println(csvErr)
		return csvErr
	}

	for _, row := range data {
		if err := csvWriter.Write(row); err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}
