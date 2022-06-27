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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimSiteGroupsDeleteReader is a Reader for the DcimSiteGroupsDelete structure.
type DcimSiteGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSiteGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimSiteGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimSiteGroupsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimSiteGroupsDeleteNoContent creates a DcimSiteGroupsDeleteNoContent with default headers values
func NewDcimSiteGroupsDeleteNoContent() *DcimSiteGroupsDeleteNoContent {
	return &DcimSiteGroupsDeleteNoContent{}
}

/* DcimSiteGroupsDeleteNoContent describes a response with status code 204, with default header values.

DcimSiteGroupsDeleteNoContent dcim site groups delete no content
*/
type DcimSiteGroupsDeleteNoContent struct {
}

func (o *DcimSiteGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/site-groups/{id}/][%d] dcimSiteGroupsDeleteNoContent ", 204)
}

func (o *DcimSiteGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimSiteGroupsDeleteDefault creates a DcimSiteGroupsDeleteDefault with default headers values
func NewDcimSiteGroupsDeleteDefault(code int) *DcimSiteGroupsDeleteDefault {
	return &DcimSiteGroupsDeleteDefault{
		_statusCode: code,
	}
}

/* DcimSiteGroupsDeleteDefault describes a response with status code -1, with default header values.

DcimSiteGroupsDeleteDefault dcim site groups delete default
*/
type DcimSiteGroupsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim site groups delete default response
func (o *DcimSiteGroupsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimSiteGroupsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/site-groups/{id}/][%d] dcim_site-groups_delete default  %+v", o._statusCode, o.Payload)
}
func (o *DcimSiteGroupsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSiteGroupsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
