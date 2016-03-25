package main

import (
	"compress/gzip"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"html/template"
	"lim.com"
	"log"
	"net/http"
	"os"
)

func init() {
	lim_com.R = render.New(render.Options{
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
}

func main() {

	mux := mux.NewRouter()

	AppRouter(mux)

	//plugins
	loggedRouter := handlers.LoggingHandler(os.Stdout, mux)
	gzipHandle := handlers.CompressHandlerLevel(loggedRouter, gzip.BestCompression)

	err := http.ListenAndServe(":8080", gzipHandle)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
