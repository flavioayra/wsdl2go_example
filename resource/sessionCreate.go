package sessioncreatesoapbinding

import (
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "https://webservices.sabre.com/websvc"

// NewSessionCreatePortType creates an initializes a SessionCreatePortType.
func NewSessionCreatePortType(cli *soap.Client) SessionCreatePortType {
	return &sessionCreatePortType{cli}
}

// SessionCreatePortType was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type SessionCreatePortType interface {
	// SessionCreateRQ was auto-generated from WSDL.
	SessionCreateRQ(header *MessageHeader, header2 *Security, body *SessionCreateRQ) (*MessageHeader, *Security, *SessionCreateRS, error)
}

// DateTime in WSDL format.
type DateTime string

// CryptoBinary was auto-generated from WSDL.
type CryptoBinary []byte

// DigestValueType was auto-generated from WSDL.
type DigestValueType []byte

// HMACOutputLengthType was auto-generated from WSDL.
type HMACOutputLengthType int64

// MessageStatusType was auto-generated from WSDL.
type messageStatusType *NMTOKEN

// Validate validates messageStatusType.
func (v messageStatusType) Validate() bool {
	for _, vv := range []*NMTOKEN{
		UnAuthorized,
		NotRecognized,
		Received,
		Processed,
		Forwarded,
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// NomEmptyString was auto-generated from WSDL.
type nomEmptyString string

// SeverityType was auto-generated from WSDL.
type severityType *NMTOKEN

// Validate validates severityType.
func (v severityType) Validate() bool {
	for _, vv := range []*NMTOKEN{
		Warning,
		Error,
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// StatusType was auto-generated from WSDL.
type statusType *NMTOKEN

// Validate validates statusType.
func (v statusType) Validate() bool {
	for _, vv := range []*NMTOKEN{
		Reset,
		Continue,
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AckRequested was auto-generated from WSDL.
type AckRequested []interface{}

// Acknowledgment was auto-generated from WSDL.
type Acknowledgment []interface{}

// Body was auto-generated from WSDL.
type Body []interface{}

// CanonicalizationMethodType was auto-generated from WSDL.
type CanonicalizationMethodType []interface{}

// DSAKeyValueType was auto-generated from WSDL.
type DSAKeyValueType struct {
	G *CryptoBinary `xml:"G,omitempty" json:"G,omitempty" yaml:"G,omitempty"`
	Y *CryptoBinary `xml:"Y,omitempty" json:"Y,omitempty" yaml:"Y,omitempty"`
	J *CryptoBinary `xml:"J,omitempty" json:"J,omitempty" yaml:"J,omitempty"`
}

// Description was auto-generated from WSDL.
type Description struct {
}

// DigestMethodType was auto-generated from WSDL.
type DigestMethodType []interface{}

// Envelope was auto-generated from WSDL.
type Envelope []interface{}

// Error was auto-generated from WSDL.
type Error []interface{}

// ErrorList was auto-generated from WSDL.
type ErrorList []interface{}

// 	    Fault reporting structure
type Fault struct {
	Faultcode   *string `xml:"faultcode,omitempty" json:"faultcode,omitempty" yaml:"faultcode,omitempty"`
	Faultstring *string `xml:"faultstring,omitempty" json:"faultstring,omitempty" yaml:"faultstring,omitempty"`
	Faultactor  *string `xml:"faultactor,omitempty" json:"faultactor,omitempty" yaml:"faultactor,omitempty"`
	Detail      *Detail `xml:"detail,omitempty" json:"detail,omitempty" yaml:"detail,omitempty"`
}

// From was auto-generated from WSDL.
type From struct {
	Role *nomEmptyString `xml:"Role,omitempty" json:"Role,omitempty" yaml:"Role,omitempty"`
}

// Header was auto-generated from WSDL.
type Header []interface{}

// KeyInfoType was auto-generated from WSDL.
type KeyInfoType struct {
}

// KeyValueType was auto-generated from WSDL.
type KeyValueType struct {
}

// Manifest was auto-generated from WSDL.
type Manifest []interface{}

// ManifestType was auto-generated from WSDL.
type ManifestType struct {
	Reference *ReferenceType `xml:"Reference,omitempty" json:"Reference,omitempty" yaml:"Reference,omitempty"`
	Id        *ID            `xml:"Id,attr,omitempty" json:"Id,attr,omitempty" yaml:"Id,attr,omitempty"`
}

// MessageData was auto-generated from WSDL.
type MessageData struct {
	MessageId      *nomEmptyString `xml:"MessageId,omitempty" json:"MessageId,omitempty" yaml:"MessageId,omitempty"`
	Timestamp      *string         `xml:"Timestamp,omitempty" json:"Timestamp,omitempty" yaml:"Timestamp,omitempty"`
	RefToMessageId *nomEmptyString `xml:"RefToMessageId,omitempty" json:"RefToMessageId,omitempty" yaml:"RefToMessageId,omitempty"`
	TimeToLive     *DateTime       `xml:"TimeToLive,omitempty" json:"TimeToLive,omitempty" yaml:"TimeToLive,omitempty"`
	Timeout        *uint           `xml:"Timeout,omitempty" json:"Timeout,omitempty" yaml:"Timeout,omitempty"`
}

// MessageHeader was auto-generated from WSDL.
type MessageHeader []interface{}

// MessageOrder was auto-generated from WSDL.
type MessageOrder []interface{}

// ObjectType was auto-generated from WSDL.
type ObjectType []interface{}

// PGPDataType was auto-generated from WSDL.
type PGPDataType []interface{}

// PartyId was auto-generated from WSDL.
type PartyId struct {
}

// RSAKeyValueType was auto-generated from WSDL.
type RSAKeyValueType struct {
	Modulus  *CryptoBinary `xml:"Modulus,omitempty" json:"Modulus,omitempty" yaml:"Modulus,omitempty"`
	Exponent *CryptoBinary `xml:"Exponent,omitempty" json:"Exponent,omitempty" yaml:"Exponent,omitempty"`
}

// Reference was auto-generated from WSDL.
type Reference []interface{}

// ReferenceType was auto-generated from WSDL.
type ReferenceType struct {
	Transforms   *TransformsType   `xml:"Transforms,omitempty" json:"Transforms,omitempty" yaml:"Transforms,omitempty"`
	DigestMethod *DigestMethodType `xml:"DigestMethod,omitempty" json:"DigestMethod,omitempty" yaml:"DigestMethod,omitempty"`
	DigestValue  *DigestValueType  `xml:"DigestValue,omitempty" json:"DigestValue,omitempty" yaml:"DigestValue,omitempty"`
	Id           *ID               `xml:"Id,attr,omitempty" json:"Id,attr,omitempty" yaml:"Id,attr,omitempty"`
	URI          string            `xml:"URI,attr,omitempty" json:"URI,attr,omitempty" yaml:"URI,attr,omitempty"`
	Type         string            `xml:"Type,attr,omitempty" json:"Type,attr,omitempty" yaml:"Type,attr,omitempty"`
}

// RetrievalMethodType was auto-generated from WSDL.
type RetrievalMethodType struct {
	Transforms *TransformsType `xml:"Transforms,omitempty" json:"Transforms,omitempty" yaml:"Transforms,omitempty"`
	URI        string          `xml:"URI,attr,omitempty" json:"URI,attr,omitempty" yaml:"URI,attr,omitempty"`
	Type       string          `xml:"Type,attr,omitempty" json:"Type,attr,omitempty" yaml:"Type,attr,omitempty"`
}

// SPKIDataType was auto-generated from WSDL.
type SPKIDataType []interface{}

// Schema was auto-generated from WSDL.
type Schema struct {
}

// Security was auto-generated from WSDL.
type Security struct {
	UsernameToken       *string `xml:"UsernameToken,omitempty" json:"UsernameToken,omitempty" yaml:"UsernameToken,omitempty"`
	SabreAth            *string `xml:"SabreAth,omitempty" json:"SabreAth,omitempty" yaml:"SabreAth,omitempty"`
	BinarySecurityToken *string `xml:"BinarySecurityToken,omitempty" json:"BinarySecurityToken,omitempty" yaml:"BinarySecurityToken,omitempty"`
}

// Service was auto-generated from WSDL.
type Service struct {
}

// SignatureMethodType was auto-generated from WSDL.
type SignatureMethodType []interface{}

// SignaturePropertiesType was auto-generated from WSDL.
type SignaturePropertiesType struct {
	SignatureProperty *SignaturePropertyType `xml:"SignatureProperty,omitempty" json:"SignatureProperty,omitempty" yaml:"SignatureProperty,omitempty"`
	Id                *ID                    `xml:"Id,attr,omitempty" json:"Id,attr,omitempty" yaml:"Id,attr,omitempty"`
}

// SignaturePropertyType was auto-generated from WSDL.
type SignaturePropertyType struct {
}

// SignatureType was auto-generated from WSDL.
type SignatureType struct {
	SignedInfo     *SignedInfoType     `xml:"SignedInfo,omitempty" json:"SignedInfo,omitempty" yaml:"SignedInfo,omitempty"`
	SignatureValue *SignatureValueType `xml:"SignatureValue,omitempty" json:"SignatureValue,omitempty" yaml:"SignatureValue,omitempty"`
	KeyInfo        *KeyInfoType        `xml:"KeyInfo,omitempty" json:"KeyInfo,omitempty" yaml:"KeyInfo,omitempty"`
	Object         *ObjectType         `xml:"Object,omitempty" json:"Object,omitempty" yaml:"Object,omitempty"`
	Id             *ID                 `xml:"Id,attr,omitempty" json:"Id,attr,omitempty" yaml:"Id,attr,omitempty"`
}

// SignatureValueType was auto-generated from WSDL.
type SignatureValueType struct {
}

// SignedInfoType was auto-generated from WSDL.
type SignedInfoType struct {
	CanonicalizationMethod *CanonicalizationMethodType `xml:"CanonicalizationMethod,omitempty" json:"CanonicalizationMethod,omitempty" yaml:"CanonicalizationMethod,omitempty"`
	SignatureMethod        *SignatureMethodType        `xml:"SignatureMethod,omitempty" json:"SignatureMethod,omitempty" yaml:"SignatureMethod,omitempty"`
	Reference              *ReferenceType              `xml:"Reference,omitempty" json:"Reference,omitempty" yaml:"Reference,omitempty"`
	Id                     *ID                         `xml:"Id,attr,omitempty" json:"Id,attr,omitempty" yaml:"Id,attr,omitempty"`
}

// StatusRequest was auto-generated from WSDL.
type StatusRequest []interface{}

// StatusResponse was auto-generated from WSDL.
type StatusResponse []interface{}

// SyncReply was auto-generated from WSDL.
type SyncReply []interface{}

// To was auto-generated from WSDL.
type To struct {
	Role *nomEmptyString `xml:"Role,omitempty" json:"Role,omitempty" yaml:"Role,omitempty"`
}

// TransformType was auto-generated from WSDL.
type TransformType struct {
}

// TransformsType was auto-generated from WSDL.
type TransformsType struct {
	Transform *TransformType `xml:"Transform,omitempty" json:"Transform,omitempty" yaml:"Transform,omitempty"`
}

// X509DataType was auto-generated from WSDL.
type X509DataType struct {
}

// X509IssuerSerialType was auto-generated from WSDL.
type X509IssuerSerialType struct {
	X509IssuerName   *string `xml:"X509IssuerName,omitempty" json:"X509IssuerName,omitempty" yaml:"X509IssuerName,omitempty"`
	X509SerialNumber *int64  `xml:"X509SerialNumber,omitempty" json:"X509SerialNumber,omitempty" yaml:"X509SerialNumber,omitempty"`
}

// Detail was auto-generated from WSDL.
type Detail []interface{}

// SequenceNumberType was auto-generated from WSDL.
type SequenceNumberType struct {
}

// Operation wrapper for SessionCreateRQ.
// OperationGetSessionCreateInput was auto-generated from WSDL.
type OperationGetSessionCreateInput struct {
	Header  *MessageHeader   `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	Header2 *Security        `xml:"header2,omitempty" json:"header2,omitempty" yaml:"header2,omitempty"`
	Body    *SessionCreateRQ `xml:"body,omitempty" json:"body,omitempty" yaml:"body,omitempty"`
}

// Operation wrapper for SessionCreateRQ.
// OperationGetSessionCreateOutput was auto-generated from WSDL.
type OperationGetSessionCreateOutput struct {
	Header  *MessageHeader   `xml:"header,omitempty" json:"header,omitempty" yaml:"header,omitempty"`
	Header2 *Security        `xml:"header2,omitempty" json:"header2,omitempty" yaml:"header2,omitempty"`
	Body    *SessionCreateRS `xml:"body,omitempty" json:"body,omitempty" yaml:"body,omitempty"`
}

// sessionCreatePortType implements the SessionCreatePortType interface.
type sessionCreatePortType struct {
	cli *soap.Client
}

// SessionCreateRQ was auto-generated from WSDL.
func (p *sessionCreatePortType) SessionCreateRQ(header *MessageHeader, header2 *Security, body *SessionCreateRQ) (*MessageHeader, *Security, *SessionCreateRS, error) {
	α := struct {
		OperationGetSessionCreateInput `xml:"-"`
	}{
		OperationGetSessionCreateInput{
			header,
			header2,
			body,
		},
	}

	γ := struct {
		OperationGetSessionCreateOutput `xml:"-"`
	}{}
	if err := p.cli.RoundTripWithAction("OTA", α, &γ); err != nil {
		return nil, nil, nil, err
	}
	return γ.Header, γ.Header2γ.Bodynil
}
