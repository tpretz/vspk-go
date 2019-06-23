/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// FlowSecurityPolicyIdentity represents the Identity of the object
var FlowSecurityPolicyIdentity = bambou.Identity{
	Name:     "flowsecuritypolicy",
	Category: "flowsecuritypolicies",
}

// FlowSecurityPoliciesList represents a list of FlowSecurityPolicies
type FlowSecurityPoliciesList []*FlowSecurityPolicy

// FlowSecurityPoliciesAncestor is the interface that an ancestor of a FlowSecurityPolicy must implement.
// An Ancestor is defined as an entity that has FlowSecurityPolicy as a descendant.
// An Ancestor can get a list of its child FlowSecurityPolicies, but not necessarily create one.
type FlowSecurityPoliciesAncestor interface {
	FlowSecurityPolicies(*bambou.FetchingInfo) (FlowSecurityPoliciesList, *bambou.Error)
}

// FlowSecurityPoliciesParent is the interface that a parent of a FlowSecurityPolicy must implement.
// A Parent is defined as an entity that has FlowSecurityPolicy as a child.
// A Parent is an Ancestor which can create a FlowSecurityPolicy.
type FlowSecurityPoliciesParent interface {
	FlowSecurityPoliciesAncestor
	CreateFlowSecurityPolicy(*FlowSecurityPolicy) *bambou.Error
}

// FlowSecurityPolicy represents the model of a flowsecuritypolicy
type FlowSecurityPolicy struct {
	ID                             string `json:"ID,omitempty"`
	ParentID                       string `json:"parentID,omitempty"`
	ParentType                     string `json:"parentType,omitempty"`
	Owner                          string `json:"owner,omitempty"`
	Action                         string `json:"action,omitempty"`
	DestinationAddressOverwrite    string `json:"destinationAddressOverwrite,omitempty"`
	FlowID                         string `json:"flowID,omitempty"`
	EntityScope                    string `json:"entityScope,omitempty"`
	SourceAddressOverwrite         string `json:"sourceAddressOverwrite,omitempty"`
	Priority                       int    `json:"priority,omitempty"`
	AssociatedApplicationServiceID string `json:"associatedApplicationServiceID,omitempty"`
	AssociatedNetworkObjectID      string `json:"associatedNetworkObjectID,omitempty"`
	AssociatedNetworkObjectType    string `json:"associatedNetworkObjectType,omitempty"`
	ExternalID                     string `json:"externalID,omitempty"`
}

// NewFlowSecurityPolicy returns a new *FlowSecurityPolicy
func NewFlowSecurityPolicy() *FlowSecurityPolicy {

	return &FlowSecurityPolicy{}
}

// Identity returns the Identity of the object.
func (o *FlowSecurityPolicy) Identity() bambou.Identity {

	return FlowSecurityPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FlowSecurityPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FlowSecurityPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the FlowSecurityPolicy from the server
func (o *FlowSecurityPolicy) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the FlowSecurityPolicy into the server
func (o *FlowSecurityPolicy) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the FlowSecurityPolicy from the server
func (o *FlowSecurityPolicy) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the FlowSecurityPolicy
func (o *FlowSecurityPolicy) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the FlowSecurityPolicy
func (o *FlowSecurityPolicy) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the FlowSecurityPolicy
func (o *FlowSecurityPolicy) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the FlowSecurityPolicy
func (o *FlowSecurityPolicy) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the FlowSecurityPolicy
func (o *FlowSecurityPolicy) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
