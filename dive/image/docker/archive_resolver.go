package docker

import (
	"fmt"
	"io"
	"os"

	"github.com/wagoodman/dive/dive/image"
)

type archiveResolver struct{}

func NewResolverFromArchive() *archiveResolver {
	return &archiveResolver{}
}

func (r *archiveResolver) Fetch(path string) (*image.Image, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	img, err := NewImageArchive(reader)
	if err != nil {
		return nil, err
	}
	return img.ToImage()
}

func (r *archiveResolver) Build(args []string) (*image.Image, error) {
	return nil, fmt.Errorf("build option not supported for docker archive resolver")
}

func (r *archiveResolver) Extract(id string, l string, p string) error {

	reader, err := os.Open(id)
	if err != nil {
		return err
	}

	if err := ExtractFromImage(io.NopCloser(reader), l, p); err == nil {
		return nil
	}

	return fmt.Errorf("unable to extract from image '%s': %+v", id, err)
}
