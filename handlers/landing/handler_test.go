package landing_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/imgproxy/imgproxy/v4/handlers/landing"
	"github.com/imgproxy/imgproxy/v4/httpheaders"
	"github.com/imgproxy/imgproxy/v4/server/responsewriter"
)

func TestLandingHandler(t *testing.T) {
	rwConf := responsewriter.NewDefaultConfig()
	rwf, err := responsewriter.NewFactory(&rwConf)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	h := landing.New()

	h.Execute("test-req-id", rwf.NewWriter(rr), nil)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "text/html", rr.Header().Get(httpheaders.ContentType))

	body := rr.Body.String()
	assert.Contains(t, body, "imgproxy Playground")
	assert.Contains(t, body, "URL Builder")
	assert.Contains(t, body, "/{signature}/{processing_options}/plain/{source_url}@{ext}")
}
