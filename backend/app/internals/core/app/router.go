package app

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"learn.zone01dakar.sn/forum-rest-api/app/lib"
)

type Route struct {
	Path    string         `route:"path"`
	Method  string         `route:"method"`
	Handler lib.Handler    `route:"-"`
	Pattern *regexp.Regexp `route:"-"` //
	Params  []string       `route:"params"`
}

func (r *Route) Init() {
	// Split the path into segments
	segments := strings.Split(r.Path, "/")

	// Define regex patterns for types
	typePatterns := map[string]string{
		"string": `([^/]+)`, // Matches any characters except '/'
		"int":    `(\d+)`,   // Matches one or more digits
	}

	// Iterate over segments and replace placeholders with regex patterns
	for i, segment := range segments {
		if strings.Contains(segment, "{") && strings.Contains(segment, "}") {
			// Extract type and name from the placeholder
			placeholder := strings.Trim(segment, "{}")
			parts := strings.Split(placeholder, ":")
			if len(parts) == 2 {
				r.Params = append(r.Params, parts[0])
				typ := parts[1]
				if pattern, ok := typePatterns[typ]; ok {
					// Replace placeholder with regex pattern
					segments[i] = fmt.Sprintf(`%s`, pattern)
				}
			}
		}
	}

	// Join segments back into a regex pattern
	r.Pattern = regexp.MustCompile("^/api" + strings.Join(segments, "/") + "$")
}

func (r *Route) RebuildURLWithParams(path *url.URL) {
	// Parse existing query parameters
	existingParams, err := url.ParseQuery(path.RawQuery)
	if err != nil {
		fmt.Println("Error parsing existing query parameters:", err)
		return
	}

	// Find submatches in the path
	values := r.Pattern.FindStringSubmatch(path.Path)

	if len(values) == len(r.Params)+1 {
		// Create a new URL values map
		params := url.Values{}

		// Add existing query parameters to the new map
		for k, v := range existingParams {
			params[k] = v
		}

		// Add matched parameters to the new map
		for k, v := range values[1:] {
			params.Add(r.Params[k], v)
		}

		// Encode the URL values and set the RawQuery field of the URL
		path.RawQuery = params.Encode()
	}
}
