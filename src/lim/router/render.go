package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/unrolled/render"
	"html/template"
	"net/http"
)

var Render = render.New(render.Options{
	Directory:                 "templates",                                      // Specify what path to load the templates from.
	Layout:                    "layout/default",                                 // Specify a layout template. Layouts can call {{ yield }} to render the current template or {{ partial "css" }} to render a partial from the current template.
	Extensions:                []string{".tmpl", ".html"},                       // Specify extensions to load for templates.
	Funcs:                     []template.FuncMap{},                             // Specify helper function maps for templates to access.
	Delims:                    render.Delims{"{{", "}}"},                        // Sets delimiters to the specified strings.
	Charset:                   "UTF-8",                                          // Sets encoding for json and html content-types. Default is "UTF-8".
	IndentJSON:                true,                                             // Output human readable JSON.
	IndentXML:                 true,                                             // Output human readable XML.
	PrefixJSON:                []byte(")]}',\n"),                                // Prefixes JSON responses with the given bytes.
	PrefixXML:                 []byte("<?xml version='1.0' encoding='UTF-8'?>"), // Prefixes XML responses with the given bytes.
	HTMLContentType:           "text/html",                                      // Output XHTML content type instead of default "text/html".
	IsDevelopment:             true,                                             // Render will now recompile the templates on every HTML response.
	UnEscapeHTML:              true,                                             // Replace ensure '&<>' are output correctly (JSON only).
	StreamingJSON:             true,                                             // Streams the JSON response via json.Encoder.
	RequirePartials:           true,                                             // Return an error if a template is missing a partial used in a layout.
	DisableHTTPErrorRendering: false,                                            // Disables automatic rendering of http.StatusInternalServerError when an error occurs.
})

// HTML builds up the response from the specified template and bindings.
func HTML(c echo.Context, name string, binding interface{}, htmlOpt ...render.HTMLOptions) error {
	// `http.ResponseWriter`
	w := c.Response().(*standard.Response).ResponseWriter
	return Render.HTML(w, http.StatusOK, name, binding, htmlOpt...)
}
