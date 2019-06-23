/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// InfrastructureAccessProfileIdentity represents the Identity of the object
var InfrastructureAccessProfileIdentity = bambou.Identity{
	Name:     "infrastructureaccessprofile",
	Category: "infrastructureaccessprofiles",
}

// InfrastructureAccessProfilesList represents a list of InfrastructureAccessProfiles
type InfrastructureAccessProfilesList []*InfrastructureAccessProfile

// InfrastructureAccessProfilesAncestor is the interface that an ancestor of a InfrastructureAccessProfile must implement.
// An Ancestor is defined as an entity that has InfrastructureAccessProfile as a descendant.
// An Ancestor can get a list of its child InfrastructureAccessProfiles, but not necessarily create one.
type InfrastructureAccessProfilesAncestor interface {
	InfrastructureAccessProfiles(*bambou.FetchingInfo) (InfrastructureAccessProfilesList, *bambou.Error)
}

// InfrastructureAccessProfilesParent is the interface that a parent of a InfrastructureAccessProfile must implement.
// A Parent is defined as an entity that has InfrastructureAccessProfile as a child.
// A Parent is an Ancestor which can create a InfrastructureAccessProfile.
type InfrastructureAccessProfilesParent interface {
	InfrastructureAccessProfilesAncestor
	CreateInfrastructureAccessProfile(*InfrastructureAccessProfile) *bambou.Error
}

// InfrastructureAccessProfile represents the model of a infrastructureaccessprofile
type InfrastructureAccessProfile struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewInfrastructureAccessProfile returns a new *InfrastructureAccessProfile
func NewInfrastructureAccessProfile() *InfrastructureAccessProfile {

	return &InfrastructureAccessProfile{}
}

// Identity returns the Identity of the object.
func (o *InfrastructureAccessProfile) Identity() bambou.Identity {

	return InfrastructureAccessProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *InfrastructureAccessProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *InfrastructureAccessProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the InfrastructureAccessProfile from the server
func (o *InfrastructureAccessProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the InfrastructureAccessProfile into the server
func (o *InfrastructureAccessProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the InfrastructureAccessProfile from the server
func (o *InfrastructureAccessProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the InfrastructureAccessProfile
func (o *InfrastructureAccessProfile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the InfrastructureAccessProfile
func (o *InfrastructureAccessProfile) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the InfrastructureAccessProfile
func (o *InfrastructureAccessProfile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the InfrastructureAccessProfile
func (o *InfrastructureAccessProfile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
