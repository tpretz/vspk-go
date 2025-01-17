/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VSDIdentity represents the Identity of the object
var VSDIdentity = bambou.Identity{
	Name:     "vsd",
	Category: "vsds",
}

// VSDsList represents a list of VSDs
type VSDsList []*VSD

// VSDsAncestor is the interface that an ancestor of a VSD must implement.
// An Ancestor is defined as an entity that has VSD as a descendant.
// An Ancestor can get a list of its child VSDs, but not necessarily create one.
type VSDsAncestor interface {
	VSDs(*bambou.FetchingInfo) (VSDsList, *bambou.Error)
}

// VSDsParent is the interface that a parent of a VSD must implement.
// A Parent is defined as an entity that has VSD as a child.
// A Parent is an Ancestor which can create a VSD.
type VSDsParent interface {
	VSDsAncestor
	CreateVSD(*VSD) *bambou.Error
}

// VSD represents the model of a vsd
type VSD struct {
	ID                          string        `json:"ID,omitempty"`
	ParentID                    string        `json:"parentID,omitempty"`
	ParentType                  string        `json:"parentType,omitempty"`
	Owner                       string        `json:"owner,omitempty"`
	URL                         string        `json:"URL,omitempty"`
	Name                        string        `json:"name,omitempty"`
	ManagementIP                string        `json:"managementIP,omitempty"`
	LastStateChange             int           `json:"lastStateChange,omitempty"`
	LastUpdatedBy               string        `json:"lastUpdatedBy,omitempty"`
	Address                     string        `json:"address,omitempty"`
	PeakCPUUsage                float64       `json:"peakCPUUsage,omitempty"`
	PeakMemoryUsage             float64       `json:"peakMemoryUsage,omitempty"`
	PeerAddresses               string        `json:"peerAddresses,omitempty"`
	Description                 string        `json:"description,omitempty"`
	Messages                    []interface{} `json:"messages,omitempty"`
	Disks                       []interface{} `json:"disks,omitempty"`
	AlreadyMarkedForUnavailable bool          `json:"alreadyMarkedForUnavailable"`
	UnavailableTimestamp        int           `json:"unavailableTimestamp,omitempty"`
	EntityScope                 string        `json:"entityScope,omitempty"`
	Location                    string        `json:"location,omitempty"`
	Mode                        string        `json:"mode,omitempty"`
	ProductVersion              string        `json:"productVersion,omitempty"`
	Status                      string        `json:"status,omitempty"`
	CurrentCPUUsage             float64       `json:"currentCPUUsage,omitempty"`
	CurrentMemoryUsage          float64       `json:"currentMemoryUsage,omitempty"`
	AverageCPUUsage             float64       `json:"averageCPUUsage,omitempty"`
	AverageMemoryUsage          float64       `json:"averageMemoryUsage,omitempty"`
	ExternalID                  string        `json:"externalID,omitempty"`
}

// NewVSD returns a new *VSD
func NewVSD() *VSD {

	return &VSD{}
}

// Identity returns the Identity of the object.
func (o *VSD) Identity() bambou.Identity {

	return VSDIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VSD) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VSD) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VSD from the server
func (o *VSD) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VSD into the server
func (o *VSD) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VSD from the server
func (o *VSD) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the VSD
func (o *VSD) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VSD
func (o *VSD) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Alarms retrieves the list of child Alarms of the VSD
func (o *VSD) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

	var list AlarmsList
	err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
	return list, err
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VSD
func (o *VSD) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VSD
func (o *VSD) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// CreateJob creates a new child Job under the VSD
func (o *VSD) CreateJob(child *Job) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VSDComponents retrieves the list of child VSDComponents of the VSD
func (o *VSD) VSDComponents(info *bambou.FetchingInfo) (VSDComponentsList, *bambou.Error) {

	var list VSDComponentsList
	err := bambou.CurrentSession().FetchChildren(o, VSDComponentIdentity, &list, info)
	return list, err
}

// EventLogs retrieves the list of child EventLogs of the VSD
func (o *VSD) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
