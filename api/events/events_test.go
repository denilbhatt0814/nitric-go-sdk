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

package events

import (
	"fmt"
	"os"

	"github.com/golang/mock/gomock"
	v1 "github.com/nitrictech/go-sdk/interfaces/nitric/v1"
	mock_v1 "github.com/nitrictech/go-sdk/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Events", func() {
	ctrl := gomock.NewController(GinkgoT())
	// Skip for now
	Context("New", func() {
		When("calling with no membrane available", func() {
			os.Setenv("NITRIC_SERVICE_DIAL_TIMEOUT", "10")
			client, err := New()

			It("should return a nil client", func() {
				Expect(client).To(BeNil())
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Context("Topics", func() {
		When("the membrane server returns an error", func() {
			mockEvt := mock_v1.NewMockEventServiceClient(ctrl)
			mockTopic := mock_v1.NewMockTopicServiceClient(ctrl)

			mockTopic.EXPECT().List(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("mock error"))

			evt := &eventsImpl{
				ec: mockEvt,
				tc: mockTopic,
			}

			It("Should pass through the error", func() {
				_, err := evt.Topics()

				Expect(err).To(HaveOccurred())
			})
		})

		When("topics are available", func() {
			mockEvt := mock_v1.NewMockEventServiceClient(ctrl)
			mockTopic := mock_v1.NewMockTopicServiceClient(ctrl)

			mockTopic.EXPECT().List(gomock.Any(), gomock.Any()).Return(&v1.TopicListResponse{
				Topics: []*v1.NitricTopic{
					{
						Name: "test-topic",
					},
				},
			}, nil)

			evt := &eventsImpl{
				ec: mockEvt,
				tc: mockTopic,
			}

			topics, _ := evt.Topics()

			It("should return the same number of available topics", func() {
				Expect(topics).To(HaveLen(1))
			})

			topic, ok := topics[0].(*topicImpl)

			It("should return an instance of topicImpl", func() {
				Expect(ok).To(BeTrue())
			})

			It("should copy through the topics name", func() {
				Expect(topic.name).To(Equal("test-topic"))
			})

			It("should have the same EventServiceClient as the base client", func() {
				Expect(topic.ec).To(Equal(mockEvt))
			})
		})
	})

	Context("Topic", func() {
		When("directly creating a topic reference", func() {
			mockEvt := mock_v1.NewMockEventServiceClient(ctrl)
			mockTopic := mock_v1.NewMockTopicServiceClient(ctrl)

			evt := &eventsImpl{
				ec: mockEvt,
				tc: mockTopic,
			}

			topic := evt.Topic("test-topic")

			topicI, ok := topic.(*topicImpl)

			It("should return an instance of topicImpl", func() {
				Expect(ok).To(BeTrue())
			})

			It("should have the provided name", func() {
				Expect(topicI.name).To(Equal("test-topic"))
			})

			It("should share the same event client as the root eventing client", func() {
				Expect(topicI.ec).To(Equal(mockEvt))
			})
		})
	})

	Context("Topic.Publish", func() {
		When("the topic exists", func() {
			mockEvt := mock_v1.NewMockEventServiceClient(ctrl)
			mockTopic := mock_v1.NewMockTopicServiceClient(ctrl)

			mockEvt.EXPECT().Publish(gomock.Any(), gomock.Any()).Return(&v1.EventPublishResponse{
				Id: "1234",
			}, nil)

			evt := &eventsImpl{
				ec: mockEvt,
				tc: mockTopic,
			}

			returnEvt, err := evt.Topic("test-topic").Publish(&Event{
				Payload: map[string]interface{}{
					"test": "test",
				},
				PayloadType: "test-payload",
			})

			It("should not return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should have the ID generated by the server", func() {
				Expect(returnEvt.ID).To(Equal("1234"))
			})
		})

		When("the topic does not exist", func() {
			mockEvt := mock_v1.NewMockEventServiceClient(ctrl)
			mockTopic := mock_v1.NewMockTopicServiceClient(ctrl)

			mockEvt.EXPECT().Publish(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("mock error"))

			evt := &eventsImpl{
				ec: mockEvt,
				tc: mockTopic,
			}

			_, err := evt.Topic("test-topic").Publish(&Event{
				Payload: map[string]interface{}{
					"test": "test",
				},
				PayloadType: "test-payload",
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
