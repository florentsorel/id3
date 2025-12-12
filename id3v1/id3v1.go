package id3v1

import (
	"errors"
	"io"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

// Version represents the ID3v1 tag version.
type Version string

const (
	// ID3v1_0 represents ID3v1.0 specification (without track number).
	ID3v1_0 Version = "ID3v1_0"
	// ID3v1_1 represents ID3v1.1 specification (with track number).
	ID3v1_1 Version = "ID3v1_1"
)

const id3v1TagSize = 128

// ID3v1 represents an ID3v1 metadata tag containing information about an audio file.
type ID3v1 struct {
	Title   string
	Artist  string
	Album   string
	Year    string
	Comment string
	Track   uint8
	Genre   byte
	Version Version
}

// String returns the string representation of the Version.
func (v Version) String() string {
	if v == ID3v1_0 {
		return "ID3v1.0"
	}
	return "ID3v1.1"
}

// Open reads and parses an ID3v1 tag from the specified audio file.
// It returns an ID3v1 struct containing the parsed metadata or an error if the file cannot be read or parsed.
func Open(filename string) (*ID3v1, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return Read(b)
}

// Read parses an ID3v1 tag from the provided byte slice.
// It expects the ID3v1 tag to be in the last 128 bytes of the data.
// It returns an ID3v1 struct containing the parsed metadata or an error if the data is too short or no tag is found.
func Read(data []byte) (*ID3v1, error) {
	if len(data) < id3v1TagSize {
		return nil, errors.New("data too short to contain ID3v1 tag")
	}

	tagData := data[len(data)-id3v1TagSize:]

	return parse(tagData)
}

func parse(data []byte) (*ID3v1, error) {
	if string(data[0:3]) != "TAG" {
		return nil, errors.New("no ID3v1 tag found")
	}

	comment, track := parseCommentAndTrack(data[97:127])

	version := ID3v1_1
	if track == 0 {
		version = ID3v1_0
	}

	id3v1 := &ID3v1{
		Title:   trim(data[3:33]),
		Artist:  trim(data[33:63]),
		Album:   trim(data[63:93]),
		Year:    trim(data[93:97]),
		Comment: comment,
		Track:   track,
		Genre:   data[127],
		Version: version,
	}

	return id3v1, nil
}

func parseCommentAndTrack(b []byte) (string, uint8) {
	if b[28] == 0 && b[29] != 0 {
		return trim(b[0:28]), b[29]
	}

	return trim(b[0:30]), 0

}

func trim(b []byte) string {
	s := decodeLatin1(b)
	s = strings.TrimRight(s, "\x00")
	s = strings.TrimSpace(s)
	return s
}

func decodeLatin1(b []byte) string {
	s, _ := charmap.ISO8859_1.NewDecoder().Bytes(b)
	return string(s)
}
