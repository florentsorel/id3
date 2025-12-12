package id3

import "github.com/florentsorel/id3/id3v1"

// ID3 represents ID3 metadata tags from an audio file.
type ID3 struct {
	ID3v1 *id3v1.ID3v1
}

// HasID3v1 returns true if the ID3 tag contains ID3v1 metadata.
func (t *ID3) HasID3v1() bool {
	return t.ID3v1 != nil
}

// Open reads and parses ID3 tags from the specified audio file.
// It returns an ID3 struct containing the parsed metadata or an error if the file cannot be read or parsed.
func Open(filename string) (*ID3, error) {
	id3v1Tag, err := id3v1.Open(filename)
	if err != nil {
		return nil, err
	}

	return &ID3{
		ID3v1: id3v1Tag,
	}, nil
}

// Read parses ID3 tags from the provided byte slice.
// It returns an ID3 struct containing the parsed metadata or an error if the data cannot be parsed.
func Read(data []byte) (*ID3, error) {
	id3v1Tag, err := id3v1.Read(data)
	if err != nil {
		return nil, err
	}

	return &ID3{
		ID3v1: id3v1Tag,
	}, nil
}
