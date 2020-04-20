package main

import (
	"fmt"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/mutate"
	"github.com/google/go-containerregistry/pkg/v1/random"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"log"
)

func main() {
	image, err := random.Image(9999999, 5)
	if err != nil {
		log.Fatal(err)
	}

	reference, err := name.ParseReference("kpack/bars")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("attemptint upload\n")

	get, err := remote.Image(reference, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		log.Fatal(err)
	}

	layers, err := get.Layers()
	if err != nil {
		log.Fatal(err)
	}


	image, err = mutate.AppendLayers(image, layers[0:5]...)
	if err != nil {
		log.Fatal(err)
	}

	runRef, err := name.ParseReference("cloudfoundry/run:base-cnb")
	if err != nil {
		log.Fatal(err)
	}
	runImage, err := remote.Image(runRef, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		log.Fatal(err)
	}

	layers, err = runImage.Layers()
	if err != nil {
		log.Fatal(err)
	}


	image, err = mutate.AppendLayers(image, layers...)
	if err != nil {
		log.Fatal(err)
	}

	l, _ := image.Layers()
	for _, _ = range l {
		remote.LayerNames = append(remote.LayerNames, "some-name")

	}

	err = remote.Write(reference, image, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		log.Fatal(err)
	}
}
