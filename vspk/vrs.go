/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VRSIdentity represents the Identity of the object
var VRSIdentity = bambou.Identity{
	Name:     "vrs",
	Category: "vrss",
}

// VRSsList represents a list of VRSs
type VRSsList []*VRS

// VRSsAncestor is the interface that an ancestor of a VRS must implement.
// An Ancestor is defined as an entity that has VRS as a descendant.
// An Ancestor can get a list of its child VRSs, but not necessarily create one.
type VRSsAncestor interface {
	VRSs(*bambou.FetchingInfo) (VRSsList, *bambou.Error)
}

// VRSsParent is the interface that a parent of a VRS must implement.
// A Parent is defined as an entity that has VRS as a child.
// A Parent is an Ancestor which can create a VRS.
type VRSsParent interface {
	VRSsAncestor
	CreateVRS(*VRS) *bambou.Error
}

// VRS represents the model of a vrs
type VRS struct {
	ID                        string        `json:"ID,omitempty"`
	ParentID                  string        `json:"parentID,omitempty"`
	ParentType                string        `json:"parentType,omitempty"`
	Owner                     string        `json:"owner,omitempty"`
	JSONRPCConnectionState    string        `json:"JSONRPCConnectionState,omitempty"`
	Name                      string        `json:"name,omitempty"`
	ManagementIP              string        `json:"managementIP,omitempty"`
	ParentIDs                 []interface{} `json:"parentIDs,omitempty"`
	LastEventName             string        `json:"lastEventName,omitempty"`
	LastEventObject           string        `json:"lastEventObject,omitempty"`
	LastEventTimestamp        int           `json:"lastEventTimestamp,omitempty"`
	LastStateChange           int           `json:"lastStateChange,omitempty"`
	LastUpdatedBy             string        `json:"lastUpdatedBy,omitempty"`
	GatewayUUID               string        `json:"gatewayUUID,omitempty"`
	DbSynced                  bool          `json:"dbSynced"`
	Address                   string        `json:"address,omitempty"`
	PeakCPUUsage              float64       `json:"peakCPUUsage,omitempty"`
	PeakMemoryUsage           float64       `json:"peakMemoryUsage,omitempty"`
	Peer                      string        `json:"peer,omitempty"`
	Personality               string        `json:"personality,omitempty"`
	Description               string        `json:"description,omitempty"`
	Messages                  []interface{} `json:"messages,omitempty"`
	RevertBehaviorEnabled     bool          `json:"revertBehaviorEnabled"`
	RevertCompleted           bool          `json:"revertCompleted"`
	RevertCount               int           `json:"revertCount,omitempty"`
	RevertFailedCount         int           `json:"revertFailedCount,omitempty"`
	LicensedState             string        `json:"licensedState,omitempty"`
	Disks                     []interface{} `json:"disks,omitempty"`
	ClusterNodeRole           string        `json:"clusterNodeRole,omitempty"`
	EntityScope               string        `json:"entityScope,omitempty"`
	Location                  string        `json:"location,omitempty"`
	Role                      string        `json:"role,omitempty"`
	Uptime                    int           `json:"uptime,omitempty"`
	PrimaryVSCConnectionLost  bool          `json:"primaryVSCConnectionLost"`
	ProductVersion            string        `json:"productVersion,omitempty"`
	IsResilient               bool          `json:"isResilient"`
	VscConfigState            string        `json:"vscConfigState,omitempty"`
	VscCurrentState           string        `json:"vscCurrentState,omitempty"`
	Status                    string        `json:"status,omitempty"`
	MultiNICVPortEnabled      bool          `json:"multiNICVPortEnabled"`
	NumberOfBridgeInterfaces  int           `json:"numberOfBridgeInterfaces,omitempty"`
	NumberOfContainers        int           `json:"numberOfContainers,omitempty"`
	NumberOfHostInterfaces    int           `json:"numberOfHostInterfaces,omitempty"`
	NumberOfVirtualMachines   int           `json:"numberOfVirtualMachines,omitempty"`
	CurrentCPUUsage           float64       `json:"currentCPUUsage,omitempty"`
	CurrentMemoryUsage        float64       `json:"currentMemoryUsage,omitempty"`
	AverageCPUUsage           float64       `json:"averageCPUUsage,omitempty"`
	AverageMemoryUsage        float64       `json:"averageMemoryUsage,omitempty"`
	ExternalID                string        `json:"externalID,omitempty"`
	Dynamic                   bool          `json:"dynamic"`
	HypervisorConnectionState string        `json:"hypervisorConnectionState,omitempty"`
	HypervisorIdentifier      string        `json:"hypervisorIdentifier,omitempty"`
	HypervisorName            string        `json:"hypervisorName,omitempty"`
	HypervisorType            string        `json:"hypervisorType,omitempty"`
}

// NewVRS returns a new *VRS
func NewVRS() *VRS {

	return &VRS{}
}

// Identity returns the Identity of the object.
func (o *VRS) Identity() bambou.Identity {

	return VRSIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VRS) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VRS) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VRS from the server
func (o *VRS) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VRS into the server
func (o *VRS) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VRS from the server
func (o *VRS) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the VRS
func (o *VRS) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VRS
func (o *VRS) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Alarms retrieves the list of child Alarms of the VRS
func (o *VRS) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

	var list AlarmsList
	err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
	return list, err
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VRS
func (o *VRS) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VRS
func (o *VRS) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VMs retrieves the list of child VMs of the VRS
func (o *VRS) VMs(info *bambou.FetchingInfo) (VMsList, *bambou.Error) {

	var list VMsList
	err := bambou.CurrentSession().FetchChildren(o, VMIdentity, &list, info)
	return list, err
}

// CreateJob creates a new child Job under the VRS
func (o *VRS) CreateJob(child *Job) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// MonitoringPorts retrieves the list of child MonitoringPorts of the VRS
func (o *VRS) MonitoringPorts(info *bambou.FetchingInfo) (MonitoringPortsList, *bambou.Error) {

	var list MonitoringPortsList
	err := bambou.CurrentSession().FetchChildren(o, MonitoringPortIdentity, &list, info)
	return list, err
}

// Containers retrieves the list of child Containers of the VRS
func (o *VRS) Containers(info *bambou.FetchingInfo) (ContainersList, *bambou.Error) {

	var list ContainersList
	err := bambou.CurrentSession().FetchChildren(o, ContainerIdentity, &list, info)
	return list, err
}

// VPorts retrieves the list of child VPorts of the VRS
func (o *VRS) VPorts(info *bambou.FetchingInfo) (VPortsList, *bambou.Error) {

	var list VPortsList
	err := bambou.CurrentSession().FetchChildren(o, VPortIdentity, &list, info)
	return list, err
}

// Statistics retrieves the list of child Statistics of the VRS
func (o *VRS) Statistics(info *bambou.FetchingInfo) (StatisticsList, *bambou.Error) {

	var list StatisticsList
	err := bambou.CurrentSession().FetchChildren(o, StatisticsIdentity, &list, info)
	return list, err
}

// MultiNICVPorts retrieves the list of child MultiNICVPorts of the VRS
func (o *VRS) MultiNICVPorts(info *bambou.FetchingInfo) (MultiNICVPortsList, *bambou.Error) {

	var list MultiNICVPortsList
	err := bambou.CurrentSession().FetchChildren(o, MultiNICVPortIdentity, &list, info)
	return list, err
}

// EventLogs retrieves the list of child EventLogs of the VRS
func (o *VRS) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
