/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// FlowIdentity represents the Identity of the object
var FlowIdentity = bambou.Identity{
	Name:     "flow",
	Category: "flows",
}

// FlowsList represents a list of Flows
type FlowsList []*Flow

// FlowsAncestor is the interface that an ancestor of a Flow must implement.
// An Ancestor is defined as an entity that has Flow as a descendant.
// An Ancestor can get a list of its child Flows, but not necessarily create one.
type FlowsAncestor interface {
	Flows(*bambou.FetchingInfo) (FlowsList, *bambou.Error)
}

// FlowsParent is the interface that a parent of a Flow must implement.
// A Parent is defined as an entity that has Flow as a child.
// A Parent is an Ancestor which can create a Flow.
type FlowsParent interface {
	FlowsAncestor
	CreateFlow(*Flow) *bambou.Error
}

// Flow represents the model of a flow
type Flow struct {
	ID                string `json:"ID,omitempty"`
	ParentID          string `json:"parentID,omitempty"`
	ParentType        string `json:"parentType,omitempty"`
	Owner             string `json:"owner,omitempty"`
	Name              string `json:"name,omitempty"`
	LastUpdatedBy     string `json:"lastUpdatedBy,omitempty"`
	Description       string `json:"description,omitempty"`
	DestinationTierID string `json:"destinationTierID,omitempty"`
	Metadata          string `json:"metadata,omitempty"`
	EntityScope       string `json:"entityScope,omitempty"`
	OriginTierID      string `json:"originTierID,omitempty"`
	ExternalID        string `json:"externalID,omitempty"`
}

// NewFlow returns a new *Flow
func NewFlow() *Flow {

	return &Flow{}
}

// Identity returns the Identity of the object.
func (o *Flow) Identity() bambou.Identity {

	return FlowIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Flow) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Flow) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Flow from the server
func (o *Flow) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Flow into the server
func (o *Flow) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Flow from the server
func (o *Flow) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the Flow
func (o *Flow) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Flow
func (o *Flow) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Flow
func (o *Flow) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Flow
func (o *Flow) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// FlowForwardingPolicies retrieves the list of child FlowForwardingPolicies of the Flow
func (o *Flow) FlowForwardingPolicies(info *bambou.FetchingInfo) (FlowForwardingPoliciesList, *bambou.Error) {

	var list FlowForwardingPoliciesList
	err := bambou.CurrentSession().FetchChildren(o, FlowForwardingPolicyIdentity, &list, info)
	return list, err
}

// CreateFlowForwardingPolicy creates a new child FlowForwardingPolicy under the Flow
func (o *Flow) CreateFlowForwardingPolicy(child *FlowForwardingPolicy) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// FlowSecurityPolicies retrieves the list of child FlowSecurityPolicies of the Flow
func (o *Flow) FlowSecurityPolicies(info *bambou.FetchingInfo) (FlowSecurityPoliciesList, *bambou.Error) {

	var list FlowSecurityPoliciesList
	err := bambou.CurrentSession().FetchChildren(o, FlowSecurityPolicyIdentity, &list, info)
	return list, err
}

// CreateFlowSecurityPolicy creates a new child FlowSecurityPolicy under the Flow
func (o *Flow) CreateFlowSecurityPolicy(child *FlowSecurityPolicy) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the Flow
func (o *Flow) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
