/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// ApplicationperformancemanagementbindingIdentity represents the Identity of the object
var ApplicationperformancemanagementbindingIdentity = bambou.Identity{
	Name:     "applicationperformancemanagementbinding",
	Category: "applicationperformancemanagementbindings",
}

// ApplicationperformancemanagementbindingsList represents a list of Applicationperformancemanagementbindings
type ApplicationperformancemanagementbindingsList []*Applicationperformancemanagementbinding

// ApplicationperformancemanagementbindingsAncestor is the interface that an ancestor of a Applicationperformancemanagementbinding must implement.
// An Ancestor is defined as an entity that has Applicationperformancemanagementbinding as a descendant.
// An Ancestor can get a list of its child Applicationperformancemanagementbindings, but not necessarily create one.
type ApplicationperformancemanagementbindingsAncestor interface {
	Applicationperformancemanagementbindings(*bambou.FetchingInfo) (ApplicationperformancemanagementbindingsList, *bambou.Error)
}

// ApplicationperformancemanagementbindingsParent is the interface that a parent of a Applicationperformancemanagementbinding must implement.
// A Parent is defined as an entity that has Applicationperformancemanagementbinding as a child.
// A Parent is an Ancestor which can create a Applicationperformancemanagementbinding.
type ApplicationperformancemanagementbindingsParent interface {
	ApplicationperformancemanagementbindingsAncestor
	CreateApplicationperformancemanagementbinding(*Applicationperformancemanagementbinding) *bambou.Error
}

// Applicationperformancemanagementbinding represents the model of a applicationperformancemanagementbinding
type Applicationperformancemanagementbinding struct {
	ID                                           string `json:"ID,omitempty"`
	ParentID                                     string `json:"parentID,omitempty"`
	ParentType                                   string `json:"parentType,omitempty"`
	Owner                                        string `json:"owner,omitempty"`
	ReadOnly                                     bool   `json:"readOnly"`
	Priority                                     int    `json:"priority,omitempty"`
	AssociatedApplicationPerformanceManagementID string `json:"associatedApplicationPerformanceManagementID,omitempty"`
}

// NewApplicationperformancemanagementbinding returns a new *Applicationperformancemanagementbinding
func NewApplicationperformancemanagementbinding() *Applicationperformancemanagementbinding {

	return &Applicationperformancemanagementbinding{
		ReadOnly: false,
	}
}

// Identity returns the Identity of the object.
func (o *Applicationperformancemanagementbinding) Identity() bambou.Identity {

	return ApplicationperformancemanagementbindingIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Applicationperformancemanagementbinding) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Applicationperformancemanagementbinding) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Applicationperformancemanagementbinding from the server
func (o *Applicationperformancemanagementbinding) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Applicationperformancemanagementbinding into the server
func (o *Applicationperformancemanagementbinding) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Applicationperformancemanagementbinding from the server
func (o *Applicationperformancemanagementbinding) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
