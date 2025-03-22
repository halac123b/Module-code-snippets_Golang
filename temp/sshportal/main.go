package main

import (
	"math/rand"
	"os"
	"path"

	"github.com/urfave/cli"
	"moul.io/srand"
)

var (
	// GitTag will be overwritten automatically by the build system
	GitTag = "n/a"
	// GitSha will be overwritten automatically by the build system
	GitSha = "n/a"
)

func main() {
	rand.Seed(srand.MustSecure())

	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Author = "River"
	app.Version = GitTag + " (" + GitSha + ")"
	app.Email = "email@gmail.com"
	app.Commands = []cli.Command{}
}
