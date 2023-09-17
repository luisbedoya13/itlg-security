package dynamo_db

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"itlg_security/create-token/internal/repository"
	"itlg_security/create-token/pkg/model"
	"os"
)

type Reposiroty struct {
	client *dynamodb.Client
}

func New() (*Reposiroty, error) {
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
		return nil, err
	} else {
		return &Reposiroty{client: dynamodb.NewFromConfig(cfg)}, nil
	}
}

func (r *Reposiroty) GetUser(ctx context.Context, email string) (*model.DdbUser, error) {
	output, err := r.client.GetItem(ctx, &dynamodb.GetItemInput{
		Key:       map[string]types.AttributeValue{"email": &types.AttributeValueMemberS{Value: email}},
		TableName: aws.String(fmt.Sprintf("itlg-%s-users", os.Getenv("STAGE"))),
	})
	if err != nil {
		return nil, err
	} else {
		user := model.DdbUser{}
		if len(output.Item) == 0 {
			return nil, repository.ErrNotFound
		}
		if err := attributevalue.UnmarshalMap(output.Item, &user); err != nil {
			return nil, err
		}
		return &user, nil
	}
}
