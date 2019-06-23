/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// ViaIdentity represents the Identity of the object
var ViaIdentity = bambou.Identity{
	Name:     "via",
	Category: "vias",
}

// ViasList represents a list of Vias
type ViasList []*Via

// ViasAncestor is the interface that an ancestor of a Via must implement.
// An Ancestor is defined as an entity that has Via as a descendant.
// An Ancestor can get a list of its child Vias, but not necessarily create one.
type ViasAncestor interface {
	Vias(*bambou.FetchingInfo) (ViasList, *bambou.Error)
}

// ViasParent is the interface that a parent of a Via must implement.
// A Parent is defined as an entity that has Via as a child.
// A Parent is an Ancestor which can create a Via.
type ViasParent interface {
	ViasAncestor
	CreateVia(*Via) *bambou.Error
}

// Via represents the model of a via
type Via struct {
	ID         string        `json:"ID,omitempty"`
	ParentID   string        `json:"parentID,omitempty"`
	ParentType string        `json:"parentType,omitempty"`
	Owner      string        `json:"owner,omitempty"`
	NextHops   []interface{} `json:"nextHops,omitempty"`
}

// NewVia returns a new *Via
func NewVia() *Via {

	return &Via{}
}

// Identity returns the Identity of the object.
func (o *Via) Identity() bambou.Identity {

	return ViaIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Via) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Via) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Via from the server
func (o *Via) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Via into the server
func (o *Via) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Via from the server
func (o *Via) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
