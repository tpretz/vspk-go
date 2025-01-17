/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// EnterpriseNetworkIdentity represents the Identity of the object
var EnterpriseNetworkIdentity = bambou.Identity{
	Name:     "enterprisenetwork",
	Category: "enterprisenetworks",
}

// EnterpriseNetworksList represents a list of EnterpriseNetworks
type EnterpriseNetworksList []*EnterpriseNetwork

// EnterpriseNetworksAncestor is the interface that an ancestor of a EnterpriseNetwork must implement.
// An Ancestor is defined as an entity that has EnterpriseNetwork as a descendant.
// An Ancestor can get a list of its child EnterpriseNetworks, but not necessarily create one.
type EnterpriseNetworksAncestor interface {
	EnterpriseNetworks(*bambou.FetchingInfo) (EnterpriseNetworksList, *bambou.Error)
}

// EnterpriseNetworksParent is the interface that a parent of a EnterpriseNetwork must implement.
// A Parent is defined as an entity that has EnterpriseNetwork as a child.
// A Parent is an Ancestor which can create a EnterpriseNetwork.
type EnterpriseNetworksParent interface {
	EnterpriseNetworksAncestor
	CreateEnterpriseNetwork(*EnterpriseNetwork) *bambou.Error
}

// EnterpriseNetwork represents the model of a enterprisenetwork
type EnterpriseNetwork struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	IPType        string `json:"IPType,omitempty"`
	Name          string `json:"name,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	Address       string `json:"address,omitempty"`
	Netmask       string `json:"netmask,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewEnterpriseNetwork returns a new *EnterpriseNetwork
func NewEnterpriseNetwork() *EnterpriseNetwork {

	return &EnterpriseNetwork{}
}

// Identity returns the Identity of the object.
func (o *EnterpriseNetwork) Identity() bambou.Identity {

	return EnterpriseNetworkIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EnterpriseNetwork) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EnterpriseNetwork) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the EnterpriseNetwork from the server
func (o *EnterpriseNetwork) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the EnterpriseNetwork into the server
func (o *EnterpriseNetwork) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the EnterpriseNetwork from the server
func (o *EnterpriseNetwork) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the EnterpriseNetwork
func (o *EnterpriseNetwork) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the EnterpriseNetwork
func (o *EnterpriseNetwork) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// NetworkMacroGroups retrieves the list of child NetworkMacroGroups of the EnterpriseNetwork
func (o *EnterpriseNetwork) NetworkMacroGroups(info *bambou.FetchingInfo) (NetworkMacroGroupsList, *bambou.Error) {

	var list NetworkMacroGroupsList
	err := bambou.CurrentSession().FetchChildren(o, NetworkMacroGroupIdentity, &list, info)
	return list, err
}

// AssignNetworkMacroGroups assigns the list of NetworkMacroGroups to the EnterpriseNetwork
func (o *EnterpriseNetwork) AssignNetworkMacroGroups(children NetworkMacroGroupsList) *bambou.Error {

	list := []bambou.Identifiable{}
	for _, c := range children {
		list = append(list, c)
	}

	return bambou.CurrentSession().AssignChildren(o, list, NetworkMacroGroupIdentity)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the EnterpriseNetwork
func (o *EnterpriseNetwork) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the EnterpriseNetwork
func (o *EnterpriseNetwork) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the EnterpriseNetwork
func (o *EnterpriseNetwork) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
