/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VCenterClusterIdentity represents the Identity of the object
var VCenterClusterIdentity = bambou.Identity{
	Name:     "vcentercluster",
	Category: "vcenterclusters",
}

// VCenterClustersList represents a list of VCenterClusters
type VCenterClustersList []*VCenterCluster

// VCenterClustersAncestor is the interface that an ancestor of a VCenterCluster must implement.
// An Ancestor is defined as an entity that has VCenterCluster as a descendant.
// An Ancestor can get a list of its child VCenterClusters, but not necessarily create one.
type VCenterClustersAncestor interface {
	VCenterClusters(*bambou.FetchingInfo) (VCenterClustersList, *bambou.Error)
}

// VCenterClustersParent is the interface that a parent of a VCenterCluster must implement.
// A Parent is defined as an entity that has VCenterCluster as a child.
// A Parent is an Ancestor which can create a VCenterCluster.
type VCenterClustersParent interface {
	VCenterClustersAncestor
	CreateVCenterCluster(*VCenterCluster) *bambou.Error
}

// VCenterCluster represents the model of a vcentercluster
type VCenterCluster struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewVCenterCluster returns a new *VCenterCluster
func NewVCenterCluster() *VCenterCluster {

	return &VCenterCluster{}
}

// Identity returns the Identity of the object.
func (o *VCenterCluster) Identity() bambou.Identity {

	return VCenterClusterIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VCenterCluster) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VCenterCluster) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VCenterCluster from the server
func (o *VCenterCluster) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VCenterCluster into the server
func (o *VCenterCluster) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VCenterCluster from the server
func (o *VCenterCluster) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// VCenterHypervisors retrieves the list of child VCenterHypervisors of the VCenterCluster
func (o *VCenterCluster) VCenterHypervisors(info *bambou.FetchingInfo) (VCenterHypervisorsList, *bambou.Error) {

	var list VCenterHypervisorsList
	err := bambou.CurrentSession().FetchChildren(o, VCenterHypervisorIdentity, &list, info)
	return list, err
}

// CreateVCenterHypervisor creates a new child VCenterHypervisor under the VCenterCluster
func (o *VCenterCluster) CreateVCenterHypervisor(child *VCenterHypervisor) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the VCenterCluster
func (o *VCenterCluster) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VCenterCluster
func (o *VCenterCluster) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VCenterCluster
func (o *VCenterCluster) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VCenterCluster
func (o *VCenterCluster) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Jobs retrieves the list of child Jobs of the VCenterCluster
func (o *VCenterCluster) Jobs(info *bambou.FetchingInfo) (JobsList, *bambou.Error) {

	var list JobsList
	err := bambou.CurrentSession().FetchChildren(o, JobIdentity, &list, info)
	return list, err
}

// CreateJob creates a new child Job under the VCenterCluster
func (o *VCenterCluster) CreateJob(child *Job) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VRSAddressRanges retrieves the list of child VRSAddressRanges of the VCenterCluster
func (o *VCenterCluster) VRSAddressRanges(info *bambou.FetchingInfo) (VRSAddressRangesList, *bambou.Error) {

	var list VRSAddressRangesList
	err := bambou.CurrentSession().FetchChildren(o, VRSAddressRangeIdentity, &list, info)
	return list, err
}

// CreateVRSAddressRange creates a new child VRSAddressRange under the VCenterCluster
func (o *VCenterCluster) CreateVRSAddressRange(child *VRSAddressRange) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// AutoDiscoverHypervisorFromClusters retrieves the list of child AutoDiscoverHypervisorFromClusters of the VCenterCluster
func (o *VCenterCluster) AutoDiscoverHypervisorFromClusters(info *bambou.FetchingInfo) (AutoDiscoverHypervisorFromClustersList, *bambou.Error) {

	var list AutoDiscoverHypervisorFromClustersList
	err := bambou.CurrentSession().FetchChildren(o, AutoDiscoverHypervisorFromClusterIdentity, &list, info)
	return list, err
}
