// Copyright 2016 Mender Software AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package view

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

// Headers
const (
	HttpHeaderLocation = "Location"
)

// Errors
var (
	ErrNotFound = errors.New("Resource not found")
)

type RESTView struct {
}

func (p *RESTView) RenderSuccessPost(w rest.ResponseWriter, r *rest.Request, id string) {
	w.Header().Add(HttpHeaderLocation, fmt.Sprintf("%s/%s", r.URL.String(), id))
	w.WriteHeader(http.StatusCreated)
}

func (p *RESTView) RenderSuccessGet(w rest.ResponseWriter, object interface{}) {
	w.WriteJson(object)
}

func (p *RESTView) RenderError(w rest.ResponseWriter, err error, status int) {
	rest.Error(w, err.Error(), status)
}

func (p *RESTView) RenderErrorNotFound(w rest.ResponseWriter) {
	p.RenderError(w, ErrNotFound, http.StatusNotFound)
}

func (p *RESTView) RenderSuccessDelete(w rest.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func (p *RESTView) RenderSuccessPut(w rest.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}