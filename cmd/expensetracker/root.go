package expensetracker

import (
	"os"

	"github.com/spf13/cobra"
)

func Execute(){
	addCmd := &cobra.Command{ Use: "add", Short: "", Run: AddExpense }
	addCmd.Flags().Float64P("amount", "a", 10, "Amount of expense to track.")
	addCmd.Flags().StringP("description", "d", "New Expense", "Description of the expense to track")
	addCmd.Flags().StringP("category", "c", "Default", "Category of the expense to track")
	
	updateCmd := &cobra.Command{ Use: "update", Short: "", Run: UpdateExpense }
	updateCmd.Flags().String("id", "0", "ID of the expense to update")
	updateCmd.Flags().Float64P("amount", "a", 0, "Amount of expense to update.")
	updateCmd.Flags().StringP("description", "d", "", "Description of the expense to update")
	updateCmd.Flags().StringP("category", "c", "", "Category of the expense to update")

	listCmd := &cobra.Command{ Use: "list", Short: "", Run: ListExpenses }

	summaryCmd := &cobra.Command{ Use: "summary", Short: "", Run: ExpensesSummary }
	summaryCmd.Flags().Int64P("month", "m", 0, "Month to get an expense summary of.")

	deleteCmd := &cobra.Command{ Use: "delete", Short: "", Run: DeleteExpense }
	deleteCmd.Flags().String("id", "0", "ID of the expense to delete")

	rootCmd := &cobra.Command{
		Use: "expensetracker",
		Short: "",
		Long: "",
	}
	rootCmd.AddCommand(addCmd,updateCmd, listCmd, summaryCmd, deleteCmd)
	err := rootCmd.Execute()
	if err != nil{
		os.Exit(1)
	}
}

