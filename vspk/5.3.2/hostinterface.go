/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// HostInterfaceIdentity represents the Identity of the object
var HostInterfaceIdentity = bambou.Identity{
	Name:     "hostinterface",
	Category: "hostinterfaces",
}

// HostInterfacesList represents a list of HostInterfaces
type HostInterfacesList []*HostInterface

// HostInterfacesAncestor is the interface that an ancestor of a HostInterface must implement.
// An Ancestor is defined as an entity that has HostInterface as a descendant.
// An Ancestor can get a list of its child HostInterfaces, but not necessarily create one.
type HostInterfacesAncestor interface {
	HostInterfaces(*bambou.FetchingInfo) (HostInterfacesList, *bambou.Error)
}

// HostInterfacesParent is the interface that a parent of a HostInterface must implement.
// A Parent is defined as an entity that has HostInterface as a child.
// A Parent is an Ancestor which can create a HostInterface.
type HostInterfacesParent interface {
	HostInterfacesAncestor
	CreateHostInterface(*HostInterface) *bambou.Error
}

// HostInterface represents the model of a hostinterface
type HostInterface struct {
	ID                          string `json:"ID,omitempty"`
	ParentID                    string `json:"parentID,omitempty"`
	ParentType                  string `json:"parentType,omitempty"`
	Owner                       string `json:"owner,omitempty"`
	MAC                         string `json:"MAC,omitempty"`
	IPAddress                   string `json:"IPAddress,omitempty"`
	VPortID                     string `json:"VPortID,omitempty"`
	VPortName                   string `json:"VPortName,omitempty"`
	IPv6Address                 string `json:"IPv6Address,omitempty"`
	IPv6Gateway                 string `json:"IPv6Gateway,omitempty"`
	Name                        string `json:"name,omitempty"`
	LastUpdatedBy               string `json:"lastUpdatedBy,omitempty"`
	Gateway                     string `json:"gateway,omitempty"`
	Netmask                     string `json:"netmask,omitempty"`
	NetworkName                 string `json:"networkName,omitempty"`
	TierID                      string `json:"tierID,omitempty"`
	EntityScope                 string `json:"entityScope,omitempty"`
	PolicyDecisionID            string `json:"policyDecisionID,omitempty"`
	DomainID                    string `json:"domainID,omitempty"`
	DomainName                  string `json:"domainName,omitempty"`
	ZoneID                      string `json:"zoneID,omitempty"`
	ZoneName                    string `json:"zoneName,omitempty"`
	AssociatedFloatingIPAddress string `json:"associatedFloatingIPAddress,omitempty"`
	AttachedNetworkID           string `json:"attachedNetworkID,omitempty"`
	AttachedNetworkType         string `json:"attachedNetworkType,omitempty"`
	ExternalID                  string `json:"externalID,omitempty"`
}

// NewHostInterface returns a new *HostInterface
func NewHostInterface() *HostInterface {

	return &HostInterface{}
}

// Identity returns the Identity of the object.
func (o *HostInterface) Identity() bambou.Identity {

	return HostInterfaceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *HostInterface) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *HostInterface) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the HostInterface from the server
func (o *HostInterface) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the HostInterface into the server
func (o *HostInterface) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the HostInterface from the server
func (o *HostInterface) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// TCAs retrieves the list of child TCAs of the HostInterface
func (o *HostInterface) TCAs(info *bambou.FetchingInfo) (TCAsList, *bambou.Error) {

	var list TCAsList
	err := bambou.CurrentSession().FetchChildren(o, TCAIdentity, &list, info)
	return list, err
}

// RedirectionTargets retrieves the list of child RedirectionTargets of the HostInterface
func (o *HostInterface) RedirectionTargets(info *bambou.FetchingInfo) (RedirectionTargetsList, *bambou.Error) {

	var list RedirectionTargetsList
	err := bambou.CurrentSession().FetchChildren(o, RedirectionTargetIdentity, &list, info)
	return list, err
}

// Metadatas retrieves the list of child Metadatas of the HostInterface
func (o *HostInterface) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the HostInterface
func (o *HostInterface) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// DHCPOptions retrieves the list of child DHCPOptions of the HostInterface
func (o *HostInterface) DHCPOptions(info *bambou.FetchingInfo) (DHCPOptionsList, *bambou.Error) {

	var list DHCPOptionsList
	err := bambou.CurrentSession().FetchChildren(o, DHCPOptionIdentity, &list, info)
	return list, err
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the HostInterface
func (o *HostInterface) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the HostInterface
func (o *HostInterface) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// PolicyDecisions retrieves the list of child PolicyDecisions of the HostInterface
func (o *HostInterface) PolicyDecisions(info *bambou.FetchingInfo) (PolicyDecisionsList, *bambou.Error) {

	var list PolicyDecisionsList
	err := bambou.CurrentSession().FetchChildren(o, PolicyDecisionIdentity, &list, info)
	return list, err
}

// PolicyGroups retrieves the list of child PolicyGroups of the HostInterface
func (o *HostInterface) PolicyGroups(info *bambou.FetchingInfo) (PolicyGroupsList, *bambou.Error) {

	var list PolicyGroupsList
	err := bambou.CurrentSession().FetchChildren(o, PolicyGroupIdentity, &list, info)
	return list, err
}

// QOSs retrieves the list of child QOSs of the HostInterface
func (o *HostInterface) QOSs(info *bambou.FetchingInfo) (QOSsList, *bambou.Error) {

	var list QOSsList
	err := bambou.CurrentSession().FetchChildren(o, QOSIdentity, &list, info)
	return list, err
}

// CreateQOS creates a new child QOS under the HostInterface
func (o *HostInterface) CreateQOS(child *QOS) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// StaticRoutes retrieves the list of child StaticRoutes of the HostInterface
func (o *HostInterface) StaticRoutes(info *bambou.FetchingInfo) (StaticRoutesList, *bambou.Error) {

	var list StaticRoutesList
	err := bambou.CurrentSession().FetchChildren(o, StaticRouteIdentity, &list, info)
	return list, err
}

// Statistics retrieves the list of child Statistics of the HostInterface
func (o *HostInterface) Statistics(info *bambou.FetchingInfo) (StatisticsList, *bambou.Error) {

	var list StatisticsList
	err := bambou.CurrentSession().FetchChildren(o, StatisticsIdentity, &list, info)
	return list, err
}

// MultiCastChannelMaps retrieves the list of child MultiCastChannelMaps of the HostInterface
func (o *HostInterface) MultiCastChannelMaps(info *bambou.FetchingInfo) (MultiCastChannelMapsList, *bambou.Error) {

	var list MultiCastChannelMapsList
	err := bambou.CurrentSession().FetchChildren(o, MultiCastChannelMapIdentity, &list, info)
	return list, err
}

// EventLogs retrieves the list of child EventLogs of the HostInterface
func (o *HostInterface) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
