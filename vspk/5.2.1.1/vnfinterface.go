/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VNFInterfaceIdentity represents the Identity of the object
var VNFInterfaceIdentity = bambou.Identity{
	Name:     "vnfinterface",
	Category: "vnfinterfaces",
}

// VNFInterfacesList represents a list of VNFInterfaces
type VNFInterfacesList []*VNFInterface

// VNFInterfacesAncestor is the interface that an ancestor of a VNFInterface must implement.
// An Ancestor is defined as an entity that has VNFInterface as a descendant.
// An Ancestor can get a list of its child VNFInterfaces, but not necessarily create one.
type VNFInterfacesAncestor interface {
	VNFInterfaces(*bambou.FetchingInfo) (VNFInterfacesList, *bambou.Error)
}

// VNFInterfacesParent is the interface that a parent of a VNFInterface must implement.
// A Parent is defined as an entity that has VNFInterface as a child.
// A Parent is an Ancestor which can create a VNFInterface.
type VNFInterfacesParent interface {
	VNFInterfacesAncestor
	CreateVNFInterface(*VNFInterface) *bambou.Error
}

// VNFInterface represents the model of a vnfinterface
type VNFInterface struct {
	ID                    string `json:"ID,omitempty"`
	ParentID              string `json:"parentID,omitempty"`
	ParentType            string `json:"parentType,omitempty"`
	Owner                 string `json:"owner,omitempty"`
	MAC                   string `json:"MAC,omitempty"`
	VNFUUID               string `json:"VNFUUID,omitempty"`
	IPAddress             string `json:"IPAddress,omitempty"`
	VPortID               string `json:"VPortID,omitempty"`
	VPortName             string `json:"VPortName,omitempty"`
	Name                  string `json:"name,omitempty"`
	Gateway               string `json:"gateway,omitempty"`
	Netmask               string `json:"netmask,omitempty"`
	NetworkName           string `json:"networkName,omitempty"`
	PolicyDecisionID      string `json:"policyDecisionID,omitempty"`
	DomainID              string `json:"domainID,omitempty"`
	DomainName            string `json:"domainName,omitempty"`
	ZoneID                string `json:"zoneID,omitempty"`
	ZoneName              string `json:"zoneName,omitempty"`
	IsManagementInterface bool   `json:"isManagementInterface"`
	AttachedNetworkID     string `json:"attachedNetworkID,omitempty"`
	AttachedNetworkType   string `json:"attachedNetworkType,omitempty"`
}

// NewVNFInterface returns a new *VNFInterface
func NewVNFInterface() *VNFInterface {

	return &VNFInterface{
		IsManagementInterface: false,
	}
}

// Identity returns the Identity of the object.
func (o *VNFInterface) Identity() bambou.Identity {

	return VNFInterfaceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VNFInterface) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VNFInterface) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VNFInterface from the server
func (o *VNFInterface) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VNFInterface into the server
func (o *VNFInterface) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VNFInterface from the server
func (o *VNFInterface) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
