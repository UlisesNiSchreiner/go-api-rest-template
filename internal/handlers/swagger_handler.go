package handlers

import (
	"net/http"
	"os"
)

type SwaggerHandler struct {
	specPath string
}

func NewSwaggerHandler(specPath string) *SwaggerHandler {
	return &SwaggerHandler{specPath: specPath}
}

func (h *SwaggerHandler) UI(w http.ResponseWriter, r *http.Request) {
	// Minimal Swagger UI that pulls assets from a CDN.
	// This keeps the template dependency-light; teams can vendor swagger-ui later if required.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	specURL := scheme + "://" + host + "/swagger/openapi.yaml"

	_, _ = w.Write([]byte(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>API Docs</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" />
  <style>
    body { margin: 0; }
  </style>
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-standalone-preset.js"></script>
  <script>
    window.onload = () => {
      SwaggerUIBundle({
        url: "` + specURL + `",
        dom_id: "#swagger-ui",
        presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
        layout: "StandaloneLayout"
      });
    };
  </script>
</body>
</html>`))
}

func (h *SwaggerHandler) Spec(w http.ResponseWriter, _ *http.Request) {
	b, err := os.ReadFile(h.specPath)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "unable to load openapi spec")
		return
	}
	w.Header().Set("Content-Type", "application/yaml; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)
}
