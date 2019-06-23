/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VCenterIdentity represents the Identity of the object
var VCenterIdentity = bambou.Identity{
	Name:     "vcenter",
	Category: "vcenters",
}

// VCentersList represents a list of VCenters
type VCentersList []*VCenter

// VCentersAncestor is the interface that an ancestor of a VCenter must implement.
// An Ancestor is defined as an entity that has VCenter as a descendant.
// An Ancestor can get a list of its child VCenters, but not necessarily create one.
type VCentersAncestor interface {
	VCenters(*bambou.FetchingInfo) (VCentersList, *bambou.Error)
}

// VCentersParent is the interface that a parent of a VCenter must implement.
// A Parent is defined as an entity that has VCenter as a child.
// A Parent is an Ancestor which can create a VCenter.
type VCentersParent interface {
	VCentersAncestor
	CreateVCenter(*VCenter) *bambou.Error
}

// VCenter represents the model of a vcenter
type VCenter struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewVCenter returns a new *VCenter
func NewVCenter() *VCenter {

	return &VCenter{}
}

// Identity returns the Identity of the object.
func (o *VCenter) Identity() bambou.Identity {

	return VCenterIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VCenter) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VCenter) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VCenter from the server
func (o *VCenter) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VCenter into the server
func (o *VCenter) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VCenter from the server
func (o *VCenter) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// VCenterDataCenters retrieves the list of child VCenterDataCenters of the VCenter
func (o *VCenter) VCenterDataCenters(info *bambou.FetchingInfo) (VCenterDataCentersList, *bambou.Error) {

	var list VCenterDataCentersList
	err := bambou.CurrentSession().FetchChildren(o, VCenterDataCenterIdentity, &list, info)
	return list, err
}

// CreateVCenterDataCenter creates a new child VCenterDataCenter under the VCenter
func (o *VCenter) CreateVCenterDataCenter(child *VCenterDataCenter) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the VCenter
func (o *VCenter) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VCenter
func (o *VCenter) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VCenter
func (o *VCenter) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VCenter
func (o *VCenter) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VRSAddressRanges retrieves the list of child VRSAddressRanges of the VCenter
func (o *VCenter) VRSAddressRanges(info *bambou.FetchingInfo) (VRSAddressRangesList, *bambou.Error) {

	var list VRSAddressRangesList
	err := bambou.CurrentSession().FetchChildren(o, VRSAddressRangeIdentity, &list, info)
	return list, err
}

// CreateVRSAddressRange creates a new child VRSAddressRange under the VCenter
func (o *VCenter) CreateVRSAddressRange(child *VRSAddressRange) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
