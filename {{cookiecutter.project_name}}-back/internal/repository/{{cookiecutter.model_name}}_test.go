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
	{{cookiecutter.model_name.capitalize()}}TableName = "{{cookiecutter.project_name}}-{{cookiecutter.model_name}}s"
)

func TestNew{{cookiecutter.model_name.capitalize()}}Repository(t *testing.T) {
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
		want {{cookiecutter.model_name.capitalize()}}Repository
	}{
		{
			name: "instantiates ok",
			args: args{
				lgr:     logger,
				dynClt:  dynCltMock,
				tblName: {{cookiecutter.model_name.capitalize()}}TableName,
			},
			want: &{{cookiecutter.model_name}}Repository{
				lgr:     logger,
				dynClt:  dynCltMock,
				tblName: {{cookiecutter.model_name.capitalize()}}TableName,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New{{cookiecutter.model_name.capitalize()}}Repository(tt.args.lgr, tt.args.dynClt, tt.args.tblName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New{{cookiecutter.model_name.capitalize()}}Repository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_{{cookiecutter.model_name}}Repository_Upsert(t *testing.T) {

	ctrl := gomock.NewController(t)
	dynCltMock := mock.NewMockServiceDynamo(ctrl)
	lg := utillogger.NewEmpty()

	//arguments for create {{cookiecutter.model_name}} repository
	argsRepository := {{cookiecutter.model_name}}Repository{
		lgr:     lg,
		dynClt:  dynCltMock,
		tblName: {{cookiecutter.model_name.capitalize()}}TableName,
	}

	//arguments for function to testing Insert
	type args struct {
		ctx    context.Context
		{{cookiecutter.model_name}} model.{{cookiecutter.model_name.capitalize()}}
	}

	//Test suite
	tests := []struct {
		name           string
		argsRepo       {{cookiecutter.model_name}}Repository
		configureMocks func()
		args           args
		wantAssert     func(got model.{{cookiecutter.model_name.capitalize()}}) bool
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
				{{cookiecutter.model_name}}: getTestData().{{cookiecutter.model_name}}New,
			},
			wantAssert: func(got model.{{cookiecutter.model_name.capitalize()}}) bool { return true },
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
				{{cookiecutter.model_name}}: getTestData().{{cookiecutter.model_name}}New,
			},
			wantAssert: func(got model.{{cookiecutter.model_name.capitalize()}}) bool {
				want := getTestData().{{cookiecutter.model_name}}New
				return compare{{cookiecutter.model_name.capitalize()}}sWithoutDates(&got, &want)
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
				{{cookiecutter.model_name}}: getTestData().{{cookiecutter.model_name}}New,
			},
			wantAssert: func(got model.{{cookiecutter.model_name.capitalize()}}) bool {
				want := getTestData().{{cookiecutter.model_name}}Found
				return compare{{cookiecutter.model_name.capitalize()}}sWithoutUpdatedAt(&got, &want)
			},
			wantErr: false,
		},
	}
	//Execute test suit
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dr := New{{cookiecutter.model_name.capitalize()}}Repository(tt.argsRepo.lgr, tt.argsRepo.dynClt, tt.argsRepo.tblName)

			//Configure mock
			tt.configureMocks()

			//Testing Upsert
			got, err := dr.Upsert(tt.args.ctx, tt.args.{{cookiecutter.model_name}})

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

func Test_{{cookiecutter.model_name}}Repository_FilterByID(t *testing.T) {

	//Create mock dynamoClient (ServiceDynamo)
	ctrl := gomock.NewController(t)
	dynCltMock := mock.NewMockServiceDynamo(ctrl)

	//arguments for create {{cookiecutter.model_name}} repository
	argsRepository := {{cookiecutter.model_name}}Repository{
		lgr:     utillogger.NewEmpty(),
		dynClt:  dynCltMock,
		tblName: {{cookiecutter.model_name.capitalize()}}TableName,
	}

	//arguments for function to testing FilterByID
	type args struct {
		ctx      context.Context
		{{cookiecutter.model_name}}Id string
	}

	//Test suite
	tests := []struct {
		name           string
		argsRepo       {{cookiecutter.model_name}}Repository
		configureMocks func()
		args           args
		want           model.{{cookiecutter.model_name.capitalize()}}
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
				{{cookiecutter.model_name}}Id: "81a0aabc-7fe1-4b42-a387-d9f685a212e3",
				ctx:      context.TODO(),
			},
			want:    getTestData().{{cookiecutter.model_name}}Empty,
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
				{{cookiecutter.model_name}}Id: "b5c43ec4-4f23-4118-a497-09563f2ddf30",
				ctx:      context.TODO(),
			},
			want:    getTestData().{{cookiecutter.model_name}}Empty,
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
				{{cookiecutter.model_name}}Id: "b5c43ec4-4f23-4118-a497-09563f2ddf30",
				ctx:      context.TODO(),
			},
			want:    getTestData().{{cookiecutter.model_name}}Empty,
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
				{{cookiecutter.model_name}}Id: "81a0aabc-7fe1-4b42-a387-d9f685a212e3",
				ctx:      context.TODO(),
			},
			want:    getTestData().{{cookiecutter.model_name}}Found,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dr := New{{cookiecutter.model_name.capitalize()}}Repository(tt.argsRepo.lgr, tt.argsRepo.dynClt, tt.argsRepo.tblName)

			//Configure mock
			tt.configureMocks()

			//Testing FilterByID
			got, err := dr.FilterByID(tt.args.ctx, tt.args.{{cookiecutter.model_name}}Id)

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
	{{cookiecutter.model_name}}Empty              model.{{cookiecutter.model_name.capitalize()}}
	{{cookiecutter.model_name}}New                model.{{cookiecutter.model_name.capitalize()}}
	{{cookiecutter.model_name}}Found              model.{{cookiecutter.model_name.capitalize()}}
	queryInputExist          *dynamodb.QueryInput
	queryInputNotExist       *dynamodb.QueryInput
	queryOutputEmpty         *dynamodb.QueryOutput
	queryOutputItemFound     *dynamodb.QueryOutput
	queryOutputTwoItemsFound *dynamodb.QueryOutput
}

