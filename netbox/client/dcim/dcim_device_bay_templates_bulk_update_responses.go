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

// DcimDeviceBayTemplatesBulkUpdateReader is a Reader for the DcimDeviceBayTemplatesBulkUpdate structure.
type DcimDeviceBayTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBayTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceBayTemplatesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceBayTemplatesBulkUpdateOK creates a DcimDeviceBayTemplatesBulkUpdateOK with default headers values
func NewDcimDeviceBayTemplatesBulkUpdateOK() *DcimDeviceBayTemplatesBulkUpdateOK {
	return &DcimDeviceBayTemplatesBulkUpdateOK{}
}

/* DcimDeviceBayTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

DcimDeviceBayTemplatesBulkUpdateOK dcim device bay templates bulk update o k
*/
type DcimDeviceBayTemplatesBulkUpdateOK struct {
	Payload *models.DeviceBayTemplate
}

func (o *DcimDeviceBayTemplatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-bay-templates/][%d] dcimDeviceBayTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceBayTemplatesBulkUpdateOK) GetPayload() *models.DeviceBayTemplate {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBayTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceBayTemplatesBulkUpdateDefault creates a DcimDeviceBayTemplatesBulkUpdateDefault with default headers values
func NewDcimDeviceBayTemplatesBulkUpdateDefault(code int) *DcimDeviceBayTemplatesBulkUpdateDefault {
	return &DcimDeviceBayTemplatesBulkUpdateDefault{
		_statusCode: code,
	}
}

/* DcimDeviceBayTemplatesBulkUpdateDefault describes a response with status code -1, with default header values.

DcimDeviceBayTemplatesBulkUpdateDefault dcim device bay templates bulk update default
*/
type DcimDeviceBayTemplatesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device bay templates bulk update default response
func (o *DcimDeviceBayTemplatesBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimDeviceBayTemplatesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-bay-templates/][%d] dcim_device-bay-templates_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimDeviceBayTemplatesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
