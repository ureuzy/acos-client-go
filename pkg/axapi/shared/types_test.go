package shared_test

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"github.com/ureuzy/acos-client-go/pkg/axapi/shared"
	"testing"
)

func Test_UnmarshalBooleanTrue(t *testing.T) {
	RegisterTestingT(t)

	dat := []byte("{ \"bool\": 1 }")
	b := new(test)
	err := json.Unmarshal(dat, b)

	Ω(err).ShouldNot(HaveOccurred())
	var a shared.Boolean = true
	Ω(b.Bool).Should(BeIdenticalTo(a))
}

func Test_UnmarshalBooleanFalse(t *testing.T) {
	RegisterTestingT(t)

	dat := []byte("{ \"bool\": 0 }")
	b := new(test)
	err := json.Unmarshal(dat, b)

	Ω(err).ShouldNot(HaveOccurred())
	var a shared.Boolean = false
	Ω(b.Bool).Should(BeIdenticalTo(a))
}

func Test_UnmarshalBooleanTrueToo(t *testing.T) {
	RegisterTestingT(t)

	dat := []byte("{ \"bool\": 222 }")
	b := new(test)
	err := json.Unmarshal(dat, b)

	Ω(err).ShouldNot(HaveOccurred())
	var a shared.Boolean = true
	Ω(b.Bool).Should(BeIdenticalTo(a))
}

type test struct {
	Bool shared.Boolean `json:"bool"`
}
