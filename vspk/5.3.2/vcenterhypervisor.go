/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VCenterHypervisorIdentity represents the Identity of the object
var VCenterHypervisorIdentity = bambou.Identity{
	Name:     "vcenterhypervisor",
	Category: "vcenterhypervisors",
}

// VCenterHypervisorsList represents a list of VCenterHypervisors
type VCenterHypervisorsList []*VCenterHypervisor

// VCenterHypervisorsAncestor is the interface that an ancestor of a VCenterHypervisor must implement.
// An Ancestor is defined as an entity that has VCenterHypervisor as a descendant.
// An Ancestor can get a list of its child VCenterHypervisors, but not necessarily create one.
type VCenterHypervisorsAncestor interface {
	VCenterHypervisors(*bambou.FetchingInfo) (VCenterHypervisorsList, *bambou.Error)
}

// VCenterHypervisorsParent is the interface that a parent of a VCenterHypervisor must implement.
// A Parent is defined as an entity that has VCenterHypervisor as a child.
// A Parent is an Ancestor which can create a VCenterHypervisor.
type VCenterHypervisorsParent interface {
	VCenterHypervisorsAncestor
	CreateVCenterHypervisor(*VCenterHypervisor) *bambou.Error
}

// VCenterHypervisor represents the model of a vcenterhypervisor
type VCenterHypervisor struct {
	ID                                        string        `json:"ID,omitempty"`
	ParentID                                  string        `json:"parentID,omitempty"`
	ParentType                                string        `json:"parentType,omitempty"`
	Owner                                     string        `json:"owner,omitempty"`
	VCenterIP                                 string        `json:"vCenterIP,omitempty"`
	VCenterPassword                           string        `json:"vCenterPassword,omitempty"`
	VCenterUser                               string        `json:"vCenterUser,omitempty"`
	VRSAgentMOID                              string        `json:"VRSAgentMOID,omitempty"`
	VRSAgentName                              string        `json:"VRSAgentName,omitempty"`
	VRSConfigurationTimeLimit                 int           `json:"VRSConfigurationTimeLimit,omitempty"`
	VRSMetricsID                              string        `json:"VRSMetricsID,omitempty"`
	VRSMgmtHostname                           string        `json:"VRSMgmtHostname,omitempty"`
	VRSState                                  string        `json:"VRSState,omitempty"`
	VRequireNuageMetadata                     bool          `json:"vRequireNuageMetadata"`
	Name                                      string        `json:"name,omitempty"`
	ManageVRSAvailability                     bool          `json:"manageVRSAvailability"`
	ManagedObjectID                           string        `json:"managedObjectID,omitempty"`
	LastUpdatedBy                             string        `json:"lastUpdatedBy,omitempty"`
	LastVRSDeployedDate                       float64       `json:"lastVRSDeployedDate,omitempty"`
	DataDNS1                                  string        `json:"dataDNS1,omitempty"`
	DataDNS2                                  string        `json:"dataDNS2,omitempty"`
	DataGateway                               string        `json:"dataGateway,omitempty"`
	DataIPAddress                             string        `json:"dataIPAddress,omitempty"`
	DataNetmask                               string        `json:"dataNetmask,omitempty"`
	DataNetworkPortgroup                      string        `json:"dataNetworkPortgroup,omitempty"`
	DatapathSyncTimeout                       int           `json:"datapathSyncTimeout,omitempty"`
	Scope                                     bool          `json:"scope"`
	SecondaryDataUplinkDHCPEnabled            bool          `json:"secondaryDataUplinkDHCPEnabled"`
	SecondaryDataUplinkEnabled                bool          `json:"secondaryDataUplinkEnabled"`
	SecondaryDataUplinkIP                     string        `json:"secondaryDataUplinkIP,omitempty"`
	SecondaryDataUplinkInterface              string        `json:"secondaryDataUplinkInterface,omitempty"`
	SecondaryDataUplinkMTU                    int           `json:"secondaryDataUplinkMTU,omitempty"`
	SecondaryDataUplinkNetmask                string        `json:"secondaryDataUplinkNetmask,omitempty"`
	SecondaryDataUplinkPrimaryController      string        `json:"secondaryDataUplinkPrimaryController,omitempty"`
	SecondaryDataUplinkSecondaryController    string        `json:"secondaryDataUplinkSecondaryController,omitempty"`
	SecondaryDataUplinkUnderlayID             int           `json:"secondaryDataUplinkUnderlayID,omitempty"`
	SecondaryDataUplinkVDFControlVLAN         int           `json:"secondaryDataUplinkVDFControlVLAN,omitempty"`
	SecondaryNuageController                  string        `json:"secondaryNuageController,omitempty"`
	MemorySizeInGB                            string        `json:"memorySizeInGB,omitempty"`
	RemoteSyslogServerIP                      string        `json:"remoteSyslogServerIP,omitempty"`
	RemoteSyslogServerPort                    int           `json:"remoteSyslogServerPort,omitempty"`
	RemoteSyslogServerType                    string        `json:"remoteSyslogServerType,omitempty"`
	RemovedFromVCenterInventory               bool          `json:"removedFromVCenterInventory"`
	GenericSplitActivation                    bool          `json:"genericSplitActivation"`
	SeparateDataNetwork                       bool          `json:"separateDataNetwork"`
	DeploymentCount                           int           `json:"deploymentCount,omitempty"`
	Personality                               string        `json:"personality,omitempty"`
	Description                               string        `json:"description,omitempty"`
	DestinationMirrorPort                     string        `json:"destinationMirrorPort,omitempty"`
	MetadataServerIP                          string        `json:"metadataServerIP,omitempty"`
	MetadataServerListenPort                  int           `json:"metadataServerListenPort,omitempty"`
	MetadataServerPort                        int           `json:"metadataServerPort,omitempty"`
	MetadataServiceEnabled                    bool          `json:"metadataServiceEnabled"`
	NetworkUplinkInterface                    string        `json:"networkUplinkInterface,omitempty"`
	NetworkUplinkInterfaceGateway             string        `json:"networkUplinkInterfaceGateway,omitempty"`
	NetworkUplinkInterfaceIp                  string        `json:"networkUplinkInterfaceIp,omitempty"`
	NetworkUplinkInterfaceNetmask             string        `json:"networkUplinkInterfaceNetmask,omitempty"`
	RevertiveControllerEnabled                bool          `json:"revertiveControllerEnabled"`
	RevertiveTimer                            int           `json:"revertiveTimer,omitempty"`
	NfsLogServer                              string        `json:"nfsLogServer,omitempty"`
	NfsMountPath                              string        `json:"nfsMountPath,omitempty"`
	MgmtDNS1                                  string        `json:"mgmtDNS1,omitempty"`
	MgmtDNS2                                  string        `json:"mgmtDNS2,omitempty"`
	MgmtGateway                               string        `json:"mgmtGateway,omitempty"`
	MgmtIPAddress                             string        `json:"mgmtIPAddress,omitempty"`
	MgmtNetmask                               string        `json:"mgmtNetmask,omitempty"`
	MgmtNetworkPortgroup                      string        `json:"mgmtNetworkPortgroup,omitempty"`
	DhcpRelayServer                           string        `json:"dhcpRelayServer,omitempty"`
	MirrorNetworkPortgroup                    string        `json:"mirrorNetworkPortgroup,omitempty"`
	DisableGROOnDatapath                      bool          `json:"disableGROOnDatapath"`
	DisableLROOnDatapath                      bool          `json:"disableLROOnDatapath"`
	SiteId                                    string        `json:"siteId,omitempty"`
	AllowDataDHCP                             bool          `json:"allowDataDHCP"`
	AllowMgmtDHCP                             bool          `json:"allowMgmtDHCP"`
	FlowEvictionThreshold                     int           `json:"flowEvictionThreshold,omitempty"`
	VmNetworkPortgroup                        string        `json:"vmNetworkPortgroup,omitempty"`
	EnableVRSResourceReservation              bool          `json:"enableVRSResourceReservation"`
	EntityScope                               string        `json:"entityScope,omitempty"`
	ConfiguredMetricsPushInterval             int           `json:"configuredMetricsPushInterval,omitempty"`
	ToolboxDeploymentMode                     bool          `json:"toolboxDeploymentMode"`
	ToolboxGroup                              string        `json:"toolboxGroup,omitempty"`
	ToolboxIP                                 string        `json:"toolboxIP,omitempty"`
	ToolboxPassword                           string        `json:"toolboxPassword,omitempty"`
	ToolboxUserName                           string        `json:"toolboxUserName,omitempty"`
	PortgroupMetadata                         bool          `json:"portgroupMetadata"`
	NovaClientVersion                         int           `json:"novaClientVersion,omitempty"`
	NovaIdentityURLVersion                    string        `json:"novaIdentityURLVersion,omitempty"`
	NovaMetadataServiceAuthUrl                string        `json:"novaMetadataServiceAuthUrl,omitempty"`
	NovaMetadataServiceEndpoint               string        `json:"novaMetadataServiceEndpoint,omitempty"`
	NovaMetadataServicePassword               string        `json:"novaMetadataServicePassword,omitempty"`
	NovaMetadataServiceTenant                 string        `json:"novaMetadataServiceTenant,omitempty"`
	NovaMetadataServiceUsername               string        `json:"novaMetadataServiceUsername,omitempty"`
	NovaMetadataSharedSecret                  string        `json:"novaMetadataSharedSecret,omitempty"`
	NovaOSKeystoneUsername                    string        `json:"novaOSKeystoneUsername,omitempty"`
	NovaProjectDomainName                     string        `json:"novaProjectDomainName,omitempty"`
	NovaProjectName                           string        `json:"novaProjectName,omitempty"`
	NovaRegionName                            string        `json:"novaRegionName,omitempty"`
	NovaUserDomainName                        string        `json:"novaUserDomainName,omitempty"`
	UpgradePackagePassword                    string        `json:"upgradePackagePassword,omitempty"`
	UpgradePackageURL                         string        `json:"upgradePackageURL,omitempty"`
	UpgradePackageUsername                    string        `json:"upgradePackageUsername,omitempty"`
	UpgradeScriptTimeLimit                    int           `json:"upgradeScriptTimeLimit,omitempty"`
	UpgradeStatus                             string        `json:"upgradeStatus,omitempty"`
	UpgradeTimedout                           bool          `json:"upgradeTimedout"`
	CpuCount                                  string        `json:"cpuCount,omitempty"`
	PrimaryDataUplinkUnderlayID               int           `json:"primaryDataUplinkUnderlayID,omitempty"`
	PrimaryDataUplinkVDFControlVLAN           int           `json:"primaryDataUplinkVDFControlVLAN,omitempty"`
	PrimaryNuageController                    string        `json:"primaryNuageController,omitempty"`
	VrsId                                     string        `json:"vrsId,omitempty"`
	VrsMarkedAsAvailable                      bool          `json:"vrsMarkedAsAvailable"`
	VrsPassword                               string        `json:"vrsPassword,omitempty"`
	VrsUserName                               string        `json:"vrsUserName,omitempty"`
	StaticRoute                               string        `json:"staticRoute,omitempty"`
	StaticRouteGateway                        string        `json:"staticRouteGateway,omitempty"`
	StaticRouteNetmask                        string        `json:"staticRouteNetmask,omitempty"`
	NtpServer1                                string        `json:"ntpServer1,omitempty"`
	NtpServer2                                string        `json:"ntpServer2,omitempty"`
	Mtu                                       int           `json:"mtu,omitempty"`
	SuccessfullyAppliedUpgradePackagePassword string        `json:"successfullyAppliedUpgradePackagePassword,omitempty"`
	SuccessfullyAppliedUpgradePackageURL      string        `json:"successfullyAppliedUpgradePackageURL,omitempty"`
	SuccessfullyAppliedUpgradePackageUsername string        `json:"successfullyAppliedUpgradePackageUsername,omitempty"`
	SuccessfullyAppliedVersion                string        `json:"successfullyAppliedVersion,omitempty"`
	MultiVMSsupport                           bool          `json:"multiVMSsupport"`
	MulticastReceiveInterface                 string        `json:"multicastReceiveInterface,omitempty"`
	MulticastReceiveInterfaceIP               string        `json:"multicastReceiveInterfaceIP,omitempty"`
	MulticastReceiveInterfaceNetmask          string        `json:"multicastReceiveInterfaceNetmask,omitempty"`
	MulticastReceiveRange                     string        `json:"multicastReceiveRange,omitempty"`
	MulticastSendInterface                    string        `json:"multicastSendInterface,omitempty"`
	MulticastSendInterfaceIP                  string        `json:"multicastSendInterfaceIP,omitempty"`
	MulticastSendInterfaceNetmask             string        `json:"multicastSendInterfaceNetmask,omitempty"`
	MulticastSourcePortgroup                  string        `json:"multicastSourcePortgroup,omitempty"`
	CustomizedScriptURL                       string        `json:"customizedScriptURL,omitempty"`
	AvailableNetworks                         []interface{} `json:"availableNetworks,omitempty"`
	OvfURL                                    string        `json:"ovfURL,omitempty"`
	AvrsEnabled                               bool          `json:"avrsEnabled"`
	AvrsProfile                               string        `json:"avrsProfile,omitempty"`
	ExternalID                                string        `json:"externalID,omitempty"`
	HypervisorIP                              string        `json:"hypervisorIP,omitempty"`
	HypervisorPassword                        string        `json:"hypervisorPassword,omitempty"`
	HypervisorUser                            string        `json:"hypervisorUser,omitempty"`
}

