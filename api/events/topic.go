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
	"time"

	"github.com/nitrictech/go-sdk/api/errors"
	"github.com/nitrictech/go-sdk/api/errors/codes"
	v1 "github.com/nitrictech/nitric/core/pkg/api/nitric/v1"
	"github.com/nitrictech/protoutils"
)

type PublishOption = func(*v1.EventPublishRequest)

// Topic for pub/sub async messaging.
type Topic interface {
	// Name returns the Topic name.
	Name() string

	// Publish will publish the provided events on the topic.
	Publish(context.Context, *Event, ...PublishOption) (*Event, error)
}

type topicImpl struct {
	name string
	eventClient   v1.EventServiceClient
}

func (s *topicImpl) Name() string {
	return s.name
}

// WithDelay - Delay event publishing by the given duration
func WithDelay(duration time.Duration) func(*v1.EventPublishRequest) {
	return func(epr *v1.EventPublishRequest) {
		epr.Delay = uint32(duration.Seconds())
	}
}

func (s *topicImpl) Publish(ctx context.Context, evt *Event, opts ...PublishOption) (*Event, error) {
	// Convert payload to Protobuf Struct
	payloadStruct, err := protoutils.NewStruct(evt.Payload)
	if err != nil {
		return nil, errors.NewWithCause(codes.InvalidArgument, "Topic.Publish", err)
	}

	event := &v1.EventPublishRequest{
		Topic: s.name,
		Event: &v1.NitricEvent{
			Id:          evt.ID,
			Payload:     payloadStruct,
			PayloadType: evt.PayloadType,
		},
	}

	// Apply options to the event payload
	for _, opt := range opts {
		opt(event)
	}

	r, err := s.eventClient.Publish(ctx, event)
	if err != nil {
		return nil, errors.FromGrpcError(err)
	}

	// Return a reference to a new event with a populated ID
	return &Event{
		ID:          r.GetId(),
		Payload:     evt.Payload,
		PayloadType: evt.PayloadType,
	}, nil
}
