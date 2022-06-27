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

	"github.com/netbox-community/go-netbox/netbox/models"
)

// WirelessWirelessLanGroupsUpdateReader is a Reader for the WirelessWirelessLanGroupsUpdate structure.
type WirelessWirelessLanGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLanGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLanGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWirelessWirelessLanGroupsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWirelessWirelessLanGroupsUpdateOK creates a WirelessWirelessLanGroupsUpdateOK with default headers values
func NewWirelessWirelessLanGroupsUpdateOK() *WirelessWirelessLanGroupsUpdateOK {
	return &WirelessWirelessLanGroupsUpdateOK{}
}

/* WirelessWirelessLanGroupsUpdateOK describes a response with status code 200, with default header values.

WirelessWirelessLanGroupsUpdateOK wireless wireless lan groups update o k
*/
type WirelessWirelessLanGroupsUpdateOK struct {
	Payload *models.WirelessLANGroup
}

func (o *WirelessWirelessLanGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /wireless/wireless-lan-groups/{id}/][%d] wirelessWirelessLanGroupsUpdateOK  %+v", 200, o.Payload)
}
func (o *WirelessWirelessLanGroupsUpdateOK) GetPayload() *models.WirelessLANGroup {
	return o.Payload
}

func (o *WirelessWirelessLanGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWirelessWirelessLanGroupsUpdateDefault creates a WirelessWirelessLanGroupsUpdateDefault with default headers values
func NewWirelessWirelessLanGroupsUpdateDefault(code int) *WirelessWirelessLanGroupsUpdateDefault {
	return &WirelessWirelessLanGroupsUpdateDefault{
		_statusCode: code,
	}
}

/* WirelessWirelessLanGroupsUpdateDefault describes a response with status code -1, with default header values.

WirelessWirelessLanGroupsUpdateDefault wireless wireless lan groups update default
*/
type WirelessWirelessLanGroupsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the wireless wireless lan groups update default response
func (o *WirelessWirelessLanGroupsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *WirelessWirelessLanGroupsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /wireless/wireless-lan-groups/{id}/][%d] wireless_wireless-lan-groups_update default  %+v", o._statusCode, o.Payload)
}
func (o *WirelessWirelessLanGroupsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *WirelessWirelessLanGroupsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
