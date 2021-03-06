package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/acsaba22/go/grpc/mempb"

	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) <= 1 {
		panic("Give one argument e.g. 42")
	}
	k := os.Args[1]

	channel, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer channel.Close()
	client := mempb.NewMemServerClient(channel)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Get(ctx, &mempb.GetRequest{Key: k})
	if err != nil {
		panic(err)
	}
	v := r.Value
	if !r.Exists {
		v = "NOTFOUND"
	}
	fmt.Printf("%s: %s\n", r.RequestedKey, v)
}
