/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// NSRedundantGatewayGroupIdentity represents the Identity of the object
var NSRedundantGatewayGroupIdentity = bambou.Identity{
	Name:     "nsgredundancygroup",
	Category: "nsgredundancygroups",
}

// NSRedundantGatewayGroupsList represents a list of NSRedundantGatewayGroups
type NSRedundantGatewayGroupsList []*NSRedundantGatewayGroup

// NSRedundantGatewayGroupsAncestor is the interface that an ancestor of a NSRedundantGatewayGroup must implement.
// An Ancestor is defined as an entity that has NSRedundantGatewayGroup as a descendant.
// An Ancestor can get a list of its child NSRedundantGatewayGroups, but not necessarily create one.
type NSRedundantGatewayGroupsAncestor interface {
	NSRedundantGatewayGroups(*bambou.FetchingInfo) (NSRedundantGatewayGroupsList, *bambou.Error)
}

// NSRedundantGatewayGroupsParent is the interface that a parent of a NSRedundantGatewayGroup must implement.
// A Parent is defined as an entity that has NSRedundantGatewayGroup as a child.
// A Parent is an Ancestor which can create a NSRedundantGatewayGroup.
type NSRedundantGatewayGroupsParent interface {
	NSRedundantGatewayGroupsAncestor
	CreateNSRedundantGatewayGroup(*NSRedundantGatewayGroup) *bambou.Error
}

// NSRedundantGatewayGroup represents the model of a nsgredundancygroup
type NSRedundantGatewayGroup struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewNSRedundantGatewayGroup returns a new *NSRedundantGatewayGroup
func NewNSRedundantGatewayGroup() *NSRedundantGatewayGroup {

	return &NSRedundantGatewayGroup{}
}

// Identity returns the Identity of the object.
func (o *NSRedundantGatewayGroup) Identity() bambou.Identity {

	return NSRedundantGatewayGroupIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NSRedundantGatewayGroup) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NSRedundantGatewayGroup) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NSRedundantGatewayGroup from the server
func (o *NSRedundantGatewayGroup) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NSRedundantGatewayGroup into the server
func (o *NSRedundantGatewayGroup) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NSRedundantGatewayGroup from the server
func (o *NSRedundantGatewayGroup) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the NSRedundantGatewayGroup
func (o *NSRedundantGatewayGroup) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the NSRedundantGatewayGroup
func (o *NSRedundantGatewayGroup) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// ShuntLinks retrieves the list of child ShuntLinks of the NSRedundantGatewayGroup
func (o *NSRedundantGatewayGroup) ShuntLinks(info *bambou.FetchingInfo) (ShuntLinksList, *bambou.Error) {

	var list ShuntLinksList
	err := bambou.CurrentSession().FetchChildren(o, ShuntLinkIdentity, &list, info)
	return list, err
}

// CreateShuntLink creates a new child ShuntLink under the NSRedundantGatewayGroup
func (o *NSRedundantGatewayGroup) CreateShuntLink(child *ShuntLink) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Alarms retrieves the list of child Alarms of the NSRedundantGatewayGroup
func (o *NSRedundantGatewayGroup) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

	var list AlarmsList
	err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
	return list, err
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the NSRedundantGatewayGroup
func (o *NSRedundantGatewayGroup) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the NSRedundantGatewayGroup
func (o *NSRedundantGatewayGroup) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// NSGateways retrieves the list of child NSGateways of the NSRedundantGatewayGroup
func (o *NSRedundantGatewayGroup) NSGateways(info *bambou.FetchingInfo) (NSGatewaysList, *bambou.Error) {

	var list NSGatewaysList
	err := bambou.CurrentSession().FetchChildren(o, NSGatewayIdentity, &list, info)
	return list, err
}

// RedundantPorts retrieves the list of child RedundantPorts of the NSRedundantGatewayGroup
func (o *NSRedundantGatewayGroup) RedundantPorts(info *bambou.FetchingInfo) (RedundantPortsList, *bambou.Error) {

	var list RedundantPortsList
	err := bambou.CurrentSession().FetchChildren(o, RedundantPortIdentity, &list, info)
	return list, err
}

// CreateRedundantPort creates a new child RedundantPort under the NSRedundantGatewayGroup
func (o *NSRedundantGatewayGroup) CreateRedundantPort(child *RedundantPort) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the NSRedundantGatewayGroup
func (o *NSRedundantGatewayGroup) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
