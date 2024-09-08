package expensetracker

import (
	"encoding/csv"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func ExpensesSummary(cmd *cobra.Command, args []string){
	file, new := persistenceFile(false)
	if new {
		fmt.Println("No saved Expenses:( Run 'add' to track a new expense")
		return
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil{
		panic(err)
	}

	var total float64 = 0
	for index, val := range records{
		if index != 0{
			amount, err := strconv.ParseFloat(val[4], 64)
			if err != nil{
				panic(err)
			}
			total += amount
		}
	}

	fmt.Printf("Total Expenses: $%v\n", total)

}