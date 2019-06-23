/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// RedundancyGroupIdentity represents the Identity of the object
var RedundancyGroupIdentity = bambou.Identity{
	Name:     "redundancygroup",
	Category: "redundancygroups",
}

// RedundancyGroupsList represents a list of RedundancyGroups
type RedundancyGroupsList []*RedundancyGroup

// RedundancyGroupsAncestor is the interface that an ancestor of a RedundancyGroup must implement.
// An Ancestor is defined as an entity that has RedundancyGroup as a descendant.
// An Ancestor can get a list of its child RedundancyGroups, but not necessarily create one.
type RedundancyGroupsAncestor interface {
	RedundancyGroups(*bambou.FetchingInfo) (RedundancyGroupsList, *bambou.Error)
}

// RedundancyGroupsParent is the interface that a parent of a RedundancyGroup must implement.
// A Parent is defined as an entity that has RedundancyGroup as a child.
// A Parent is an Ancestor which can create a RedundancyGroup.
type RedundancyGroupsParent interface {
	RedundancyGroupsAncestor
	CreateRedundancyGroup(*RedundancyGroup) *bambou.Error
}

// RedundancyGroup represents the model of a redundancygroup
type RedundancyGroup struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewRedundancyGroup returns a new *RedundancyGroup
func NewRedundancyGroup() *RedundancyGroup {

	return &RedundancyGroup{}
}

// Identity returns the Identity of the object.
func (o *RedundancyGroup) Identity() bambou.Identity {

	return RedundancyGroupIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *RedundancyGroup) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *RedundancyGroup) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the RedundancyGroup from the server
func (o *RedundancyGroup) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the RedundancyGroup into the server
func (o *RedundancyGroup) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the RedundancyGroup from the server
func (o *RedundancyGroup) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Gateways retrieves the list of child Gateways of the RedundancyGroup
func (o *RedundancyGroup) Gateways(info *bambou.FetchingInfo) (GatewaysList, *bambou.Error) {

	var list GatewaysList
	err := bambou.CurrentSession().FetchChildren(o, GatewayIdentity, &list, info)
	return list, err
}

// Permissions retrieves the list of child Permissions of the RedundancyGroup
func (o *RedundancyGroup) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

	var list PermissionsList
	err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
	return list, err
}

// CreatePermission creates a new child Permission under the RedundancyGroup
func (o *RedundancyGroup) CreatePermission(child *Permission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// WANServices retrieves the list of child WANServices of the RedundancyGroup
func (o *RedundancyGroup) WANServices(info *bambou.FetchingInfo) (WANServicesList, *bambou.Error) {

	var list WANServicesList
	err := bambou.CurrentSession().FetchChildren(o, WANServiceIdentity, &list, info)
	return list, err
}

// CreateWANService creates a new child WANService under the RedundancyGroup
func (o *RedundancyGroup) CreateWANService(child *WANService) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the RedundancyGroup
func (o *RedundancyGroup) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the RedundancyGroup
func (o *RedundancyGroup) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Alarms retrieves the list of child Alarms of the RedundancyGroup
func (o *RedundancyGroup) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

	var list AlarmsList
	err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
	return list, err
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the RedundancyGroup
func (o *RedundancyGroup) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the RedundancyGroup
func (o *RedundancyGroup) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EnterprisePermissions retrieves the list of child EnterprisePermissions of the RedundancyGroup
func (o *RedundancyGroup) EnterprisePermissions(info *bambou.FetchingInfo) (EnterprisePermissionsList, *bambou.Error) {

	var list EnterprisePermissionsList
	err := bambou.CurrentSession().FetchChildren(o, EnterprisePermissionIdentity, &list, info)
	return list, err
}

// CreateEnterprisePermission creates a new child EnterprisePermission under the RedundancyGroup
func (o *RedundancyGroup) CreateEnterprisePermission(child *EnterprisePermission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Ports retrieves the list of child Ports of the RedundancyGroup
func (o *RedundancyGroup) Ports(info *bambou.FetchingInfo) (PortsList, *bambou.Error) {

	var list PortsList
	err := bambou.CurrentSession().FetchChildren(o, PortIdentity, &list, info)
	return list, err
}

// CreatePort creates a new child Port under the RedundancyGroup
func (o *RedundancyGroup) CreatePort(child *Port) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VsgRedundantPorts retrieves the list of child VsgRedundantPorts of the RedundancyGroup
func (o *RedundancyGroup) VsgRedundantPorts(info *bambou.FetchingInfo) (VsgRedundantPortsList, *bambou.Error) {

	var list VsgRedundantPortsList
	err := bambou.CurrentSession().FetchChildren(o, VsgRedundantPortIdentity, &list, info)
	return list, err
}

// CreateVsgRedundantPort creates a new child VsgRedundantPort under the RedundancyGroup
func (o *RedundancyGroup) CreateVsgRedundantPort(child *VsgRedundantPort) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the RedundancyGroup
func (o *RedundancyGroup) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
