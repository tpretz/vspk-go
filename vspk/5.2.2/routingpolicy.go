/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// RoutingPolicyIdentity represents the Identity of the object
var RoutingPolicyIdentity = bambou.Identity{
	Name:     "routingpolicy",
	Category: "routingpolicies",
}

// RoutingPoliciesList represents a list of RoutingPolicies
type RoutingPoliciesList []*RoutingPolicy

// RoutingPoliciesAncestor is the interface that an ancestor of a RoutingPolicy must implement.
// An Ancestor is defined as an entity that has RoutingPolicy as a descendant.
// An Ancestor can get a list of its child RoutingPolicies, but not necessarily create one.
type RoutingPoliciesAncestor interface {
	RoutingPolicies(*bambou.FetchingInfo) (RoutingPoliciesList, *bambou.Error)
}

// RoutingPoliciesParent is the interface that a parent of a RoutingPolicy must implement.
// A Parent is defined as an entity that has RoutingPolicy as a child.
// A Parent is an Ancestor which can create a RoutingPolicy.
type RoutingPoliciesParent interface {
	RoutingPoliciesAncestor
	CreateRoutingPolicy(*RoutingPolicy) *bambou.Error
}

// RoutingPolicy represents the model of a routingpolicy
type RoutingPolicy struct {
	ID               string `json:"ID,omitempty"`
	ParentID         string `json:"parentID,omitempty"`
	ParentType       string `json:"parentType,omitempty"`
	Owner            string `json:"owner,omitempty"`
	Name             string `json:"name,omitempty"`
	DefaultAction    string `json:"defaultAction,omitempty"`
	Description      string `json:"description,omitempty"`
	EntityScope      string `json:"entityScope,omitempty"`
	PolicyDefinition string `json:"policyDefinition,omitempty"`
	ExternalID       string `json:"externalID,omitempty"`
}

// NewRoutingPolicy returns a new *RoutingPolicy
func NewRoutingPolicy() *RoutingPolicy {

	return &RoutingPolicy{}
}

// Identity returns the Identity of the object.
func (o *RoutingPolicy) Identity() bambou.Identity {

	return RoutingPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *RoutingPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *RoutingPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the RoutingPolicy from the server
func (o *RoutingPolicy) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the RoutingPolicy into the server
func (o *RoutingPolicy) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the RoutingPolicy from the server
func (o *RoutingPolicy) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the RoutingPolicy
func (o *RoutingPolicy) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the RoutingPolicy
func (o *RoutingPolicy) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the RoutingPolicy
func (o *RoutingPolicy) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the RoutingPolicy
func (o *RoutingPolicy) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
