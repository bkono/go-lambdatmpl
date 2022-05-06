//go:build !test

package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

func main() {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatalf("FATAL error getting an AWS session: %+v\n", err)
	}

	h := &handler{
		s3svc: s3.New(sess),
	}

	lambda.Start(h.handle)
}
