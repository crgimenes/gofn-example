package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gofn/gofn"
	"github.com/gofn/gofn/provision"
)

func main() {
	imgName := flag.String("image", "", "Docker image name")
	p1 := flag.String("p1", "", "parameter 1")
	p2 := flag.String("p2", "", "parameter 2")
	flag.Parse()
	var imageName, stdin string
	switch *imgName {
	case "upper":
		imageName = "upper"
		stdin = fmt.Sprintf(`%v`, *p1)
	case "encode":
		imageName = "encode"
		stdin = fmt.Sprintf(`{"phrase": %q}`, *p1)
	case "sum":
		imageName = "sum"
		stdin = fmt.Sprintf(`{"factor1": %v, "factor2": %v}`, *p1, *p2)
	default:
		fmt.Printf("invalid parameter %q\n", *imgName)
		os.Exit(-1)
	}
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
