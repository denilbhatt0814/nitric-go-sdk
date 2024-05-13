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

package handler

import (
	"github.com/nitrictech/go-sdk/api/storage"
	http "github.com/nitrictech/nitric/core/pkg/proto/apis/v1"
	storagepb "github.com/nitrictech/nitric/core/pkg/proto/storage/v1"
)

type HttpContext struct {
	id       string
	Request  HttpRequest
	Response *HttpResponse
	Extras   map[string]interface{}
}

func (c *HttpContext) ToClientMessage() *http.ClientMessage {
	headers := make(map[string]*http.HeaderValue)
	for k, v := range c.Response.Headers {
		headers[k] = &http.HeaderValue{
			Value: v,
		}
	}

	return &http.ClientMessage{
		Id: c.id,
		Content: &http.ClientMessage_HttpResponse{
			HttpResponse: &http.HttpResponse{
				Status:  int32(c.Response.Status),
				Headers: headers,
				Body:    c.Response.Body,
			},
		},
	}
}

func NewHttpContext(msg *http.ServerMessage) *HttpContext {
	req := msg.GetHttpRequest()

	headers := make(map[string][]string)
	for k, v := range req.Headers {
		headers[k] = v.GetValue()
	}

	query := make(map[string][]string)
	for k, v := range req.QueryParams {
		query[k] = v.GetValue()
	}

	return &HttpContext{
		id: msg.Id,
		Request: &httpRequestImpl{
			method:     req.Method,
			path:       req.Path,
			pathParams: req.PathParams,
			query:      query,
			headers:    headers,
		},
		Response: &HttpResponse{
			Status:  200,
			Headers: map[string][]string{},
			Body:    nil,
		},
	}
}

func (c *HttpContext) WithError(err error) {
	c.Response = &HttpResponse{
		Status: 500,
		Headers: map[string][]string{
			"Content-Type": {"text/plain"},
		},
		Body: []byte("Internal Server Error"),
	}
}

type MessageContext struct {
	Request  MessageRequest
	Response *MessageResponse
	Extras   map[string]interface{}
}

type IntervalContext struct {
	Request  IntervalRequest
	Response *IntervalResponse
	Extras   map[string]interface{}
}

type BlobEventContext struct {
	id 		string
	Request  BlobEventRequest
	Response *BlobEventResponse
	Extras   map[string]interface{}
}

func (c *BlobEventContext) ToClientMessage() *storagepb.ClientMessage {
	return &storagepb.ClientMessage{
		Id: c.id,
		Content: &storagepb.ClientMessage_BlobEventResponse{
			BlobEventResponse: &storagepb.BlobEventResponse{
				Success: c.Response.Success,
			},
		},
	}
}

func NewBlobEventContext(msg *storagepb.ServerMessage) *BlobEventContext {
	req := msg.GetBlobEventRequest()
	
	return &BlobEventContext{
		id: msg.Id,
		Request: &blobEventRequestImpl{
			key: req.GetBlobEvent().Key,
		},
		Response: &BlobEventResponse{
			Success: true,
		},
	}
}

func (c *BlobEventContext) WithError(err error){
	c.Response = &BlobEventResponse{
		Success: false,
	}
}

type FileEventContext struct {
	id 		string
	Request  FileEventRequest
	Response *FileEventResponse
	Extras   map[string]interface{}
}

func (c *FileEventContext) ToClientMessage() *storagepb.ClientMessage {
	return &storagepb.ClientMessage{
		Id: c.id,
		Content: &storagepb.ClientMessage_BlobEventResponse{
			BlobEventResponse: &storagepb.BlobEventResponse{
				Success: c.Response.Success,
			},
		},
	}
}

func NewFileEventContext(msg *storagepb.ServerMessage, bucket *storage.Bucket) *FileEventContext {
	_ = msg.GetBlobEventRequest()
	
	return &FileEventContext{
		id: msg.Id,
		Request: &fileEventRequestImpl{
			bucket: bucket, //TODO:
		},
		Response: &FileEventResponse{
			Success: true,
		},
	}
}

func (c *FileEventContext) WithError(err error){
	c.Response = &FileEventResponse{
		Success: false,
	}
}

type WebsocketContext struct {
	Request  WebsocketRequest
	Response *WebsocketResponse
	Extras   map[string]interface{}
}
