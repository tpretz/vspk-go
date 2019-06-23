/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VCenterVRSConfigIdentity represents the Identity of the object
var VCenterVRSConfigIdentity = bambou.Identity{
	Name:     "vrsconfig",
	Category: "vrsconfigs",
}

// VCenterVRSConfigsList represents a list of VCenterVRSConfigs
type VCenterVRSConfigsList []*VCenterVRSConfig

// VCenterVRSConfigsAncestor is the interface that an ancestor of a VCenterVRSConfig must implement.
// An Ancestor is defined as an entity that has VCenterVRSConfig as a descendant.
// An Ancestor can get a list of its child VCenterVRSConfigs, but not necessarily create one.
type VCenterVRSConfigsAncestor interface {
	VCenterVRSConfigs(*bambou.FetchingInfo) (VCenterVRSConfigsList, *bambou.Error)
}

// VCenterVRSConfigsParent is the interface that a parent of a VCenterVRSConfig must implement.
// A Parent is defined as an entity that has VCenterVRSConfig as a child.
// A Parent is an Ancestor which can create a VCenterVRSConfig.
type VCenterVRSConfigsParent interface {
	VCenterVRSConfigsAncestor
	CreateVCenterVRSConfig(*VCenterVRSConfig) *bambou.Error
}

// VCenterVRSConfig represents the model of a vrsconfig
type VCenterVRSConfig struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewVCenterVRSConfig returns a new *VCenterVRSConfig
func NewVCenterVRSConfig() *VCenterVRSConfig {

	return &VCenterVRSConfig{}
}

// Identity returns the Identity of the object.
func (o *VCenterVRSConfig) Identity() bambou.Identity {

	return VCenterVRSConfigIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VCenterVRSConfig) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VCenterVRSConfig) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VCenterVRSConfig from the server
func (o *VCenterVRSConfig) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VCenterVRSConfig into the server
func (o *VCenterVRSConfig) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VCenterVRSConfig from the server
func (o *VCenterVRSConfig) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the VCenterVRSConfig
func (o *VCenterVRSConfig) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VCenterVRSConfig
func (o *VCenterVRSConfig) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VCenterVRSConfig
func (o *VCenterVRSConfig) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VCenterVRSConfig
func (o *VCenterVRSConfig) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VRSAddressRanges retrieves the list of child VRSAddressRanges of the VCenterVRSConfig
func (o *VCenterVRSConfig) VRSAddressRanges(info *bambou.FetchingInfo) (VRSAddressRangesList, *bambou.Error) {

	var list VRSAddressRangesList
	err := bambou.CurrentSession().FetchChildren(o, VRSAddressRangeIdentity, &list, info)
	return list, err
}

// CreateVRSAddressRange creates a new child VRSAddressRange under the VCenterVRSConfig
func (o *VCenterVRSConfig) CreateVRSAddressRange(child *VRSAddressRange) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VRSRedeploymentpolicies retrieves the list of child VRSRedeploymentpolicies of the VCenterVRSConfig
func (o *VCenterVRSConfig) VRSRedeploymentpolicies(info *bambou.FetchingInfo) (VRSRedeploymentpoliciesList, *bambou.Error) {

	var list VRSRedeploymentpoliciesList
	err := bambou.CurrentSession().FetchChildren(o, VRSRedeploymentpolicyIdentity, &list, info)
	return list, err
}

// CreateVRSRedeploymentpolicy creates a new child VRSRedeploymentpolicy under the VCenterVRSConfig
func (o *VCenterVRSConfig) CreateVRSRedeploymentpolicy(child *VRSRedeploymentpolicy) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
