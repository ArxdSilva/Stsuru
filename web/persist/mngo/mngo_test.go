package mngo

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	link, linkShort, hash string
	isURL                 bool
}{
	{
		"abcdef",
		"",
		"",
		false,
	},
	{
		"www.globo.com",
		"",
		"",
		false,
	},
	{
		"www.notvalidurl.netscape",
		"",
		"",
		false,
	},
	{
		"http://www.gorillatoolkit.org/pkg/mux",
		"",
		"70df8650c03c9fdfc959f04a64ecd956",
		true,
	},
	{
		"https://mail.google.com/mail/u/0/#inbox",
		"",
		"2122c5656da3d86d77c08f7af48c0268",
		true,
	},
	{
		"https://mail.google.com/mail/u/0/#inbox",
		"",
		"2122c5656da3d86d77c08f7af48c0268",
		true,
	},
	{
		"https://www.youtube.com/watch?v=grwx4OMfAn4",
		"",
		"678989a28d9b88ada6cc6678df8e6aa1",
		true,
	},
}

var s = MongoStorage{
	URL:        "localhost",
	DB:         "tsuru",
	Collection: "links",
}

func TestInsert(t *testing.T) {
	fmt.Print("Test Insert: ")
	for _, test := range testCases {
		err := s.Save(test.link, test.linkShort, test.hash)
		checkError(err)
		fmt.Print(".")
	}
	fmt.Println()
}

func TestFindHash(t *testing.T) {
	fmt.Print("Test FindHash: ")
	for _, test := range testCases {
		link, err := s.FindHash(test.hash)
		if err != nil && test.isURL == false {
			fmt.Print(".")
			continue
		}
		if test.link == link {
			fmt.Print(".")
		}
		continue
	}
	fmt.Println()
}

func TestGetAll(t *testing.T) {
	fmt.Print("Test GetAll: ")
	a, err := s.GetAll()
	checkError(err)
	if len(a) == 3 {
		fmt.Print("...")
	} else {
		fmt.Println()
		fmt.Printf("Array bigger than expected, len == %v", len(a))
	}
	fmt.Println()
}

func TestDelete(t *testing.T) {
	fmt.Print("Test Delete: ")
	for _, test := range testCases {
		err := s.Remove(test.hash)
		if err != nil {
			continue
		}
		fmt.Print(".")
	}
	fmt.Println()
}