// NewVCenterHypervisor returns a new *VCenterHypervisor
func NewVCenterHypervisor() *VCenterHypervisor {

	return &VCenterHypervisor{
		VRSState:                          "NOT_DEPLOYED",
		ManageVRSAvailability:             false,
		SecondaryDataUplinkDHCPEnabled:    false,
		SecondaryDataUplinkEnabled:        false,
		SecondaryDataUplinkMTU:            1500,
		SecondaryDataUplinkUnderlayID:     1,
		SecondaryDataUplinkVDFControlVLAN: 0,
		MemorySizeInGB:                    "DEFAULT_4",
		RemoteSyslogServerPort:            514,
		RemoteSyslogServerType:            "NONE",
		Personality:                       "VRS",
		DestinationMirrorPort:             "no_mirror",
		RevertiveControllerEnabled:        false,
		RevertiveTimer:                    300,
		DisableGROOnDatapath:              false,
		DisableLROOnDatapath:              false,
		EnableVRSResourceReservation:      false,
		ConfiguredMetricsPushInterval:     60,
		CpuCount:                          "DEFAULT_2",
		PrimaryDataUplinkUnderlayID:       0,
		PrimaryDataUplinkVDFControlVLAN:   0,
		VrsMarkedAsAvailable:              false,
		AvrsEnabled:                       false,
		AvrsProfile:                       "AVRS_25G",
	}
}

