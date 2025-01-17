/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// AllRedundancyGroupIdentity represents the Identity of the object
var AllRedundancyGroupIdentity = bambou.Identity{
	Name:     "allredundancygroup",
	Category: "allredundancygroups",
}

// AllRedundancyGroupsList represents a list of AllRedundancyGroups
type AllRedundancyGroupsList []*AllRedundancyGroup

// AllRedundancyGroupsAncestor is the interface that an ancestor of a AllRedundancyGroup must implement.
// An Ancestor is defined as an entity that has AllRedundancyGroup as a descendant.
// An Ancestor can get a list of its child AllRedundancyGroups, but not necessarily create one.
type AllRedundancyGroupsAncestor interface {
	AllRedundancyGroups(*bambou.FetchingInfo) (AllRedundancyGroupsList, *bambou.Error)
}

// AllRedundancyGroupsParent is the interface that a parent of a AllRedundancyGroup must implement.
// A Parent is defined as an entity that has AllRedundancyGroup as a child.
// A Parent is an Ancestor which can create a AllRedundancyGroup.
type AllRedundancyGroupsParent interface {
	AllRedundancyGroupsAncestor
	CreateAllRedundancyGroup(*AllRedundancyGroup) *bambou.Error
}

// AllRedundancyGroup represents the model of a allredundancygroup
type AllRedundancyGroup struct {
	ID                                  string `json:"ID,omitempty"`
	ParentID                            string `json:"parentID,omitempty"`
	ParentType                          string `json:"parentType,omitempty"`
	Owner                               string `json:"owner,omitempty"`
	Name                                string `json:"name,omitempty"`
	LastUpdatedBy                       string `json:"lastUpdatedBy,omitempty"`
	GatewayPeer1AutodiscoveredGatewayID string `json:"gatewayPeer1AutodiscoveredGatewayID,omitempty"`
	GatewayPeer1Connected               bool   `json:"gatewayPeer1Connected"`
	GatewayPeer1ID                      string `json:"gatewayPeer1ID,omitempty"`
	GatewayPeer1Name                    string `json:"gatewayPeer1Name,omitempty"`
	GatewayPeer2AutodiscoveredGatewayID string `json:"gatewayPeer2AutodiscoveredGatewayID,omitempty"`
	GatewayPeer2Connected               bool   `json:"gatewayPeer2Connected"`
	GatewayPeer2Name                    string `json:"gatewayPeer2Name,omitempty"`
	RedundantGatewayStatus              string `json:"redundantGatewayStatus,omitempty"`
	PermittedAction                     string `json:"permittedAction,omitempty"`
	Personality                         string `json:"personality,omitempty"`
	Description                         string `json:"description,omitempty"`
	EnterpriseID                        string `json:"enterpriseID,omitempty"`
	EntityScope                         string `json:"entityScope,omitempty"`
	Vtep                                string `json:"vtep,omitempty"`
	ExternalID                          string `json:"externalID,omitempty"`
}

// NewAllRedundancyGroup returns a new *AllRedundancyGroup
func NewAllRedundancyGroup() *AllRedundancyGroup {

	return &AllRedundancyGroup{}
}

// Identity returns the Identity of the object.
func (o *AllRedundancyGroup) Identity() bambou.Identity {

	return AllRedundancyGroupIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AllRedundancyGroup) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AllRedundancyGroup) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the AllRedundancyGroup from the server
func (o *AllRedundancyGroup) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the AllRedundancyGroup into the server
func (o *AllRedundancyGroup) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the AllRedundancyGroup from the server
func (o *AllRedundancyGroup) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the AllRedundancyGroup
func (o *AllRedundancyGroup) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the AllRedundancyGroup
func (o *AllRedundancyGroup) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the AllRedundancyGroup
func (o *AllRedundancyGroup) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the AllRedundancyGroup
func (o *AllRedundancyGroup) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
