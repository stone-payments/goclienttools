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

//Serializing represents a serialization error
type Serializing struct {
	*errorWrapper
}

//NewSerializing constructs a Serializing error
func NewSerializing(messages ...string) *Serializing {
	return &Serializing{
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