// Identity returns the Identity of the object.
func (o *VCenterHypervisor) Identity() bambou.Identity {

	return VCenterHypervisorIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VCenterHypervisor) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VCenterHypervisor) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VCenterHypervisor from the server
func (o *VCenterHypervisor) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VCenterHypervisor into the server
func (o *VCenterHypervisor) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VCenterHypervisor from the server
func (o *VCenterHypervisor) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the VCenterHypervisor
func (o *VCenterHypervisor) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VCenterHypervisor
func (o *VCenterHypervisor) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VCenterHypervisor
func (o *VCenterHypervisor) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VCenterHypervisor
func (o *VCenterHypervisor) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// CreateJob creates a new child Job under the VCenterHypervisor
func (o *VCenterHypervisor) CreateJob(child *Job) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VRSAddressRanges retrieves the list of child VRSAddressRanges of the VCenterHypervisor
func (o *VCenterHypervisor) VRSAddressRanges(info *bambou.FetchingInfo) (VRSAddressRangesList, *bambou.Error) {

	var list VRSAddressRangesList
	err := bambou.CurrentSession().FetchChildren(o, VRSAddressRangeIdentity, &list, info)
	return list, err
}

// CreateVRSAddressRange creates a new child VRSAddressRange under the VCenterHypervisor
func (o *VCenterHypervisor) CreateVRSAddressRange(child *VRSAddressRange) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VRSMetrics retrieves the list of child VRSMetrics of the VCenterHypervisor
func (o *VCenterHypervisor) VRSMetrics(info *bambou.FetchingInfo) (VRSMetricsList, *bambou.Error) {

	var list VRSMetricsList
	err := bambou.CurrentSession().FetchChildren(o, VRSMetricsIdentity, &list, info)
	return list, err
}

// VRSRedeploymentpolicies retrieves the list of child VRSRedeploymentpolicies of the VCenterHypervisor
func (o *VCenterHypervisor) VRSRedeploymentpolicies(info *bambou.FetchingInfo) (VRSRedeploymentpoliciesList, *bambou.Error) {

	var list VRSRedeploymentpoliciesList
	err := bambou.CurrentSession().FetchChildren(o, VRSRedeploymentpolicyIdentity, &list, info)
	return list, err
}

// CreateVRSRedeploymentpolicy creates a new child VRSRedeploymentpolicy under the VCenterHypervisor
func (o *VCenterHypervisor) CreateVRSRedeploymentpolicy(child *VRSRedeploymentpolicy) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
