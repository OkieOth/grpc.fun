// Attention, this file is generated. Manual changes get lost with the next
// run of the code generation.
// created by yacg (template: golang_types.mako v1.1.0)
package types


import (
    uuid "github.com/google/uuid"
    optional "github.com/okieoth/goptional"
)


type DummyType struct {

    Guid uuid.UUID

    Date time.Date

    Timestamp optional.Optional[time.Date]

    Time optional.Optional[time.Time]

    EnumValue DummyTypeEnumValueEnum

    ArrayValue []string

    Complex OptionalSubType
}

// Creates a DummyType object
func MakeDummyType() DummyType {
    var ret DummyType
    // TODO: initialize default values
    return ret
}

type OptionalDummyType struct {
	Value DummyType
	IsSet bool
}

// Creates a OptionalDummyType object
func MakeOptionalDummyType() OptionalDummyType {
    var ret OptionalDummyType
    // TODO: initialize default values
    return ret
}

func (m *OptionalDummyType) Set(v DummyType) {
	m.Value = v
	m.IsSet = true
}

func (m *OptionalDummyType) UnSet() {
	m.IsSet = false
}


type DummyTypeEnumValueEnum string

const (
    DummyTypeEnumValueEnum_option1 DummyTypeEnumValueEnum = "option1"
    DummyTypeEnumValueEnum_option2 = "option2"
    DummyTypeEnumValueEnum_option3 = "option3"
)


type OptionalDummyTypeEnumValueEnum struct {
	Value DummyTypeEnumValueEnum
	IsSet bool
}

// Creates a OptionalDummyTypeEnumValueEnum object
func MakeOptionalDummyTypeEnumValueEnum() OptionalDummyTypeEnumValueEnum {
    var ret OptionalDummyTypeEnumValueEnum
    // TODO: initialize default values
    return ret
}

func (m *OptionalDummyTypeEnumValueEnum) Set(v DummyTypeEnumValueEnum) {
	m.Value = v
	m.IsSet = true
}

func (m *OptionalDummyTypeEnumValueEnum) UnSet() {
	m.IsSet = false
}



type SubType struct {

    StringValue string

    IntValue int32

    FloatValue float64
}

// Creates a SubType object
func MakeSubType() SubType {
    var ret SubType
    // TODO: initialize default values
    return ret
}

type OptionalSubType struct {
	Value SubType
	IsSet bool
}

// Creates a OptionalSubType object
func MakeOptionalSubType() OptionalSubType {
    var ret OptionalSubType
    // TODO: initialize default values
    return ret
}

func (m *OptionalSubType) Set(v SubType) {
	m.Value = v
	m.IsSet = true
}

func (m *OptionalSubType) UnSet() {
	m.IsSet = false
}


