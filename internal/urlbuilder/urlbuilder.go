/*
Package urlbuilder to build url
*/
package urlbuilder

import (
	"fmt"
	"strconv"
)

// buildQueryParamsAndSuffix builds query parameters and filename suffix based on image options
func buildQueryParamsAndSuffix(grayscale, blur bool, blurLevel int) (queryParams, filenameSuffix string) {
	if grayscale && blurLevel > 0 {
		queryParams = fmt.Sprintf("?grayscale&blur=%d", blurLevel)
		filenameSuffix = fmt.Sprintf("_gray_blur%d", blurLevel)
	} else if grayscale && blur {
		queryParams = "?grayscale&blur"
		filenameSuffix = "_gray_blur"
	} else if grayscale {
		queryParams = "?grayscale"
		filenameSuffix = "_gray"
	} else if blurLevel > 0 {
		queryParams = fmt.Sprintf("?blur=%d", blurLevel)
		filenameSuffix = fmt.Sprintf("_blur%d", blurLevel)
	} else if blur {
		queryParams = "?blur"
		filenameSuffix = "_blur"
	}
	return queryParams, filenameSuffix
}

// BuildURL constructs the picsum.photos URL and filename based on arguments and options
func BuildURL(args []string, imageID, seed string, grayscale, blur bool, blurLevel int) (url, filename string, err error) {
	subPath := ""
	filePrefix := ""

	if seed != "" {
		subPath = fmt.Sprintf("seed/%s/", seed)
		filePrefix = "seed_"
	} else if imageID != "" {
		subPath = fmt.Sprintf("id/%s/", imageID)
		filePrefix = "id_"
	}

	if len(args) == 1 {
		// Parse single number
		num1, err := strconv.Atoi(args[0])
		if err != nil {
			return "", "", fmt.Errorf("invalid number: %s", args[0])
		}
		url = fmt.Sprintf("https://picsum.photos/%s%d", subPath, num1)
		filename = fmt.Sprintf("%s%d", filePrefix, num1)

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
		url = fmt.Sprintf("https://picsum.photos/%s%d/%d", subPath, num1, num2)
		filename = fmt.Sprintf("%s%dx%d", filePrefix, num1, num2)
	}

	queryParams, filenameSuffix := buildQueryParamsAndSuffix(grayscale, blur, blurLevel)
	url += queryParams
	filename += filenameSuffix + ".jpg"

	return url, filename, nil
}
