// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package http_after_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v3/models"
)

// DeleteHTTPAfterResponseRuleAcceptedCode is the HTTP code returned for type DeleteHTTPAfterResponseRuleAccepted
const DeleteHTTPAfterResponseRuleAcceptedCode int = 202

/*DeleteHTTPAfterResponseRuleAccepted Configuration change accepted and reload requested

swagger:response deleteHttpAfterResponseRuleAccepted
*/
type DeleteHTTPAfterResponseRuleAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteHTTPAfterResponseRuleAccepted creates DeleteHTTPAfterResponseRuleAccepted with default headers values
func NewDeleteHTTPAfterResponseRuleAccepted() *DeleteHTTPAfterResponseRuleAccepted {

	return &DeleteHTTPAfterResponseRuleAccepted{}
}

// WithReloadID adds the reloadId to the delete Http after response rule accepted response
func (o *DeleteHTTPAfterResponseRuleAccepted) WithReloadID(reloadID string) *DeleteHTTPAfterResponseRuleAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete Http after response rule accepted response
func (o *DeleteHTTPAfterResponseRuleAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteHTTPAfterResponseRuleAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteHTTPAfterResponseRuleNoContentCode is the HTTP code returned for type DeleteHTTPAfterResponseRuleNoContent
const DeleteHTTPAfterResponseRuleNoContentCode int = 204

/*DeleteHTTPAfterResponseRuleNoContent HTTP After Response Rule deleted

swagger:response deleteHttpAfterResponseRuleNoContent
*/
type DeleteHTTPAfterResponseRuleNoContent struct {
}

// NewDeleteHTTPAfterResponseRuleNoContent creates DeleteHTTPAfterResponseRuleNoContent with default headers values
func NewDeleteHTTPAfterResponseRuleNoContent() *DeleteHTTPAfterResponseRuleNoContent {

	return &DeleteHTTPAfterResponseRuleNoContent{}
}

// WriteResponse to the client
func (o *DeleteHTTPAfterResponseRuleNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteHTTPAfterResponseRuleNotFoundCode is the HTTP code returned for type DeleteHTTPAfterResponseRuleNotFound
const DeleteHTTPAfterResponseRuleNotFoundCode int = 404

/*DeleteHTTPAfterResponseRuleNotFound The specified resource was not found

swagger:response deleteHttpAfterResponseRuleNotFound
*/
type DeleteHTTPAfterResponseRuleNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteHTTPAfterResponseRuleNotFound creates DeleteHTTPAfterResponseRuleNotFound with default headers values
func NewDeleteHTTPAfterResponseRuleNotFound() *DeleteHTTPAfterResponseRuleNotFound {

	return &DeleteHTTPAfterResponseRuleNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete Http after response rule not found response
func (o *DeleteHTTPAfterResponseRuleNotFound) WithConfigurationVersion(configurationVersion string) *DeleteHTTPAfterResponseRuleNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete Http after response rule not found response
func (o *DeleteHTTPAfterResponseRuleNotFound) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete Http after response rule not found response
func (o *DeleteHTTPAfterResponseRuleNotFound) WithPayload(payload *models.Error) *DeleteHTTPAfterResponseRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Http after response rule not found response
func (o *DeleteHTTPAfterResponseRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteHTTPAfterResponseRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteHTTPAfterResponseRuleDefault General Error

swagger:response deleteHttpAfterResponseRuleDefault
*/
type DeleteHTTPAfterResponseRuleDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteHTTPAfterResponseRuleDefault creates DeleteHTTPAfterResponseRuleDefault with default headers values
func NewDeleteHTTPAfterResponseRuleDefault(code int) *DeleteHTTPAfterResponseRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteHTTPAfterResponseRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete HTTP after response rule default response
func (o *DeleteHTTPAfterResponseRuleDefault) WithStatusCode(code int) *DeleteHTTPAfterResponseRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete HTTP after response rule default response
func (o *DeleteHTTPAfterResponseRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete HTTP after response rule default response
func (o *DeleteHTTPAfterResponseRuleDefault) WithConfigurationVersion(configurationVersion string) *DeleteHTTPAfterResponseRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete HTTP after response rule default response
func (o *DeleteHTTPAfterResponseRuleDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete HTTP after response rule default response
func (o *DeleteHTTPAfterResponseRuleDefault) WithPayload(payload *models.Error) *DeleteHTTPAfterResponseRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete HTTP after response rule default response
func (o *DeleteHTTPAfterResponseRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteHTTPAfterResponseRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
