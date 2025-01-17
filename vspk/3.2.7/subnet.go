/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// SubnetIdentity represents the Identity of the object
var SubnetIdentity = bambou.Identity{
	Name:     "subnet",
	Category: "subnets",
}

// SubnetsList represents a list of Subnets
type SubnetsList []*Subnet

// SubnetsAncestor is the interface that an ancestor of a Subnet must implement.
// An Ancestor is defined as an entity that has Subnet as a descendant.
// An Ancestor can get a list of its child Subnets, but not necessarily create one.
type SubnetsAncestor interface {
	Subnets(*bambou.FetchingInfo) (SubnetsList, *bambou.Error)
}

// SubnetsParent is the interface that a parent of a Subnet must implement.
// A Parent is defined as an entity that has Subnet as a child.
// A Parent is an Ancestor which can create a Subnet.
type SubnetsParent interface {
	SubnetsAncestor
	CreateSubnet(*Subnet) *bambou.Error
}

// Subnet represents the model of a subnet
type Subnet struct {
	ID                                string `json:"ID,omitempty"`
	ParentID                          string `json:"parentID,omitempty"`
	ParentType                        string `json:"parentType,omitempty"`
	Owner                             string `json:"owner,omitempty"`
	PATEnabled                        string `json:"PATEnabled,omitempty"`
	IPType                            string `json:"IPType,omitempty"`
	MaintenanceMode                   string `json:"maintenanceMode,omitempty"`
	Name                              string `json:"name,omitempty"`
	LastUpdatedBy                     string `json:"lastUpdatedBy,omitempty"`
	Gateway                           string `json:"gateway,omitempty"`
	GatewayMACAddress                 string `json:"gatewayMACAddress,omitempty"`
	Address                           string `json:"address,omitempty"`
	TemplateID                        string `json:"templateID,omitempty"`
	ServiceID                         int    `json:"serviceID,omitempty"`
	Description                       string `json:"description,omitempty"`
	Netmask                           string `json:"netmask,omitempty"`
	VnId                              int    `json:"vnId,omitempty"`
	Encryption                        string `json:"encryption,omitempty"`
	UnderlayEnabled                   string `json:"underlayEnabled,omitempty"`
	EntityScope                       string `json:"entityScope,omitempty"`
	PolicyGroupID                     int    `json:"policyGroupID,omitempty"`
	RouteDistinguisher                string `json:"routeDistinguisher,omitempty"`
	RouteTarget                       string `json:"routeTarget,omitempty"`
	SplitSubnet                       bool   `json:"splitSubnet"`
	ProxyARP                          bool   `json:"proxyARP"`
	AssociatedApplicationID           string `json:"associatedApplicationID,omitempty"`
	AssociatedApplicationObjectID     string `json:"associatedApplicationObjectID,omitempty"`
	AssociatedApplicationObjectType   string `json:"associatedApplicationObjectType,omitempty"`
	AssociatedMulticastChannelMapID   string `json:"associatedMulticastChannelMapID,omitempty"`
	AssociatedSharedNetworkResourceID string `json:"associatedSharedNetworkResourceID,omitempty"`
	Public                            bool   `json:"public"`
	Multicast                         string `json:"multicast,omitempty"`
	ExternalID                        string `json:"externalID,omitempty"`
}

// NewSubnet returns a new *Subnet
func NewSubnet() *Subnet {

	return &Subnet{}
}

// Identity returns the Identity of the object.
func (o *Subnet) Identity() bambou.Identity {

	return SubnetIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Subnet) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Subnet) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Subnet from the server
func (o *Subnet) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Subnet into the server
func (o *Subnet) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Subnet from the server
func (o *Subnet) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// TCAs retrieves the list of child TCAs of the Subnet
func (o *Subnet) TCAs(info *bambou.FetchingInfo) (TCAsList, *bambou.Error) {

	var list TCAsList
	err := bambou.CurrentSession().FetchChildren(o, TCAIdentity, &list, info)
	return list, err
}

// CreateTCA creates a new child TCA under the Subnet
func (o *Subnet) CreateTCA(child *TCA) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// AddressRanges retrieves the list of child AddressRanges of the Subnet
func (o *Subnet) AddressRanges(info *bambou.FetchingInfo) (AddressRangesList, *bambou.Error) {

	var list AddressRangesList
	err := bambou.CurrentSession().FetchChildren(o, AddressRangeIdentity, &list, info)
	return list, err
}

