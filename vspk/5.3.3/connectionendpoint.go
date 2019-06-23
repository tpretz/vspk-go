/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// ConnectionendpointIdentity represents the Identity of the object
var ConnectionendpointIdentity = bambou.Identity{
	Name:     "connectionendpoint",
	Category: "connectionendpoints",
}

// ConnectionendpointsList represents a list of Connectionendpoints
type ConnectionendpointsList []*Connectionendpoint

// ConnectionendpointsAncestor is the interface that an ancestor of a Connectionendpoint must implement.
// An Ancestor is defined as an entity that has Connectionendpoint as a descendant.
// An Ancestor can get a list of its child Connectionendpoints, but not necessarily create one.
type ConnectionendpointsAncestor interface {
	Connectionendpoints(*bambou.FetchingInfo) (ConnectionendpointsList, *bambou.Error)
}

// ConnectionendpointsParent is the interface that a parent of a Connectionendpoint must implement.
// A Parent is defined as an entity that has Connectionendpoint as a child.
// A Parent is an Ancestor which can create a Connectionendpoint.
type ConnectionendpointsParent interface {
	ConnectionendpointsAncestor
	CreateConnectionendpoint(*Connectionendpoint) *bambou.Error
}

// Connectionendpoint represents the model of a connectionendpoint
type Connectionendpoint struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewConnectionendpoint returns a new *Connectionendpoint
func NewConnectionendpoint() *Connectionendpoint {

	return &Connectionendpoint{}
}

// Identity returns the Identity of the object.
func (o *Connectionendpoint) Identity() bambou.Identity {

	return ConnectionendpointIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Connectionendpoint) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Connectionendpoint) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Connectionendpoint from the server
func (o *Connectionendpoint) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Connectionendpoint into the server
func (o *Connectionendpoint) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Connectionendpoint from the server
func (o *Connectionendpoint) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the Connectionendpoint
func (o *Connectionendpoint) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Connectionendpoint
func (o *Connectionendpoint) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Connectionendpoint
func (o *Connectionendpoint) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Connectionendpoint
func (o *Connectionendpoint) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
