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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// IpamPrefixesBulkUpdateReader is a Reader for the IpamPrefixesBulkUpdate structure.
type IpamPrefixesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamPrefixesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamPrefixesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamPrefixesBulkUpdateOK creates a IpamPrefixesBulkUpdateOK with default headers values
func NewIpamPrefixesBulkUpdateOK() *IpamPrefixesBulkUpdateOK {
	return &IpamPrefixesBulkUpdateOK{}
}

/* IpamPrefixesBulkUpdateOK describes a response with status code 200, with default header values.

IpamPrefixesBulkUpdateOK ipam prefixes bulk update o k
*/
type IpamPrefixesBulkUpdateOK struct {
	Payload *models.Prefix
}

func (o *IpamPrefixesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/prefixes/][%d] ipamPrefixesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamPrefixesBulkUpdateOK) GetPayload() *models.Prefix {
	return o.Payload
}

func (o *IpamPrefixesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefix)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamPrefixesBulkUpdateDefault creates a IpamPrefixesBulkUpdateDefault with default headers values
func NewIpamPrefixesBulkUpdateDefault(code int) *IpamPrefixesBulkUpdateDefault {
	return &IpamPrefixesBulkUpdateDefault{
		_statusCode: code,
	}
}

/* IpamPrefixesBulkUpdateDefault describes a response with status code -1, with default header values.

IpamPrefixesBulkUpdateDefault ipam prefixes bulk update default
*/
type IpamPrefixesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam prefixes bulk update default response
func (o *IpamPrefixesBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamPrefixesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/prefixes/][%d] ipam_prefixes_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *IpamPrefixesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamPrefixesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
