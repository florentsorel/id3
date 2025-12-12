package id3v1

import (
	"io"
	"os"
	"testing"
)

func TestID3v1_Open(t *testing.T) {
	id3v1, err := Open("./../testdata/id3v1_0.mp3")
	if err != nil {
		t.Fatalf("Failed to open ID3v1 tag: %v", err)
	}

	expectedTitle := "Title éÃ©"
	if id3v1.Title != expectedTitle {
		t.Errorf("Expected title %q, got %q", expectedTitle, id3v1.Title)
	}

	expectedArtist := "Artist"
	if id3v1.Artist != expectedArtist {
		t.Errorf("Expected artist %q, got %q", expectedArtist, id3v1.Artist)
	}

	expectedAlbum := "Album"
	if id3v1.Album != expectedAlbum {
		t.Errorf("Expected album %q, got %q", expectedAlbum, id3v1.Album)
	}

	expectedYear := "2025"
	if id3v1.Year != expectedYear {
		t.Errorf("Expected year %q, got %q", expectedYear, id3v1.Year)
	}

	expectedComment := "Comment with 30 characters !!!"
	if id3v1.Comment != expectedComment {
		t.Errorf("Expected comment %q, got %q", expectedComment, id3v1.Comment)
	}

	expectedTrack := uint8(0)
	if id3v1.Track != expectedTrack {
		t.Errorf("Expected track %d, got %d", expectedTrack, id3v1.Track)
	}

	expectedGenre := byte(15)
	if id3v1.Genre != expectedGenre {
		t.Errorf("Expected genre %d, got %d", expectedGenre, id3v1.Genre)
	}

	expectedVersion := ID3v1_0
	if id3v1.Version != expectedVersion {
		t.Errorf("Expected version %q, got %q", expectedVersion, id3v1.Version)
	}
}

func TestRead_ID3v1_1(t *testing.T) {
	f, err := os.Open("./../testdata/id3v1_1.mp3")
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	id3v1, err := Read(data)
	if err != nil {
		t.Fatalf("Failed to read ID3v1 tag: %v", err)
	}

	expectedTitle := "Title"
	if id3v1.Title != expectedTitle {
		t.Errorf("Expected title %q, got %q", expectedTitle, id3v1.Title)
	}

	expectedArtist := "Artist"
	if id3v1.Artist != expectedArtist {
		t.Errorf("Expected artist %q, got %q", expectedArtist, id3v1.Artist)
	}

	expectedAlbum := "Album"
	if id3v1.Album != expectedAlbum {
		t.Errorf("Expected album %q, got %q", expectedAlbum, id3v1.Album)
	}

	expectedYear := "2025"
	if id3v1.Year != expectedYear {
		t.Errorf("Expected year %q, got %q", expectedYear, id3v1.Year)
	}

	expectedComment := "Comment with 28 characters !"
	if id3v1.Comment != expectedComment {
		t.Errorf("Expected comment %q, got %q", expectedComment, id3v1.Comment)
	}

	expectedTrack := uint8(3)
	if id3v1.Track != expectedTrack {
		t.Errorf("Expected track %d, got %d", expectedTrack, id3v1.Track)
	}

	expectedGenre := byte(255)
	if id3v1.Genre != expectedGenre {
		t.Errorf("Expected genre %d, got %d", expectedGenre, id3v1.Genre)
	}

	expectedVersion := ID3v1_1
	if id3v1.Version != expectedVersion {
		t.Errorf("Expected version %q, got %q", expectedVersion, id3v1.Version)
	}
}

func TestParse_NoTag(t *testing.T) {
	f, err := os.Open("./../testdata/id3_no_tags.mp3")
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	_, err = parse(data)
	if err == nil {
		t.Fatal("Expected error for missing ID3v1 tag, got nil")
	}

	exceptedError := "no ID3v1 tag found"
	if err.Error() != exceptedError {
		t.Errorf("Expected error %q, got %q", exceptedError, err.Error())
	}
}
