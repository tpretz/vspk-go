/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VCenterDataCenterIdentity represents the Identity of the object
var VCenterDataCenterIdentity = bambou.Identity{
	Name:     "vcenterdatacenter",
	Category: "vcenterdatacenters",
}

// VCenterDataCentersList represents a list of VCenterDataCenters
type VCenterDataCentersList []*VCenterDataCenter

// VCenterDataCentersAncestor is the interface that an ancestor of a VCenterDataCenter must implement.
// An Ancestor is defined as an entity that has VCenterDataCenter as a descendant.
// An Ancestor can get a list of its child VCenterDataCenters, but not necessarily create one.
type VCenterDataCentersAncestor interface {
	VCenterDataCenters(*bambou.FetchingInfo) (VCenterDataCentersList, *bambou.Error)
}

// VCenterDataCentersParent is the interface that a parent of a VCenterDataCenter must implement.
// A Parent is defined as an entity that has VCenterDataCenter as a child.
// A Parent is an Ancestor which can create a VCenterDataCenter.
type VCenterDataCentersParent interface {
	VCenterDataCentersAncestor
	CreateVCenterDataCenter(*VCenterDataCenter) *bambou.Error
}

// VCenterDataCenter represents the model of a vcenterdatacenter
type VCenterDataCenter struct {
	ID                                     string `json:"ID,omitempty"`
	ParentID                               string `json:"parentID,omitempty"`
	ParentType                             string `json:"parentType,omitempty"`
	Owner                                  string `json:"owner,omitempty"`
	VRSConfigurationTimeLimit              int    `json:"VRSConfigurationTimeLimit,omitempty"`
	VRequireNuageMetadata                  bool   `json:"vRequireNuageMetadata"`
	Name                                   string `json:"name,omitempty"`
	ManagedObjectID                        string `json:"managedObjectID,omitempty"`
	LastUpdatedBy                          string `json:"lastUpdatedBy,omitempty"`
	DataDNS1                               string `json:"dataDNS1,omitempty"`
	DataDNS2                               string `json:"dataDNS2,omitempty"`
	DataGateway                            string `json:"dataGateway,omitempty"`
	DataNetworkPortgroup                   string `json:"dataNetworkPortgroup,omitempty"`
	DatapathSyncTimeout                    int    `json:"datapathSyncTimeout,omitempty"`
	SecondaryDataUplinkDHCPEnabled         bool   `json:"secondaryDataUplinkDHCPEnabled"`
	SecondaryDataUplinkEnabled             bool   `json:"secondaryDataUplinkEnabled"`
	SecondaryDataUplinkInterface           string `json:"secondaryDataUplinkInterface,omitempty"`
	SecondaryDataUplinkMTU                 int    `json:"secondaryDataUplinkMTU,omitempty"`
	SecondaryDataUplinkPrimaryController   string `json:"secondaryDataUplinkPrimaryController,omitempty"`
	SecondaryDataUplinkSecondaryController string `json:"secondaryDataUplinkSecondaryController,omitempty"`
	SecondaryDataUplinkUnderlayID          int    `json:"secondaryDataUplinkUnderlayID,omitempty"`
	SecondaryNuageController               string `json:"secondaryNuageController,omitempty"`
	DeletedFromVCenter                     bool   `json:"deletedFromVCenter"`
	MemorySizeInGB                         string `json:"memorySizeInGB,omitempty"`
	RemoteSyslogServerIP                   string `json:"remoteSyslogServerIP,omitempty"`
	RemoteSyslogServerPort                 int    `json:"remoteSyslogServerPort,omitempty"`
	RemoteSyslogServerType                 string `json:"remoteSyslogServerType,omitempty"`
	GenericSplitActivation                 bool   `json:"genericSplitActivation"`
	SeparateDataNetwork                    bool   `json:"separateDataNetwork"`
	Personality                            string `json:"personality,omitempty"`
	Description                            string `json:"description,omitempty"`
	DestinationMirrorPort                  string `json:"destinationMirrorPort,omitempty"`
	MetadataServerIP                       string `json:"metadataServerIP,omitempty"`
	MetadataServerListenPort               int    `json:"metadataServerListenPort,omitempty"`
	MetadataServerPort                     int    `json:"metadataServerPort,omitempty"`
	MetadataServiceEnabled                 bool   `json:"metadataServiceEnabled"`
	NetworkUplinkInterface                 string `json:"networkUplinkInterface,omitempty"`
	NetworkUplinkInterfaceGateway          string `json:"networkUplinkInterfaceGateway,omitempty"`
	NetworkUplinkInterfaceIp               string `json:"networkUplinkInterfaceIp,omitempty"`
	NetworkUplinkInterfaceNetmask          string `json:"networkUplinkInterfaceNetmask,omitempty"`
	RevertiveControllerEnabled             bool   `json:"revertiveControllerEnabled"`
	RevertiveTimer                         int    `json:"revertiveTimer,omitempty"`
	NfsLogServer                           string `json:"nfsLogServer,omitempty"`
	NfsMountPath                           string `json:"nfsMountPath,omitempty"`
	MgmtDNS1                               string `json:"mgmtDNS1,omitempty"`
	MgmtDNS2                               string `json:"mgmtDNS2,omitempty"`
	MgmtGateway                            string `json:"mgmtGateway,omitempty"`
	MgmtNetworkPortgroup                   string `json:"mgmtNetworkPortgroup,omitempty"`
	DhcpRelayServer                        string `json:"dhcpRelayServer,omitempty"`
	MirrorNetworkPortgroup                 string `json:"mirrorNetworkPortgroup,omitempty"`
	DisableGROOnDatapath                   bool   `json:"disableGROOnDatapath"`
	DisableLROOnDatapath                   bool   `json:"disableLROOnDatapath"`
	SiteId                                 string `json:"siteId,omitempty"`
	AllowDataDHCP                          bool   `json:"allowDataDHCP"`
	AllowMgmtDHCP                          bool   `json:"allowMgmtDHCP"`
	FlowEvictionThreshold                  int    `json:"flowEvictionThreshold,omitempty"`
	VmNetworkPortgroup                     string `json:"vmNetworkPortgroup,omitempty"`
	EnableVRSResourceReservation           bool   `json:"enableVRSResourceReservation"`
	EntityScope                            string `json:"entityScope,omitempty"`
	ConfiguredMetricsPushInterval          int    `json:"configuredMetricsPushInterval,omitempty"`
	PortgroupMetadata                      bool   `json:"portgroupMetadata"`
	NovaClientVersion                      int    `json:"novaClientVersion,omitempty"`
	NovaIdentityURLVersion                 string `json:"novaIdentityURLVersion,omitempty"`
	NovaMetadataServiceAuthUrl             string `json:"novaMetadataServiceAuthUrl,omitempty"`
	NovaMetadataServiceEndpoint            string `json:"novaMetadataServiceEndpoint,omitempty"`
	NovaMetadataServicePassword            string `json:"novaMetadataServicePassword,omitempty"`
	NovaMetadataServiceTenant              string `json:"novaMetadataServiceTenant,omitempty"`
	NovaMetadataServiceUsername            string `json:"novaMetadataServiceUsername,omitempty"`
	NovaMetadataSharedSecret               string `json:"novaMetadataSharedSecret,omitempty"`
	NovaOSKeystoneUsername                 string `json:"novaOSKeystoneUsername,omitempty"`
	NovaProjectDomainName                  string `json:"novaProjectDomainName,omitempty"`
	NovaProjectName                        string `json:"novaProjectName,omitempty"`
	NovaRegionName                         string `json:"novaRegionName,omitempty"`
	NovaUserDomainName                     string `json:"novaUserDomainName,omitempty"`
	UpgradePackagePassword                 string `json:"upgradePackagePassword,omitempty"`
	UpgradePackageURL                      string `json:"upgradePackageURL,omitempty"`
	UpgradePackageUsername                 string `json:"upgradePackageUsername,omitempty"`
	UpgradeScriptTimeLimit                 int    `json:"upgradeScriptTimeLimit,omitempty"`
	CpuCount                               string `json:"cpuCount,omitempty"`
	PrimaryDataUplinkUnderlayID            int    `json:"primaryDataUplinkUnderlayID,omitempty"`
	PrimaryNuageController                 string `json:"primaryNuageController,omitempty"`
	VrsPassword                            string `json:"vrsPassword,omitempty"`
	VrsUserName                            string `json:"vrsUserName,omitempty"`
	AssociatedVCenterID                    string `json:"associatedVCenterID,omitempty"`
	StaticRoute                            string `json:"staticRoute,omitempty"`
	StaticRouteGateway                     string `json:"staticRouteGateway,omitempty"`
	StaticRouteNetmask                     string `json:"staticRouteNetmask,omitempty"`
	NtpServer1                             string `json:"ntpServer1,omitempty"`
	NtpServer2                             string `json:"ntpServer2,omitempty"`
	Mtu                                    int    `json:"mtu,omitempty"`
	MultiVMSsupport                        bool   `json:"multiVMSsupport"`
	MulticastReceiveInterface              string `json:"multicastReceiveInterface,omitempty"`
	MulticastReceiveInterfaceIP            string `json:"multicastReceiveInterfaceIP,omitempty"`
	MulticastReceiveInterfaceNetmask       string `json:"multicastReceiveInterfaceNetmask,omitempty"`
	MulticastReceiveRange                  string `json:"multicastReceiveRange,omitempty"`
	MulticastSendInterface                 string `json:"multicastSendInterface,omitempty"`
	MulticastSendInterfaceIP               string `json:"multicastSendInterfaceIP,omitempty"`
	MulticastSendInterfaceNetmask          string `json:"multicastSendInterfaceNetmask,omitempty"`
	MulticastSourcePortgroup               string `json:"multicastSourcePortgroup,omitempty"`
	CustomizedScriptURL                    string `json:"customizedScriptURL,omitempty"`
	OvfURL                                 string `json:"ovfURL,omitempty"`
	AvrsEnabled                            bool   `json:"avrsEnabled"`
	AvrsProfile                            string `json:"avrsProfile,omitempty"`
	ExternalID                             string `json:"externalID,omitempty"`
}

