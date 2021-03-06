// Copyright 2019 Michael Li <dyc92@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package chi_test

import (
	"bytes"
	"github.com/dyc92/mir"
	"github.com/go-chi/chi"
	"net/http/httptest"

	. "github.com/dyc92/mir/module/chi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Core", func() {
	var (
		router chi.Router
		w      *httptest.ResponseRecorder
		err    error
	)

	JustBeforeEach(func() {
		w = httptest.NewRecorder()
	})

	Context("check Mir function", func() {
		BeforeEach(func() {
			router = chi.NewRouter()
			mirE := Mir(router)
			err = mir.Register(mirE, &entry{Chain: mirChain()})
		})

		It("no error", func() {
			Expect(err).Should(BeNil())
		})

		It("no nil", func() {
			Expect(router).ShouldNot(BeNil())
		})

		It("handle add", func() {
			body := bytes.NewReader([]byte("hello"))
			r := httptest.NewRequest(mir.MethodPost, "/v1/add/10086/", body)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Add:10086:hello"))
		})

		It("handler index", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v1/index/", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Index"))
		})

		It("handle articles", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v1/articles/golang/10086", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("GetArticles:golang:10086"))
		})

		It("handle chain func1", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v1/chainfunc1", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("ChainFunc1"))
		})

		It("handle chain func2", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v1/chainfunc2", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("GetChainFunc2"))
		})
	})

	Context("check Register function", func() {
		BeforeEach(func() {
			router = chi.NewRouter()
			err = Register(router, &entry{Group: "/v2", Chain: mirChain()})
		})

		It("no error", func() {
			Expect(err).Should(BeNil())
		})

		It("no nil", func() {
			Expect(router).ShouldNot(BeNil())
		})

		It("handle add", func() {
			body := bytes.NewReader([]byte("hello"))
			r := httptest.NewRequest(mir.MethodPost, "/v2/add/10086/", body)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Add:10086:hello"))
		})

		It("handler index", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/index/", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Index"))
		})

		It("handle articles", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/articles/golang/10086", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("GetArticles:golang:10086"))
		})

		It("handle chain func1", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/chainfunc1", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("ChainFunc1"))
		})

		It("handle chain func2", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/chainfunc2", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("GetChainFunc2"))
		})
	})

	Context("check Register entries", func() {
		BeforeEach(func() {
			router = chi.NewRouter()
			err = Register(router, &entry{}, &entry{Group: "v2", Chain: mirChain()})
		})

		It("no error", func() {
			Expect(err).Should(BeNil())
		})

		It("no nil", func() {
			Expect(router).ShouldNot(BeNil())
		})

		It("handle v1 add", func() {
			body := bytes.NewReader([]byte("hello"))
			r := httptest.NewRequest(mir.MethodPost, "/v1/add/10086/", body)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Add:10086:hello"))
		})

		It("handler v1 index", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v1/index/", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Index"))
		})

		It("handle v1 articles", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v1/articles/golang/10086", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("GetArticles:golang:10086"))
		})

		It("handle v1 chain func1", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v1/chainfunc1", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("ChainFunc1"))
		})

		It("handle v1 chain func2", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v1/chainfunc2", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("GetChainFunc2"))
		})

		It("handle v2 add", func() {
			body := bytes.NewReader([]byte("hello"))
			r := httptest.NewRequest(mir.MethodPost, "/v2/add/10086/", body)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Add:10086:hello"))
		})

		It("handler v2 index", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/index/", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Index"))
		})

		It("handle v2 articles", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/articles/golang/10086", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("GetArticles:golang:10086"))
		})

		It("handle v2 chain func1", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/chainfunc1", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("ChainFunc1"))
		})

		It("handle v2 chain func2", func() {
			r := httptest.NewRequest(mir.MethodGet, "/v2/chainfunc2", nil)
			router.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("GetChainFunc2"))
		})
	})
})
