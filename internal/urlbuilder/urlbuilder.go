/*
Package urlbuilder to build url
*/
package urlbuilder

import (
	"fmt"
	"strconv"
)

// BuildURL constructs the picsum.photos URL and filename based on arguments and options
func BuildURL(args []string, imageID, seed string, grayscale bool) (url, filename string, err error) {
	if len(args) == 1 {
		// Parse single number
		num1, err := strconv.Atoi(args[0])
		if err != nil {
			return "", "", fmt.Errorf("invalid number: %s", args[0])
		}
		if imageID != "" {
			url = fmt.Sprintf("https://picsum.photos/id/%s/%d", imageID, num1)
			filename = fmt.Sprintf("id_%s_%d", imageID, num1)
		} else if seed != "" {
			url = fmt.Sprintf("https://picsum.photos/seed/%s/%d", seed, num1)
			filename = fmt.Sprintf("seed_%s_%d", seed, num1)
		} else {
			url = fmt.Sprintf("https://picsum.photos/%d", num1)
			filename = fmt.Sprintf("%d", num1)
		}
	} else {
		// Parse two numbers
		num1, err := strconv.Atoi(args[0])
		if err != nil {
			return "", "", fmt.Errorf("invalid first number: %s", args[0])
		}
		num2, err := strconv.Atoi(args[1])
		if err != nil {
			return "", "", fmt.Errorf("invalid second number: %s", args[1])
		}
		if imageID != "" {
			url = fmt.Sprintf("https://picsum.photos/id/%s/%d/%d", imageID, num1, num2)
			filename = fmt.Sprintf("id_%s_%dx%d", imageID, num1, num2)
		} else if seed != "" {
			url = fmt.Sprintf("https://picsum.photos/seed/%s/%d/%d", seed, num1, num2)
			filename = fmt.Sprintf("seed_%s_%dx%d", seed, num1, num2)
		} else {
			url = fmt.Sprintf("https://picsum.photos/%d/%d", num1, num2)
			filename = fmt.Sprintf("%dx%d", num1, num2)
		}
	}

	if grayscale {
		url += "?grayscale"
		filename += "_gray"
	}
	filename += ".jpg"

	return url, filename, nil
}
