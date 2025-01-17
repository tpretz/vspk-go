/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// HSCIdentity represents the Identity of the object
var HSCIdentity = bambou.Identity{
	Name:     "hsc",
	Category: "hscs",
}

// HSCsList represents a list of HSCs
type HSCsList []*HSC

// HSCsAncestor is the interface that an ancestor of a HSC must implement.
// An Ancestor is defined as an entity that has HSC as a descendant.
// An Ancestor can get a list of its child HSCs, but not necessarily create one.
type HSCsAncestor interface {
	HSCs(*bambou.FetchingInfo) (HSCsList, *bambou.Error)
}

// HSCsParent is the interface that a parent of a HSC must implement.
// A Parent is defined as an entity that has HSC as a child.
// A Parent is an Ancestor which can create a HSC.
type HSCsParent interface {
	HSCsAncestor
	CreateHSC(*HSC) *bambou.Error
}

// HSC represents the model of a hsc
type HSC struct {
	ID                          string        `json:"ID,omitempty"`
	ParentID                    string        `json:"parentID,omitempty"`
	ParentType                  string        `json:"parentType,omitempty"`
	Owner                       string        `json:"owner,omitempty"`
	Name                        string        `json:"name,omitempty"`
	ManagementIP                string        `json:"managementIP,omitempty"`
	LastStateChange             int           `json:"lastStateChange,omitempty"`
	LastUpdatedBy               string        `json:"lastUpdatedBy,omitempty"`
	Address                     string        `json:"address,omitempty"`
	PeakCPUUsage                float64       `json:"peakCPUUsage,omitempty"`
	PeakMemoryUsage             float64       `json:"peakMemoryUsage,omitempty"`
	Description                 string        `json:"description,omitempty"`
	Messages                    []interface{} `json:"messages,omitempty"`
	Disks                       []interface{} `json:"disks,omitempty"`
	AlreadyMarkedForUnavailable bool          `json:"alreadyMarkedForUnavailable"`
	UnavailableTimestamp        int           `json:"unavailableTimestamp,omitempty"`
	EntityScope                 string        `json:"entityScope,omitempty"`
	Location                    string        `json:"location,omitempty"`
	Model                       string        `json:"model,omitempty"`
	ProductVersion              string        `json:"productVersion,omitempty"`
	Vsds                        []interface{} `json:"vsds,omitempty"`
	Status                      string        `json:"status,omitempty"`
	CurrentCPUUsage             float64       `json:"currentCPUUsage,omitempty"`
	CurrentMemoryUsage          float64       `json:"currentMemoryUsage,omitempty"`
	AverageCPUUsage             float64       `json:"averageCPUUsage,omitempty"`
	AverageMemoryUsage          float64       `json:"averageMemoryUsage,omitempty"`
	ExternalID                  string        `json:"externalID,omitempty"`
	Type                        string        `json:"type,omitempty"`
}

// NewHSC returns a new *HSC
func NewHSC() *HSC {

	return &HSC{}
}

// Identity returns the Identity of the object.
func (o *HSC) Identity() bambou.Identity {

	return HSCIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *HSC) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *HSC) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the HSC from the server
func (o *HSC) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the HSC into the server
func (o *HSC) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the HSC from the server
func (o *HSC) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the HSC
func (o *HSC) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the HSC
func (o *HSC) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// BGPPeers retrieves the list of child BGPPeers of the HSC
func (o *HSC) BGPPeers(info *bambou.FetchingInfo) (BGPPeersList, *bambou.Error) {

	var list BGPPeersList
	err := bambou.CurrentSession().FetchChildren(o, BGPPeerIdentity, &list, info)
	return list, err
}

// Alarms retrieves the list of child Alarms of the HSC
func (o *HSC) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

	var list AlarmsList
	err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
	return list, err
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the HSC
func (o *HSC) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the HSC
func (o *HSC) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// CreateJob creates a new child Job under the HSC
func (o *HSC) CreateJob(child *Job) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// MonitoringPorts retrieves the list of child MonitoringPorts of the HSC
func (o *HSC) MonitoringPorts(info *bambou.FetchingInfo) (MonitoringPortsList, *bambou.Error) {

	var list MonitoringPortsList
	err := bambou.CurrentSession().FetchChildren(o, MonitoringPortIdentity, &list, info)
	return list, err
}

// VRSs retrieves the list of child VRSs of the HSC
func (o *HSC) VRSs(info *bambou.FetchingInfo) (VRSsList, *bambou.Error) {

	var list VRSsList
	err := bambou.CurrentSession().FetchChildren(o, VRSIdentity, &list, info)
	return list, err
}

// EventLogs retrieves the list of child EventLogs of the HSC
func (o *HSC) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
