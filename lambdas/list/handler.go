package main

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"log"
	"strings"
)

type handler struct {
	s3svc s3iface.S3API
}
type ListObjectsEvent struct {
	Bucket string
}
type Response struct {
	Contents string
}

func (h *handler) handle(ctx context.Context, evt ListObjectsEvent) (*Response, error) {
	out, err := h.s3svc.ListObjectsV2WithContext(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(evt.Bucket),
	})

	if err != nil {
		log.Printf("handle: err listing objects: %+v\n", err)
		return nil, err
	}

	var sb strings.Builder
	for _, obj := range out.Contents {
		sb.WriteString(aws.StringValue(obj.Key) + "\n")
	}

	return &Response{Contents: sb.String()}, nil
}
