/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// NSPortIdentity represents the Identity of the object
var NSPortIdentity = bambou.Identity{
	Name:     "nsport",
	Category: "nsports",
}

// NSPortsList represents a list of NSPorts
type NSPortsList []*NSPort

// NSPortsAncestor is the interface that an ancestor of a NSPort must implement.
// An Ancestor is defined as an entity that has NSPort as a descendant.
// An Ancestor can get a list of its child NSPorts, but not necessarily create one.
type NSPortsAncestor interface {
	NSPorts(*bambou.FetchingInfo) (NSPortsList, *bambou.Error)
}

// NSPortsParent is the interface that a parent of a NSPort must implement.
// A Parent is defined as an entity that has NSPort as a child.
// A Parent is an Ancestor which can create a NSPort.
type NSPortsParent interface {
	NSPortsAncestor
	CreateNSPort(*NSPort) *bambou.Error
}

// NSPort represents the model of a nsport
type NSPort struct {
	ID                          string `json:"ID,omitempty"`
	ParentID                    string `json:"parentID,omitempty"`
	ParentType                  string `json:"parentType,omitempty"`
	Owner                       string `json:"owner,omitempty"`
	NATTraversal                string `json:"NATTraversal,omitempty"`
	VLANRange                   string `json:"VLANRange,omitempty"`
	Name                        string `json:"name,omitempty"`
	LastUpdatedBy               string `json:"lastUpdatedBy,omitempty"`
	TemplateID                  string `json:"templateID,omitempty"`
	PermittedAction             string `json:"permittedAction,omitempty"`
	Description                 string `json:"description,omitempty"`
	PhysicalName                string `json:"physicalName,omitempty"`
	EnableNATProbes             bool   `json:"enableNATProbes"`
	EntityScope                 string `json:"entityScope,omitempty"`
	PortType                    string `json:"portType,omitempty"`
	Speed                       string `json:"speed,omitempty"`
	TrafficThroughUBROnly       bool   `json:"TrafficThroughUBROnly"`
	UseUserMnemonic             bool   `json:"useUserMnemonic"`
	UserMnemonic                string `json:"userMnemonic,omitempty"`
	AssociatedEgressQOSPolicyID string `json:"associatedEgressQOSPolicyID,omitempty"`
	AssociatedRedundantPortID   string `json:"associatedRedundantPortID,omitempty"`
	Status                      string `json:"status,omitempty"`
	Mtu                         int    `json:"mtu,omitempty"`
	ExternalID                  string `json:"externalID,omitempty"`
}

// NewNSPort returns a new *NSPort
func NewNSPort() *NSPort {

	return &NSPort{
		NATTraversal:          "NONE",
		EnableNATProbes:       true,
		TrafficThroughUBROnly: false,
		Mtu: 1500,
	}
}

// Identity returns the Identity of the object.
func (o *NSPort) Identity() bambou.Identity {

	return NSPortIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NSPort) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NSPort) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NSPort from the server
func (o *NSPort) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NSPort into the server
func (o *NSPort) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NSPort from the server
func (o *NSPort) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Permissions retrieves the list of child Permissions of the NSPort
func (o *NSPort) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

	var list PermissionsList
	err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
	return list, err
}

// CreatePermission creates a new child Permission under the NSPort
func (o *NSPort) CreatePermission(child *Permission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the NSPort
func (o *NSPort) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the NSPort
func (o *NSPort) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VLANs retrieves the list of child VLANs of the NSPort
func (o *NSPort) VLANs(info *bambou.FetchingInfo) (VLANsList, *bambou.Error) {

	var list VLANsList
	err := bambou.CurrentSession().FetchChildren(o, VLANIdentity, &list, info)
	return list, err
}

// CreateVLAN creates a new child VLAN under the NSPort
func (o *NSPort) CreateVLAN(child *VLAN) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Alarms retrieves the list of child Alarms of the NSPort
func (o *NSPort) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

	var list AlarmsList
	err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
	return list, err
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the NSPort
func (o *NSPort) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the NSPort
func (o *NSPort) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EnterprisePermissions retrieves the list of child EnterprisePermissions of the NSPort
func (o *NSPort) EnterprisePermissions(info *bambou.FetchingInfo) (EnterprisePermissionsList, *bambou.Error) {

	var list EnterprisePermissionsList
	err := bambou.CurrentSession().FetchChildren(o, EnterprisePermissionIdentity, &list, info)
	return list, err
}

// CreateEnterprisePermission creates a new child EnterprisePermission under the NSPort
func (o *NSPort) CreateEnterprisePermission(child *EnterprisePermission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Statistics retrieves the list of child Statistics of the NSPort
func (o *NSPort) Statistics(info *bambou.FetchingInfo) (StatisticsList, *bambou.Error) {

	var list StatisticsList
	err := bambou.CurrentSession().FetchChildren(o, StatisticsIdentity, &list, info)
	return list, err
}

// StatisticsPolicies retrieves the list of child StatisticsPolicies of the NSPort
func (o *NSPort) StatisticsPolicies(info *bambou.FetchingInfo) (StatisticsPoliciesList, *bambou.Error) {

	var list StatisticsPoliciesList
	err := bambou.CurrentSession().FetchChildren(o, StatisticsPolicyIdentity, &list, info)
	return list, err
}

// CreateStatisticsPolicy creates a new child StatisticsPolicy under the NSPort
func (o *NSPort) CreateStatisticsPolicy(child *StatisticsPolicy) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// LTEInformations retrieves the list of child LTEInformations of the NSPort
func (o *NSPort) LTEInformations(info *bambou.FetchingInfo) (LTEInformationsList, *bambou.Error) {

	var list LTEInformationsList
	err := bambou.CurrentSession().FetchChildren(o, LTEInformationIdentity, &list, info)
	return list, err
}

// EventLogs retrieves the list of child EventLogs of the NSPort
func (o *NSPort) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