// NewVCenterDataCenter returns a new *VCenterDataCenter
func NewVCenterDataCenter() *VCenterDataCenter {

	return &VCenterDataCenter{
		SecondaryDataUplinkDHCPEnabled: false,
		SecondaryDataUplinkEnabled:     false,
		SecondaryDataUplinkMTU:         1500,
		SecondaryDataUplinkUnderlayID:  1,
		MemorySizeInGB:                 "DEFAULT_4",
		RemoteSyslogServerPort:         514,
		RemoteSyslogServerType:         "NONE",
		DestinationMirrorPort:          "no_mirror",
		RevertiveControllerEnabled:     false,
		RevertiveTimer:                 300,
		DisableGROOnDatapath:           false,
		DisableLROOnDatapath:           false,
		EnableVRSResourceReservation:   false,
		ConfiguredMetricsPushInterval:  60,
		CpuCount:                       "DEFAULT_2",
		PrimaryDataUplinkUnderlayID:    0,
		AvrsEnabled:                    false,
		AvrsProfile:                    "AVRS_25G",
	}
}

// Identity returns the Identity of the object.
func (o *VCenterDataCenter) Identity() bambou.Identity {

	return VCenterDataCenterIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VCenterDataCenter) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VCenterDataCenter) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VCenterDataCenter from the server
