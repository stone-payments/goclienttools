package http

import (
	"net/http"
	"reflect"

	"github.com/bouk/monkey"
	"github.com/crowleyfelix/base-client-go/errors"
	"github.com/levigross/grequests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response", func() {

	var (
		resp *response
	)

	BeforeEach(func() {
		resp = new(response)
		resp.Response = new(grequests.Response)
	})
	AfterEach(func() {
		monkey.UnpatchAll()
	})

	Describe("Ok", func() {
		It("should call grequest response Ok method", func() {
			resp.Response.Ok = true
			Expect(resp.Ok()).To(BeTrue())
		})
	})
	Describe("Error", func() {
		It("should call grequest response Error method", func() {
			resp.Response.Error = errors.Build(0)
			Expect(resp.Error()).To(Equal(resp.Response.Error))
		})
	})
	Describe("RawResponse", func() {
		It("should call grequest response RawResponse method", func() {
			resp.Response.RawResponse = new(http.Response)
			Expect(resp.RawResponse()).To(Equal(resp.Response.RawResponse))
		})
	})
	Describe("StatusCode", func() {
		It("should call grequest response StatusCode method", func() {
			resp.Response.StatusCode = 500
			Expect(resp.StatusCode()).To(Equal(resp.Response.StatusCode))
		})
	})
	Describe("Header", func() {
		It("should call grequest response Header method", func() {
			resp.Response.Header = make(http.Header)
			Expect(resp.Header()).To(Equal(resp.Response.Header))
		})
	})
	Describe("JSON", func() {
		var (
			err errors.Error
		)
		BeforeEach(func() {
			monkey.PatchInstanceMethod(reflect.TypeOf(new(grequests.Response)), "JSON", func(_ *grequests.Response, obj interface{}) error {
				return err
			})
		})
		It("should call grequest response JSON method", func() {
			err = errors.NewSerializing()
			Expect(resp.JSON(make(map[string]string))).To(Equal(err))
		})
	})
})