func getTestData() TestData {

	{{cookiecutter.model_name}}New := create{{cookiecutter.model_name.capitalize()}}New()
	{{cookiecutter.model_name}}Found := create{{cookiecutter.model_name.capitalize()}}Found()
	queryInputExist := createQueryInputExist()
	queryInputNotExist := createQueryInputNotExist()
	queryOutputEmpty := createQueryOutputEmpty()
	queryOutputItemFound := createQueryOutputItemFound()
	queryOutputTwoItemsFound := createQueryOutputTwoItemsFound()

	testData := TestData{
		{{cookiecutter.model_name}}Empty:              model.{{cookiecutter.model_name.capitalize()}}{},
		{{cookiecutter.model_name}}New:                {{cookiecutter.model_name}}New,
		{{cookiecutter.model_name}}Found:              {{cookiecutter.model_name}}Found,
		queryInputExist:          &queryInputExist,
		queryInputNotExist:       &queryInputNotExist,
		queryOutputEmpty:         &queryOutputEmpty,
		queryOutputItemFound:     &queryOutputItemFound,
		queryOutputTwoItemsFound: &queryOutputTwoItemsFound,
	}
	return testData
}

func create{{cookiecutter.model_name.capitalize()}}New() model.{{cookiecutter.model_name.capitalize()}} {
	{{cookiecutter.model_name}}New := model.{{cookiecutter.model_name.capitalize()}}{
		{{cookiecutter.model_name.capitalize()}}Id:      "81a0aabc-7fe1-4b42-a387-d9f685a212e3",
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
	return {{cookiecutter.model_name}}New
}

func create{{cookiecutter.model_name.capitalize()}}Found() model.{{cookiecutter.model_name.capitalize()}} {
	{{cookiecutter.model_name}}Found := model.{{cookiecutter.model_name.capitalize()}}{
		{{cookiecutter.model_name.capitalize()}}Id:      "81a0aabc-7fe1-4b42-a387-d9f685a212e3",
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
	return {{cookiecutter.model_name}}Found
}

func createQueryInput({{cookiecutter.model_name}}Id string) dynamodb.QueryInput {
	{{cookiecutter.model_name}} := model.{{cookiecutter.model_name.capitalize()}}{}
	{{cookiecutter.model_name}}.{{cookiecutter.model_name.capitalize()}}Id = {{cookiecutter.model_name}}Id
	cond := expression.Key(TableColumn{{cookiecutter.model_name.capitalize()}}Id).Equal(expression.Value(&types.AttributeValueMemberS{Value: {{cookiecutter.model_name}}.{{cookiecutter.model_name.capitalize()}}Id}))
	expr, _ := expression.NewBuilder().WithKeyCondition(cond).Build()
	tableName := {{cookiecutter.model_name.capitalize()}}TableName
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
		"{{cookiecutter.model_name}}Id": &types.AttributeValueMemberS{Value: "81a0aabc-7fe1-4b42-a387-d9f685a212e3"},
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
		if namesKeys[i] != TableColumn{{cookiecutter.model_name.capitalize()}}Id {
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
		Key("{{cookiecutter.model_name}}Id", gomock.Eq(&types.AttributeValueMemberS{Value: "81a0aabc-7fe1-4b42-a387-d9f685a212e3"}))

	updateItemMatcher := extra.StructMatcher().
		Field("Key", keyValue).
		Field("TableName", gomock.Eq(aws.String({{cookiecutter.model_name.capitalize()}}TableName))).
		Field("ExpressionAttributeNames", gomock.Eq(expr.Names())).
		Field("ExpressionAttributeValues", expressionAttributeValues).
		Field("ReturnValues", types.ReturnValueAllOld).
		Field("UpdateExpression", gomock.Eq(expr.Update()))

	return updateItemMatcher.Matches(x)
}

func (m updateMatcher) String() string {
	return "UpdateItemInput"
}

// ASSERTS compare results {{cookiecutter.model_name.capitalize()}}s
//compare without dates (no compare dates because change in every insert)
func compare{{cookiecutter.model_name.capitalize()}}sWithoutDates(a, b *model.{{cookiecutter.model_name.capitalize()}}) bool {
	aux := new(model.{{cookiecutter.model_name.capitalize()}})
	*aux = *a
	aux.CreatedAt = b.CreatedAt
	aux.UpdatedAt = b.UpdatedAt
	return *aux == *b
}

//compare without updatedat (no compare updatedAt because change in every update)
func compare{{cookiecutter.model_name.capitalize()}}sWithoutUpdatedAt(a, b *model.{{cookiecutter.model_name.capitalize()}}) bool {
	aux := new(model.{{cookiecutter.model_name.capitalize()}})
	*aux = *a
	aux.UpdatedAt = b.UpdatedAt
	return *aux == *b
}
