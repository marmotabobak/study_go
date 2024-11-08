package versionpage

import (
	"fmt"
	"net/http"
)

type VersionPage struct {
	version string
}

func (v *VersionPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", v.version)
}

func NewVersionPage(version string) http.Handler {
	return &VersionPage{version: version}
}
