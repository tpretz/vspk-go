/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// DefaultGatewayIdentity represents the Identity of the object
var DefaultGatewayIdentity = bambou.Identity{
	Name:     "defaultgateway",
	Category: "defaultgateways",
}

// DefaultGatewaysList represents a list of DefaultGateways
type DefaultGatewaysList []*DefaultGateway

// DefaultGatewaysAncestor is the interface that an ancestor of a DefaultGateway must implement.
// An Ancestor is defined as an entity that has DefaultGateway as a descendant.
// An Ancestor can get a list of its child DefaultGateways, but not necessarily create one.
type DefaultGatewaysAncestor interface {
	DefaultGateways(*bambou.FetchingInfo) (DefaultGatewaysList, *bambou.Error)
}

// DefaultGatewaysParent is the interface that a parent of a DefaultGateway must implement.
// A Parent is defined as an entity that has DefaultGateway as a child.
// A Parent is an Ancestor which can create a DefaultGateway.
type DefaultGatewaysParent interface {
	DefaultGatewaysAncestor
	CreateDefaultGateway(*DefaultGateway) *bambou.Error
}

// DefaultGateway represents the model of a defaultgateway
type DefaultGateway struct {
	ID                string `json:"ID,omitempty"`
	ParentID          string `json:"parentID,omitempty"`
	ParentType        string `json:"parentType,omitempty"`
	Owner             string `json:"owner,omitempty"`
	Name              string `json:"name,omitempty"`
	LastUpdatedBy     string `json:"lastUpdatedBy,omitempty"`
	GatewayIPAddress  string `json:"gatewayIPAddress,omitempty"`
	GatewayMACAddress string `json:"gatewayMACAddress,omitempty"`
	EntityScope       string `json:"entityScope,omitempty"`
	ExternalID        string `json:"externalID,omitempty"`
}

// NewDefaultGateway returns a new *DefaultGateway
func NewDefaultGateway() *DefaultGateway {

	return &DefaultGateway{}
}

// Identity returns the Identity of the object.
func (o *DefaultGateway) Identity() bambou.Identity {

	return DefaultGatewayIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DefaultGateway) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DefaultGateway) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the DefaultGateway from the server
func (o *DefaultGateway) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the DefaultGateway into the server
func (o *DefaultGateway) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the DefaultGateway from the server
func (o *DefaultGateway) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the DefaultGateway
func (o *DefaultGateway) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the DefaultGateway
func (o *DefaultGateway) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the DefaultGateway
func (o *DefaultGateway) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the DefaultGateway
func (o *DefaultGateway) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
