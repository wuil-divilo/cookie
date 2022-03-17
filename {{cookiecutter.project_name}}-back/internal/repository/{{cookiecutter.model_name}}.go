package repository

//go:generate mockgen -source ./device.go -destination ../testing/repository/mock/device.go -package mock

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

//DeviceRepository repository interface for device
type DeviceRepository interface {
	Upsert(ctx context.Context, device model.Device) (model.Device, error)
	FilterByID(ctx context.Context, dvcID string) (model.Device, error)
}

type deviceRepository struct {
	lgr     *zap.SugaredLogger
	dynClt  clientDynamo.ServiceDynamo
	tblName string
}

//NewDeviceRepository creates a new device repository
func NewDeviceRepository(lgr *zap.SugaredLogger, dynClt clientDynamo.ServiceDynamo, tblName string) DeviceRepository {
	return &deviceRepository{
		lgr:     lgr,
		dynClt:  dynClt,
		tblName: tblName,
	}
}

func (dr *deviceRepository) FilterByID(ctx context.Context, deviceId string) (model.Device, error) {

	device := model.Device{}
	device.DeviceId = deviceId

	cond := expression.Key(TableColumnDeviceId).Equal(expression.Value(&types.AttributeValueMemberS{Value: device.DeviceId}))

	expr, err := expression.NewBuilder().WithKeyCondition(cond).Build()
	if err != nil {
		return model.Device{}, err
	}

	input := &dynamodb.QueryInput{
		TableName:                 aws.String(dr.tblName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	}

	res, err := dr.dynClt.Query(ctx, input)
	if err != nil {
		return model.Device{}, err
	}

	if len(res.Items) == 0 {
		return model.Device{}, errors.New(MsgErrorNotFound)
	}
	if len(res.Items) > 1 {
		dr.lgr.Errorw(MsgErrorMoreThanOneFound, zap.Any("items", res.Items))
		return model.Device{}, errors.New(MsgErrorMoreThanOneFound)
	}

	// unMarshall item found to  deviceTable
	var deviceTableFound deviceTable
	if err := attributevalue.UnmarshalMap(res.Items[0], &deviceTableFound); err != nil {
		dr.lgr.Errorw(MsgErrorNotAbleMarshal, zap.Error(err))
		return model.Device{}, err
	}

	//mapping to Device (model)
	device = *toModelDevice(&deviceTableFound)
	return device, nil
}

//Create new or update if exist
func (dr *deviceRepository) Upsert(ctx context.Context, device model.Device) (model.Device, error) {

	//Create updateItemInput of device
	updateItemInput, err := dr.createUpdateItemInput(device)
	if err != nil {
		dr.lgr.Errorw(MsgErrorNotCreatedUpdateItem, zap.Error(err))
		return model.Device{}, err
	}

	//Excecute Update Item _updateItemOutput
	var updateItemOutput *dynamodb.UpdateItemOutput
	updateItemOutput, err = dr.dynClt.UpdateItem(ctx, &updateItemInput)
	if err != nil {
		dr.lgr.Errorw(MsgErrorDynamoClientUpdateItem, zap.Error(err))
		return model.Device{}, err
	}

	var createdAt int64
	_ = attributevalue.Unmarshal(updateItemOutput.Attributes[TableColumnCreatedAt], &createdAt)
	if createdAt > 0 {
		device.CreatedAt = createdAt
	}
	return device, nil
}

func (dr *deviceRepository) createUpdateItemInput(device model.Device) (dynamodb.UpdateItemInput, error) {

	device.CreatedAt = time.Now().Unix()
	device.UpdatedAt = time.Now().Unix()

	data, _ := attributevalue.MarshalMap(fromModelDevice(&device))
	var updateBuilder expression.UpdateBuilder
	for k, v := range data {
		if k == TableColumnCreatedAt {
			updateBuilder = updateBuilder.Set(expression.Name(k), expression.IfNotExists(expression.Name(k), expression.Value(v)))
			continue
		}
		if k != TableColumnDeviceId {
			updateBuilder = updateBuilder.Set(expression.Name(k), expression.Value(v))
		}
	}

	expr, err := expression.NewBuilder().WithUpdate(updateBuilder).Build()
	if err != nil {
		return dynamodb.UpdateItemInput{}, err
	}

	updateItemInput := &dynamodb.UpdateItemInput{
		Key:                       map[string]types.AttributeValue{TableColumnDeviceId: &types.AttributeValueMemberS{Value: device.DeviceId}},
		TableName:                 aws.String(dr.tblName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ReturnValues:              types.ReturnValueAllOld,
		UpdateExpression:          expr.Update(),
	}
	return *updateItemInput, nil
}
