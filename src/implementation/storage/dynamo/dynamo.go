package dynamo

import (
	"go-clean-app-project/src/domain/models"
	"math/rand"
	"time"
)

// DynamoStorage implements the Storage interface for DynamoDB.
type DynamoStorage struct {
	// Add DynamoDB client or session here
}

// NewDynamoStorage creates a new instance of DynamoStorage.
func NewDynamoStorage() (*DynamoStorage, error) {
	return &DynamoStorage{
		// Initialize DynamoDB client or session here
	}, nil
}

// SaveTask saves a task to the DynamoDB database.
func (s *DynamoStorage) SaveTask(task *models.Task) error {
	task.ID = uint64(rand.Intn(999999))
	task.CreatedAt = time.Now()

	// Implement the logic to save the task to the DynamoDB database
	// This could involve using the AWS SDK for Go to put an item in a DynamoDB table
	// Example: _, err := s.dynamoClient.PutItem(&dynamodb.PutItemInput{
	//     TableName: aws.String("Tasks"),
	//     Item: map[string]*dynamodb.AttributeValue{
	//         "ID":          {S: aws.String(task.ID)},
	//         "Title":       {S: aws.String(task.Title)},
	//         "Description": {S: aws.String(task.Description)},
	//         "CreatedAt":   {S: aws.String(task.CreatedAt)},
	//     },
	// })

	return nil
}
