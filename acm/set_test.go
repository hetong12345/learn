package main

import (
	"fmt"
	"github.com/go-ego/gse"
	"testing"
)

func TestSet(t *testing.T) {

}

func main() {
	// Load the dictionary
	var seg gse.Segmenter
	// Loading the default dictionary
	seg.LoadDict()
	// seg.LoadDict("your gopath"+"/src/github.com/go-ego/gse/data/dict/dictionary.txt")

	// Text Segmentation
	text := []byte("你好世界, Hello world.")
	fmt.Println(segmenter.String(text, true))

	segments := segmenter.Segment(text)

	// Handle word segmentation results
	// Support for normal mode and search mode two participle,
	// see the comments in the code ToString function.
	// The search mode is mainly used to provide search engines
	// with as many keywords as possible
	fmt.Println(gse.ToString(segments, true))
}
