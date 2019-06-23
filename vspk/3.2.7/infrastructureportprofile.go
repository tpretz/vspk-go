/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// InfrastructurePortProfileIdentity represents the Identity of the object
var InfrastructurePortProfileIdentity = bambou.Identity{
	Name:     "infrastructureportprofile",
	Category: "infrastructureportprofiles",
}

// InfrastructurePortProfilesList represents a list of InfrastructurePortProfiles
type InfrastructurePortProfilesList []*InfrastructurePortProfile

// InfrastructurePortProfilesAncestor is the interface that an ancestor of a InfrastructurePortProfile must implement.
// An Ancestor is defined as an entity that has InfrastructurePortProfile as a descendant.
// An Ancestor can get a list of its child InfrastructurePortProfiles, but not necessarily create one.
type InfrastructurePortProfilesAncestor interface {
	InfrastructurePortProfiles(*bambou.FetchingInfo) (InfrastructurePortProfilesList, *bambou.Error)
}

// InfrastructurePortProfilesParent is the interface that a parent of a InfrastructurePortProfile must implement.
// A Parent is defined as an entity that has InfrastructurePortProfile as a child.
// A Parent is an Ancestor which can create a InfrastructurePortProfile.
type InfrastructurePortProfilesParent interface {
	InfrastructurePortProfilesAncestor
	CreateInfrastructurePortProfile(*InfrastructurePortProfile) *bambou.Error
}

// InfrastructurePortProfile represents the model of a infrastructureportprofile
type InfrastructurePortProfile struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewInfrastructurePortProfile returns a new *InfrastructurePortProfile
func NewInfrastructurePortProfile() *InfrastructurePortProfile {

	return &InfrastructurePortProfile{}
}

// Identity returns the Identity of the object.
func (o *InfrastructurePortProfile) Identity() bambou.Identity {

	return InfrastructurePortProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *InfrastructurePortProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *InfrastructurePortProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the InfrastructurePortProfile from the server
func (o *InfrastructurePortProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the InfrastructurePortProfile into the server
func (o *InfrastructurePortProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the InfrastructurePortProfile from the server
func (o *InfrastructurePortProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the InfrastructurePortProfile
func (o *InfrastructurePortProfile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the InfrastructurePortProfile
func (o *InfrastructurePortProfile) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the InfrastructurePortProfile
func (o *InfrastructurePortProfile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the InfrastructurePortProfile
func (o *InfrastructurePortProfile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
