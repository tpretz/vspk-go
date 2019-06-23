/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// AllGatewayIdentity represents the Identity of the object
var AllGatewayIdentity = bambou.Identity{
	Name:     "allgateway",
	Category: "allgateways",
}

// AllGatewaysList represents a list of AllGateways
type AllGatewaysList []*AllGateway

// AllGatewaysAncestor is the interface that an ancestor of a AllGateway must implement.
// An Ancestor is defined as an entity that has AllGateway as a descendant.
// An Ancestor can get a list of its child AllGateways, but not necessarily create one.
type AllGatewaysAncestor interface {
	AllGateways(*bambou.FetchingInfo) (AllGatewaysList, *bambou.Error)
}

// AllGatewaysParent is the interface that a parent of a AllGateway must implement.
// A Parent is defined as an entity that has AllGateway as a child.
// A Parent is an Ancestor which can create a AllGateway.
type AllGatewaysParent interface {
	AllGatewaysAncestor
	CreateAllGateway(*AllGateway) *bambou.Error
}

// AllGateway represents the model of a allgateway
type AllGateway struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewAllGateway returns a new *AllGateway
func NewAllGateway() *AllGateway {

	return &AllGateway{}
}

// Identity returns the Identity of the object.
func (o *AllGateway) Identity() bambou.Identity {

	return AllGatewayIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AllGateway) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AllGateway) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the AllGateway from the server
func (o *AllGateway) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the AllGateway into the server
func (o *AllGateway) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the AllGateway from the server
func (o *AllGateway) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the AllGateway
func (o *AllGateway) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the AllGateway
func (o *AllGateway) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the AllGateway
func (o *AllGateway) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the AllGateway
func (o *AllGateway) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
