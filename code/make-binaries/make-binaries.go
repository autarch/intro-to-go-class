package main

import (
	"log"
	"os"
	"os/exec"
	"path"
)

type target struct {
	os   string
	arch string
}

func main() {
	if err := os.Chdir("../class-setup"); err != nil {
		log.Fatal(err)
	}

	targets := []target{}
	for _, os := range []string{"darwin", "linux", "windows"} {
		for _, arch := range []string{"386", "amd64"} {
			targets = append(targets, target{os: os, arch: arch})
		}
	}

	for _, t := range targets {
		if err := os.Setenv("GOOS", t.os); err != nil {
			log.Fatalf("Setenv GOOS=%s failed: %s", t.os, err)
		}
		if err := os.Setenv("GOARCH", t.arch); err != nil {
			log.Fatalf("Setenv GOARCH=%s failed: %s", t.arch, err)
		}

		log.Printf("%s %s", t.os, t.arch)
		build := exec.Command("go", "build")
		if o, err := build.CombinedOutput(); err != nil {
			log.Fatalf("go build failed: %s\n%s\n", err, string(o))
		}

		from := "class-setup"
		to := "class-setup"

		if t.os == "windows" {
			from = from + ".exe"
			to = to + "." + t.arch + ".exe"
		} else {
			to = to + "." + t.os + "." + t.arch
		}
		to = path.Join("/home/autarch/projects/intro-to-go-class/setup", to)

		if o, err := exec.Command("mv", from, to).CombinedOutput(); err != nil {
			log.Fatalf("mv %s => %s failed: %s\n%s\n", from, to, err, string(o))
		}
	}
}
