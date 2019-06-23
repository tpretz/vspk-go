/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// MetadataTagIdentity represents the Identity of the object
var MetadataTagIdentity = bambou.Identity{
	Name:     "metadatatag",
	Category: "metadatatags",
}

// MetadataTagsList represents a list of MetadataTags
type MetadataTagsList []*MetadataTag

// MetadataTagsAncestor is the interface that an ancestor of a MetadataTag must implement.
// An Ancestor is defined as an entity that has MetadataTag as a descendant.
// An Ancestor can get a list of its child MetadataTags, but not necessarily create one.
type MetadataTagsAncestor interface {
	MetadataTags(*bambou.FetchingInfo) (MetadataTagsList, *bambou.Error)
}

// MetadataTagsParent is the interface that a parent of a MetadataTag must implement.
// A Parent is defined as an entity that has MetadataTag as a child.
// A Parent is an Ancestor which can create a MetadataTag.
type MetadataTagsParent interface {
	MetadataTagsAncestor
	CreateMetadataTag(*MetadataTag) *bambou.Error
}

// MetadataTag represents the model of a metadatatag
type MetadataTag struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewMetadataTag returns a new *MetadataTag
func NewMetadataTag() *MetadataTag {

	return &MetadataTag{}
}

// Identity returns the Identity of the object.
func (o *MetadataTag) Identity() bambou.Identity {

	return MetadataTagIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *MetadataTag) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *MetadataTag) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the MetadataTag from the server
func (o *MetadataTag) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the MetadataTag into the server
func (o *MetadataTag) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the MetadataTag from the server
func (o *MetadataTag) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the MetadataTag
func (o *MetadataTag) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the MetadataTag
func (o *MetadataTag) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// EventLogs retrieves the list of child EventLogs of the MetadataTag
func (o *MetadataTag) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