func (o *VCenterDataCenter) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VCenterDataCenter into the server
func (o *VCenterDataCenter) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VCenterDataCenter from the server
func (o *VCenterDataCenter) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// VCenterClusters retrieves the list of child VCenterClusters of the VCenterDataCenter
func (o *VCenterDataCenter) VCenterClusters(info *bambou.FetchingInfo) (VCenterClustersList, *bambou.Error) {

	var list VCenterClustersList
	err := bambou.CurrentSession().FetchChildren(o, VCenterClusterIdentity, &list, info)
	return list, err
}

// CreateVCenterCluster creates a new child VCenterCluster under the VCenterDataCenter
func (o *VCenterDataCenter) CreateVCenterCluster(child *VCenterCluster) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VCenterHypervisors retrieves the list of child VCenterHypervisors of the VCenterDataCenter
func (o *VCenterDataCenter) VCenterHypervisors(info *bambou.FetchingInfo) (VCenterHypervisorsList, *bambou.Error) {

	var list VCenterHypervisorsList
	err := bambou.CurrentSession().FetchChildren(o, VCenterHypervisorIdentity, &list, info)
	return list, err
}

// CreateVCenterHypervisor creates a new child VCenterHypervisor under the VCenterDataCenter
func (o *VCenterDataCenter) CreateVCenterHypervisor(child *VCenterHypervisor) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the VCenterDataCenter
func (o *VCenterDataCenter) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VCenterDataCenter
func (o *VCenterDataCenter) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VCenterDataCenter
func (o *VCenterDataCenter) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VCenterDataCenter
func (o *VCenterDataCenter) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VRSAddressRanges retrieves the list of child VRSAddressRanges of the VCenterDataCenter
func (o *VCenterDataCenter) VRSAddressRanges(info *bambou.FetchingInfo) (VRSAddressRangesList, *bambou.Error) {

	var list VRSAddressRangesList
	err := bambou.CurrentSession().FetchChildren(o, VRSAddressRangeIdentity, &list, info)
	return list, err
}

