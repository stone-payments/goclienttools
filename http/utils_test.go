package http

import (
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {

	PDescribe("SwitchBody", func() {

	})

	Describe("ToMap", func() {
		i := 0
		obj := struct {
			FieldString   string  `param:"field_string"`
			FieldPtString *string `param:"field_ptstring"`
			FieldInt      int     `param:"field_int"`
			FieldPtInt    *int    `param:"field_ptint"`
		}{"value", nil, 1, &i}

		It("should turn object into map", func() {
			Expect(ToMap(obj)).To(Equal(map[string]string{
				"field_string": "value",
				"field_int":    "1",
				"field_ptint":  strconv.Itoa(i),
			}))
		})
	})
})
