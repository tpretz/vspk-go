/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// FlowForwardingPolicyIdentity represents the Identity of the object
var FlowForwardingPolicyIdentity = bambou.Identity{
	Name:     "flowforwardingpolicy",
	Category: "flowforwardingpolicies",
}

// FlowForwardingPoliciesList represents a list of FlowForwardingPolicies
type FlowForwardingPoliciesList []*FlowForwardingPolicy

// FlowForwardingPoliciesAncestor is the interface that an ancestor of a FlowForwardingPolicy must implement.
// An Ancestor is defined as an entity that has FlowForwardingPolicy as a descendant.
// An Ancestor can get a list of its child FlowForwardingPolicies, but not necessarily create one.
type FlowForwardingPoliciesAncestor interface {
	FlowForwardingPolicies(*bambou.FetchingInfo) (FlowForwardingPoliciesList, *bambou.Error)
}

// FlowForwardingPoliciesParent is the interface that a parent of a FlowForwardingPolicy must implement.
// A Parent is defined as an entity that has FlowForwardingPolicy as a child.
// A Parent is an Ancestor which can create a FlowForwardingPolicy.
type FlowForwardingPoliciesParent interface {
	FlowForwardingPoliciesAncestor
	CreateFlowForwardingPolicy(*FlowForwardingPolicy) *bambou.Error
}

// FlowForwardingPolicy represents the model of a flowforwardingpolicy
type FlowForwardingPolicy struct {
	ID                             string `json:"ID,omitempty"`
	ParentID                       string `json:"parentID,omitempty"`
	ParentType                     string `json:"parentType,omitempty"`
	Owner                          string `json:"owner,omitempty"`
	RedirectTargetID               string `json:"redirectTargetID,omitempty"`
	DestinationAddressOverwrite    string `json:"destinationAddressOverwrite,omitempty"`
	FlowID                         string `json:"flowID,omitempty"`
	EntityScope                    string `json:"entityScope,omitempty"`
	SourceAddressOverwrite         string `json:"sourceAddressOverwrite,omitempty"`
	AssociatedApplicationServiceID string `json:"associatedApplicationServiceID,omitempty"`
	AssociatedNetworkObjectID      string `json:"associatedNetworkObjectID,omitempty"`
	AssociatedNetworkObjectType    string `json:"associatedNetworkObjectType,omitempty"`
	ExternalID                     string `json:"externalID,omitempty"`
	Type                           string `json:"type,omitempty"`
}

// NewFlowForwardingPolicy returns a new *FlowForwardingPolicy
func NewFlowForwardingPolicy() *FlowForwardingPolicy {

	return &FlowForwardingPolicy{}
}

// Identity returns the Identity of the object.
func (o *FlowForwardingPolicy) Identity() bambou.Identity {

	return FlowForwardingPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FlowForwardingPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FlowForwardingPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the FlowForwardingPolicy from the server
func (o *FlowForwardingPolicy) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the FlowForwardingPolicy into the server
func (o *FlowForwardingPolicy) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the FlowForwardingPolicy from the server
func (o *FlowForwardingPolicy) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the FlowForwardingPolicy
func (o *FlowForwardingPolicy) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the FlowForwardingPolicy
func (o *FlowForwardingPolicy) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the FlowForwardingPolicy
func (o *FlowForwardingPolicy) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the FlowForwardingPolicy
func (o *FlowForwardingPolicy) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the FlowForwardingPolicy
func (o *FlowForwardingPolicy) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
