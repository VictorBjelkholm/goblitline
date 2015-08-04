// Goblitline is a golang client for the Blitline API.
//
// It provide a DSL for building Jobs, Functions and Containers (function saving).
//
// The DSL try to mimic Blitline API, so if you are familiar with it, use goblitline should be easy enough.
// For more information about blitline api, please refer to http://www.blitline.com/docs/api.
package goblitline

import (
	"fmt"
	"math/rand"
)

// Returns a JobBuilder
func Job(AppId string) JobBuilder {
	return JobBuilder{}.ApplicationID(AppId)
}

// Returns a FunctionBuilder
func Function(name string) FunctionBuilder {
	return FunctionBuilder{}.Name(name)
}

// Returns a ContainerBuilder
func Container(imageId string, destination *S3Destination) ContainerBuilder {
	container_builder := ContainerBuilder{}.
		ImageIdentifier(imageId).
		Quality(75)

	if destination != nil {
		if destination.Key == "" {
			destination.Key = imageId + "-" + randString(10)
		}
		if destination.Bucket == "" {
			panic("You need to set S3Destination.Bucket")
		}
		container_builder.S3Destination(destination)
	}

	fmt.Printf("%#v\n", container_builder.S3Destination)

	return container_builder
}

var alphanum = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"

func randString(size int) string {
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = alphanum[rand.Intn(len(alphanum))]
	}
	return string(buf)
}
