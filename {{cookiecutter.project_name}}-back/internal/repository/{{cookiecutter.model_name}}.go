package repository

//go:generate mockgen -source ./{{cookiecutter.model_name}}.go -destination ../testing/repository/mock/{{cookiecutter.model_name}}.go -package mock

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	clientDynamo "github.com/divilo/aws-go/service/dynamodb"
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/model"
	"go.uber.org/zap"
	"time"
)

//{{cookiecutter.model_name.capitalize()}}Repository repository interface for {{cookiecutter.model_name}}
type {{cookiecutter.model_name.capitalize()}}Repository interface {
	Upsert(ctx context.Context, {{cookiecutter.model_name}} model.{{cookiecutter.model_name.capitalize()}}) (model.{{cookiecutter.model_name.capitalize()}}, error)
	FilterByID(ctx context.Context, dvcID string) (model.{{cookiecutter.model_name.capitalize()}}, error)
}

type {{cookiecutter.model_name}}Repository struct {
	lgr     *zap.SugaredLogger
	dynClt  clientDynamo.ServiceDynamo
	tblName string
}

//New{{cookiecutter.model_name.capitalize()}}Repository creates a new {{cookiecutter.model_name}} repository
func New{{cookiecutter.model_name.capitalize()}}Repository(lgr *zap.SugaredLogger, dynClt clientDynamo.ServiceDynamo, tblName string) {{cookiecutter.model_name.capitalize()}}Repository {
	return &{{cookiecutter.model_name}}Repository{
		lgr:     lgr,
		dynClt:  dynClt,
		tblName: tblName,
	}
}

func (dr *{{cookiecutter.model_name}}Repository) FilterByID(ctx context.Context, {{cookiecutter.model_name}}Id string) (model.{{cookiecutter.model_name.capitalize()}}, error) {

	{{cookiecutter.model_name}} := model.{{cookiecutter.model_name.capitalize()}}{}
	{{cookiecutter.model_name}}.{{cookiecutter.model_name.capitalize()}}Id = {{cookiecutter.model_name}}Id

	cond := expression.Key(TableColumn{{cookiecutter.model_name.capitalize()}}Id).Equal(expression.Value(&types.AttributeValueMemberS{Value: {{cookiecutter.model_name}}.{{cookiecutter.model_name.capitalize()}}Id}))

	expr, err := expression.NewBuilder().WithKeyCondition(cond).Build()
	if err != nil {
		return model.{{cookiecutter.model_name.capitalize()}}{}, err
	}

	input := &dynamodb.QueryInput{
		TableName:                 aws.String(dr.tblName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	}

	res, err := dr.dynClt.Query(ctx, input)
	if err != nil {
		return model.{{cookiecutter.model_name.capitalize()}}{}, err
	}

	if len(res.Items) == 0 {
		return model.{{cookiecutter.model_name.capitalize()}}{}, errors.New(MsgErrorNotFound)
	}
	if len(res.Items) > 1 {
		dr.lgr.Errorw(MsgErrorMoreThanOneFound, zap.Any("items", res.Items))
		return model.{{cookiecutter.model_name.capitalize()}}{}, errors.New(MsgErrorMoreThanOneFound)
	}

	// unMarshall item found to  {{cookiecutter.model_name}}Table
	var {{cookiecutter.model_name}}TableFound {{cookiecutter.model_name}}Table
	if err := attributevalue.UnmarshalMap(res.Items[0], &{{cookiecutter.model_name}}TableFound); err != nil {
		dr.lgr.Errorw(MsgErrorNotAbleMarshal, zap.Error(err))
		return model.{{cookiecutter.model_name.capitalize()}}{}, err
	}

	//mapping to {{cookiecutter.model_name.capitalize()}} (model)
	{{cookiecutter.model_name}} = *toModel{{cookiecutter.model_name.capitalize()}}(&{{cookiecutter.model_name}}TableFound)
	return {{cookiecutter.model_name}}, nil
}

//Create new or update if exist
func (dr *{{cookiecutter.model_name}}Repository) Upsert(ctx context.Context, {{cookiecutter.model_name}} model.{{cookiecutter.model_name.capitalize()}}) (model.{{cookiecutter.model_name.capitalize()}}, error) {

	//Create updateItemInput of {{cookiecutter.model_name}}
	updateItemInput, err := dr.createUpdateItemInput({{cookiecutter.model_name}})
	if err != nil {
		dr.lgr.Errorw(MsgErrorNotCreatedUpdateItem, zap.Error(err))
		return model.{{cookiecutter.model_name.capitalize()}}{}, err
	}

	//Excecute Update Item _updateItemOutput
	var updateItemOutput *dynamodb.UpdateItemOutput
	updateItemOutput, err = dr.dynClt.UpdateItem(ctx, &updateItemInput)
	if err != nil {
		dr.lgr.Errorw(MsgErrorDynamoClientUpdateItem, zap.Error(err))
		return model.{{cookiecutter.model_name.capitalize()}}{}, err
	}

	var createdAt int64
	_ = attributevalue.Unmarshal(updateItemOutput.Attributes[TableColumnCreatedAt], &createdAt)
	if createdAt > 0 {
		{{cookiecutter.model_name}}.CreatedAt = createdAt
	}
	return {{cookiecutter.model_name}}, nil
}

func (dr *{{cookiecutter.model_name}}Repository) createUpdateItemInput({{cookiecutter.model_name}} model.{{cookiecutter.model_name.capitalize()}}) (dynamodb.UpdateItemInput, error) {

	{{cookiecutter.model_name}}.CreatedAt = time.Now().Unix()
	{{cookiecutter.model_name}}.UpdatedAt = time.Now().Unix()

	data, _ := attributevalue.MarshalMap(fromModel{{cookiecutter.model_name.capitalize()}}(&{{cookiecutter.model_name}}))
	var updateBuilder expression.UpdateBuilder
	for k, v := range data {
		if k == TableColumnCreatedAt {
			updateBuilder = updateBuilder.Set(expression.Name(k), expression.IfNotExists(expression.Name(k), expression.Value(v)))
			continue
		}
		if k != TableColumn{{cookiecutter.model_name.capitalize()}}Id {
			updateBuilder = updateBuilder.Set(expression.Name(k), expression.Value(v))
		}
	}

	expr, err := expression.NewBuilder().WithUpdate(updateBuilder).Build()
	if err != nil {
		return dynamodb.UpdateItemInput{}, err
	}

	updateItemInput := &dynamodb.UpdateItemInput{
		Key:                       map[string]types.AttributeValue{TableColumn{{cookiecutter.model_name.capitalize()}}Id: &types.AttributeValueMemberS{Value: {{cookiecutter.model_name}}.{{cookiecutter.model_name.capitalize()}}Id}},
		TableName:                 aws.String(dr.tblName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ReturnValues:              types.ReturnValueAllOld,
		UpdateExpression:          expr.Update(),
	}
	return *updateItemInput, nil
}
