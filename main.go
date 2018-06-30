package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gofn/gofn"
	"github.com/gofn/gofn/provision"
)

func main() {
	imageName := "sum"
	stdin := `{"factor1": 10, "factor2": 10}`
	stdout, stderr, err := run(imageName, stdin)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("stderr: %q\n", stderr)
	fmt.Printf("stdout: %q\n", stdout)
}

func run(imageName, input string) (stdout, stderr string, err error) {
	buildOpts := &provision.BuildOptions{
		ImageName: imageName,
		StdIN:     input,
		DoNotUsePrefixImageName: true,
	}
	containerOpts := &provision.ContainerOptions{}

	ctx, cancelFunc := context.WithTimeout(
		context.Background(),
		time.Duration(5)*time.Second)
	stdout, stderr, err = gofn.Run(ctx, buildOpts, containerOpts)
	cancelFunc()
	return
}
