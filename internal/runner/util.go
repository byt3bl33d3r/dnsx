package runner

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	stdinMarker = "-"
	Comma       = ","
	NewLine     = "\n"
)

func linesInFile(fileName string) ([]string, error) {
	result := []string{}
	f, err := os.Open(fileName)
	if err != nil {
		return result, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result, nil
}

// isURL tests a string to determine if it is a well-structured url or not.
func isURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func extractDomain(URL string) string {
	u, err := url.Parse(URL)
	if err != nil {
		return ""
	}

	return u.Hostname()
}

func prepareResolver(resolver string) string {
	resolver = strings.TrimSpace(resolver)
	if !strings.Contains(resolver, ":") {
		resolver += ":53"
	}
	return resolver
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%d:%02d:%02d", h, m, s)
}
