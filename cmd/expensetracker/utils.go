package expensetracker

import (
	"errors"
	"os"
	"strconv"
)

func persistenceFile (create bool) (*os.File, bool) {
	file, err := os.OpenFile(FILE_NAME, os.O_RDWR, 0644)
	if err != nil {
		if errors.Is(err, os.ErrNotExist){
			if create {
				file, err = os.Create(FILE_NAME)
				if err != nil {
					panic(err)
				}
			}
			return file, true
		} else {
			panic(err)
		}
	}
	return file, false
}

func nextId (records *[][]string) int{
	var lastId string = "0";
	if len(*records) != 0{
		lastId = (*records)[len(*records)-1][0]
	}
	id, err := strconv.Atoi(lastId)
	if err != nil{
		id = 0
	}
	return id + 1
}