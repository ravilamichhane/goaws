package database

import (
	"lambda-func/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBClient struct {
	databaseStore *dynamodb.DynamoDB
}

const (
	TABLE_NAME = "users"
)

type UserStore interface {
	DoesUserExist(username string) (bool, error)
	InsertUser(user types.RegisterUserRequest) error
}

func NewDynamoDBClient() *DynamoDBClient {
	dbSession := session.Must(session.NewSession())
	db := dynamodb.New(dbSession)
	return &DynamoDBClient{databaseStore: db}
}

func (d DynamoDBClient) DoesUserExist(username string) (bool, error) {

	result, err := d.databaseStore.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	})

	if err != nil {
		return false, err
	}

	if result.Item == nil {
		return false, nil
	}

	return true, nil
}

func (d DynamoDBClient) InsertUser(user types.RegisterUserRequest) error {

	item := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(user.Username),
			},
			"password": {
				S: aws.String(user.Password),
			},
		},
	}

	_, err := d.databaseStore.PutItem(item)

	if err != nil {
		return err
	}

	return nil
}
