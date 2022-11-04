package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/matryer/is"
)

func TestHandle(t *testing.T) {

	tests := []struct {
		name     string
		contents string
		err      error
	}{
		{name: "happy path", contents: "foo"},
		{name: "bubbles errors", err: fmt.Errorf("boom")},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)
			h := &handler{&mockS3Client{contents: tc.contents, err: tc.err}}
			rsp, err := h.handle(context.Background(), ListObjectsEvent{Bucket: "foo"})

			if tc.err != nil {
				is.True(rsp == nil)
				is.True(err != nil)
			} else {
				is.NoErr(err)
				is.Equal(rsp.Contents, tc.contents+"\n")
			}

		})
	}
}

type mockS3Client struct {
	s3iface.S3API
	contents string
	err      error
}

func (m *mockS3Client) ListObjectsV2WithContext(ctx aws.Context, in *s3.ListObjectsV2Input, opts ...request.Option) (*s3.ListObjectsV2Output, error) {
	if m.err != nil {
		return nil, m.err
	}

	return &s3.ListObjectsV2Output{Contents: []*s3.Object{{Key: aws.String(m.contents)}}}, nil
}
