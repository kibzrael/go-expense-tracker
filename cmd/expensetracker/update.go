package expensetracker

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func UpdateExpense(cmd *cobra.Command, args []string){
	id, err := cmd.Flags().GetString("id")
	if err != nil{
		panic(err)
	}
	description, _ := cmd.Flags().GetString("description")
	category, _ := cmd.Flags().GetString("category")
	amount, err := cmd.Flags().GetFloat64("amount")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
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

	for _, record := range records{
		if record[0] == id{
			if description != "" {
				record[2] = description
			} else if category != "" {
				record[3] = category
			} else if amount != 0 {
				record[4] = fmt.Sprint(amount)
			}
		}
	}
	writer.WriteAll(records)
	fmt.Printf("Expense added updated (ID: %v)\n", id)
}