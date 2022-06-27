// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WirelessWirelessLansBulkDeleteReader is a Reader for the WirelessWirelessLansBulkDelete structure.
type WirelessWirelessLansBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLansBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewWirelessWirelessLansBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWirelessWirelessLansBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWirelessWirelessLansBulkDeleteNoContent creates a WirelessWirelessLansBulkDeleteNoContent with default headers values
func NewWirelessWirelessLansBulkDeleteNoContent() *WirelessWirelessLansBulkDeleteNoContent {
	return &WirelessWirelessLansBulkDeleteNoContent{}
}

/* WirelessWirelessLansBulkDeleteNoContent describes a response with status code 204, with default header values.

WirelessWirelessLansBulkDeleteNoContent wireless wireless lans bulk delete no content
*/
type WirelessWirelessLansBulkDeleteNoContent struct {
}

func (o *WirelessWirelessLansBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /wireless/wireless-lans/][%d] wirelessWirelessLansBulkDeleteNoContent ", 204)
}

func (o *WirelessWirelessLansBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWirelessWirelessLansBulkDeleteDefault creates a WirelessWirelessLansBulkDeleteDefault with default headers values
func NewWirelessWirelessLansBulkDeleteDefault(code int) *WirelessWirelessLansBulkDeleteDefault {
	return &WirelessWirelessLansBulkDeleteDefault{
		_statusCode: code,
	}
}

/* WirelessWirelessLansBulkDeleteDefault describes a response with status code -1, with default header values.

WirelessWirelessLansBulkDeleteDefault wireless wireless lans bulk delete default
*/
type WirelessWirelessLansBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the wireless wireless lans bulk delete default response
func (o *WirelessWirelessLansBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *WirelessWirelessLansBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /wireless/wireless-lans/][%d] wireless_wireless-lans_bulk_delete default  %+v", o._statusCode, o.Payload)
}
func (o *WirelessWirelessLansBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *WirelessWirelessLansBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
