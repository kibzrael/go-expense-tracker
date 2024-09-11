package expensetracker

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"time"

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

	month, err := cmd.Flags().GetInt64("month")
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
			date, err:= time.Parse("2006-01-02", val[1])
			if err != nil{
				panic(err)
			}
			isThisMonth := date.Year() == time.Now().Year() && int64(date.Month()) == month
			if isThisMonth || month == 0 {
				total += amount
            }
		}
	}

	if month == 0 {
		fmt.Printf("Total Expenses: $%v\n", total)
	} else {
        fmt.Printf("Total Expenses for %s: $%v\n", time.Month(month).String(), total)
    }
}