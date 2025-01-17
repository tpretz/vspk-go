/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// SharedNetworkResourceIdentity represents the Identity of the object
var SharedNetworkResourceIdentity = bambou.Identity{
	Name:     "sharednetworkresource",
	Category: "sharednetworkresources",
}

// SharedNetworkResourcesList represents a list of SharedNetworkResources
type SharedNetworkResourcesList []*SharedNetworkResource

// SharedNetworkResourcesAncestor is the interface that an ancestor of a SharedNetworkResource must implement.
// An Ancestor is defined as an entity that has SharedNetworkResource as a descendant.
// An Ancestor can get a list of its child SharedNetworkResources, but not necessarily create one.
type SharedNetworkResourcesAncestor interface {
	SharedNetworkResources(*bambou.FetchingInfo) (SharedNetworkResourcesList, *bambou.Error)
}

// SharedNetworkResourcesParent is the interface that a parent of a SharedNetworkResource must implement.
// A Parent is defined as an entity that has SharedNetworkResource as a child.
// A Parent is an Ancestor which can create a SharedNetworkResource.
type SharedNetworkResourcesParent interface {
	SharedNetworkResourcesAncestor
	CreateSharedNetworkResource(*SharedNetworkResource) *bambou.Error
}

// SharedNetworkResource represents the model of a sharednetworkresource
type SharedNetworkResource struct {
	ID                          string `json:"ID,omitempty"`
	ParentID                    string `json:"parentID,omitempty"`
	ParentType                  string `json:"parentType,omitempty"`
	Owner                       string `json:"owner,omitempty"`
	ECMPCount                   int    `json:"ECMPCount,omitempty"`
	DHCPManaged                 bool   `json:"DHCPManaged"`
	BackHaulRouteDistinguisher  string `json:"backHaulRouteDistinguisher,omitempty"`
	BackHaulRouteTarget         string `json:"backHaulRouteTarget,omitempty"`
	BackHaulVNID                int    `json:"backHaulVNID,omitempty"`
	Name                        string `json:"name,omitempty"`
	LastUpdatedBy               string `json:"lastUpdatedBy,omitempty"`
	Gateway                     string `json:"gateway,omitempty"`
	GatewayMACAddress           string `json:"gatewayMACAddress,omitempty"`
	AccessRestrictionEnabled    bool   `json:"accessRestrictionEnabled"`
	Address                     string `json:"address,omitempty"`
	PermittedActionType         string `json:"permittedActionType,omitempty"`
	Description                 string `json:"description,omitempty"`
	Netmask                     string `json:"netmask,omitempty"`
	SharedResourceParentID      string `json:"sharedResourceParentID,omitempty"`
	VnID                        int    `json:"vnID,omitempty"`
	Underlay                    bool   `json:"underlay"`
	EnterpriseID                string `json:"enterpriseID,omitempty"`
	EntityScope                 string `json:"entityScope,omitempty"`
	DomainRouteDistinguisher    string `json:"domainRouteDistinguisher,omitempty"`
	DomainRouteTarget           string `json:"domainRouteTarget,omitempty"`
	UplinkGWVlanAttachmentID    string `json:"uplinkGWVlanAttachmentID,omitempty"`
	UplinkInterfaceIP           string `json:"uplinkInterfaceIP,omitempty"`
	UplinkInterfaceMAC          string `json:"uplinkInterfaceMAC,omitempty"`
	UplinkVPortName             string `json:"uplinkVPortName,omitempty"`
	UseGlobalMAC                string `json:"useGlobalMAC,omitempty"`
	AssociatedPATMapperID       string `json:"associatedPATMapperID,omitempty"`
	SubnetRouteDistinguisher    string `json:"subnetRouteDistinguisher,omitempty"`
	SubnetRouteTarget           string `json:"subnetRouteTarget,omitempty"`
	ExternalID                  string `json:"externalID,omitempty"`
	DynamicPATAllocationEnabled bool   `json:"dynamicPATAllocationEnabled"`
	Type                        string `json:"type,omitempty"`
}

