package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"os"
	"io"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: log_rotate <file> [file.tar.gz]")
		os.Exit(1)
	}

	input := os.Args[1]
	var output string

	if len(os.Args) >= 3 {
		output = os.Args[2]
	} else {
		base := filepath.Base(input)
		output = base + ".tar.gz"
	}

	if err := TarGzFile(input, output); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Println("Compressed to:", output)
}

func TarGzFile(inputPath, outputPath string) error {
	in, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer in.Close()

	info, err := in.Stat()
	if err != nil {
		return err
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	gz := gzip.NewWriter(out)
	defer gz.Close()

	tw := tar.NewWriter(gz)
	defer tw.Close()

	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}

	header.Name = filepath.Base(inputPath)

	if err := tw.WriteHeader(header); err != nil {
		return err
	}

	_, err = io.Copy (tw, in)

	return err
}