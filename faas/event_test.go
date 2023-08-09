// Copyright 2021 Nitric Technologies Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package faas

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Event", func() {
	Context("ComposeEventMiddleware", func() {
		When("Creating a function with a single handler", func() {
			hndlr := ComposeEventMiddleware(func(ctx *EventContext, next EventHandler) (*EventContext, error) {
				ctx.Response.Success = false

				return next(ctx)
			})

			It("Should call the provided function", func() {
				ctx, err := hndlr(&EventContext{
					Response: &EventResponse{},
				}, nil)

				Expect(err).ToNot(HaveOccurred())
				Expect(ctx.Response.Success).To(BeEquivalentTo(false))
			})
		})

		When("Creating a function from multiple handlers", func() {
			callOrder := make([]string, 0)

			hndlr := ComposeEventMiddleware(
				func(ctx *EventContext, next EventHandler) (*EventContext, error) {
					callOrder = append(callOrder, "1")
					return next(ctx)
				},
				func(ctx *EventContext, next EventHandler) (*EventContext, error) {
					callOrder = append(callOrder, "2")
					return ctx, nil
				},
			)

			It("Should call the functions in the provided order", func() {
				_, err := hndlr(&EventContext{
					Response: &EventResponse{},
				}, nil)
				Expect(err).To(BeNil())

				Expect(callOrder).To(BeEquivalentTo([]string{"1", "2"}))
			})
		})

		When("Creating a function from multiple nested middlewares", func() {
			callOrder := make([]string, 0)

			hndlr := ComposeEventMiddleware(ComposeEventMiddleware(
				func(ctx *EventContext, next EventHandler) (*EventContext, error) {
					callOrder = append(callOrder, "1")
					return next(ctx)
				},
				func(ctx *EventContext, next EventHandler) (*EventContext, error) {
					callOrder = append(callOrder, "2")
					return next(ctx)
				},
			), ComposeEventMiddleware(
				func(ctx *EventContext, next EventHandler) (*EventContext, error) {
					callOrder = append(callOrder, "3")
					return ctx, nil
				},
			))

			It("Should call the functions in the provided order", func() {
				_, err := hndlr(&EventContext{
					Response: &EventResponse{},
				}, nil)
				Expect(err).To(BeNil())

				Expect(callOrder).To(BeEquivalentTo([]string{"1", "2", "3"}))
			})
		})
	})
})
