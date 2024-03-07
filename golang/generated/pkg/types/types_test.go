// Attention, this file is generated. Manual changes get lost with the next
// run of the code generation.
// created by yacg (template: golang.mako v1.0.1)
package types

import (
    "testing"
)

// how to test enums???

func TestDummyType(t *testing.T) {
    l := MakeDummyType()
    if l.Timestamp.IsSet {
        t.Errorf("Optional field 'Timestamp' is set, after creation, but should be unset")
    }

    if l.Time.IsSet {
        t.Errorf("Optional field 'Time' is set, after creation, but should be unset")
    }

    if len(l.ArrayValue) != 0 {
        t.Errorf("Array field 'ArrayValue' hasn't a 0 length, after creation")
    }

    if l.Complex.IsSet {
        t.Errorf("Optional field 'Complex' is set, after creation, but should be unset")
    }

}

func TestSubType(t *testing.T) {
    l := MakeSubType()
}

