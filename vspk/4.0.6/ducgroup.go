/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// DUCGroupIdentity represents the Identity of the object
var DUCGroupIdentity = bambou.Identity{
	Name:     "ducgroup",
	Category: "ducgroups",
}

// DUCGroupsList represents a list of DUCGroups
type DUCGroupsList []*DUCGroup

// DUCGroupsAncestor is the interface that an ancestor of a DUCGroup must implement.
// An Ancestor is defined as an entity that has DUCGroup as a descendant.
// An Ancestor can get a list of its child DUCGroups, but not necessarily create one.
type DUCGroupsAncestor interface {
	DUCGroups(*bambou.FetchingInfo) (DUCGroupsList, *bambou.Error)
}

// DUCGroupsParent is the interface that a parent of a DUCGroup must implement.
// A Parent is defined as an entity that has DUCGroup as a child.
// A Parent is an Ancestor which can create a DUCGroup.
type DUCGroupsParent interface {
	DUCGroupsAncestor
	CreateDUCGroup(*DUCGroup) *bambou.Error
}

// DUCGroup represents the model of a ducgroup
type DUCGroup struct {
	ID                             string        `json:"ID,omitempty"`
	ParentID                       string        `json:"parentID,omitempty"`
	ParentType                     string        `json:"parentType,omitempty"`
	Owner                          string        `json:"owner,omitempty"`
	Name                           string        `json:"name,omitempty"`
	Id                             string        `json:"id,omitempty"`
	Description                    string        `json:"Description ,omitempty"`
	AssociatedDUCs                 []interface{} `json:"associatedDUCs,omitempty"`
	AssociatedPerformanceMonitorID string        `json:"associatedPerformanceMonitorID,omitempty"`
}

// NewDUCGroup returns a new *DUCGroup
func NewDUCGroup() *DUCGroup {

	return &DUCGroup{}
}

// Identity returns the Identity of the object.
func (o *DUCGroup) Identity() bambou.Identity {

	return DUCGroupIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DUCGroup) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DUCGroup) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the DUCGroup from the server
func (o *DUCGroup) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the DUCGroup into the server
func (o *DUCGroup) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the DUCGroup from the server
func (o *DUCGroup) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// NSGateways retrieves the list of child NSGateways of the DUCGroup
func (o *DUCGroup) NSGateways(info *bambou.FetchingInfo) (NSGatewaysList, *bambou.Error) {

	var list NSGatewaysList
	err := bambou.CurrentSession().FetchChildren(o, NSGatewayIdentity, &list, info)
	return list, err
}

// AssignNSGateways assigns the list of NSGateways to the DUCGroup
func (o *DUCGroup) AssignNSGateways(children NSGatewaysList) *bambou.Error {

	list := []bambou.Identifiable{}
	for _, c := range children {
		list = append(list, c)
	}

	return bambou.CurrentSession().AssignChildren(o, list, NSGatewayIdentity)
}
