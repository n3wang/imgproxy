package landing

import (
	_ "embed"
	"net/http"

	"github.com/imgproxy/imgproxy/v4/httpheaders"
	"github.com/imgproxy/imgproxy/v4/server"
)

//go:embed body.html
var landingBody []byte

// Handler handles landing requests
type Handler struct{}

// New creates new handler object
func New() *Handler {
	return &Handler{}
}

// Execute handles the landing request
func (h *Handler) Execute(
	reqID string,
	rw server.ResponseWriter,
	req *http.Request,
) *server.Error {
	// Set route-specific CSP in response-writer result headers to override
	// the default sandbox policy used for image responses.
	csp := http.Header{}
	csp.Set(
		httpheaders.ContentSecurityPolicy,
		"default-src 'self'; script-src 'unsafe-inline'; style-src 'unsafe-inline'; img-src 'self' https: http: data:; connect-src 'self' https: http:; object-src 'none'; base-uri 'none'; frame-ancestors 'self'",
	)
	rw.CopyFrom(csp, []string{httpheaders.ContentSecurityPolicy})
	rw.Header().Set(httpheaders.ContentType, "text/html")
	rw.WriteHeader(http.StatusOK)
	rw.Write(landingBody)
	return nil
}
