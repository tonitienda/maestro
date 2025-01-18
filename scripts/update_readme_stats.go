package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
)

// TODO: Fix colors (green good, red bad)
func makeSVG(label string, value string) string {
	// Create a Markdown table
	return fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="110" height="20" role="img" aria-label="%s: %s">
  <title>%s: %s</title>
  <rect width="63" height="20" fill="#555" />
  <rect x="63" width="47" height="20" fill="#4c1" />
  <text x="31.5" y="14" fill="#fff" font-family="Verdana, sans-serif" font-size="10" text-anchor="middle">%s</text>
  <text x="86.5" y="14" fill="#fff" font-family="Verdana, sans-serif" font-size="10" text-anchor="middle">%s</text>
</svg>
`, label, value, label, value, label, value)

}

func makeSVGPercentage(label string, percentage float64) string {
	roundedPercentage := math.Round(percentage*10) / 10

	return makeSVG(label, fmt.Sprintf("%.1f%%", roundedPercentage))

}

func makeSVGNumber(label string, number int) string {

	return makeSVG(label, fmt.Sprintf("%d", number))

}

func main() {
	// Define command-line arguments
	coverage := flag.Float64("coverage", 0.0, "Coverage percentage (float)")
	tests := flag.Int("tests", 0, "Number of tests")

	// Parse the arguments
	flag.Parse()

	// Validate the required arguments
	if *coverage <= 0 || *tests <= 0 {
		fmt.Println("Usage: update_readme -coverage=<coverage> -tests=<tests>")
		os.Exit(1)
	}

	// Round the coverage to one decimal place

	// Create a Markdown table
	table := makeSVGPercentage("Coverage", *coverage) + makeSVGNumber("Tests run", *tests)

	// File path for the README
	filePath := "README.md"

	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Convert to string for processing
	contentStr := string(content)

	// Regex to find the <stats> block
	statsRegex := regexp.MustCompile(`(?s)<stats>.*?</stats>`)

	if statsRegex.MatchString(contentStr) {
		// Replace the <stats> block
		contentStr = statsRegex.ReplaceAllString(contentStr, "<stats>\n"+table+"\n</stats>")
	} else {
		// Append <stats> block if it doesn't exist
		contentStr += "\n<stats>\n" + table + "\n</stats>"
	}

	// Write the updated content back to the file
	err = os.WriteFile(filePath, []byte(contentStr), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("README.md updated successfully!")
}