// CreateVRSAddressRange creates a new child VRSAddressRange under the VCenterDataCenter
func (o *VCenterDataCenter) CreateVRSAddressRange(child *VRSAddressRange) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VRSRedeploymentpolicies retrieves the list of child VRSRedeploymentpolicies of the VCenterDataCenter
func (o *VCenterDataCenter) VRSRedeploymentpolicies(info *bambou.FetchingInfo) (VRSRedeploymentpoliciesList, *bambou.Error) {

	var list VRSRedeploymentpoliciesList
	err := bambou.CurrentSession().FetchChildren(o, VRSRedeploymentpolicyIdentity, &list, info)
	return list, err
}

// CreateVRSRedeploymentpolicy creates a new child VRSRedeploymentpolicy under the VCenterDataCenter
func (o *VCenterDataCenter) CreateVRSRedeploymentpolicy(child *VRSRedeploymentpolicy) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// AutoDiscoverClusters retrieves the list of child AutoDiscoverClusters of the VCenterDataCenter
func (o *VCenterDataCenter) AutoDiscoverClusters(info *bambou.FetchingInfo) (AutoDiscoverClustersList, *bambou.Error) {

	var list AutoDiscoverClustersList
	err := bambou.CurrentSession().FetchChildren(o, AutoDiscoverClusterIdentity, &list, info)
	return list, err
}

// AutoDiscoverHypervisorFromClusters retrieves the list of child AutoDiscoverHypervisorFromClusters of the VCenterDataCenter
func (o *VCenterDataCenter) AutoDiscoverHypervisorFromClusters(info *bambou.FetchingInfo) (AutoDiscoverHypervisorFromClustersList, *bambou.Error) {

	var list AutoDiscoverHypervisorFromClustersList
	err := bambou.CurrentSession().FetchChildren(o, AutoDiscoverHypervisorFromClusterIdentity, &list, info)
	return list, err
}
