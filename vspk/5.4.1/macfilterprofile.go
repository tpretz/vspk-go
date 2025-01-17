/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// MACFilterProfileIdentity represents the Identity of the object
var MACFilterProfileIdentity = bambou.Identity{
	Name:     "macfilterprofile",
	Category: "macfilterprofiles",
}

// MACFilterProfilesList represents a list of MACFilterProfiles
type MACFilterProfilesList []*MACFilterProfile

// MACFilterProfilesAncestor is the interface that an ancestor of a MACFilterProfile must implement.
// An Ancestor is defined as an entity that has MACFilterProfile as a descendant.
// An Ancestor can get a list of its child MACFilterProfiles, but not necessarily create one.
type MACFilterProfilesAncestor interface {
	MACFilterProfiles(*bambou.FetchingInfo) (MACFilterProfilesList, *bambou.Error)
}

// MACFilterProfilesParent is the interface that a parent of a MACFilterProfile must implement.
// A Parent is defined as an entity that has MACFilterProfile as a child.
// A Parent is an Ancestor which can create a MACFilterProfile.
type MACFilterProfilesParent interface {
	MACFilterProfilesAncestor
	CreateMACFilterProfile(*MACFilterProfile) *bambou.Error
}

// MACFilterProfile represents the model of a macfilterprofile
type MACFilterProfile struct {
	ID              string `json:"ID,omitempty"`
	ParentID        string `json:"parentID,omitempty"`
	ParentType      string `json:"parentType,omitempty"`
	Owner           string `json:"owner,omitempty"`
	Name            string `json:"name,omitempty"`
	LastUpdatedBy   string `json:"lastUpdatedBy,omitempty"`
	Description     string `json:"description,omitempty"`
	EntityScope     string `json:"entityScope,omitempty"`
	AssocEntityType string `json:"assocEntityType,omitempty"`
	ExternalID      string `json:"externalID,omitempty"`
}

// NewMACFilterProfile returns a new *MACFilterProfile
func NewMACFilterProfile() *MACFilterProfile {

	return &MACFilterProfile{}
}

// Identity returns the Identity of the object.
func (o *MACFilterProfile) Identity() bambou.Identity {

	return MACFilterProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *MACFilterProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *MACFilterProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the MACFilterProfile from the server
func (o *MACFilterProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the MACFilterProfile into the server
func (o *MACFilterProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the MACFilterProfile from the server
func (o *MACFilterProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the MACFilterProfile
func (o *MACFilterProfile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the MACFilterProfile
func (o *MACFilterProfile) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the MACFilterProfile
func (o *MACFilterProfile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the MACFilterProfile
func (o *MACFilterProfile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
