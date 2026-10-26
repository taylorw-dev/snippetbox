package main

import (
	"bytes"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-playground/form/v4"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	app.logger.Error(
		err.Error(),
		slog.String("method", method),
		slog.String("uri", uri),
		slog.String("trace", trace),
	)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) error {
	ts, ok := app.templateCache[page]
	if !ok {
		return fmt.Errorf("the template %s does not exist in the cache", page)
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		return err
	}

	w.WriteHeader(status)
	_, err = buf.WriteTo(w)
	return err
}

func (app *application) newTemplateData(r *http.Request) templateData {
	return templateData{
		CurrentYear: time.Now().Year(),
	}
}

func (app *application) decodePostForm(r *http.Request, dst any) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	err = app.formDecoder.Decode(dst, r.PostForm)
	if err != nil {
		if _, ok := errors.AsType[*form.InvalidDecoderError](err); ok {
			panic(err)
		}

		return err
	}

	return nil
}