// NewSharedNetworkResource returns a new *SharedNetworkResource
func NewSharedNetworkResource() *SharedNetworkResource {

	return &SharedNetworkResource{}
}

// Identity returns the Identity of the object.
func (o *SharedNetworkResource) Identity() bambou.Identity {

	return SharedNetworkResourceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SharedNetworkResource) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SharedNetworkResource) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the SharedNetworkResource from the server
func (o *SharedNetworkResource) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the SharedNetworkResource into the server
func (o *SharedNetworkResource) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the SharedNetworkResource from the server
func (o *SharedNetworkResource) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// PATIPEntries retrieves the list of child PATIPEntries of the SharedNetworkResource
func (o *SharedNetworkResource) PATIPEntries(info *bambou.FetchingInfo) (PATIPEntriesList, *bambou.Error) {

	var list PATIPEntriesList
	err := bambou.CurrentSession().FetchChildren(o, PATIPEntryIdentity, &list, info)
	return list, err
}

// CreatePATIPEntry creates a new child PATIPEntry under the SharedNetworkResource
func (o *SharedNetworkResource) CreatePATIPEntry(child *PATIPEntry) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// AddressRanges retrieves the list of child AddressRanges of the SharedNetworkResource
func (o *SharedNetworkResource) AddressRanges(info *bambou.FetchingInfo) (AddressRangesList, *bambou.Error) {

	var list AddressRangesList
	err := bambou.CurrentSession().FetchChildren(o, AddressRangeIdentity, &list, info)
	return list, err
}

// CreateAddressRange creates a new child AddressRange under the SharedNetworkResource
func (o *SharedNetworkResource) CreateAddressRange(child *AddressRange) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the SharedNetworkResource
func (o *SharedNetworkResource) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the SharedNetworkResource
func (o *SharedNetworkResource) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// DHCPOptions retrieves the list of child DHCPOptions of the SharedNetworkResource
func (o *SharedNetworkResource) DHCPOptions(info *bambou.FetchingInfo) (DHCPOptionsList, *bambou.Error) {

	var list DHCPOptionsList
	err := bambou.CurrentSession().FetchChildren(o, DHCPOptionIdentity, &list, info)
	return list, err
}

// CreateDHCPOption creates a new child DHCPOption under the SharedNetworkResource
func (o *SharedNetworkResource) CreateDHCPOption(child *DHCPOption) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the SharedNetworkResource
func (o *SharedNetworkResource) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the SharedNetworkResource
func (o *SharedNetworkResource) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EnterprisePermissions retrieves the list of child EnterprisePermissions of the SharedNetworkResource
func (o *SharedNetworkResource) EnterprisePermissions(info *bambou.FetchingInfo) (EnterprisePermissionsList, *bambou.Error) {

	var list EnterprisePermissionsList
	err := bambou.CurrentSession().FetchChildren(o, EnterprisePermissionIdentity, &list, info)
	return list, err
}

// CreateEnterprisePermission creates a new child EnterprisePermission under the SharedNetworkResource
func (o *SharedNetworkResource) CreateEnterprisePermission(child *EnterprisePermission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VPNConnections retrieves the list of child VPNConnections of the SharedNetworkResource
func (o *SharedNetworkResource) VPNConnections(info *bambou.FetchingInfo) (VPNConnectionsList, *bambou.Error) {

	var list VPNConnectionsList
	err := bambou.CurrentSession().FetchChildren(o, VPNConnectionIdentity, &list, info)
	return list, err
}

// CreateVPNConnection creates a new child VPNConnection under the SharedNetworkResource
func (o *SharedNetworkResource) CreateVPNConnection(child *VPNConnection) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// StaticRoutes retrieves the list of child StaticRoutes of the SharedNetworkResource
func (o *SharedNetworkResource) StaticRoutes(info *bambou.FetchingInfo) (StaticRoutesList, *bambou.Error) {

	var list StaticRoutesList
	err := bambou.CurrentSession().FetchChildren(o, StaticRouteIdentity, &list, info)
	return list, err
}

// CreateStaticRoute creates a new child StaticRoute under the SharedNetworkResource
func (o *SharedNetworkResource) CreateStaticRoute(child *StaticRoute) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
