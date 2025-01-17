/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// PolicyGroupIdentity represents the Identity of the object
var PolicyGroupIdentity = bambou.Identity{
	Name:     "policygroup",
	Category: "policygroups",
}

// PolicyGroupsList represents a list of PolicyGroups
type PolicyGroupsList []*PolicyGroup

// PolicyGroupsAncestor is the interface that an ancestor of a PolicyGroup must implement.
// An Ancestor is defined as an entity that has PolicyGroup as a descendant.
// An Ancestor can get a list of its child PolicyGroups, but not necessarily create one.
type PolicyGroupsAncestor interface {
	PolicyGroups(*bambou.FetchingInfo) (PolicyGroupsList, *bambou.Error)
}

// PolicyGroupsParent is the interface that a parent of a PolicyGroup must implement.
// A Parent is defined as an entity that has PolicyGroup as a child.
// A Parent is an Ancestor which can create a PolicyGroup.
type PolicyGroupsParent interface {
	PolicyGroupsAncestor
	CreatePolicyGroup(*PolicyGroup) *bambou.Error
}

// PolicyGroup represents the model of a policygroup
type PolicyGroup struct {
	ID                           string `json:"ID,omitempty"`
	ParentID                     string `json:"parentID,omitempty"`
	ParentType                   string `json:"parentType,omitempty"`
	Owner                        string `json:"owner,omitempty"`
	EVPNCommunityTag             string `json:"EVPNCommunityTag,omitempty"`
	Name                         string `json:"name,omitempty"`
	LastUpdatedBy                string `json:"lastUpdatedBy,omitempty"`
	TemplateID                   string `json:"templateID,omitempty"`
	Description                  string `json:"description,omitempty"`
	EntityScope                  string `json:"entityScope,omitempty"`
	PolicyGroupID                int    `json:"policyGroupID,omitempty"`
	AssocPolicyGroupCategoryID   string `json:"assocPolicyGroupCategoryID,omitempty"`
	AssocPolicyGroupCategoryName string `json:"assocPolicyGroupCategoryName,omitempty"`
	External                     bool   `json:"external"`
	ExternalID                   string `json:"externalID,omitempty"`
	Type                         string `json:"type,omitempty"`
}

// NewPolicyGroup returns a new *PolicyGroup
func NewPolicyGroup() *PolicyGroup {

	return &PolicyGroup{}
}

// Identity returns the Identity of the object.
func (o *PolicyGroup) Identity() bambou.Identity {

	return PolicyGroupIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PolicyGroup) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PolicyGroup) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PolicyGroup from the server
func (o *PolicyGroup) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PolicyGroup into the server
func (o *PolicyGroup) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PolicyGroup from the server
func (o *PolicyGroup) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the PolicyGroup
func (o *PolicyGroup) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the PolicyGroup
func (o *PolicyGroup) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the PolicyGroup
func (o *PolicyGroup) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the PolicyGroup
func (o *PolicyGroup) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VPorts retrieves the list of child VPorts of the PolicyGroup
func (o *PolicyGroup) VPorts(info *bambou.FetchingInfo) (VPortsList, *bambou.Error) {

	var list VPortsList
	err := bambou.CurrentSession().FetchChildren(o, VPortIdentity, &list, info)
	return list, err
}

// AssignVPorts assigns the list of VPorts to the PolicyGroup
func (o *PolicyGroup) AssignVPorts(children VPortsList) *bambou.Error {

	list := []bambou.Identifiable{}
	for _, c := range children {
		list = append(list, c)
	}

	return bambou.CurrentSession().AssignChildren(o, list, VPortIdentity)
}

// EventLogs retrieves the list of child EventLogs of the PolicyGroup
func (o *PolicyGroup) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
