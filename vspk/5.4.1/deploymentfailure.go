/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// DeploymentFailureIdentity represents the Identity of the object
var DeploymentFailureIdentity = bambou.Identity{
	Name:     "deploymentfailure",
	Category: "deploymentfailures",
}

// DeploymentFailuresList represents a list of DeploymentFailures
type DeploymentFailuresList []*DeploymentFailure

// DeploymentFailuresAncestor is the interface that an ancestor of a DeploymentFailure must implement.
// An Ancestor is defined as an entity that has DeploymentFailure as a descendant.
// An Ancestor can get a list of its child DeploymentFailures, but not necessarily create one.
type DeploymentFailuresAncestor interface {
	DeploymentFailures(*bambou.FetchingInfo) (DeploymentFailuresList, *bambou.Error)
}

// DeploymentFailuresParent is the interface that a parent of a DeploymentFailure must implement.
// A Parent is defined as an entity that has DeploymentFailure as a child.
// A Parent is an Ancestor which can create a DeploymentFailure.
type DeploymentFailuresParent interface {
	DeploymentFailuresAncestor
	CreateDeploymentFailure(*DeploymentFailure) *bambou.Error
}

// DeploymentFailure represents the model of a deploymentfailure
type DeploymentFailure struct {
	ID                          string `json:"ID,omitempty"`
	ParentID                    string `json:"parentID,omitempty"`
	ParentType                  string `json:"parentType,omitempty"`
	Owner                       string `json:"owner,omitempty"`
	LastFailureReason           string `json:"lastFailureReason,omitempty"`
	LastKnownError              string `json:"lastKnownError,omitempty"`
	LastUpdatedBy               string `json:"lastUpdatedBy,omitempty"`
	AffectedEntityID            string `json:"affectedEntityID,omitempty"`
	AffectedEntityType          string `json:"affectedEntityType,omitempty"`
	DiffMap                     string `json:"diffMap,omitempty"`
	EntityScope                 string `json:"entityScope,omitempty"`
	ErrorCondition              int    `json:"errorCondition,omitempty"`
	AssocEntityId               string `json:"assocEntityId,omitempty"`
	AssocEntityType             string `json:"assocEntityType,omitempty"`
	AssociatedNetworkEntityID   string `json:"associatedNetworkEntityID,omitempty"`
	AssociatedNetworkEntityType string `json:"associatedNetworkEntityType,omitempty"`
	NumberOfOccurences          int    `json:"numberOfOccurences,omitempty"`
	EventType                   string `json:"eventType,omitempty"`
	ExternalID                  string `json:"externalID,omitempty"`
}

// NewDeploymentFailure returns a new *DeploymentFailure
func NewDeploymentFailure() *DeploymentFailure {

	return &DeploymentFailure{}
}

// Identity returns the Identity of the object.
func (o *DeploymentFailure) Identity() bambou.Identity {

	return DeploymentFailureIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DeploymentFailure) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DeploymentFailure) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the DeploymentFailure from the server
func (o *DeploymentFailure) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the DeploymentFailure into the server
func (o *DeploymentFailure) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the DeploymentFailure from the server
func (o *DeploymentFailure) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the DeploymentFailure
func (o *DeploymentFailure) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the DeploymentFailure
func (o *DeploymentFailure) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the DeploymentFailure
func (o *DeploymentFailure) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the DeploymentFailure
func (o *DeploymentFailure) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
