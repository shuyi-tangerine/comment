package top

import (
	"context"
	"fmt"
	"testing"
)

func TestComment_ToListImages(t *testing.T) {
	images := `["a", "b"]`
	c := &Comment{Images: &images}
	listImages, err := c.ToListImages(context.Background())
	fmt.Println(err, listImages)
}
