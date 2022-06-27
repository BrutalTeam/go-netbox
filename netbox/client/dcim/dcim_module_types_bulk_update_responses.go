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

	"github.com/netbox-community/go-netbox/netbox/models"
)

// DcimModuleTypesBulkUpdateReader is a Reader for the DcimModuleTypesBulkUpdate structure.
type DcimModuleTypesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleTypesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimModuleTypesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleTypesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleTypesBulkUpdateOK creates a DcimModuleTypesBulkUpdateOK with default headers values
func NewDcimModuleTypesBulkUpdateOK() *DcimModuleTypesBulkUpdateOK {
	return &DcimModuleTypesBulkUpdateOK{}
}

/* DcimModuleTypesBulkUpdateOK describes a response with status code 200, with default header values.

DcimModuleTypesBulkUpdateOK dcim module types bulk update o k
*/
type DcimModuleTypesBulkUpdateOK struct {
	Payload *models.ModuleType
}

func (o *DcimModuleTypesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/module-types/][%d] dcimModuleTypesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimModuleTypesBulkUpdateOK) GetPayload() *models.ModuleType {
	return o.Payload
}

func (o *DcimModuleTypesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModuleType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleTypesBulkUpdateDefault creates a DcimModuleTypesBulkUpdateDefault with default headers values
func NewDcimModuleTypesBulkUpdateDefault(code int) *DcimModuleTypesBulkUpdateDefault {
	return &DcimModuleTypesBulkUpdateDefault{
		_statusCode: code,
	}
}

/* DcimModuleTypesBulkUpdateDefault describes a response with status code -1, with default header values.

DcimModuleTypesBulkUpdateDefault dcim module types bulk update default
*/
type DcimModuleTypesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim module types bulk update default response
func (o *DcimModuleTypesBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimModuleTypesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/module-types/][%d] dcim_module-types_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimModuleTypesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleTypesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
