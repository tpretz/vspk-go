/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// NSGatewayIdentity represents the Identity of the object
var NSGatewayIdentity = bambou.Identity{
	Name:     "nsgateway",
	Category: "nsgateways",
}

// NSGatewaysList represents a list of NSGateways
type NSGatewaysList []*NSGateway

// NSGatewaysAncestor is the interface that an ancestor of a NSGateway must implement.
// An Ancestor is defined as an entity that has NSGateway as a descendant.
// An Ancestor can get a list of its child NSGateways, but not necessarily create one.
type NSGatewaysAncestor interface {
	NSGateways(*bambou.FetchingInfo) (NSGatewaysList, *bambou.Error)
}

// NSGatewaysParent is the interface that a parent of a NSGateway must implement.
// A Parent is defined as an entity that has NSGateway as a child.
// A Parent is an Ancestor which can create a NSGateway.
type NSGatewaysParent interface {
	NSGatewaysAncestor
	CreateNSGateway(*NSGateway) *bambou.Error
}

// NSGateway represents the model of a nsgateway
type NSGateway struct {
	ID                                 string `json:"ID,omitempty"`
	ParentID                           string `json:"parentID,omitempty"`
	ParentType                         string `json:"parentType,omitempty"`
	Owner                              string `json:"owner,omitempty"`
	MACAddress                         string `json:"MACAddress,omitempty"`
	NATTraversalEnabled                bool   `json:"NATTraversalEnabled"`
	TCPMSSEnabled                      bool   `json:"TCPMSSEnabled"`
	TCPMaximumSegmentSize              int    `json:"TCPMaximumSegmentSize,omitempty"`
	BIOSReleaseDate                    string `json:"BIOSReleaseDate,omitempty"`
	BIOSVersion                        string `json:"BIOSVersion,omitempty"`
	SKU                                string `json:"SKU,omitempty"`
	TPMStatus                          string `json:"TPMStatus,omitempty"`
	CPUType                            string `json:"CPUType,omitempty"`
	NSGVersion                         string `json:"NSGVersion,omitempty"`
	SSHService                         string `json:"SSHService,omitempty"`
	UUID                               string `json:"UUID,omitempty"`
	Name                               string `json:"name,omitempty"`
	Family                             string `json:"family,omitempty"`
	LastConfigurationReloadTimestamp   int    `json:"lastConfigurationReloadTimestamp,omitempty"`
	LastUpdatedBy                      string `json:"lastUpdatedBy,omitempty"`
	DatapathID                         string `json:"datapathID,omitempty"`
	RedundancyGroupID                  string `json:"redundancyGroupID,omitempty"`
	TemplateID                         string `json:"templateID,omitempty"`
	Pending                            bool   `json:"pending"`
	SerialNumber                       string `json:"serialNumber,omitempty"`
	DerivedSSHServiceState             string `json:"derivedSSHServiceState,omitempty"`
	PermittedAction                    string `json:"permittedAction,omitempty"`
	Personality                        string `json:"personality,omitempty"`
	Description                        string `json:"description,omitempty"`
	Libraries                          string `json:"libraries,omitempty"`
	InheritedSSHServiceState           string `json:"inheritedSSHServiceState,omitempty"`
	EnterpriseID                       string `json:"enterpriseID,omitempty"`
	EntityScope                        string `json:"entityScope,omitempty"`
	LocationID                         string `json:"locationID,omitempty"`
	ConfigurationReloadState           string `json:"configurationReloadState,omitempty"`
	ConfigurationStatus                string `json:"configurationStatus,omitempty"`
	ControlTrafficCOSValue             int    `json:"controlTrafficCOSValue,omitempty"`
	ControlTrafficDSCPValue            int    `json:"controlTrafficDSCPValue,omitempty"`
	BootstrapID                        string `json:"bootstrapID,omitempty"`
	BootstrapStatus                    string `json:"bootstrapStatus,omitempty"`
	OperationMode                      string `json:"operationMode,omitempty"`
	OperationStatus                    string `json:"operationStatus,omitempty"`
	ProductName                        string `json:"productName,omitempty"`
	AssociatedGatewaySecurityID        string `json:"associatedGatewaySecurityID,omitempty"`
	AssociatedGatewaySecurityProfileID string `json:"associatedGatewaySecurityProfileID,omitempty"`
	AssociatedNSGInfoID                string `json:"associatedNSGInfoID,omitempty"`
	AssociatedNSGUpgradeProfileID      string `json:"associatedNSGUpgradeProfileID,omitempty"`
	AutoDiscGatewayID                  string `json:"autoDiscGatewayID,omitempty"`
	ExternalID                         string `json:"externalID,omitempty"`
	SystemID                           string `json:"systemID,omitempty"`
}

