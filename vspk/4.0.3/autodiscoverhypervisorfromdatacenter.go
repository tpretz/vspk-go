/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// AutoDiscoverHypervisorFromDatacenterIdentity represents the Identity of the object
var AutoDiscoverHypervisorFromDatacenterIdentity = bambou.Identity{
	Name:     "autodiscoveredcomputeresource",
	Category: "autodiscoveredcomputeresources",
}

// AutoDiscoverHypervisorFromDatacentersList represents a list of AutoDiscoverHypervisorFromDatacenters
type AutoDiscoverHypervisorFromDatacentersList []*AutoDiscoverHypervisorFromDatacenter

// AutoDiscoverHypervisorFromDatacentersAncestor is the interface that an ancestor of a AutoDiscoverHypervisorFromDatacenter must implement.
// An Ancestor is defined as an entity that has AutoDiscoverHypervisorFromDatacenter as a descendant.
// An Ancestor can get a list of its child AutoDiscoverHypervisorFromDatacenters, but not necessarily create one.
type AutoDiscoverHypervisorFromDatacentersAncestor interface {
	AutoDiscoverHypervisorFromDatacenters(*bambou.FetchingInfo) (AutoDiscoverHypervisorFromDatacentersList, *bambou.Error)
}

// AutoDiscoverHypervisorFromDatacentersParent is the interface that a parent of a AutoDiscoverHypervisorFromDatacenter must implement.
// A Parent is defined as an entity that has AutoDiscoverHypervisorFromDatacenter as a child.
// A Parent is an Ancestor which can create a AutoDiscoverHypervisorFromDatacenter.
type AutoDiscoverHypervisorFromDatacentersParent interface {
	AutoDiscoverHypervisorFromDatacentersAncestor
	CreateAutoDiscoverHypervisorFromDatacenter(*AutoDiscoverHypervisorFromDatacenter) *bambou.Error
}

// AutoDiscoverHypervisorFromDatacenter represents the model of a autodiscoveredcomputeresource
type AutoDiscoverHypervisorFromDatacenter struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewAutoDiscoverHypervisorFromDatacenter returns a new *AutoDiscoverHypervisorFromDatacenter
func NewAutoDiscoverHypervisorFromDatacenter() *AutoDiscoverHypervisorFromDatacenter {

	return &AutoDiscoverHypervisorFromDatacenter{}
}

// Identity returns the Identity of the object.
func (o *AutoDiscoverHypervisorFromDatacenter) Identity() bambou.Identity {

	return AutoDiscoverHypervisorFromDatacenterIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AutoDiscoverHypervisorFromDatacenter) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AutoDiscoverHypervisorFromDatacenter) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the AutoDiscoverHypervisorFromDatacenter from the server
func (o *AutoDiscoverHypervisorFromDatacenter) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the AutoDiscoverHypervisorFromDatacenter into the server
func (o *AutoDiscoverHypervisorFromDatacenter) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the AutoDiscoverHypervisorFromDatacenter from the server
func (o *AutoDiscoverHypervisorFromDatacenter) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
