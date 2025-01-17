/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// BootstrapActivationIdentity represents the Identity of the object
var BootstrapActivationIdentity = bambou.Identity{
	Name:     "bootstrapactivation",
	Category: "bootstrapactivations",
}

// BootstrapActivationsList represents a list of BootstrapActivations
type BootstrapActivationsList []*BootstrapActivation

// BootstrapActivationsAncestor is the interface that an ancestor of a BootstrapActivation must implement.
// An Ancestor is defined as an entity that has BootstrapActivation as a descendant.
// An Ancestor can get a list of its child BootstrapActivations, but not necessarily create one.
type BootstrapActivationsAncestor interface {
	BootstrapActivations(*bambou.FetchingInfo) (BootstrapActivationsList, *bambou.Error)
}

// BootstrapActivationsParent is the interface that a parent of a BootstrapActivation must implement.
// A Parent is defined as an entity that has BootstrapActivation as a child.
// A Parent is an Ancestor which can create a BootstrapActivation.
type BootstrapActivationsParent interface {
	BootstrapActivationsAncestor
	CreateBootstrapActivation(*BootstrapActivation) *bambou.Error
}

// BootstrapActivation represents the model of a bootstrapactivation
type BootstrapActivation struct {
	ID                   string `json:"ID,omitempty"`
	ParentID             string `json:"parentID,omitempty"`
	ParentType           string `json:"parentType,omitempty"`
	Owner                string `json:"owner,omitempty"`
	Cacert               string `json:"cacert,omitempty"`
	Hash                 string `json:"hash,omitempty"`
	LastUpdatedBy        string `json:"lastUpdatedBy,omitempty"`
	Action               string `json:"action,omitempty"`
	Seed                 string `json:"seed,omitempty"`
	Cert                 string `json:"cert,omitempty"`
	EntityScope          string `json:"entityScope,omitempty"`
	ConfigURL            string `json:"configURL,omitempty"`
	TpmOwnerPassword     string `json:"tpmOwnerPassword,omitempty"`
	TpmState             int    `json:"tpmState,omitempty"`
	SrkPassword          string `json:"srkPassword,omitempty"`
	VsdTime              int    `json:"vsdTime,omitempty"`
	Csr                  string `json:"csr,omitempty"`
	AssociatedEntityType string `json:"associatedEntityType,omitempty"`
	Status               string `json:"status,omitempty"`
	AutoBootstrap        bool   `json:"autoBootstrap"`
	ExternalID           string `json:"externalID,omitempty"`
}

// NewBootstrapActivation returns a new *BootstrapActivation
func NewBootstrapActivation() *BootstrapActivation {

	return &BootstrapActivation{
		TpmState: 0,
	}
}

// Identity returns the Identity of the object.
func (o *BootstrapActivation) Identity() bambou.Identity {

	return BootstrapActivationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *BootstrapActivation) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *BootstrapActivation) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the BootstrapActivation from the server
func (o *BootstrapActivation) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the BootstrapActivation into the server
func (o *BootstrapActivation) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the BootstrapActivation from the server
func (o *BootstrapActivation) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the BootstrapActivation
func (o *BootstrapActivation) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the BootstrapActivation
func (o *BootstrapActivation) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the BootstrapActivation
func (o *BootstrapActivation) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the BootstrapActivation
func (o *BootstrapActivation) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