// NewNSGateway returns a new *NSGateway
func NewNSGateway() *NSGateway {

	return &NSGateway{
		TCPMSSEnabled:                    false,
		TCPMaximumSegmentSize:            1330,
		TPMStatus:                        "UNKNOWN",
		SSHService:                       "INHERITED",
		LastConfigurationReloadTimestamp: -1,
		InheritedSSHServiceState:         "ENABLED",
		ConfigurationReloadState:         "UNKNOWN",
		ConfigurationStatus:              "UNKNOWN",
		ControlTrafficCOSValue:           7,
		ControlTrafficDSCPValue:          56,
	}
}

// Identity returns the Identity of the object.
func (o *NSGateway) Identity() bambou.Identity {

	return NSGatewayIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NSGateway) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NSGateway) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NSGateway from the server
func (o *NSGateway) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NSGateway into the server
func (o *NSGateway) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NSGateway from the server
func (o *NSGateway) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// GatewaySecurities retrieves the list of child GatewaySecurities of the NSGateway
func (o *NSGateway) GatewaySecurities(info *bambou.FetchingInfo) (GatewaySecuritiesList, *bambou.Error) {

	var list GatewaySecuritiesList
	err := bambou.CurrentSession().FetchChildren(o, GatewaySecurityIdentity, &list, info)
	return list, err
}

// PATNATPools retrieves the list of child PATNATPools of the NSGateway
func (o *NSGateway) PATNATPools(info *bambou.FetchingInfo) (PATNATPoolsList, *bambou.Error) {

	var list PATNATPoolsList
	err := bambou.CurrentSession().FetchChildren(o, PATNATPoolIdentity, &list, info)
	return list, err
}

// AssignPATNATPools assigns the list of PATNATPools to the NSGateway
func (o *NSGateway) AssignPATNATPools(children PATNATPoolsList) *bambou.Error {

	list := []bambou.Identifiable{}
	for _, c := range children {
		list = append(list, c)
	}

	return bambou.CurrentSession().AssignChildren(o, list, PATNATPoolIdentity)
}

// Permissions retrieves the list of child Permissions of the NSGateway
func (o *NSGateway) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

	var list PermissionsList
	err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
	return list, err
}

