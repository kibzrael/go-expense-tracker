package expensetracker

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func DeleteExpense(cmd *cobra.Command, args []string){
	id, err := cmd.Flags().GetString("id")
	if err != nil{
		panic(err)
	}

	file, err := os.Open(FILE_NAME)
	if err != nil {
		fmt.Println("No saved Expenses :( Run 'add' to track a new expense")
		return
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil{
		panic(err)
	}

	file, err = os.Create(FILE_NAME)
	if err != nil{
		panic(err)
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	var newRecords [][]string
	for _, record := range records{
		if record[0] != id{
			newRecords = append(newRecords, record)
		}
	}
	writer.WriteAll(newRecords)
	fmt.Printf("Expense added deleted (ID: %v)\n", id)
}