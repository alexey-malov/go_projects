package font_test

import (
	"fmt"
	"github.com/alexey-malov/go_projects/workshop02/font"
	"testing"
)

func TestFontConstruction(t *testing.T) {
	expectSuccessfulConstruction("f", 5, t)
	expectSuccessfulConstruction("f", 144, t)
	expectFontConstruction("fnt", "fnt", 4, 5, t)
	expectFontConstruction("fnt", "fnt", 145, 144, t)
	expectFontConstruction("", "Arial", 10, 10, t)
}

func TestFont(t *testing.T) {
	bodyFont := font.New("Nimbus Sans", 10)
	titleFont := font.New("serif", 11)
	f1(bodyFont, titleFont, t)
}

func TestFont_SetSize(t *testing.T) {
	f := font.New("Family", 10)

	f.SetSize(font.MIN_SIZE)
	if f.Size() != 5 {
		t.Fatal("Failed to set font size to 5")
	}

	f.SetSize(font.MIN_SIZE - 1)
	if f.Size() != font.MIN_SIZE {
		t.Fatal("Font size must be set to font.MIN_SIZE")
	}

	f.SetSize(font.MAX_SIZE)
	if f.Size() != font.MAX_SIZE {
		t.Fatal("Failed to set max font size")
	}

	f.SetSize(font.MAX_SIZE + 1)
	if f.Size() != font.MAX_SIZE {
		t.Fatal("Font size must be set to font.MAX_SIZE")
	}
}

func expectSuccessfulConstruction(family string, size int, t *testing.T) {
	expectFontConstruction(family, family, size, size, t)
}

func expectFontConstruction(family, expectedFamily string, size, expectedSize int, t *testing.T) {
	fnt := font.New(family, size)

	gotFamily := fnt.Family()
	if fnt.Family() != expectedFamily {
		t.Errorf("Expected font.Family() to be equal to '%s' but got '%s'", expectedFamily, gotFamily)
		return
	}

	gotSize := fnt.Size()
	if fnt.Size() != expectedSize {
		t.Errorf("Expected font.Size() to be equal to '%d' but got '%d'", expectedSize, gotSize)
		return
	}
}

func f1(bodyFont, titleFont font.Font, t *testing.T) {
	if bodyFont.String() !=
		`{font-family: "Nimbus Sans"; font-size: 10pt;}` {
		t.Fatal("#1 bodyFont invalid CSS")
	}
	if bodyFont.Size() != 10 || bodyFont.Family() != "Nimbus Sans" {
		t.Fatal("#2 bodyFont invalid attributes")
	}
	bodyFont.SetSize(12)
	if bodyFont.Size() != 12 || bodyFont.Family() != "Nimbus Sans" {
		t.Fatal("#3 bodyFont invalid attributes")
	}
	if bodyFont.String() !=
		`{font-family: "Nimbus Sans"; font-size: 12pt;}` {
		t.Fatal("#4 bodyFont invalid CSS")
	}
	bodyFont.SetFamily("")
	if bodyFont.Size() != 12 || bodyFont.Family() != "Nimbus Sans" {
		t.Fatal("#5 bodyFont invalid attributes")
	}

	if titleFont.String() != `{font-family: "serif"; font-size: 11pt;}` {
		t.Fatal("#6 titleFont invalid CSS")
	}
	if titleFont.Size() != 11 || titleFont.Family() != "serif" {
		t.Fatal("#7 titleFont invalid attributes")
	}
	titleFont.SetFamily("Helvetica")
	titleFont.SetSize(20)
	if titleFont.Size() != 20 || titleFont.Family() != "Helvetica" {
		t.Fatal("#8 titleFont invalid attributes")
	}

	f2(bodyFont, titleFont)
}

func f2(bodyFont, titleFont font.Font) {
	fmt.Println(bodyFont)
	fmt.Println(titleFont)
}
