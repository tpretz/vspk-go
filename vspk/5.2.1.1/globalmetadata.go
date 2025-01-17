/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// GlobalMetadataIdentity represents the Identity of the object
var GlobalMetadataIdentity = bambou.Identity{
	Name:     "globalmetadata",
	Category: "globalmetadatas",
}

// GlobalMetadatasList represents a list of GlobalMetadatas
type GlobalMetadatasList []*GlobalMetadata

// GlobalMetadatasAncestor is the interface that an ancestor of a GlobalMetadata must implement.
// An Ancestor is defined as an entity that has GlobalMetadata as a descendant.
// An Ancestor can get a list of its child GlobalMetadatas, but not necessarily create one.
type GlobalMetadatasAncestor interface {
	GlobalMetadatas(*bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error)
}

// GlobalMetadatasParent is the interface that a parent of a GlobalMetadata must implement.
// A Parent is defined as an entity that has GlobalMetadata as a child.
// A Parent is an Ancestor which can create a GlobalMetadata.
type GlobalMetadatasParent interface {
	GlobalMetadatasAncestor
	CreateGlobalMetadata(*GlobalMetadata) *bambou.Error
}

// GlobalMetadata represents the model of a globalmetadata
type GlobalMetadata struct {
	ID                          string        `json:"ID,omitempty"`
	ParentID                    string        `json:"parentID,omitempty"`
	ParentType                  string        `json:"parentType,omitempty"`
	Owner                       string        `json:"owner,omitempty"`
	Name                        string        `json:"name,omitempty"`
	LastUpdatedBy               string        `json:"lastUpdatedBy,omitempty"`
	Description                 string        `json:"description,omitempty"`
	MetadataTagIDs              []interface{} `json:"metadataTagIDs,omitempty"`
	NetworkNotificationDisabled bool          `json:"networkNotificationDisabled"`
	Blob                        string        `json:"blob,omitempty"`
	Global                      bool          `json:"global"`
	EntityScope                 string        `json:"entityScope,omitempty"`
	ExternalID                  string        `json:"externalID,omitempty"`
}

// NewGlobalMetadata returns a new *GlobalMetadata
func NewGlobalMetadata() *GlobalMetadata {

	return &GlobalMetadata{}
}

// Identity returns the Identity of the object.
func (o *GlobalMetadata) Identity() bambou.Identity {

	return GlobalMetadataIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *GlobalMetadata) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *GlobalMetadata) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the GlobalMetadata from the server
func (o *GlobalMetadata) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the GlobalMetadata into the server
func (o *GlobalMetadata) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the GlobalMetadata from the server
func (o *GlobalMetadata) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the GlobalMetadata
func (o *GlobalMetadata) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the GlobalMetadata
func (o *GlobalMetadata) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the GlobalMetadata
func (o *GlobalMetadata) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the GlobalMetadata
func (o *GlobalMetadata) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
