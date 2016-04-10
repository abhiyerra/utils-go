package acksin_dynamodb

import (
	"testing"
)

func TestMapToDynamoDBItem(t *testing.T) {
	in := make(map[string]interface{})

	in["Number"] = 1
	in["NumberList"] = []int{1, 2, 3}
	in["Float"] = 0.123
	in["FloatList"] = []float64{1.1, 2.2, 3.3}
	in["Null"] = nil
	in["String"] = "hey world"
	in["StringList"] = []string{"hey", "world"}
	in["Bool"] = false
	in["Bytes"] = []byte("meh")
	in["Bytes"] = [][]byte{[]byte("meow"), []byte("cat")}

	innerMap := make(map[string]interface{})
	innerMap["Number"] = 1
	innerMap["NumberList"] = []int{1, 2, 3}
	innerMap["Float"] = 0.123
	innerMap["FloatList"] = []float64{1.1, 2.2, 3.3}
	innerMap["Null"] = nil
	innerMap["String"] = "hey world"
	innerMap["StringList"] = []string{"hey", "world"}
	innerMap["Bool"] = false
	innerMap["Bytes"] = []byte("meh")
	innerMap["Bytes"] = [][]byte{[]byte("meow"), []byte("cat")}

	in["Map"] = innerMap

	results := MapToDynamoDBItem(in)

	if *results["Number"].N != "1" {
		t.Errorf("Invalid object for Number %s\n", results["Number"])
	}

	if *results["Map"].M["Float"].N != "0.123" {
		t.Errorf("Invalid object for Number %v\n", *results["Map"])
	}
}
