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

// DcimInventoryItemTemplatesBulkPartialUpdateReader is a Reader for the DcimInventoryItemTemplatesBulkPartialUpdate structure.
type DcimInventoryItemTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemTemplatesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemTemplatesBulkPartialUpdateOK creates a DcimInventoryItemTemplatesBulkPartialUpdateOK with default headers values
func NewDcimInventoryItemTemplatesBulkPartialUpdateOK() *DcimInventoryItemTemplatesBulkPartialUpdateOK {
	return &DcimInventoryItemTemplatesBulkPartialUpdateOK{}
}

/* DcimInventoryItemTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimInventoryItemTemplatesBulkPartialUpdateOK dcim inventory item templates bulk partial update o k
*/
type DcimInventoryItemTemplatesBulkPartialUpdateOK struct {
	Payload *models.InventoryItemTemplate
}

func (o *DcimInventoryItemTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-templates/][%d] dcimInventoryItemTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimInventoryItemTemplatesBulkPartialUpdateOK) GetPayload() *models.InventoryItemTemplate {
	return o.Payload
}

func (o *DcimInventoryItemTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItemTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemTemplatesBulkPartialUpdateDefault creates a DcimInventoryItemTemplatesBulkPartialUpdateDefault with default headers values
func NewDcimInventoryItemTemplatesBulkPartialUpdateDefault(code int) *DcimInventoryItemTemplatesBulkPartialUpdateDefault {
	return &DcimInventoryItemTemplatesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimInventoryItemTemplatesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimInventoryItemTemplatesBulkPartialUpdateDefault dcim inventory item templates bulk partial update default
*/
type DcimInventoryItemTemplatesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim inventory item templates bulk partial update default response
func (o *DcimInventoryItemTemplatesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimInventoryItemTemplatesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-templates/][%d] dcim_inventory-item-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimInventoryItemTemplatesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemTemplatesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
