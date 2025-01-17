/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// MonitorscopeIdentity represents the Identity of the object
var MonitorscopeIdentity = bambou.Identity{
	Name:     "monitorscope",
	Category: "monitorscopes",
}

// MonitorscopesList represents a list of Monitorscopes
type MonitorscopesList []*Monitorscope

// MonitorscopesAncestor is the interface that an ancestor of a Monitorscope must implement.
// An Ancestor is defined as an entity that has Monitorscope as a descendant.
// An Ancestor can get a list of its child Monitorscopes, but not necessarily create one.
type MonitorscopesAncestor interface {
	Monitorscopes(*bambou.FetchingInfo) (MonitorscopesList, *bambou.Error)
}

// MonitorscopesParent is the interface that a parent of a Monitorscope must implement.
// A Parent is defined as an entity that has Monitorscope as a child.
// A Parent is an Ancestor which can create a Monitorscope.
type MonitorscopesParent interface {
	MonitorscopesAncestor
	CreateMonitorscope(*Monitorscope) *bambou.Error
}

// Monitorscope represents the model of a monitorscope
type Monitorscope struct {
	ID              string        `json:"ID,omitempty"`
	ParentID        string        `json:"parentID,omitempty"`
	ParentType      string        `json:"parentType,omitempty"`
	Owner           string        `json:"owner,omitempty"`
	Name            string        `json:"name,omitempty"`
	ReadOnly        bool          `json:"readOnly"`
	DestinationNSGs []interface{} `json:"destinationNSGs,omitempty"`
	SourceNSGs      []interface{} `json:"sourceNSGs,omitempty"`
}

// NewMonitorscope returns a new *Monitorscope
func NewMonitorscope() *Monitorscope {

	return &Monitorscope{
		ReadOnly: false,
	}
}

// Identity returns the Identity of the object.
func (o *Monitorscope) Identity() bambou.Identity {

	return MonitorscopeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Monitorscope) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Monitorscope) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Monitorscope from the server
func (o *Monitorscope) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Monitorscope into the server
func (o *Monitorscope) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Monitorscope from the server
func (o *Monitorscope) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
