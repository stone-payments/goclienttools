package http

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stone-payments/goclienttools/errors"
)

type mockInterceptor struct {
	err              errors.Error
	onRequestCalled  bool
	onResponseCalled bool
}

func (m *mockInterceptor) OnRequest(req Request) Request {
	m.onRequestCalled = true
	return req
}
func (m *mockInterceptor) OnResponse(resp Response) (Response, errors.Error) {
	m.onResponseCalled = true
	return resp, m.err
}

var _ = Describe("Manager", func() {

	var (
		url          = "url"
		interceptors []Interceptor
		man          *manager
	)

	BeforeEach(func() {
		man = NewManager("url", interceptors...).(*manager)
	})
	Describe("NewManager", func() {
		It("should set interceptors", func() {
			Expect(man.interceptors).To(Equal(interceptors))
		})
		It("should set base url", func() {
			Expect(man.baseURL).To(Equal(url))
		})
	})
	Describe("BuildURL", func() {
		It("should build url with params", func() {
			Expect(man.BuildURL("/%v/%v/%v", 1, 2, 3)).To(Equal(url + "/1/2/3"))
		})
	})
	Describe("Request", func() {

		var (
			req *request
		)

		var (
			resp   *response
			err    error
			method = func(_ string, _ Request) (Response, error) {
				resp = new(response)
				return resp, err
			}
		)
		var (
			externalInterceptor *mockInterceptor
			localInterceptor    *mockInterceptor
			paramInterceptor    *mockInterceptor
		)

		//output
		var (
			actualResp Response
			actualErr  errors.Error
		)

		BeforeEach(func() {
			req = new(request)
			err = nil

			externalInterceptor = new(mockInterceptor)
			localInterceptor = new(mockInterceptor)
			paramInterceptor = new(mockInterceptor)

			Interceptors = []Interceptor{externalInterceptor}
			man.interceptors = []Interceptor{localInterceptor}
		})
		JustBeforeEach(func() {
			actualResp, actualErr = man.Request(method, url, req, paramInterceptor)
		})

		Context("When request failled", func() {
			BeforeEach(func() {
				err = errors.Build(0)
			})
			It("should call OnRequest of external interceptors", func() {
				Expect(externalInterceptor.onRequestCalled).To(BeTrue())
				Expect(externalInterceptor.onResponseCalled).To(BeFalse())
			})
			It("should call OnRequest of local interceptors", func() {
				Expect(localInterceptor.onRequestCalled).To(BeTrue())
				Expect(localInterceptor.onResponseCalled).To(BeFalse())
			})
			It("should call OnRequest of parameterized interceptors", func() {
				Expect(paramInterceptor.onRequestCalled).To(BeTrue())
				Expect(paramInterceptor.onResponseCalled).To(BeFalse())
			})
			It("should return an error", func() {
				Expect(actualResp).To(BeNil())
				Expect(actualErr).To(BeAssignableToTypeOf(errors.NewCallout()))
			})
		})
		Context("When request success", func() {
			Context("and external interceptors return an error", func() {
				BeforeEach(func() {
					externalInterceptor.err = errors.Build(0)
				})
				It("should return an error", func() {
					Expect(actualErr).To(BeAssignableToTypeOf(externalInterceptor.err))
				})
				It("should not call next interceptors", func() {
					Expect(localInterceptor.onResponseCalled).To(BeFalse())
					Expect(paramInterceptor.onResponseCalled).To(BeFalse())
				})
			})
			Context("and local interceptors return an error", func() {
				BeforeEach(func() {
					localInterceptor.err = errors.Build(0)
				})
				It("should return an error", func() {
					Expect(actualErr).To(BeAssignableToTypeOf(localInterceptor.err))
				})
				It("should not call next interceptors", func() {
					Expect(paramInterceptor.onResponseCalled).To(BeFalse())
				})
			})
			Context("and parameterized interceptors return an error", func() {
				BeforeEach(func() {
					paramInterceptor.err = errors.Build(0)
				})
				It("should return an error", func() {
					Expect(actualErr).To(BeAssignableToTypeOf(paramInterceptor.err))
				})
			})
			Context("and executed with success", func() {
				It("should return response", func() {
					Expect(actualErr).To(BeNil())
					Expect(actualResp).To(Equal(resp))
				})
			})
		})
	})
	Describe("allInterceptors", func() {
		var all []Interceptor

		BeforeEach(func() {
			Interceptors = []Interceptor{new(mockInterceptor)}
			man.interceptors = []Interceptor{new(mockInterceptor)}
			all = man.allInterceptors(new(mockInterceptor))
		})
		It("should return all available interceptors", func() {
			Expect(all).To(HaveLen(3))
		})
	})
})