// CreatePermission creates a new child Permission under the NSGateway
func (o *NSGateway) CreatePermission(child *Permission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the NSGateway
func (o *NSGateway) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the NSGateway
func (o *NSGateway) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// WirelessPorts retrieves the list of child WirelessPorts of the NSGateway
func (o *NSGateway) WirelessPorts(info *bambou.FetchingInfo) (WirelessPortsList, *bambou.Error) {

	var list WirelessPortsList
	err := bambou.CurrentSession().FetchChildren(o, WirelessPortIdentity, &list, info)
	return list, err
}

// CreateWirelessPort creates a new child WirelessPort under the NSGateway
func (o *NSGateway) CreateWirelessPort(child *WirelessPort) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Alarms retrieves the list of child Alarms of the NSGateway
func (o *NSGateway) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

	var list AlarmsList
	err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
	return list, err
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the NSGateway
func (o *NSGateway) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the NSGateway
func (o *NSGateway) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// InfrastructureConfigs retrieves the list of child InfrastructureConfigs of the NSGateway
func (o *NSGateway) InfrastructureConfigs(info *bambou.FetchingInfo) (InfrastructureConfigsList, *bambou.Error) {

	var list InfrastructureConfigsList
	err := bambou.CurrentSession().FetchChildren(o, InfrastructureConfigIdentity, &list, info)
	return list, err
}

// EnterprisePermissions retrieves the list of child EnterprisePermissions of the NSGateway
func (o *NSGateway) EnterprisePermissions(info *bambou.FetchingInfo) (EnterprisePermissionsList, *bambou.Error) {

	var list EnterprisePermissionsList
	err := bambou.CurrentSession().FetchChildren(o, EnterprisePermissionIdentity, &list, info)
	return list, err
}

// Jobs retrieves the list of child Jobs of the NSGateway
func (o *NSGateway) Jobs(info *bambou.FetchingInfo) (JobsList, *bambou.Error) {

	var list JobsList
	err := bambou.CurrentSession().FetchChildren(o, JobIdentity, &list, info)
	return list, err
}

// CreateJob creates a new child Job under the NSGateway
func (o *NSGateway) CreateJob(child *Job) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Locations retrieves the list of child Locations of the NSGateway
func (o *NSGateway) Locations(info *bambou.FetchingInfo) (LocationsList, *bambou.Error) {

	var list LocationsList
	err := bambou.CurrentSession().FetchChildren(o, LocationIdentity, &list, info)
	return list, err
}

// Commands retrieves the list of child Commands of the NSGateway
func (o *NSGateway) Commands(info *bambou.FetchingInfo) (CommandsList, *bambou.Error) {

	var list CommandsList
	err := bambou.CurrentSession().FetchChildren(o, CommandIdentity, &list, info)
	return list, err
}

// CreateCommand creates a new child Command under the NSGateway
func (o *NSGateway) CreateCommand(child *Command) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Monitorscopes retrieves the list of child Monitorscopes of the NSGateway
func (o *NSGateway) Monitorscopes(info *bambou.FetchingInfo) (MonitorscopesList, *bambou.Error) {

	var list MonitorscopesList
	err := bambou.CurrentSession().FetchChildren(o, MonitorscopeIdentity, &list, info)
	return list, err
}

// Bootstraps retrieves the list of child Bootstraps of the NSGateway
func (o *NSGateway) Bootstraps(info *bambou.FetchingInfo) (BootstrapsList, *bambou.Error) {

	var list BootstrapsList
	err := bambou.CurrentSession().FetchChildren(o, BootstrapIdentity, &list, info)
	return list, err
}

// CreateBootstrapActivation creates a new child BootstrapActivation under the NSGateway
func (o *NSGateway) CreateBootstrapActivation(child *BootstrapActivation) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// UplinkConnections retrieves the list of child UplinkConnections of the NSGateway
func (o *NSGateway) UplinkConnections(info *bambou.FetchingInfo) (UplinkConnectionsList, *bambou.Error) {

	var list UplinkConnectionsList
	err := bambou.CurrentSession().FetchChildren(o, UplinkConnectionIdentity, &list, info)
	return list, err
}

// NSGInfos retrieves the list of child NSGInfos of the NSGateway
func (o *NSGateway) NSGInfos(info *bambou.FetchingInfo) (NSGInfosList, *bambou.Error) {

	var list NSGInfosList
	err := bambou.CurrentSession().FetchChildren(o, NSGInfoIdentity, &list, info)
	return list, err
}

// NSPorts retrieves the list of child NSPorts of the NSGateway
func (o *NSGateway) NSPorts(info *bambou.FetchingInfo) (NSPortsList, *bambou.Error) {

	var list NSPortsList
	err := bambou.CurrentSession().FetchChildren(o, NSPortIdentity, &list, info)
	return list, err
}

// CreateNSPort creates a new child NSPort under the NSGateway
func (o *NSGateway) CreateNSPort(child *NSPort) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Subnets retrieves the list of child Subnets of the NSGateway
func (o *NSGateway) Subnets(info *bambou.FetchingInfo) (SubnetsList, *bambou.Error) {

	var list SubnetsList
	err := bambou.CurrentSession().FetchChildren(o, SubnetIdentity, &list, info)
	return list, err
}

// EventLogs retrieves the list of child EventLogs of the NSGateway
func (o *NSGateway) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