// CreateAddressRange creates a new child AddressRange under the Subnet
func (o *Subnet) CreateAddressRange(child *AddressRange) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VMResyncs retrieves the list of child VMResyncs of the Subnet
func (o *Subnet) VMResyncs(info *bambou.FetchingInfo) (VMResyncsList, *bambou.Error) {

	var list VMResyncsList
	err := bambou.CurrentSession().FetchChildren(o, VMResyncIdentity, &list, info)
	return list, err
}

// CreateVMResync creates a new child VMResync under the Subnet
func (o *Subnet) CreateVMResync(child *VMResync) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the Subnet
func (o *Subnet) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Subnet
func (o *Subnet) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// DHCPOptions retrieves the list of child DHCPOptions of the Subnet
func (o *Subnet) DHCPOptions(info *bambou.FetchingInfo) (DHCPOptionsList, *bambou.Error) {

	var list DHCPOptionsList
	err := bambou.CurrentSession().FetchChildren(o, DHCPOptionIdentity, &list, info)
	return list, err
}

// CreateDHCPOption creates a new child DHCPOption under the Subnet
func (o *Subnet) CreateDHCPOption(child *DHCPOption) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VirtualIPs retrieves the list of child VirtualIPs of the Subnet
func (o *Subnet) VirtualIPs(info *bambou.FetchingInfo) (VirtualIPsList, *bambou.Error) {

	var list VirtualIPsList
	err := bambou.CurrentSession().FetchChildren(o, VirtualIPIdentity, &list, info)
	return list, err
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Subnet
func (o *Subnet) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Subnet
func (o *Subnet) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VMs retrieves the list of child VMs of the Subnet
func (o *Subnet) VMs(info *bambou.FetchingInfo) (VMsList, *bambou.Error) {

	var list VMsList
	err := bambou.CurrentSession().FetchChildren(o, VMIdentity, &list, info)
	return list, err
}

// VMInterfaces retrieves the list of child VMInterfaces of the Subnet
func (o *Subnet) VMInterfaces(info *bambou.FetchingInfo) (VMInterfacesList, *bambou.Error) {

	var list VMInterfacesList
	err := bambou.CurrentSession().FetchChildren(o, VMInterfaceIdentity, &list, info)
	return list, err
}

// QOSs retrieves the list of child QOSs of the Subnet
func (o *Subnet) QOSs(info *bambou.FetchingInfo) (QOSsList, *bambou.Error) {

	var list QOSsList
	err := bambou.CurrentSession().FetchChildren(o, QOSIdentity, &list, info)
	return list, err
}

// CreateQOS creates a new child QOS under the Subnet
func (o *Subnet) CreateQOS(child *QOS) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VPorts retrieves the list of child VPorts of the Subnet
func (o *Subnet) VPorts(info *bambou.FetchingInfo) (VPortsList, *bambou.Error) {

	var list VPortsList
	err := bambou.CurrentSession().FetchChildren(o, VPortIdentity, &list, info)
	return list, err
}

// CreateVPort creates a new child VPort under the Subnet
func (o *Subnet) CreateVPort(child *VPort) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// IPReservations retrieves the list of child IPReservations of the Subnet
func (o *Subnet) IPReservations(info *bambou.FetchingInfo) (IPReservationsList, *bambou.Error) {

	var list IPReservationsList
	err := bambou.CurrentSession().FetchChildren(o, IPReservationIdentity, &list, info)
	return list, err
}

// CreateIPReservation creates a new child IPReservation under the Subnet
func (o *Subnet) CreateIPReservation(child *IPReservation) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Statistics retrieves the list of child Statistics of the Subnet
func (o *Subnet) Statistics(info *bambou.FetchingInfo) (StatisticsList, *bambou.Error) {

	var list StatisticsList
	err := bambou.CurrentSession().FetchChildren(o, StatisticsIdentity, &list, info)
	return list, err
}

// StatisticsPolicies retrieves the list of child StatisticsPolicies of the Subnet
func (o *Subnet) StatisticsPolicies(info *bambou.FetchingInfo) (StatisticsPoliciesList, *bambou.Error) {

	var list StatisticsPoliciesList
	err := bambou.CurrentSession().FetchChildren(o, StatisticsPolicyIdentity, &list, info)
	return list, err
}

// CreateStatisticsPolicy creates a new child StatisticsPolicy under the Subnet
func (o *Subnet) CreateStatisticsPolicy(child *StatisticsPolicy) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the Subnet
func (o *Subnet) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
