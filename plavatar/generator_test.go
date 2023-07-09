package plavatar

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"image/png"
	"io"
	"strings"
	"testing"
)

var avatarGenerator = Generator{}

func TestAvatarSolid(t *testing.T) {
	var shaHasher = sha256.New()
	generatedAvatar, rngSeed, err := avatarGenerator.GenerateAvatar(avatarGenerator.Solid, &Options{
		Name:         "8",
		OutputSize:   512,
		OutputFormat: PNG,
		OutputShape:  Square,
	})

	if err != nil {
		t.Fatal("err != nil", err)
	}

	if rngSeed != "8" {
		t.Fatal("rngSeed mismatch")
	}

	shaHasher.Write(generatedAvatar.Bytes())
	hash := fmt.Sprintf("%x", shaHasher.Sum(nil))

	if strings.ToLower(hash) != "90c34a5824d789ef4323fcd7b4fe5260ac68a95e843629a7e922c091d0c5445e" {
		t.Error("hash missmatch. check if intentional and change hash accordingly.", hash)
	}
}

func TestAvatarLaughing(t *testing.T) {
	var shaHasher = sha256.New()
	generatedAvatar, rngSeed, err := avatarGenerator.GenerateAvatar(avatarGenerator.Laughing, &Options{
		Name:         "6",
		OutputSize:   256,
		OutputFormat: PNG,
		OutputShape:  Circle,
	})

	if err != nil {
		t.Fatal("err != nil", err)
	}

	if rngSeed != "6" {
		t.Fatal("rngSeed mismatch")
	}

	tmpGeneratedAvatarBuffer, err := io.ReadAll(generatedAvatar)
	if err != nil {
		t.Fatal("failed reading avatar buffer")
	}

	parsedPng, err := png.Decode(bytes.NewBuffer(tmpGeneratedAvatarBuffer))
	if err != nil {
		t.Fatal("failed decoding generated png")
	}
	if parsedPng.Bounds().Size().X != 256 || parsedPng.Bounds().Size().Y != 256 {
		t.Fatal("wrong output size, check image scaling logic")
	}

	shaHasher.Write(tmpGeneratedAvatarBuffer)
	hash := fmt.Sprintf("%x", shaHasher.Sum(nil))

	if strings.ToLower(hash) != "102176c91b530b457c7c3d341e34ca677ac78a7d86319981d0b88e99913c83dc" {
		t.Error("hash missmatch. check if intentional and change hash accordingly.", hash)
	}
}

// todo: check svg contents
func TestAvatarSmileySVG(t *testing.T) {
	_, rngSeed, err := avatarGenerator.GenerateAvatar(avatarGenerator.Smiley, &Options{
		Name:         "6",
		OutputSize:   256,
		OutputFormat: SVG,
		OutputShape:  Circle,
	})

	if err != nil {
		t.Fatal("err != nil", err)
	}

	if rngSeed != "6" {
		t.Fatal("rngSeed mismatch")
	}
}

func BenchmarkAvatarSmileyPNG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, err := avatarGenerator.GenerateAvatar(avatarGenerator.Smiley, &Options{
			Name:         "6",
			OutputSize:   512,
			OutputFormat: PNG,
			OutputShape:  Circle,
		})

		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAvatarSmileySVG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, err := avatarGenerator.GenerateAvatar(avatarGenerator.Smiley, &Options{
			Name:         "6",
			OutputSize:   512,
			OutputFormat: SVG,
			OutputShape:  Circle,
		})

		if err != nil {
			b.Fatal(err)
		}
	}
}
