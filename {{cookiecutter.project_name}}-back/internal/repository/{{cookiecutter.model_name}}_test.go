package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/middleware"
	clientDynamo "github.com/divilo/aws-go/service/dynamodb"
	"github.com/divilo/aws-go/testing/mock"
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/model"
	utillogger "github.com/divilo/utils-go/service/logger"
	"github.com/golang/mock/gomock"
	"github.com/oxyno-zeta/gomock-extra-matcher"
	"go.uber.org/zap"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

const (
	DeviceTableName = "{{cookiecutter.project_name}}-{{cookiecutter.model_name}}"
)

func TestNewDeviceRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	dynCltMock := mock.NewMockServiceDynamo(ctrl)
	type args struct {
		lgr     *zap.SugaredLogger
		dynClt  clientDynamo.ServiceDynamo
		tblName string
	}

	logger := utillogger.NewEmpty()

	tests := []struct {
		name string
		args args
		want DeviceRepository
	}{
		{
			name: "instantiates ok",
			args: args{
				lgr:     logger,
				dynClt:  dynCltMock,
				tblName: DeviceTableName,
			},
			want: &deviceRepository{
				lgr:     logger,
				dynClt:  dynCltMock,
				tblName: DeviceTableName,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeviceRepository(tt.args.lgr, tt.args.dynClt, tt.args.tblName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeviceRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deviceRepository_Upsert(t *testing.T) {

	ctrl := gomock.NewController(t)
	dynCltMock := mock.NewMockServiceDynamo(ctrl)
	lg := utillogger.NewEmpty()

	//arguments for create device repository
	argsRepository := deviceRepository{
		lgr:     lg,
		dynClt:  dynCltMock,
		tblName: DeviceTableName,
	}

	//arguments for function to testing Insert
	type args struct {
		ctx    context.Context
		device model.Device
	}

	//Test suite
	tests := []struct {
		name           string
		argsRepo       deviceRepository
		configureMocks func()
		args           args
		wantAssert     func(got model.Device) bool
		wantErr        bool
	}{
		{
			name:     "Upsert error client dynamo",
			argsRepo: argsRepository,
			configureMocks: func() {

				dynCltMock.EXPECT().UpdateItem(context.TODO(), new(updateMatcher)).
					Return(nil, errors.New("some error client dynamo"))
			},
			args: args{
				ctx:    context.TODO(),
				device: getTestData().deviceNew,
			},
			wantAssert: func(got model.Device) bool { return true },
			wantErr:    true,
		},
		{
			name:     "Upsert created Ok",
			argsRepo: argsRepository,
			configureMocks: func() {

				dynCltMock.EXPECT().UpdateItem(context.TODO(), new(updateMatcher)).
					Return(createUpdateItemOutputCreated(), nil)
			},
			args: args{
				ctx:    context.TODO(),
				device: getTestData().deviceNew,
			},
			wantAssert: func(got model.Device) bool {
				want := getTestData().deviceNew
				return compare{{cookiecutter.model_name}}WithoutDates(&got, &want)
			},
			wantErr: false,
		},
		{
			name:     "Upsert updated Ok",
			argsRepo: argsRepository,
			configureMocks: func() {

				dynCltMock.EXPECT().UpdateItem(context.TODO(), new(updateMatcher)).
					Return(createUpdateItemOutputUpdated(), nil)
			},
			args: args{
				ctx:    context.TODO(),
				device: getTestData().deviceNew,
			},
			wantAssert: func(got model.Device) bool {
				want := getTestData().deviceFound
				return compare{{cookiecutter.model_name}}WithoutUpdatedAt(&got, &want)
			},
			wantErr: false,
		},
	}
	//Execute test suit
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dr := NewDeviceRepository(tt.argsRepo.lgr, tt.argsRepo.dynClt, tt.argsRepo.tblName)

			//Configure mock
			tt.configureMocks()

			//Testing Upsert
			got, err := dr.Upsert(tt.args.ctx, tt.args.device)

			//Check result test
			if (err != nil) != tt.wantErr {
				t.Errorf("Upsert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantAssert(got) {
				t.Errorf("Upsert() got = %v", got)
			}
		})
	}
}

func Test_deviceRepository_FilterByID(t *testing.T) {

	//Create mock dynamoClient (ServiceDynamo)
	ctrl := gomock.NewController(t)
	dynCltMock := mock.NewMockServiceDynamo(ctrl)

	//arguments for create device repository
	argsRepository := deviceRepository{
		lgr:     utillogger.NewEmpty(),
		dynClt:  dynCltMock,
		tblName: DeviceTableName,
	}

	//arguments for function to testing FilterByID
	type args struct {
		ctx      context.Context
		deviceId string
	}

	//Test suite
	tests := []struct {
		name           string
		argsRepo       deviceRepository
		configureMocks func()
		args           args
		want           model.Device
		wantErr        bool
	}{
		{
			name:     "Query error",
			argsRepo: argsRepository,
			configureMocks: func() {

				dynCltMock.EXPECT().Query(context.TODO(), getTestData().queryInputExist).
					Return(nil, errors.New("operation error DynamoDB: Query"))
			},
			args: args{
				deviceId: "81a0aabc-7fe1-4b42-a387-d9f685a212e3",
				ctx:      context.TODO(),
			},
			want:    getTestData().deviceEmpty,
			wantErr: true,
		},
		{
			name:     "Not found",
			argsRepo: argsRepository,
			configureMocks: func() {

				dynCltMock.EXPECT().Query(context.TODO(), getTestData().queryInputNotExist).
					Return(getTestData().queryOutputEmpty, nil)
			},
			args: args{
				deviceId: "b5c43ec4-4f23-4118-a497-09563f2ddf30",
				ctx:      context.TODO(),
			},
			want:    getTestData().deviceEmpty,
			wantErr: true,
		},
		{
			name:     "More than one found",
			argsRepo: argsRepository,
			configureMocks: func() {

				dynCltMock.EXPECT().Query(context.TODO(), getTestData().queryInputNotExist).
					Return(getTestData().queryOutputTwoItemsFound, nil)
			},
			args: args{
				deviceId: "b5c43ec4-4f23-4118-a497-09563f2ddf30",
				ctx:      context.TODO(),
			},
			want:    getTestData().deviceEmpty,
			wantErr: true,
		},
		{
			name:     "Successful",
			argsRepo: argsRepository,
			configureMocks: func() {

				dynCltMock.EXPECT().Query(context.TODO(), getTestData().queryInputExist).
					Return(getTestData().queryOutputItemFound, nil)
			},
			args: args{
				deviceId: "81a0aabc-7fe1-4b42-a387-d9f685a212e3",
				ctx:      context.TODO(),
			},
			want:    getTestData().deviceFound,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dr := NewDeviceRepository(tt.argsRepo.lgr, tt.argsRepo.dynClt, tt.argsRepo.tblName)

			//Configure mock
			tt.configureMocks()

			//Testing FilterByID
			got, err := dr.FilterByID(tt.args.ctx, tt.args.deviceId)

			if (err != nil) != tt.wantErr {
				t.Errorf("FilterByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TEST DATA need for testing
type TestData struct {
	deviceEmpty              model.Device
	deviceNew                model.Device
	deviceFound              model.Device
	queryInputExist          *dynamodb.QueryInput
	queryInputNotExist       *dynamodb.QueryInput
	queryOutputEmpty         *dynamodb.QueryOutput
	queryOutputItemFound     *dynamodb.QueryOutput
	queryOutputTwoItemsFound *dynamodb.QueryOutput
}

func getTestData() TestData {

	deviceNew := createDeviceNew()
	deviceFound := createDeviceFound()
	queryInputExist := createQueryInputExist()
	queryInputNotExist := createQueryInputNotExist()
	queryOutputEmpty := createQueryOutputEmpty()
	queryOutputItemFound := createQueryOutputItemFound()
	queryOutputTwoItemsFound := createQueryOutputTwoItemsFound()

	testData := TestData{
		deviceEmpty:              model.Device{},
		deviceNew:                deviceNew,
		deviceFound:              deviceFound,
		queryInputExist:          &queryInputExist,
		queryInputNotExist:       &queryInputNotExist,
		queryOutputEmpty:         &queryOutputEmpty,
		queryOutputItemFound:     &queryOutputItemFound,
		queryOutputTwoItemsFound: &queryOutputTwoItemsFound,
	}
	return testData
}

func createDeviceNew() model.Device {
	deviceNew := model.Device{
		DeviceId:      "81a0aabc-7fe1-4b42-a387-d9f685a212e3",
		Model:         "model",
		HwVersion:     "hwversion",
		OSVersion:     "osversion",
		AppVersion:    "appversion",
		ApiLevel:      "apilevel",
		SecurityPatch: "securitypatch",
		NFCAvailable:  true,
		NFCEnabled:    true,
		CreatedAt:     0,
		UpdatedAt:     0,
	}
	return deviceNew
}

func createDeviceFound() model.Device {
	deviceFound := model.Device{
		DeviceId:      "81a0aabc-7fe1-4b42-a387-d9f685a212e3",
		Model:         "model",
		HwVersion:     "hwversion",
		OSVersion:     "osversion",
		AppVersion:    "appversion",
		ApiLevel:      "apilevel",
		SecurityPatch: "securitypatch",
		NFCAvailable:  true,
		NFCEnabled:    true,
		CreatedAt:     1643245234,
		UpdatedAt:     1643245234,
	}
	return deviceFound
}

func createQueryInput(deviceId string) dynamodb.QueryInput {
	device := model.Device{}
	device.DeviceId = deviceId
	cond := expression.Key(TableColumnDeviceId).Equal(expression.Value(&types.AttributeValueMemberS{Value: device.DeviceId}))
	expr, _ := expression.NewBuilder().WithKeyCondition(cond).Build()
	tableName := DeviceTableName
	queryInput := dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	}
	return queryInput
}

func createQueryInputExist() dynamodb.QueryInput {
	return createQueryInput("81a0aabc-7fe1-4b42-a387-d9f685a212e3")
}

func createQueryInputNotExist() dynamodb.QueryInput {
	return createQueryInput("b5c43ec4-4f23-4118-a497-09563f2ddf30")
}

func createQueryOutputEmpty() dynamodb.QueryOutput {
	var outputEmpty = dynamodb.QueryOutput{
		ConsumedCapacity: nil,
		Count:            0,
		Items:            []map[string]types.AttributeValue{},
		LastEvaluatedKey: nil,
		ScannedCount:     0,
		ResultMetadata:   middleware.Metadata{},
	}
	return outputEmpty
}

func createQueryOutputItemFound() dynamodb.QueryOutput {

	var outputItemFound = createQueryOutputEmpty()
	var itemFound = map[string]types.AttributeValue{
		"deviceId": &types.AttributeValueMemberS{Value: "81a0aabc-7fe1-4b42-a387-d9f685a212e3"},
		"Meta": &types.AttributeValueMemberM{
			Value: map[string]types.AttributeValue{
				"model":         &types.AttributeValueMemberS{Value: "model"},
				"hwVersion":     &types.AttributeValueMemberS{Value: "hwversion"},
				"osVersion":     &types.AttributeValueMemberS{Value: "osversion"},
				"appVersion":    &types.AttributeValueMemberS{Value: "appversion"},
				"apiLevel":      &types.AttributeValueMemberS{Value: "apilevel"},
				"securityPatch": &types.AttributeValueMemberS{Value: "securitypatch"},
				"nfcAvailable":  &types.AttributeValueMemberBOOL{Value: true},
				"nfcEnabled":    &types.AttributeValueMemberBOOL{Value: true},
			},
		},
		"createdAt": &types.AttributeValueMemberN{Value: "1643245234"},
		"updatedAt": &types.AttributeValueMemberN{Value: "1643245234"},
	}
	outputItemFound.Items = append(outputItemFound.Items, itemFound)
	return outputItemFound
}

func createQueryOutputTwoItemsFound() dynamodb.QueryOutput {
	//get one and add the same other one
	queryOutTwoItems := createQueryOutputItemFound()
	queryOutTwoItems.Items = append(queryOutTwoItems.Items, queryOutTwoItems.Items[0])
	return queryOutTwoItems
}

func createUpdateItemOutputCreated() *dynamodb.UpdateItemOutput {

	updateItemOutput := &dynamodb.UpdateItemOutput{
		Attributes:            map[string]types.AttributeValue{},
		ConsumedCapacity:      nil,
		ItemCollectionMetrics: nil,
		ResultMetadata:        middleware.Metadata{},
	}
	return updateItemOutput
}

func createUpdateItemOutputUpdated() *dynamodb.UpdateItemOutput {

	updateItemOutput := &dynamodb.UpdateItemOutput{
		Attributes: map[string]types.AttributeValue{
			"createdAt": &types.AttributeValueMemberN{Value: "1643245234"},
		},
		ConsumedCapacity:      nil,
		ItemCollectionMetrics: nil,
		ResultMetadata:        middleware.Metadata{},
	}
	return updateItemOutput
}

// MATCHERS need for testing
type updateMatcher struct{}

func (m updateMatcher) Matches(x interface{}) bool {
	input, ok := x.(*dynamodb.UpdateItemInput)
	if !ok {
		return false
	}

	valuesKeys := map[string]string{}
	namesKeys := map[int]string{}
	for k, field := range input.ExpressionAttributeNames {
		index, _ := strconv.Atoi(strings.Replace(k, "#", "", 1))
		valuesKeys[field] = fmt.Sprintf(":%d", index)
		namesKeys[index] = field
	}

	var updateBuilder expression.UpdateBuilder
	for i := 0; i < len(namesKeys); i++ {
		//Value is 0 because the value matcher is defined after
		if namesKeys[i] == TableColumnCreatedAt {
			updateBuilder = updateBuilder.Set(expression.Name(namesKeys[i]), expression.IfNotExists(expression.Name(namesKeys[i]), expression.Value(0)))
			continue
		}
		if namesKeys[i] != TableColumnDeviceId {
			updateBuilder = updateBuilder.Set(expression.Name(namesKeys[i]), expression.Value(0))
		}
	}

	expr, _ := expression.NewBuilder().WithUpdate(updateBuilder).Build()

	metaAttributesValues := extra.StructMatcher().
		Field("Value", extra.MapMatcher().
			Key("model", gomock.Eq(&types.AttributeValueMemberS{Value: "model"})).
			Key("hwVersion", gomock.Eq(&types.AttributeValueMemberS{Value: "hwversion"})).
			Key("osVersion", gomock.Eq(&types.AttributeValueMemberS{Value: "osversion"})).
			Key("appVersion", gomock.Eq(&types.AttributeValueMemberS{Value: "appversion"})).
			Key("apiLevel", gomock.Eq(&types.AttributeValueMemberS{Value: "apilevel"})).
			Key("securityPatch", gomock.Eq(&types.AttributeValueMemberS{Value: "securitypatch"})).
			Key("nfcAvailable", gomock.Eq(&types.AttributeValueMemberBOOL{Value: true})).
			Key("nfcEnabled", gomock.Eq(&types.AttributeValueMemberBOOL{Value: true})),
		)

	expressionAttributeValues := extra.MapMatcher().
		Key(valuesKeys["meta"], metaAttributesValues)

	keyValue := extra.MapMatcher().
		Key("deviceId", gomock.Eq(&types.AttributeValueMemberS{Value: "81a0aabc-7fe1-4b42-a387-d9f685a212e3"}))

	updateItemMatcher := extra.StructMatcher().
		Field("Key", keyValue).
		Field("TableName", gomock.Eq(aws.String(DeviceTableName))).
		Field("ExpressionAttributeNames", gomock.Eq(expr.Names())).
		Field("ExpressionAttributeValues", expressionAttributeValues).
		Field("ReturnValues", types.ReturnValueAllOld).
		Field("UpdateExpression", gomock.Eq(expr.Update()))

	return updateItemMatcher.Matches(x)
}

func (m updateMatcher) String() string {
	return "UpdateItemInput"
}

// ASSERTS compare results {{cookiecutter.model_name}}
//compare without dates (no compare dates because change in every insert)
func compare{{cookiecutter.model_name}}WithoutDates(a, b *model.Device) bool {
	aux := new(model.Device)
	*aux = *a
	aux.CreatedAt = b.CreatedAt
	aux.UpdatedAt = b.UpdatedAt
	return *aux == *b
}

//compare without updatedat (no compare updatedAt because change in every update)
func compare{{cookiecutter.model_name}}WithoutUpdatedAt(a, b *model.Device) bool {
	aux := new(model.Device)
	*aux = *a
	aux.UpdatedAt = b.UpdatedAt
	return *aux == *b
}
