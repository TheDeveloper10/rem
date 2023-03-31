package rem

import "path"

// * Borrowed from the net/http package.
// Returns the canonical path for p, eliminating . and .. elements.
func cleanPath(p string) string {
	if p == "" {
		return "/"
	}
	if p[0] != '/' {
		p = "/" + p
	}
	np := path.Clean(p)
	// path.Clean removes trailing slash except for root;
	// put the trailing slash back if necessary.
	if np[len(np)-1] != '/' {
		np += "/"
	}

	return np
}
