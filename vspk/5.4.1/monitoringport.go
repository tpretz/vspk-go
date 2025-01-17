/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// MonitoringPortIdentity represents the Identity of the object
var MonitoringPortIdentity = bambou.Identity{
	Name:     "monitoringport",
	Category: "monitoringports",
}

// MonitoringPortsList represents a list of MonitoringPorts
type MonitoringPortsList []*MonitoringPort

// MonitoringPortsAncestor is the interface that an ancestor of a MonitoringPort must implement.
// An Ancestor is defined as an entity that has MonitoringPort as a descendant.
// An Ancestor can get a list of its child MonitoringPorts, but not necessarily create one.
type MonitoringPortsAncestor interface {
	MonitoringPorts(*bambou.FetchingInfo) (MonitoringPortsList, *bambou.Error)
}

// MonitoringPortsParent is the interface that a parent of a MonitoringPort must implement.
// A Parent is defined as an entity that has MonitoringPort as a child.
// A Parent is an Ancestor which can create a MonitoringPort.
type MonitoringPortsParent interface {
	MonitoringPortsAncestor
	CreateMonitoringPort(*MonitoringPort) *bambou.Error
}

// MonitoringPort represents the model of a monitoringport
type MonitoringPort struct {
	ID              string `json:"ID,omitempty"`
	ParentID        string `json:"parentID,omitempty"`
	ParentType      string `json:"parentType,omitempty"`
	Owner           string `json:"owner,omitempty"`
	Name            string `json:"name,omitempty"`
	LastStateChange int    `json:"lastStateChange,omitempty"`
	Access          bool   `json:"access"`
	Description     string `json:"description,omitempty"`
	ResiliencyState string `json:"resiliencyState,omitempty"`
	Resilient       bool   `json:"resilient"`
	EntityScope     string `json:"entityScope,omitempty"`
	DpdkEnabled     bool   `json:"dpdkEnabled"`
	Uplink          bool   `json:"uplink"`
	State           string `json:"state,omitempty"`
	ExternalID      string `json:"externalID,omitempty"`
}

// NewMonitoringPort returns a new *MonitoringPort
func NewMonitoringPort() *MonitoringPort {

	return &MonitoringPort{}
}

// Identity returns the Identity of the object.
func (o *MonitoringPort) Identity() bambou.Identity {

	return MonitoringPortIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *MonitoringPort) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *MonitoringPort) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the MonitoringPort from the server
func (o *MonitoringPort) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the MonitoringPort into the server
func (o *MonitoringPort) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the MonitoringPort from the server
func (o *MonitoringPort) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the MonitoringPort
func (o *MonitoringPort) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the MonitoringPort
func (o *MonitoringPort) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the MonitoringPort
func (o *MonitoringPort) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the MonitoringPort
func (o *MonitoringPort) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
