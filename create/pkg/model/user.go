package model

type DdbUser struct {
	Email string `dynamodbav:"email"`
	Password string `dynamodbav:"password"`
	ID string `dynamodbav:"id"`
}