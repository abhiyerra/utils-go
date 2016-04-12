package acksin_dynamodb

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// InterfaceToDynamoDBItem takes an interface{} and returns a
// dynamodb.AttributeValue which can be used to send data to DynamoDB.
func InterfaceToDynamoDBItem(i interface{}) *dynamodb.AttributeValue {
	switch v := i.(type) {
	case nil:
		return &dynamodb.AttributeValue{
			NULL: aws.Bool(v == nil),
		}
	case int:
		return &dynamodb.AttributeValue{
			N: aws.String(fmt.Sprintf("%d", v)),
		}
	case float64:
		return &dynamodb.AttributeValue{
			N: aws.String(fmt.Sprintf("%v", v)),
		}
	case []int:
		var val []*string
		for _, s := range v {
			val = append(val, aws.String(fmt.Sprintf("%d", s)))
		}

		return &dynamodb.AttributeValue{
			NS: val,
		}
	case []float64:
		var val []*string
		for _, s := range v {
			val = append(val, aws.String(fmt.Sprintf("%v", s)))
		}

		return &dynamodb.AttributeValue{
			NS: val,
		}
	case string:
		if v != "" {
			return &dynamodb.AttributeValue{
				S: aws.String(v),
			}
		}
	case []string:
		var val []*string
		for _, s := range v {
			val = append(val, aws.String(s))
		}

		return &dynamodb.AttributeValue{
			SS: val,
		}
	case bool:
		return &dynamodb.AttributeValue{
			BOOL: aws.Bool(v),
		}
	case []byte:
		return &dynamodb.AttributeValue{
			B: v,
		}
	case [][]byte:
		return &dynamodb.AttributeValue{
			BS: v,
		}
	case []interface{}:
		var val []*dynamodb.AttributeValue

		for _, li := range v {
			val = append(val, InterfaceToDynamoDBItem(li))
		}

		return &dynamodb.AttributeValue{
			L: val,
		}
	case map[string]interface{}:
		out2 := make(map[string]*dynamodb.AttributeValue)

		for k, v2 := range v {
			if ret := InterfaceToDynamoDBItem(v2); ret != nil {
				out2[k] = ret
			}
		}

		return &dynamodb.AttributeValue{
			M: out2,
		}

	default:
		// do nothing
		return nil
	}

	return nil
}
