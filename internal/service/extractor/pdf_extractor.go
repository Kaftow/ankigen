package extractor

import (
	"fmt"
	fitz "github.com/gen2brain/go-fitz"
	"strings"
)

type PDFExtractor struct{}

func NewPDFExtractor() *PDFExtractor {
	return &PDFExtractor{}
}

func (PDFExtractor) SupportedExtensions() []string {
	return []string{".pdf"}
}

// Extract converts a PDF file into machine-friendly Markdown
func (PDFExtractor) Extract(path string) (string, error) {
	doc, err := fitz.New(path)
	if err != nil {
		return "", fmt.Errorf("failed to open pdf: %w", err)
	}
	defer doc.Close()

	var b strings.Builder

	for page := 0; page < doc.NumPage(); page++ {
		text, err := doc.Text(page)
		if err != nil {
			continue
		}

		pageText, err := normalizePageText(text)
		if err != nil || pageText == "" {
			continue
		}

		// Separate pages with horizontal rules
		if page > 0 {
			b.WriteString("\n---\n\n")
		}

		b.WriteString(pageText)
	}

	return strings.TrimSpace(b.String()), nil
}

// Extracts and normalizes text from a single PDF page
func normalizePageText(text string) (string, error) {

	var b strings.Builder

	lines := prepareLogicalLines(text)
	lines = restoreBrokenLines(lines)

	for _, line := range lines {
		if isPDFNoise(line) {
			continue
		}

		b.WriteString(line)
		b.WriteString("\n")
	}

	return strings.TrimSpace(b.String()), nil
}

// Convert raw page text into a cleaned sequence of logical lines
func prepareLogicalLines(text string) []string {
	raw := strings.Split(text, "\n")
	out := make([]string, 0, len(raw))

	lastWasEmpty := false

	for _, line := range raw {
		line = strings.TrimSpace(line)

		if line == "" {
			// Collapse multiple empty lines into a single one
			if !lastWasEmpty {
				out = append(out, "")
				lastWasEmpty = true
			}
			continue
		}

		lastWasEmpty = false
		out = append(out, line)
	}

	return out
}

// Fixes line breaks and hyphenation caused by PDF layout
func restoreBrokenLines(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}

	var out []string
	var buf string

	flush := func() {
		if strings.TrimSpace(buf) != "" {
			out = append(out, strings.TrimSpace(buf))
		}
		buf = ""
	}

	for _, line := range lines {
		// Paragraph break
		if line == "" {
			flush()
			out = append(out, "")
			continue
		}

		if buf == "" {
			buf = line
			continue
		}

		// Hyphenation: "docu-" + "ment"
		if strings.HasSuffix(buf, "-") {
			buf = strings.TrimSuffix(buf, "-") + line
			continue
		}

		// Likely wrapped line: merge with space
		if isLineWrap(buf, line) {
			buf = buf + " " + line
			continue
		}

		// Otherwise, flush buffer and start a new one
		flush()
		buf = line
	}

	flush()
	return out
}

// Determines IF two consecutive lines are part of the same logical sentence
func isLineWrap(prev, curr string) bool {
	// Sentence-ending punctuation → do not merge
	if strings.HasSuffix(prev, ".") ||
		strings.HasSuffix(prev, "?") ||
		strings.HasSuffix(prev, "!") ||
		strings.HasSuffix(prev, ":") {
		return false
	}

	// Current line starts with lowercase letter → likely continuation
	r := []rune(curr)
	if len(r) > 0 && r[0] >= 'a' && r[0] <= 'z' {
		return true
	}

	return false
}

// Applies lightweight heuristics to remove common PDF noise
func isPDFNoise(line string) bool {
	// Pure page numbers
	if len(line) <= 3 && isDigits(line) {
		return true
	}

	// Common header/footer keywords
	lower := strings.ToLower(line)
	if strings.Contains(lower, "confidential") {
		return true
	}
	if strings.Contains(lower, "page ") {
		return true
	}

	return false
}

// Check if the string consists only of numeric characters
func isDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
