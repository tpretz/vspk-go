/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// GatewayRedundantPortIdentity represents the Identity of the object
var GatewayRedundantPortIdentity = bambou.Identity{
	Name:     "gatewayredundantport",
	Category: "gatewayredundantports",
}

// GatewayRedundantPortsList represents a list of GatewayRedundantPorts
type GatewayRedundantPortsList []*GatewayRedundantPort

// GatewayRedundantPortsAncestor is the interface that an ancestor of a GatewayRedundantPort must implement.
// An Ancestor is defined as an entity that has GatewayRedundantPort as a descendant.
// An Ancestor can get a list of its child GatewayRedundantPorts, but not necessarily create one.
type GatewayRedundantPortsAncestor interface {
	GatewayRedundantPorts(*bambou.FetchingInfo) (GatewayRedundantPortsList, *bambou.Error)
}

// GatewayRedundantPortsParent is the interface that a parent of a GatewayRedundantPort must implement.
// A Parent is defined as an entity that has GatewayRedundantPort as a child.
// A Parent is an Ancestor which can create a GatewayRedundantPort.
type GatewayRedundantPortsParent interface {
	GatewayRedundantPortsAncestor
	CreateGatewayRedundantPort(*GatewayRedundantPort) *bambou.Error
}

// GatewayRedundantPort represents the model of a gatewayredundantport
type GatewayRedundantPort struct {
	ID                          string `json:"ID,omitempty"`
	ParentID                    string `json:"parentID,omitempty"`
	ParentType                  string `json:"parentType,omitempty"`
	Owner                       string `json:"owner,omitempty"`
	VLANRange                   string `json:"VLANRange,omitempty"`
	Name                        string `json:"name,omitempty"`
	LastUpdatedBy               string `json:"lastUpdatedBy,omitempty"`
	PermittedAction             string `json:"permittedAction,omitempty"`
	Description                 string `json:"description,omitempty"`
	PhysicalName                string `json:"physicalName,omitempty"`
	EntityScope                 string `json:"entityScope,omitempty"`
	PortPeer1ID                 string `json:"portPeer1ID,omitempty"`
	PortPeer2ID                 string `json:"portPeer2ID,omitempty"`
	PortType                    string `json:"portType,omitempty"`
	UseUserMnemonic             bool   `json:"useUserMnemonic"`
	UserMnemonic                string `json:"userMnemonic,omitempty"`
	AssociatedEgressQOSPolicyID string `json:"associatedEgressQOSPolicyID,omitempty"`
	Status                      string `json:"status,omitempty"`
	ExternalID                  string `json:"externalID,omitempty"`
}

// NewGatewayRedundantPort returns a new *GatewayRedundantPort
func NewGatewayRedundantPort() *GatewayRedundantPort {

	return &GatewayRedundantPort{}
}

// Identity returns the Identity of the object.
func (o *GatewayRedundantPort) Identity() bambou.Identity {

	return GatewayRedundantPortIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *GatewayRedundantPort) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *GatewayRedundantPort) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the GatewayRedundantPort from the server
func (o *GatewayRedundantPort) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the GatewayRedundantPort into the server
func (o *GatewayRedundantPort) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the GatewayRedundantPort from the server
func (o *GatewayRedundantPort) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Permissions retrieves the list of child Permissions of the GatewayRedundantPort
func (o *GatewayRedundantPort) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

	var list PermissionsList
	err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
	return list, err
}

// CreatePermission creates a new child Permission under the GatewayRedundantPort
func (o *GatewayRedundantPort) CreatePermission(child *Permission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the GatewayRedundantPort
func (o *GatewayRedundantPort) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the GatewayRedundantPort
func (o *GatewayRedundantPort) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VLANs retrieves the list of child VLANs of the GatewayRedundantPort
func (o *GatewayRedundantPort) VLANs(info *bambou.FetchingInfo) (VLANsList, *bambou.Error) {

	var list VLANsList
	err := bambou.CurrentSession().FetchChildren(o, VLANIdentity, &list, info)
	return list, err
}

// CreateVLAN creates a new child VLAN under the GatewayRedundantPort
func (o *GatewayRedundantPort) CreateVLAN(child *VLAN) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Alarms retrieves the list of child Alarms of the GatewayRedundantPort
func (o *GatewayRedundantPort) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

	var list AlarmsList
	err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
	return list, err
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the GatewayRedundantPort
func (o *GatewayRedundantPort) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the GatewayRedundantPort
func (o *GatewayRedundantPort) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EnterprisePermissions retrieves the list of child EnterprisePermissions of the GatewayRedundantPort
func (o *GatewayRedundantPort) EnterprisePermissions(info *bambou.FetchingInfo) (EnterprisePermissionsList, *bambou.Error) {

	var list EnterprisePermissionsList
	err := bambou.CurrentSession().FetchChildren(o, EnterprisePermissionIdentity, &list, info)
	return list, err
}

// CreateEnterprisePermission creates a new child EnterprisePermission under the GatewayRedundantPort
func (o *GatewayRedundantPort) CreateEnterprisePermission(child *EnterprisePermission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
