/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// PATNATPoolIdentity represents the Identity of the object
var PATNATPoolIdentity = bambou.Identity{
	Name:     "patnatpool",
	Category: "patnatpools",
}

// PATNATPoolsList represents a list of PATNATPools
type PATNATPoolsList []*PATNATPool

// PATNATPoolsAncestor is the interface that an ancestor of a PATNATPool must implement.
// An Ancestor is defined as an entity that has PATNATPool as a descendant.
// An Ancestor can get a list of its child PATNATPools, but not necessarily create one.
type PATNATPoolsAncestor interface {
	PATNATPools(*bambou.FetchingInfo) (PATNATPoolsList, *bambou.Error)
}

// PATNATPoolsParent is the interface that a parent of a PATNATPool must implement.
// A Parent is defined as an entity that has PATNATPool as a child.
// A Parent is an Ancestor which can create a PATNATPool.
type PATNATPoolsParent interface {
	PATNATPoolsAncestor
	CreatePATNATPool(*PATNATPool) *bambou.Error
}

// PATNATPool represents the model of a patnatpool
type PATNATPool struct {
	ID                    string `json:"ID,omitempty"`
	ParentID              string `json:"parentID,omitempty"`
	ParentType            string `json:"parentType,omitempty"`
	Owner                 string `json:"owner,omitempty"`
	Name                  string `json:"name,omitempty"`
	LastUpdatedBy         string `json:"lastUpdatedBy,omitempty"`
	AddressRange          string `json:"addressRange,omitempty"`
	DefaultPATIP          string `json:"defaultPATIP,omitempty"`
	PermittedAction       string `json:"permittedAction,omitempty"`
	Description           string `json:"description,omitempty"`
	EndAddressRange       string `json:"endAddressRange,omitempty"`
	EndSourceAddress      string `json:"endSourceAddress,omitempty"`
	EntityScope           string `json:"entityScope,omitempty"`
	TranslationTimeout    int    `json:"translationTimeout,omitempty"`
	AssociatedGatewayId   string `json:"associatedGatewayId,omitempty"`
	AssociatedGatewayType string `json:"associatedGatewayType,omitempty"`
	StartAddressRange     string `json:"startAddressRange,omitempty"`
	StartSourceAddress    string `json:"startSourceAddress,omitempty"`
	ExternalID            string `json:"externalID,omitempty"`
	DynamicSourceEnabled  bool   `json:"dynamicSourceEnabled"`
}

// NewPATNATPool returns a new *PATNATPool
func NewPATNATPool() *PATNATPool {

	return &PATNATPool{}
}

// Identity returns the Identity of the object.
func (o *PATNATPool) Identity() bambou.Identity {

	return PATNATPoolIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PATNATPool) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PATNATPool) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PATNATPool from the server
func (o *PATNATPool) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PATNATPool into the server
func (o *PATNATPool) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PATNATPool from the server
func (o *PATNATPool) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// NATMapEntries retrieves the list of child NATMapEntries of the PATNATPool
func (o *PATNATPool) NATMapEntries(info *bambou.FetchingInfo) (NATMapEntriesList, *bambou.Error) {

	var list NATMapEntriesList
	err := bambou.CurrentSession().FetchChildren(o, NATMapEntryIdentity, &list, info)
	return list, err
}

// CreateNATMapEntry creates a new child NATMapEntry under the PATNATPool
func (o *PATNATPool) CreateNATMapEntry(child *NATMapEntry) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// AddressMaps retrieves the list of child AddressMaps of the PATNATPool
func (o *PATNATPool) AddressMaps(info *bambou.FetchingInfo) (AddressMapsList, *bambou.Error) {

	var list AddressMapsList
	err := bambou.CurrentSession().FetchChildren(o, AddressMapIdentity, &list, info)
	return list, err
}

// CreateAddressMap creates a new child AddressMap under the PATNATPool
func (o *PATNATPool) CreateAddressMap(child *AddressMap) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the PATNATPool
func (o *PATNATPool) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the PATNATPool
func (o *PATNATPool) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the PATNATPool
func (o *PATNATPool) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the PATNATPool
func (o *PATNATPool) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EnterprisePermissions retrieves the list of child EnterprisePermissions of the PATNATPool
func (o *PATNATPool) EnterprisePermissions(info *bambou.FetchingInfo) (EnterprisePermissionsList, *bambou.Error) {

	var list EnterprisePermissionsList
	err := bambou.CurrentSession().FetchChildren(o, EnterprisePermissionIdentity, &list, info)
	return list, err
}

// CreateEnterprisePermission creates a new child EnterprisePermission under the PATNATPool
func (o *PATNATPool) CreateEnterprisePermission(child *EnterprisePermission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
