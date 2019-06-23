/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VCenterHypervisorIdentity represents the Identity of the object
var VCenterHypervisorIdentity = bambou.Identity{
	Name:     "vcenterhypervisor",
	Category: "vcenterhypervisors",
}

// VCenterHypervisorsList represents a list of VCenterHypervisors
type VCenterHypervisorsList []*VCenterHypervisor

// VCenterHypervisorsAncestor is the interface that an ancestor of a VCenterHypervisor must implement.
// An Ancestor is defined as an entity that has VCenterHypervisor as a descendant.
// An Ancestor can get a list of its child VCenterHypervisors, but not necessarily create one.
type VCenterHypervisorsAncestor interface {
	VCenterHypervisors(*bambou.FetchingInfo) (VCenterHypervisorsList, *bambou.Error)
}

// VCenterHypervisorsParent is the interface that a parent of a VCenterHypervisor must implement.
// A Parent is defined as an entity that has VCenterHypervisor as a child.
// A Parent is an Ancestor which can create a VCenterHypervisor.
type VCenterHypervisorsParent interface {
	VCenterHypervisorsAncestor
	CreateVCenterHypervisor(*VCenterHypervisor) *bambou.Error
}

// VCenterHypervisor represents the model of a vcenterhypervisor
type VCenterHypervisor struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewVCenterHypervisor returns a new *VCenterHypervisor
func NewVCenterHypervisor() *VCenterHypervisor {

	return &VCenterHypervisor{}
}

// Identity returns the Identity of the object.
func (o *VCenterHypervisor) Identity() bambou.Identity {

	return VCenterHypervisorIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VCenterHypervisor) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VCenterHypervisor) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VCenterHypervisor from the server
func (o *VCenterHypervisor) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VCenterHypervisor into the server
func (o *VCenterHypervisor) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VCenterHypervisor from the server
func (o *VCenterHypervisor) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the VCenterHypervisor
func (o *VCenterHypervisor) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VCenterHypervisor
func (o *VCenterHypervisor) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VCenterHypervisor
func (o *VCenterHypervisor) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VCenterHypervisor
func (o *VCenterHypervisor) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// CreateJob creates a new child Job under the VCenterHypervisor
func (o *VCenterHypervisor) CreateJob(child *Job) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VRSAddressRanges retrieves the list of child VRSAddressRanges of the VCenterHypervisor
func (o *VCenterHypervisor) VRSAddressRanges(info *bambou.FetchingInfo) (VRSAddressRangesList, *bambou.Error) {

	var list VRSAddressRangesList
	err := bambou.CurrentSession().FetchChildren(o, VRSAddressRangeIdentity, &list, info)
	return list, err
}

// CreateVRSAddressRange creates a new child VRSAddressRange under the VCenterHypervisor
func (o *VCenterHypervisor) CreateVRSAddressRange(child *VRSAddressRange) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
