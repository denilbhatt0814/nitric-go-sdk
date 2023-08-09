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
	"context"

	"google.golang.org/grpc"

	"github.com/nitrictech/go-sdk/api/errors"
	"github.com/nitrictech/go-sdk/api/errors/codes"
	"github.com/nitrictech/go-sdk/constants"
	v1 "github.com/nitrictech/nitric/core/pkg/api/nitric/v1"
)

// Events
type Events interface {
	// Topic - Retrieve a Topic reference
	Topic(name string) Topic
	// Topics - Retrievs a list of available Topic references
	Topics() ([]Topic, error)
}

type eventsImpl struct {
	eventClient v1.EventServiceClient
	topicClient v1.TopicServiceClient
}

func (s *eventsImpl) Topic(name string) Topic {
	// Just return the straight topic reference
	// we can fail if the topic does not exist
	return &topicImpl{
		name:        name,
		eventClient: s.eventClient,
	}
}

func (s *eventsImpl) Topics() ([]Topic, error) {
	r, err := s.topicClient.List(context.TODO(), &v1.TopicListRequest{})
	if err != nil {
		return nil, err
	}

	ts := make([]Topic, 0)
	for _, topic := range r.GetTopics() {
		ts = append(ts, s.Topic(topic.GetName()))
	}

	return ts, nil
}

// New - Construct a new Eventing Client with default options
func New() (Events, error) {
	conn, err := grpc.Dial(
		constants.NitricAddress(),
		constants.DefaultOptions()...,
	)
	if err != nil {
		return nil, errors.NewWithCause(codes.Unavailable, "Unable to dial Events service", err)
	}

	ec := v1.NewEventServiceClient(conn)
	tc := v1.NewTopicServiceClient(conn)

	return &eventsImpl{
		eventClient: ec,
		topicClient: tc,
	}, nil
}
