package expensetracker

import (
	"encoding/csv"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)



func AddExpense(cmd *cobra.Command, args []string){
	description, _ := cmd.Flags().GetString("description")
	category, _ := cmd.Flags().GetString("category")
	amount, err := cmd.Flags().GetFloat64("amount")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	} 

	file, new := persistenceFile(true)
	defer file.Close()

	reader := csv.NewReader(file)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	current, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	id := nextId(&current)

	if new {
		writer.Write(FILE_HEADERS)
	}

	// Fields {"ID", "Date", "Description", "Category", "Amount"}
	record := []string{fmt.Sprint(id), time.Now().Format("2006-01-02"), description, category, fmt.Sprint(amount)}
	writer.Write(record)
	
	fmt.Printf("Expense added successfully (ID: %v)\n", id)
}