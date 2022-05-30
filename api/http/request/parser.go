package request

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/form"
)

const (
	defaultMaxMemory = 4 << 20 // 4 MB
	defaultTagName   = "json"
)

// CustomFunc Request decoder can handle custom types via RegisterCustomTypeFunc.
type CustomFunc struct {
	Func  form.DecodeCustomTypeFunc
	Types []interface{}
}

// Parser net/http request parser interface
type Parser interface {
	Form(r *http.Request, i interface{}, fns ...CustomFunc) error
	Query(r *http.Request, i interface{}, fns ...CustomFunc) error
}

// parser request parser object
type parser struct {
	maxSize int64
	tagName string
}

// NewDefaultParser initiate request parser with default limit
func NewDefaultParser() Parser {
	return &parser{
		maxSize: defaultMaxMemory,
		tagName: defaultTagName,
	}
}

// SetMaxSize set max http body size (in byte) for multipart form data
func (c *parser) SetMaxSize(limit int64) {
	c.maxSize = limit
}

// SetTagName tag name for conform parsing
func (c *parser) SetTagName(tag string) {
	c.tagName = tag
}

// Form decode body data from *http.Request.
// It will automagically bind to given struct pointer,
// and cleanup its values based on given struct definition.
func (c *parser) Form(r *http.Request, i interface{}, fns ...CustomFunc) error {
	if r.Body == nil {
		return nil
	}
	// nolint
	defer r.Body.Close()

	ct := r.Header.Get("Content-Type")
	if strings.HasPrefix(ct, "application/json") || strings.HasPrefix(ct, "text/json") { // dana request use text/json
		if err := json.NewDecoder(r.Body).Decode(i); err != nil {
			return err
		}
	} else if strings.HasPrefix(ct, "multipart/form-data") {
		if err := r.ParseMultipartForm(c.maxSize); err != nil {
			return err
		}
	} else {
		// default parse is urlencoded form type
		if err := r.ParseForm(); err != nil {
			return err
		}
	}

	d := form.NewDecoder()
	d.SetTagName(c.tagName)
	if len(fns) > 0 {
		for _, v := range fns {
			d.RegisterCustomTypeFunc(v.Func, v.Types...)
		}
	}

	return d.Decode(i, r.Form)
}

// Query decode query data from *http.Request.
// It will automagically bind to given struct pointer,
// and cleanup its values based on given struct definition.
func (c *parser) Query(r *http.Request, i interface{}, fns ...CustomFunc) error {
	d := form.NewDecoder()
	d.SetTagName(c.tagName)
	if len(fns) > 0 {
		for _, v := range fns {
			d.RegisterCustomTypeFunc(v.Func, v.Types...)
		}
	}

	return d.Decode(i, r.URL.Query())
}
