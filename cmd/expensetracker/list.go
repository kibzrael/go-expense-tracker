package expensetracker

import (
	"encoding/csv"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func ListExpenses(cmd *cobra.Command, args []string){
	file, new := persistenceFile(false)
	if new {
		fmt.Println("No saved Expenses :( Run 'add' to track a new expense")
		return
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil{
		panic(err)
	}

	category, err := cmd.Flags().GetString("category")
	if err != nil{
		panic(err)
	}

	cliWriter := tabwriter.NewWriter(os.Stdout, 1, 1, 4, ' ', 0)
	defer cliWriter.Flush()
	for i, record := range records{
		var output string = ""
		if category != "" && category != record[3] && i != 0{
			continue
		}
		for _, val := range record{
			output += fmt.Sprintf("%v\t", val)
		}
		fmt.Fprintln(cliWriter, output)
	}
}
