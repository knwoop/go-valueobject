package valueobject_test

import (
	"fmt"
	"testing"

	"github.com/knwoop/go-valueobject"
)

// Create a different type of value object used in Equals() check
type NotEmailAddress struct {
	value string
}

func (n NotEmailAddress) String() string {
	return n.value
}

func (n NotEmailAddress) Equals(value valueobject.Value) bool {
	return false
}

func ExampleEmailAddress_String() {
	numeral, _ := valueobject.NewEmailAddress("go@test.com")
	fmt.Println(numeral.String())
}

func TestShouldntAcceptInvalidEmailAddress(t *testing.T) {
	_, err := valueobject.NewEmailAddress("invalid")
	if err == nil {
		t.Fatal("We expected an error")
	}
}

func ExampleEmailAddress_Equals() {
	a, _ := valueobject.NewEmailAddress("go@test.com")
	b, _ := valueobject.NewEmailAddress("go@test.com")

	fmt.Println(a.Equals(b))
}

func TestShouldCompareTwoEmailAddresssAsNotEqual(t *testing.T) {
	a, _ := valueobject.NewEmailAddress("go@test.com")
	b, _ := valueobject.NewEmailAddress("c++@test.com")
	if a.Equals(b) == true {
		t.Fatal("Shouldn't be same value")
	}
}

func TestShouldNotBeEqualIfNotEmailAddress(t *testing.T) {
	var notEmailAddress NotEmailAddress
	numeral, _ := valueobject.NewEmailAddress("go@test.com")

	if numeral.Equals(notEmailAddress) == true {
		t.Fatal("Different value object types can not be equal")
	}
}
