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

func DynamoDBItemToInterface(in *dynamodb.AttributeValue, o interface{}) {
	var itemConverter = func(i *dynamodb.AttributeValue) interface{} {
		fmt.Println(*i.S)
		switch {
		// 		_ struct{}
		// B []byte
		// A Binary data type.

		// B is automatically base64 encoded/decoded by the SDK.

		// BOOL *bool
		// A Boolean data type.

		// BS [][]byte
		// A Binary Set data type.

		// L []*AttributeValue
		// A List of attribute values.

		// M map[string]*AttributeValue
		// A Map of attribute values.

		// N *string
		// A Number data type.

		// NS []*string
		// A Number Set data type.

		// NULL *bool
		// A Null data type.

		case i.S != nil:
			fmt.Println(i.S)
			return *i.S

			// SS []*string
			// A String Set data type.

		}

		return nil
	}

	o = itemConverter(in)
}
