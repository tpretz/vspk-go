/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// NetworkMacroGroupIdentity represents the Identity of the object
var NetworkMacroGroupIdentity = bambou.Identity{
	Name:     "networkmacrogroup",
	Category: "networkmacrogroups",
}

// NetworkMacroGroupsList represents a list of NetworkMacroGroups
type NetworkMacroGroupsList []*NetworkMacroGroup

// NetworkMacroGroupsAncestor is the interface that an ancestor of a NetworkMacroGroup must implement.
// An Ancestor is defined as an entity that has NetworkMacroGroup as a descendant.
// An Ancestor can get a list of its child NetworkMacroGroups, but not necessarily create one.
type NetworkMacroGroupsAncestor interface {
	NetworkMacroGroups(*bambou.FetchingInfo) (NetworkMacroGroupsList, *bambou.Error)
}

// NetworkMacroGroupsParent is the interface that a parent of a NetworkMacroGroup must implement.
// A Parent is defined as an entity that has NetworkMacroGroup as a child.
// A Parent is an Ancestor which can create a NetworkMacroGroup.
type NetworkMacroGroupsParent interface {
	NetworkMacroGroupsAncestor
	CreateNetworkMacroGroup(*NetworkMacroGroup) *bambou.Error
}

// NetworkMacroGroup represents the model of a networkmacrogroup
type NetworkMacroGroup struct {
	ID            string        `json:"ID,omitempty"`
	ParentID      string        `json:"parentID,omitempty"`
	ParentType    string        `json:"parentType,omitempty"`
	Owner         string        `json:"owner,omitempty"`
	Name          string        `json:"name,omitempty"`
	LastUpdatedBy string        `json:"lastUpdatedBy,omitempty"`
	Description   string        `json:"description,omitempty"`
	NetworkMacros []interface{} `json:"networkMacros,omitempty"`
	EntityScope   string        `json:"entityScope,omitempty"`
	ExternalID    string        `json:"externalID,omitempty"`
}

// NewNetworkMacroGroup returns a new *NetworkMacroGroup
func NewNetworkMacroGroup() *NetworkMacroGroup {

	return &NetworkMacroGroup{}
}

// Identity returns the Identity of the object.
func (o *NetworkMacroGroup) Identity() bambou.Identity {

	return NetworkMacroGroupIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NetworkMacroGroup) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NetworkMacroGroup) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NetworkMacroGroup from the server
func (o *NetworkMacroGroup) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NetworkMacroGroup into the server
func (o *NetworkMacroGroup) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NetworkMacroGroup from the server
func (o *NetworkMacroGroup) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the NetworkMacroGroup
func (o *NetworkMacroGroup) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the NetworkMacroGroup
func (o *NetworkMacroGroup) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the NetworkMacroGroup
func (o *NetworkMacroGroup) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the NetworkMacroGroup
func (o *NetworkMacroGroup) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EnterpriseNetworks retrieves the list of child EnterpriseNetworks of the NetworkMacroGroup
func (o *NetworkMacroGroup) EnterpriseNetworks(info *bambou.FetchingInfo) (EnterpriseNetworksList, *bambou.Error) {

	var list EnterpriseNetworksList
	err := bambou.CurrentSession().FetchChildren(o, EnterpriseNetworkIdentity, &list, info)
	return list, err
}
