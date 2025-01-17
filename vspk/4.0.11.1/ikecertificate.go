/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// IKECertificateIdentity represents the Identity of the object
var IKECertificateIdentity = bambou.Identity{
	Name:     "ikecertificate",
	Category: "ikecertificates",
}

// IKECertificatesList represents a list of IKECertificates
type IKECertificatesList []*IKECertificate

// IKECertificatesAncestor is the interface that an ancestor of a IKECertificate must implement.
// An Ancestor is defined as an entity that has IKECertificate as a descendant.
// An Ancestor can get a list of its child IKECertificates, but not necessarily create one.
type IKECertificatesAncestor interface {
	IKECertificates(*bambou.FetchingInfo) (IKECertificatesList, *bambou.Error)
}

// IKECertificatesParent is the interface that a parent of a IKECertificate must implement.
// A Parent is defined as an entity that has IKECertificate as a child.
// A Parent is an Ancestor which can create a IKECertificate.
type IKECertificatesParent interface {
	IKECertificatesAncestor
	CreateIKECertificate(*IKECertificate) *bambou.Error
}

// IKECertificate represents the model of a ikecertificate
type IKECertificate struct {
	ID                     string  `json:"ID,omitempty"`
	ParentID               string  `json:"parentID,omitempty"`
	ParentType             string  `json:"parentType,omitempty"`
	Owner                  string  `json:"owner,omitempty"`
	PEMEncoded             string  `json:"PEMEncoded,omitempty"`
	Name                   string  `json:"name,omitempty"`
	LastUpdatedBy          string  `json:"lastUpdatedBy,omitempty"`
	SerialNumber           int     `json:"serialNumber,omitempty"`
	Description            string  `json:"description,omitempty"`
	EntityScope            string  `json:"entityScope,omitempty"`
	NotAfter               float64 `json:"notAfter,omitempty"`
	NotBefore              float64 `json:"notBefore,omitempty"`
	AssociatedEnterpriseID string  `json:"associatedEnterpriseID,omitempty"`
	IssuerDN               string  `json:"issuerDN,omitempty"`
	SubjectDN              string  `json:"subjectDN,omitempty"`
	ExternalID             string  `json:"externalID,omitempty"`
}

// NewIKECertificate returns a new *IKECertificate
func NewIKECertificate() *IKECertificate {

	return &IKECertificate{}
}

// Identity returns the Identity of the object.
func (o *IKECertificate) Identity() bambou.Identity {

	return IKECertificateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IKECertificate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IKECertificate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IKECertificate from the server
func (o *IKECertificate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IKECertificate into the server
func (o *IKECertificate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IKECertificate from the server
func (o *IKECertificate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the IKECertificate
func (o *IKECertificate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the IKECertificate
func (o *IKECertificate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IKECertificate
func (o *IKECertificate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the IKECertificate
func (o *IKECertificate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
