package acksin_dynamodb

import (
	"testing"
)

func TestInterfaceToDynamoDBItem(t *testing.T) {
	in := make(map[string]interface{})

	in["Number"] = 1
	in["NumberList"] = []int{1, 2, 3}
	in["Float"] = 0.123
	in["FloatList"] = []float64{1.1, 2.2, 3.3}
	in["Null"] = nil
	in["String"] = "hey world"
	in["StringEmpty"] = ""
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
	innerMap["String"] = "hello world"
	innerMap["StringEmpty"] = ""
	innerMap["StringList"] = []string{"hey", "world"}
	innerMap["Bool"] = false
	innerMap["Bytes"] = []byte("meh")
	innerMap["Bytes"] = [][]byte{[]byte("meow"), []byte("cat")}

	in["Map"] = innerMap

	results := InterfaceToDynamoDBItem(in)

	if *results.M["String"].S != "hey world" {
		t.Errorf("Invalid object for Number %s\n", results.M["String"].S)
	}

	if _, ok := results.M["StringEmpty"]; ok {
		t.Errorf("Invalid object for Number %s\n", results.M["StringEmpty"].S)
	}

	if *results.M["Map"].M["String"].S != "hello world" {
		t.Errorf("Invalid object for Number %s\n", *results.M["Map"].M["String"].S)
	}

	if _, ok := results.M["Map"].M["StringEmpty"]; ok {
		t.Errorf("Invalid object for Number %s\n", *results.M["Map"].M["String"].S)
	}

	if *results.M["Number"].N != "1" {
		t.Errorf("Invalid object for Number %s\n", results.M["Number"])
	}

	if *results.M["Map"].M["Float"].N != "0.123" {
		t.Errorf("Invalid object for Number %v\n", *results.M["Map"])
	}
}
