/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// KeyServerMonitorEncryptedSEKIdentity represents the Identity of the object
var KeyServerMonitorEncryptedSEKIdentity = bambou.Identity{
	Name:     "keyservermonitorencryptedsek",
	Category: "keyservermonitorencryptedseks",
}

// KeyServerMonitorEncryptedSEKsList represents a list of KeyServerMonitorEncryptedSEKs
type KeyServerMonitorEncryptedSEKsList []*KeyServerMonitorEncryptedSEK

// KeyServerMonitorEncryptedSEKsAncestor is the interface that an ancestor of a KeyServerMonitorEncryptedSEK must implement.
// An Ancestor is defined as an entity that has KeyServerMonitorEncryptedSEK as a descendant.
// An Ancestor can get a list of its child KeyServerMonitorEncryptedSEKs, but not necessarily create one.
type KeyServerMonitorEncryptedSEKsAncestor interface {
	KeyServerMonitorEncryptedSEKs(*bambou.FetchingInfo) (KeyServerMonitorEncryptedSEKsList, *bambou.Error)
}

// KeyServerMonitorEncryptedSEKsParent is the interface that a parent of a KeyServerMonitorEncryptedSEK must implement.
// A Parent is defined as an entity that has KeyServerMonitorEncryptedSEK as a child.
// A Parent is an Ancestor which can create a KeyServerMonitorEncryptedSEK.
type KeyServerMonitorEncryptedSEKsParent interface {
	KeyServerMonitorEncryptedSEKsAncestor
	CreateKeyServerMonitorEncryptedSEK(*KeyServerMonitorEncryptedSEK) *bambou.Error
}

// KeyServerMonitorEncryptedSEK represents the model of a keyservermonitorencryptedsek
type KeyServerMonitorEncryptedSEK struct {
	ID                                        string  `json:"ID,omitempty"`
	ParentID                                  string  `json:"parentID,omitempty"`
	ParentType                                string  `json:"parentType,omitempty"`
	Owner                                     string  `json:"owner,omitempty"`
	NSGCertificateSerialNumber                float64 `json:"NSGCertificateSerialNumber,omitempty"`
	LastUpdatedBy                             string  `json:"lastUpdatedBy,omitempty"`
	GatewaySecuredDataID                      string  `json:"gatewaySecuredDataID,omitempty"`
	KeyServerCertificateSerialNumber          float64 `json:"keyServerCertificateSerialNumber,omitempty"`
	EntityScope                               string  `json:"entityScope,omitempty"`
	AssociatedKeyServerMonitorSEKCreationTime float64 `json:"associatedKeyServerMonitorSEKCreationTime,omitempty"`
	AssociatedKeyServerMonitorSEKID           string  `json:"associatedKeyServerMonitorSEKID,omitempty"`
	ExternalID                                string  `json:"externalID,omitempty"`
}

// NewKeyServerMonitorEncryptedSEK returns a new *KeyServerMonitorEncryptedSEK
func NewKeyServerMonitorEncryptedSEK() *KeyServerMonitorEncryptedSEK {

	return &KeyServerMonitorEncryptedSEK{}
}

// Identity returns the Identity of the object.
func (o *KeyServerMonitorEncryptedSEK) Identity() bambou.Identity {

	return KeyServerMonitorEncryptedSEKIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *KeyServerMonitorEncryptedSEK) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *KeyServerMonitorEncryptedSEK) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the KeyServerMonitorEncryptedSEK from the server
func (o *KeyServerMonitorEncryptedSEK) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the KeyServerMonitorEncryptedSEK into the server
func (o *KeyServerMonitorEncryptedSEK) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the KeyServerMonitorEncryptedSEK from the server
func (o *KeyServerMonitorEncryptedSEK) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the KeyServerMonitorEncryptedSEK
func (o *KeyServerMonitorEncryptedSEK) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the KeyServerMonitorEncryptedSEK
func (o *KeyServerMonitorEncryptedSEK) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the KeyServerMonitorEncryptedSEK
func (o *KeyServerMonitorEncryptedSEK) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the KeyServerMonitorEncryptedSEK
func (o *KeyServerMonitorEncryptedSEK) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
