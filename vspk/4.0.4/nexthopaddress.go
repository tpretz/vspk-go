/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// NextHopAddressIdentity represents the Identity of the object
var NextHopAddressIdentity = bambou.Identity{
	Name:     "nexthop",
	Category: "nexthops",
}

// NextHopAddressList represents a list of NextHopAddress
type NextHopAddressList []*NextHopAddress

// NextHopAddressAncestor is the interface that an ancestor of a NextHopAddress must implement.
// An Ancestor is defined as an entity that has NextHopAddress as a descendant.
// An Ancestor can get a list of its child NextHopAddress, but not necessarily create one.
type NextHopAddressAncestor interface {
	NextHopAddress(*bambou.FetchingInfo) (NextHopAddressList, *bambou.Error)
}

// NextHopAddressParent is the interface that a parent of a NextHopAddress must implement.
// A Parent is defined as an entity that has NextHopAddress as a child.
// A Parent is an Ancestor which can create a NextHopAddress.
type NextHopAddressParent interface {
	NextHopAddressAncestor
	CreateNextHopAddress(*NextHopAddress) *bambou.Error
}

// NextHopAddress represents the model of a nexthop
type NextHopAddress struct {
	ID                 string `json:"ID,omitempty"`
	ParentID           string `json:"parentID,omitempty"`
	ParentType         string `json:"parentType,omitempty"`
	Owner              string `json:"owner,omitempty"`
	LastUpdatedBy      string `json:"lastUpdatedBy,omitempty"`
	EntityScope        string `json:"entityScope,omitempty"`
	RouteDistinguisher string `json:"routeDistinguisher,omitempty"`
	Ip                 string `json:"ip,omitempty"`
	ExternalID         string `json:"externalID,omitempty"`
}

// NewNextHopAddress returns a new *NextHopAddress
func NewNextHopAddress() *NextHopAddress {

	return &NextHopAddress{}
}

// Identity returns the Identity of the object.
func (o *NextHopAddress) Identity() bambou.Identity {

	return NextHopAddressIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NextHopAddress) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NextHopAddress) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NextHopAddress from the server
func (o *NextHopAddress) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NextHopAddress into the server
func (o *NextHopAddress) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NextHopAddress from the server
func (o *NextHopAddress) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the NextHopAddress
func (o *NextHopAddress) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the NextHopAddress
func (o *NextHopAddress) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the NextHopAddress
func (o *NextHopAddress) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the NextHopAddress
func (o *NextHopAddress) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
