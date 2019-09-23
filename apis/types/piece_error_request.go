// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PieceErrorRequest Peer's detailed information in supernode.
// swagger:model PieceErrorRequest
type PieceErrorRequest struct {

	// the peer ID of the target Peer.
	//
	DstIP string `json:"dstIP,omitempty"`

	// the peer ID of the target Peer.
	//
	DstPid string `json:"dstPid,omitempty"`

	// the error type when failed to download from supernode that dfget will report to supernode
	//
	// Enum: [FILE_NOT_EXIST FILE_MD5_NOT_MATCH]
	ErrorType string `json:"errorType,omitempty"`

	// the MD5 value of piece which returned by the supernode that
	// in order to verify the correctness of the piece content which
	// downloaded from the other peers.
	//
	ExpectedMd5 string `json:"expectedMd5,omitempty"`

	// the range of specific piece in the task, example "0-45565".
	//
	Range string `json:"range,omitempty"`

	// the MD5 information of piece which calculated by the piece content
	// which downloaded from the target peer.
	//
	RealMd5 string `json:"realMd5,omitempty"`

	// the CID of the src Peer.
	//
	SrcCid string `json:"srcCid,omitempty"`

	// the taskID of the piece.
	//
	TaskID string `json:"taskId,omitempty"`
}

// Validate validates this piece error request
func (m *PieceErrorRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var pieceErrorRequestTypeErrorTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FILE_NOT_EXIST","FILE_MD5_NOT_MATCH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pieceErrorRequestTypeErrorTypePropEnum = append(pieceErrorRequestTypeErrorTypePropEnum, v)
	}
}

const (

	// PieceErrorRequestErrorTypeFILENOTEXIST captures enum value "FILE_NOT_EXIST"
	PieceErrorRequestErrorTypeFILENOTEXIST string = "FILE_NOT_EXIST"

	// PieceErrorRequestErrorTypeFILEMD5NOTMATCH captures enum value "FILE_MD5_NOT_MATCH"
	PieceErrorRequestErrorTypeFILEMD5NOTMATCH string = "FILE_MD5_NOT_MATCH"
)

// prop value enum
func (m *PieceErrorRequest) validateErrorTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, pieceErrorRequestTypeErrorTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PieceErrorRequest) validateErrorType(formats strfmt.Registry) error {

	if swag.IsZero(m.ErrorType) { // not required
		return nil
	}

	// value enum
	if err := m.validateErrorTypeEnum("errorType", "body", m.ErrorType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PieceErrorRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PieceErrorRequest) UnmarshalBinary(b []byte) error {
	var res PieceErrorRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}