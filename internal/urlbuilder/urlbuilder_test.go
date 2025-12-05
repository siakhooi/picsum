package urlbuilder

import (
	"testing"
)

func TestBuildURL_SingleArgument_NoOptions(t *testing.T) {
	args := []string{"300"}
	url, filename, err := BuildURL(args, "", "", false, false, 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedURL := "https://picsum.photos/300"
	expectedFilename := "300.jpg"

	if url != expectedURL {
		t.Errorf("expected URL %q, got %q", expectedURL, url)
	}
	if filename != expectedFilename {
		t.Errorf("expected filename %q, got %q", expectedFilename, filename)
	}
}

func TestBuildURL_SingleArgument_WithImageID(t *testing.T) {
	args := []string{"300"}
	url, filename, err := BuildURL(args, "237", "", false, false, 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedURL := "https://picsum.photos/id/237/300"
	expectedFilename := "id_237_300.jpg"

	if url != expectedURL {
		t.Errorf("expected URL %q, got %q", expectedURL, url)
	}
	if filename != expectedFilename {
		t.Errorf("expected filename %q, got %q", expectedFilename, filename)
	}
}

func TestBuildURL_SingleArgument_WithSeed(t *testing.T) {
	args := []string{"300"}
	url, filename, err := BuildURL(args, "", "picsum", false, false, 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedURL := "https://picsum.photos/seed/picsum/300"
	expectedFilename := "seed_picsum_300.jpg"

	if url != expectedURL {
		t.Errorf("expected URL %q, got %q", expectedURL, url)
	}
	if filename != expectedFilename {
		t.Errorf("expected filename %q, got %q", expectedFilename, filename)
	}
}

func TestBuildURL_TwoArguments_NoOptions(t *testing.T) {
	args := []string{"300", "200"}
	url, filename, err := BuildURL(args, "", "", false, false, 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedURL := "https://picsum.photos/300/200"
	expectedFilename := "300x200.jpg"

	if url != expectedURL {
		t.Errorf("expected URL %q, got %q", expectedURL, url)
	}
	if filename != expectedFilename {
		t.Errorf("expected filename %q, got %q", expectedFilename, filename)
	}
}

func TestBuildURL_TwoArguments_WithImageID(t *testing.T) {
	args := []string{"300", "200"}
	url, filename, err := BuildURL(args, "237", "", false, false, 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedURL := "https://picsum.photos/id/237/300/200"
	expectedFilename := "id_237_300x200.jpg"

	if url != expectedURL {
		t.Errorf("expected URL %q, got %q", expectedURL, url)
	}
	if filename != expectedFilename {
		t.Errorf("expected filename %q, got %q", expectedFilename, filename)
	}
}

func TestBuildURL_TwoArguments_WithSeed(t *testing.T) {
	args := []string{"300", "200"}
	url, filename, err := BuildURL(args, "", "picsum", false, false, 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedURL := "https://picsum.photos/seed/picsum/300/200"
	expectedFilename := "seed_picsum_300x200.jpg"

	if url != expectedURL {
		t.Errorf("expected URL %q, got %q", expectedURL, url)
	}
	if filename != expectedFilename {
		t.Errorf("expected filename %q, got %q", expectedFilename, filename)
	}
}

func TestBuildURL_InvalidSingleNumber(t *testing.T) {
	args := []string{"abc"}
	_, _, err := BuildURL(args, "", "", false, false, 0)

	if err == nil {
		t.Fatal("expected error for invalid number, got nil")
	}

	expectedError := "invalid number: abc"
	if err.Error() != expectedError {
		t.Errorf("expected error %q, got %q", expectedError, err.Error())
	}
}

func TestBuildURL_InvalidFirstNumber(t *testing.T) {
	args := []string{"abc", "200"}
	_, _, err := BuildURL(args, "", "", false, false, 0)

	if err == nil {
		t.Fatal("expected error for invalid first number, got nil")
	}

	expectedError := "invalid first number: abc"
	if err.Error() != expectedError {
		t.Errorf("expected error %q, got %q", expectedError, err.Error())
	}
}

func TestBuildURL_InvalidSecondNumber(t *testing.T) {
	args := []string{"300", "xyz"}
	_, _, err := BuildURL(args, "", "", false, false, 0)

	if err == nil {
		t.Fatal("expected error for invalid second number, got nil")
	}

	expectedError := "invalid second number: xyz"
	if err.Error() != expectedError {
		t.Errorf("expected error %q, got %q", expectedError, err.Error())
	}
}

func TestBuildURL_WithGrayscale_SingleArgument(t *testing.T) {
	args := []string{"300"}
	url, filename, err := BuildURL(args, "", "", true, false, 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedURL := "https://picsum.photos/300?grayscale"
	expectedFilename := "300_gray.jpg"

	if url != expectedURL {
		t.Errorf("expected URL %q, got %q", expectedURL, url)
	}
	if filename != expectedFilename {
		t.Errorf("expected filename %q, got %q", expectedFilename, filename)
	}
}

func TestBuildURL_WithGrayscale_TwoArguments_WithImageID(t *testing.T) {
	args := []string{"300", "200"}
	url, filename, err := BuildURL(args, "237", "", true, false, 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedURL := "https://picsum.photos/id/237/300/200?grayscale"
	expectedFilename := "id_237_300x200_gray.jpg"

	if url != expectedURL {
		t.Errorf("expected URL %q, got %q", expectedURL, url)
	}
	if filename != expectedFilename {
		t.Errorf("expected filename %q, got %q", expectedFilename, filename)
	}
}

func TestBuildURL_WithGrayscale_WithSeed(t *testing.T) {
	args := []string{"300"}
	url, filename, err := BuildURL(args, "", "picsum", true, false, 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedURL := "https://picsum.photos/seed/picsum/300?grayscale"
	expectedFilename := "seed_picsum_300_gray.jpg"

	if url != expectedURL {
		t.Errorf("expected URL %q, got %q", expectedURL, url)
	}
	if filename != expectedFilename {
		t.Errorf("expected filename %q, got %q", expectedFilename, filename)
	}
}

func TestBuildURL_WithBlur_SingleArgument(t *testing.T) {
	args := []string{"300"}
	url, filename, err := BuildURL(args, "", "", false, true, 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedURL := "https://picsum.photos/300?blur"
	expectedFilename := "300_blur.jpg"

	if url != expectedURL {
		t.Errorf("expected URL %q, got %q", expectedURL, url)
	}
	if filename != expectedFilename {
		t.Errorf("expected filename %q, got %q", expectedFilename, filename)
	}
}

func TestBuildURL_WithBlur_TwoArguments(t *testing.T) {
	args := []string{"300", "200"}
	url, filename, err := BuildURL(args, "", "", false, true, 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedURL := "https://picsum.photos/300/200?blur"
	expectedFilename := "300x200_blur.jpg"

	if url != expectedURL {
		t.Errorf("expected URL %q, got %q", expectedURL, url)
	}
	if filename != expectedFilename {
		t.Errorf("expected filename %q, got %q", expectedFilename, filename)
	}
}

func TestBuildURL_WithBlur_WithImageID(t *testing.T) {
	args := []string{"300", "200"}
	url, filename, err := BuildURL(args, "237", "", false, true, 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedURL := "https://picsum.photos/id/237/300/200?blur"
	expectedFilename := "id_237_300x200_blur.jpg"

	if url != expectedURL {
		t.Errorf("expected URL %q, got %q", expectedURL, url)
	}
	if filename != expectedFilename {
		t.Errorf("expected filename %q, got %q", expectedFilename, filename)
	}
}

func TestBuildURL_WithGrayscaleAndBlur_SingleArgument(t *testing.T) {
	args := []string{"300"}
	url, filename, err := BuildURL(args, "", "", true, true, 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedURL := "https://picsum.photos/300?grayscale&blur"
	expectedFilename := "300_gray_blur.jpg"

	if url != expectedURL {
		t.Errorf("expected URL %q, got %q", expectedURL, url)
	}
	if filename != expectedFilename {
		t.Errorf("expected filename %q, got %q", expectedFilename, filename)
	}
}

func TestBuildURL_WithGrayscaleAndBlur_TwoArguments_WithSeed(t *testing.T) {
	args := []string{"300", "200"}
	url, filename, err := BuildURL(args, "", "picsum", true, true, 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedURL := "https://picsum.photos/seed/picsum/300/200?grayscale&blur"
	expectedFilename := "seed_picsum_300x200_gray_blur.jpg"

	if url != expectedURL {
		t.Errorf("expected URL %q, got %q", expectedURL, url)
	}
	if filename != expectedFilename {
		t.Errorf("expected filename %q, got %q", expectedFilename, filename)
	}
}

func TestBuildURL_WithBlurLevel_SingleArgument(t *testing.T) {
	args := []string{"300"}
	url, filename, err := BuildURL(args, "", "", false, false, 5)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedURL := "https://picsum.photos/300?blur=5"
	expectedFilename := "300_blur5.jpg"

	if url != expectedURL {
		t.Errorf("expected URL %q, got %q", expectedURL, url)
	}
	if filename != expectedFilename {
		t.Errorf("expected filename %q, got %q", expectedFilename, filename)
	}
}

func TestBuildURL_WithBlurLevel_TwoArguments(t *testing.T) {
	args := []string{"300", "200"}
	url, filename, err := BuildURL(args, "", "", false, false, 10)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedURL := "https://picsum.photos/300/200?blur=10"
	expectedFilename := "300x200_blur10.jpg"

	if url != expectedURL {
		t.Errorf("expected URL %q, got %q", expectedURL, url)
	}
	if filename != expectedFilename {
		t.Errorf("expected filename %q, got %q", expectedFilename, filename)
	}
}

func TestBuildURL_WithBlurLevel_WithImageID(t *testing.T) {
	args := []string{"300"}
	url, filename, err := BuildURL(args, "237", "", false, false, 3)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedURL := "https://picsum.photos/id/237/300?blur=3"
	expectedFilename := "id_237_300_blur3.jpg"

	if url != expectedURL {
		t.Errorf("expected URL %q, got %q", expectedURL, url)
	}
	if filename != expectedFilename {
		t.Errorf("expected filename %q, got %q", expectedFilename, filename)
	}
}

func TestBuildURL_WithGrayscaleAndBlurLevel_SingleArgument(t *testing.T) {
	args := []string{"300"}
	url, filename, err := BuildURL(args, "", "", true, false, 7)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedURL := "https://picsum.photos/300?grayscale&blur=7"
	expectedFilename := "300_gray_blur7.jpg"

	if url != expectedURL {
		t.Errorf("expected URL %q, got %q", expectedURL, url)
	}
	if filename != expectedFilename {
		t.Errorf("expected filename %q, got %q", expectedFilename, filename)
	}
}

func TestBuildURL_WithGrayscaleAndBlurLevel_WithSeed(t *testing.T) {
	args := []string{"300", "200"}
	url, filename, err := BuildURL(args, "", "picsum", true, false, 8)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedURL := "https://picsum.photos/seed/picsum/300/200?grayscale&blur=8"
	expectedFilename := "seed_picsum_300x200_gray_blur8.jpg"

	if url != expectedURL {
		t.Errorf("expected URL %q, got %q", expectedURL, url)
	}
	if filename != expectedFilename {
		t.Errorf("expected filename %q, got %q", expectedFilename, filename)
	}
}
