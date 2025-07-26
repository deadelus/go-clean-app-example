package cli

import (
	"context"
	"fmt"
	"go-clean-app-project/src/domain/dto"
	"go-clean-app-project/src/transport"

	"github.com/AlecAivazis/survey/v2"
)

func (s *SurveyController) createTaskFlow() error {
	fmt.Println("\n📝 Creating a new task...")

	var qs = []*survey.Question{
		{
			Name:     "title",
			Prompt:   &survey.Input{Message: "📝 Title:"},
			Validate: survey.Required,
		},
		{
			Name:     "description",
			Prompt:   &survey.Input{Message: "📝 Description:"},
			Validate: survey.Required,
		},
	}

	answers := struct {
		Title       string `survey:"title"`
		Description string `survey:"description"`
	}{}

	if err := survey.Ask(qs, &answers); err != nil {
		return err
	}

	// Confirmer avant création
	confirm := false
	confirmPrompt := &survey.Confirm{
		Message: fmt.Sprintf("Create Task for %s (%s)?", answers.Title, answers.Description),
	}
	if err := survey.AskOne(confirmPrompt, &confirm); err != nil {
		return err
	}

	if !confirm {
		fmt.Println("⏹️ Creation cancelled")
		return nil
	}

	// Créer via le handler
	req := transport.TransportRequest[dto.TaskRequest]{
		Data: dto.TaskRequest{
			Title:       answers.Title,
			Description: answers.Description,
		},
		Context: context.Background(),
		Source:  "interactive",
	}

	response := s.handler.HandleTask(req)

	if response.Success {
		fmt.Printf("\n✅ Task created successfully!\n")
		fmt.Printf("   🆔 ID: %d\n", response.Data.ID)
		fmt.Printf("   📝 Title: %s\n", response.Data.Title)
		fmt.Printf("   📝 Description: %s\n", response.Data.Description)
		fmt.Printf("   📅 Created: %s\n\n", response.Data.CreatedAt.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Printf("\n❌ Error: %s\n\n", response.Error)
	}

	return nil
}

func (s *SurveyController) listTasks() {
	fmt.Println("\n📋 Task List:")
	fmt.Println("   • Task_001 - 📝 Title: Task 1 - 📝 Description: Description 1")
	fmt.Println("   • Task_002 - 📝 Title: Task 2 - 📝 Description: Description 2")
	fmt.Println("   • Task_003 - 📝 Title: Task 3 - 📝 Description: Description 3")
	fmt.Println()
}
