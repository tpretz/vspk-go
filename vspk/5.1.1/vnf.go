/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VNFIdentity represents the Identity of the object
var VNFIdentity = bambou.Identity{
	Name:     "vnf",
	Category: "vnfs",
}

// VNFsList represents a list of VNFs
type VNFsList []*VNF

// VNFsAncestor is the interface that an ancestor of a VNF must implement.
// An Ancestor is defined as an entity that has VNF as a descendant.
// An Ancestor can get a list of its child VNFs, but not necessarily create one.
type VNFsAncestor interface {
	VNFs(*bambou.FetchingInfo) (VNFsList, *bambou.Error)
}

// VNFsParent is the interface that a parent of a VNF must implement.
// A Parent is defined as an entity that has VNF as a child.
// A Parent is an Ancestor which can create a VNF.
type VNFsParent interface {
	VNFsAncestor
	CreateVNF(*VNF) *bambou.Error
}

// VNF represents the model of a vnf
type VNF struct {
	ID                      string `json:"ID,omitempty"`
	ParentID                string `json:"parentID,omitempty"`
	ParentType              string `json:"parentType,omitempty"`
	Owner                   string `json:"owner,omitempty"`
	VNFDescriptorID         string `json:"VNFDescriptorID,omitempty"`
	VNFDescriptorName       string `json:"VNFDescriptorName,omitempty"`
	CPUCount                int    `json:"CPUCount,omitempty"`
	NSGName                 string `json:"NSGName,omitempty"`
	NSGSystemID             string `json:"NSGSystemID,omitempty"`
	NSGatewayID             string `json:"NSGatewayID,omitempty"`
	Name                    string `json:"name,omitempty"`
	TaskState               string `json:"taskState,omitempty"`
	LastKnownError          string `json:"lastKnownError,omitempty"`
	MemoryMB                int    `json:"memoryMB,omitempty"`
	Vendor                  string `json:"vendor,omitempty"`
	Description             string `json:"description,omitempty"`
	MetadataID              string `json:"metadataID,omitempty"`
	AllowedActions          string `json:"allowedActions,omitempty"`
	EnterpriseID            string `json:"enterpriseID,omitempty"`
	IsAttachedToDescriptor  bool   `json:"isAttachedToDescriptor"`
	AssociatedVNFMetadataID string `json:"associatedVNFMetadataID,omitempty"`
	Status                  string `json:"status,omitempty"`
	StorageGB               int    `json:"storageGB,omitempty"`
}

// NewVNF returns a new *VNF
func NewVNF() *VNF {

	return &VNF{
		IsAttachedToDescriptor: true,
	}
}

// Identity returns the Identity of the object.
func (o *VNF) Identity() bambou.Identity {

	return VNFIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VNF) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VNF) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VNF from the server
func (o *VNF) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VNF into the server
func (o *VNF) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VNF from the server
func (o *VNF) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// VNFInterfaces retrieves the list of child VNFInterfaces of the VNF
func (o *VNF) VNFInterfaces(info *bambou.FetchingInfo) (VNFInterfacesList, *bambou.Error) {

	var list VNFInterfacesList
	err := bambou.CurrentSession().FetchChildren(o, VNFInterfaceIdentity, &list, info)
	return list, err
}

// VNFMetadatas retrieves the list of child VNFMetadatas of the VNF
func (o *VNF) VNFMetadatas(info *bambou.FetchingInfo) (VNFMetadatasList, *bambou.Error) {

	var list VNFMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, VNFMetadataIdentity, &list, info)
	return list, err
}

// CreateJob creates a new child Job under the VNF
func (o *VNF) CreateJob(child *Job) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
