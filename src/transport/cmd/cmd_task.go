package cmd

import (
	"context"
	"fmt"
	"go-clean-app-project/src/domain/dto"
	"go-clean-app-project/src/transport"

	"github.com/spf13/cobra"
)

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "📝 Task command",
	Long:  `Execute the task use case with the provided email and name.`,
}

// createCmd represents the create subcommand
var createCmd = &cobra.Command{
	Use:   "create [title] [description]",
	Short: "➕ Create task",
	Long:  `Create an task with the specified title and description.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		description := args[1]

		// Créer le handler de base
		baseHandler := transport.NewBaseHandler(useCases, appLogger)

		// Créer la requête transport
		req := transport.TransportRequest[dto.TaskRequest]{
			Data: dto.TaskRequest{
				Title:       title,
				Description: description,
			},
			Context: context.Background(),
			Source:  "cli",
		}

		// Exécuter le handler
		response := baseHandler.HandleTask(req)

		// Afficher le résultat
		if response.Success {
			fmt.Printf("✅ task created successfully!\n")
			fmt.Printf("   ID: %d\n", response.Data.ID)
			fmt.Printf("   Title: %s\n", response.Data.Title)
			fmt.Printf("   Description: %s\n", response.Data.Description)
			fmt.Printf("   Created At: %s\n", response.Data.CreatedAt.Format("2006-01-02 15:04:05"))
		} else {
			fmt.Printf("❌ Error: %s\n", response.Error)
		}
	},
}

// Execute executes the root command
func init() {
	rootCmd.AddCommand(taskCmd)
	taskCmd.AddCommand(createCmd)

	// Flags pour la commande create
	createCmd.Flags().Bool("verbose", false, "Verbose output")
}
