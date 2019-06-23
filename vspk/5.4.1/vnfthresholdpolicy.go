/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// VNFThresholdPolicyIdentity represents the Identity of the object
var VNFThresholdPolicyIdentity = bambou.Identity{
	Name:     "vnfthresholdpolicy",
	Category: "vnfthresholdpolicies",
}

// VNFThresholdPoliciesList represents a list of VNFThresholdPolicies
type VNFThresholdPoliciesList []*VNFThresholdPolicy

// VNFThresholdPoliciesAncestor is the interface that an ancestor of a VNFThresholdPolicy must implement.
// An Ancestor is defined as an entity that has VNFThresholdPolicy as a descendant.
// An Ancestor can get a list of its child VNFThresholdPolicies, but not necessarily create one.
type VNFThresholdPoliciesAncestor interface {
	VNFThresholdPolicies(*bambou.FetchingInfo) (VNFThresholdPoliciesList, *bambou.Error)
}

// VNFThresholdPoliciesParent is the interface that a parent of a VNFThresholdPolicy must implement.
// A Parent is defined as an entity that has VNFThresholdPolicy as a child.
// A Parent is an Ancestor which can create a VNFThresholdPolicy.
type VNFThresholdPoliciesParent interface {
	VNFThresholdPoliciesAncestor
	CreateVNFThresholdPolicy(*VNFThresholdPolicy) *bambou.Error
}

// VNFThresholdPolicy represents the model of a vnfthresholdpolicy
type VNFThresholdPolicy struct {
	ID               string `json:"ID,omitempty"`
	ParentID         string `json:"parentID,omitempty"`
	ParentType       string `json:"parentType,omitempty"`
	Owner            string `json:"owner,omitempty"`
	CPUThreshold     int    `json:"CPUThreshold,omitempty"`
	Name             string `json:"name,omitempty"`
	LastUpdatedBy    string `json:"lastUpdatedBy,omitempty"`
	Action           string `json:"action,omitempty"`
	MemoryThreshold  int    `json:"memoryThreshold,omitempty"`
	Description      string `json:"description,omitempty"`
	MinOccurrence    int    `json:"minOccurrence,omitempty"`
	EntityScope      string `json:"entityScope,omitempty"`
	MonitInterval    int    `json:"monitInterval,omitempty"`
	AssocEntityType  string `json:"assocEntityType,omitempty"`
	StorageThreshold int    `json:"storageThreshold,omitempty"`
	ExternalID       string `json:"externalID,omitempty"`
}

// NewVNFThresholdPolicy returns a new *VNFThresholdPolicy
func NewVNFThresholdPolicy() *VNFThresholdPolicy {

	return &VNFThresholdPolicy{
		CPUThreshold:     80,
		Action:           "NONE",
		MemoryThreshold:  80,
		MinOccurrence:    5,
		MonitInterval:    10,
		StorageThreshold: 80,
	}
}

// Identity returns the Identity of the object.
func (o *VNFThresholdPolicy) Identity() bambou.Identity {

	return VNFThresholdPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VNFThresholdPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VNFThresholdPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VNFThresholdPolicy from the server
func (o *VNFThresholdPolicy) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VNFThresholdPolicy into the server
func (o *VNFThresholdPolicy) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VNFThresholdPolicy from the server
func (o *VNFThresholdPolicy) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the VNFThresholdPolicy
func (o *VNFThresholdPolicy) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VNFThresholdPolicy
func (o *VNFThresholdPolicy) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VNFThresholdPolicy
func (o *VNFThresholdPolicy) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VNFThresholdPolicy
func (o *VNFThresholdPolicy) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
