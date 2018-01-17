package errors

import (
	"fmt"
	"reflect"
)

//Callout represents a callout request error
type Callout struct {
	*errorWrapper
}

//NewCallout constructs a Callout error
func NewCallout(messages ...string) *Callout {
	return &Callout{
		new("Failed on requesting.", messages...),
	}
}

//Serialization represents a serialization error
type Serialization struct {
	*errorWrapper
}

//NewSerialization constructs a Serialization error
func NewSerialization(messages ...string) *Serialization {
	return &Serialization{
		new("Failed on serializing data.", messages...),
	}
}

//InvalidType represents a failed type assertion error
type InvalidType struct {
	*errorWrapper
}

//NewInvalidType constructs a InvalidType error
func NewInvalidType(actual, expected reflect.Type) *InvalidType {
	return &InvalidType{
		new(fmt.Sprintf("Actual type %s. Expected type %s.", actual, expected)),
	}
}
