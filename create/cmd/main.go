package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/subosito/gotenv"
	"itlg_security/create/pkg/model"
	"log"
	"os"
)

func init() {
	if err := gotenv.Load("../../.env"); err != nil {
		log.Fatalf("Error loading variables: %v\n", err)
	}
}

func main() {
	creds := credentials.StaticCredentialsProvider{
		Value: aws.Credentials{
			AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
			SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		},
	}
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("sa-east-1"),
		config.WithCredentialsProvider(creds),
	)
	if err != nil {
		log.Printf("Error loading config: %v\n", err)
	}
	ddbClient := dynamodb.NewFromConfig(cfg)
	output, err := ddbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key:       map[string]types.AttributeValue{"email": &types.AttributeValueMemberS{Value: "owner@mail.com"}},
		TableName: aws.String(fmt.Sprintf("itlg-%s-users", os.Getenv("STAGE"))),
	})
	if err != nil {
		log.Printf("Couldn't get item. Here's why: %v\n", err)
	} else {
		user := model.DdbUser{}
		if err := attributevalue.UnmarshalMap(output.Item, &user); err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)
	}
}
