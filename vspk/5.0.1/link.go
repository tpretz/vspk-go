/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// LinkIdentity represents the Identity of the object
var LinkIdentity = bambou.Identity{
	Name:     "link",
	Category: "links",
}

// LinksList represents a list of Links
type LinksList []*Link

// LinksAncestor is the interface that an ancestor of a Link must implement.
// An Ancestor is defined as an entity that has Link as a descendant.
// An Ancestor can get a list of its child Links, but not necessarily create one.
type LinksAncestor interface {
	Links(*bambou.FetchingInfo) (LinksList, *bambou.Error)
}

// LinksParent is the interface that a parent of a Link must implement.
// A Parent is defined as an entity that has Link as a child.
// A Parent is an Ancestor which can create a Link.
type LinksParent interface {
	LinksAncestor
	CreateLink(*Link) *bambou.Error
}

// Link represents the model of a link
type Link struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewLink returns a new *Link
func NewLink() *Link {

	return &Link{}
}

// Identity returns the Identity of the object.
func (o *Link) Identity() bambou.Identity {

	return LinkIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Link) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Link) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Link from the server
func (o *Link) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Link into the server
func (o *Link) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Link from the server
func (o *Link) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// DemarcationServices retrieves the list of child DemarcationServices of the Link
func (o *Link) DemarcationServices(info *bambou.FetchingInfo) (DemarcationServicesList, *bambou.Error) {

	var list DemarcationServicesList
	err := bambou.CurrentSession().FetchChildren(o, DemarcationServiceIdentity, &list, info)
	return list, err
}

// CreateDemarcationService creates a new child DemarcationService under the Link
func (o *Link) CreateDemarcationService(child *DemarcationService) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the Link
func (o *Link) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Link
func (o *Link) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// NextHopAddress retrieves the list of child NextHopAddress of the Link
func (o *Link) NextHopAddress(info *bambou.FetchingInfo) (NextHopAddressList, *bambou.Error) {

	var list NextHopAddressList
	err := bambou.CurrentSession().FetchChildren(o, NextHopAddressIdentity, &list, info)
	return list, err
}

// CreateNextHopAddress creates a new child NextHopAddress under the Link
func (o *Link) CreateNextHopAddress(child *NextHopAddress) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Link
func (o *Link) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Link
func (o *Link) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// CSNATPools retrieves the list of child CSNATPools of the Link
func (o *Link) CSNATPools(info *bambou.FetchingInfo) (CSNATPoolsList, *bambou.Error) {

	var list CSNATPoolsList
	err := bambou.CurrentSession().FetchChildren(o, CSNATPoolIdentity, &list, info)
	return list, err
}

// CreateCSNATPool creates a new child CSNATPool under the Link
func (o *Link) CreateCSNATPool(child *CSNATPool) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// PSNATPools retrieves the list of child PSNATPools of the Link
func (o *Link) PSNATPools(info *bambou.FetchingInfo) (PSNATPoolsList, *bambou.Error) {

	var list PSNATPoolsList
	err := bambou.CurrentSession().FetchChildren(o, PSNATPoolIdentity, &list, info)
	return list, err
}

// CreatePSNATPool creates a new child PSNATPool under the Link
func (o *Link) CreatePSNATPool(child *PSNATPool) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// OverlayAddressPools retrieves the list of child OverlayAddressPools of the Link
func (o *Link) OverlayAddressPools(info *bambou.FetchingInfo) (OverlayAddressPoolsList, *bambou.Error) {

	var list OverlayAddressPoolsList
	err := bambou.CurrentSession().FetchChildren(o, OverlayAddressPoolIdentity, &list, info)
	return list, err
}

// CreateOverlayAddressPool creates a new child OverlayAddressPool under the Link
func (o *Link) CreateOverlayAddressPool(child *OverlayAddressPool) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
