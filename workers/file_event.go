package workers

import (
	"context"
	"fmt"
	"io"

	"github.com/nitrictech/go-sdk/api/errors"
	"github.com/nitrictech/go-sdk/api/errors/codes"
	"github.com/nitrictech/go-sdk/api/storage"
	"github.com/nitrictech/go-sdk/constants"
	"github.com/nitrictech/go-sdk/handler"
	v1 "github.com/nitrictech/nitric/core/pkg/proto/storage/v1"
	"google.golang.org/grpc"
)


type FileEventWorker struct {
	client v1.StorageListenerClient
	bucket *storage.Bucket
	registrationRequest *v1.RegistrationRequest 
	middleware handler.FileEventMiddleware 
}


type FileEventWorkerOpts struct {
	Bucket 			*storage.Bucket
	RegistrationRequest *v1.RegistrationRequest
	Middleware handler.FileEventMiddleware
}

// Start implements Worker.
func (f *FileEventWorker) Start(ctx context.Context) error {
	
	initReq := &v1.ClientMessage{
		Content: &v1.ClientMessage_RegistrationRequest{
			RegistrationRequest: f.registrationRequest,
		},
	}
	
	// Create the request stream and send the initial request
	stream, err := f.client.Listen(ctx)
	if err != nil{
		return err
	}

	err = stream.Send(initReq)
	if err != nil{
		return err
	}


	for {
		var ctx *handler.FileEventContext

		resp, err := stream.Recv()

		if err == io.EOF {
			err = stream.CloseSend()
			if err != nil {
				return err
			}

			return nil
		} else if err == nil && resp.GetRegistrationResponse() != nil {
			// File Notification has connected with Nitric server
			fmt.Println("FileEventWorker connected with Nitric server")
		} else if err == nil && resp.GetBlobEventRequest() != nil {
			ctx = handler.NewFileEventContext(resp, f.bucket)
			ctx, err = f.middleware(ctx, handler.FileEventDummy)

			if err != nil {
				ctx.WithError(err)
			}

			err = stream.Send(ctx.ToClientMessage())
			if err != nil {
				return err
			}

		} else {
			return err
		}
	}
}

func NewFileEventWorker(opts *FileEventWorkerOpts) *FileEventWorker {
	ctx, _ := context.WithTimeout(context.TODO(), constants.NitricDialTimeout())

	conn, err := grpc.DialContext(
		ctx,
		constants.NitricAddress(),
		constants.DefaultOptions()...,
	)
	if err != nil {
		panic(errors.NewWithCause(
			codes.Unavailable,
			"NewFileEventWorker: Unable to reach StorageListenerClient",
			err,
		))
	}

	client := v1.NewStorageListenerClient(conn)
	
	return &FileEventWorker{
		client: client,
		bucket: opts.Bucket,
		registrationRequest: opts.RegistrationRequest,
		middleware: opts.Middleware,
	}
}