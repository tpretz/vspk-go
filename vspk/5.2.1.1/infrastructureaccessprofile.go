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
	ID             string `json:"ID,omitempty"`
	ParentID       string `json:"parentID,omitempty"`
	ParentType     string `json:"parentType,omitempty"`
	Owner          string `json:"owner,omitempty"`
	SSHAuthMode    string `json:"SSHAuthMode,omitempty"`
	Name           string `json:"name,omitempty"`
	Password       string `json:"password,omitempty"`
	LastUpdatedBy  string `json:"lastUpdatedBy,omitempty"`
	Description    string `json:"description,omitempty"`
	EnterpriseID   string `json:"enterpriseID,omitempty"`
	EntityScope    string `json:"entityScope,omitempty"`
	SourceIPFilter string `json:"sourceIPFilter,omitempty"`
	UserName       string `json:"userName,omitempty"`
	ExternalID     string `json:"externalID,omitempty"`
}

// NewInfrastructureAccessProfile returns a new *InfrastructureAccessProfile
func NewInfrastructureAccessProfile() *InfrastructureAccessProfile {

	return &InfrastructureAccessProfile{
		SSHAuthMode:    "PASSWORD_AND_KEY_BASED",
		SourceIPFilter: "DISABLED",
		UserName:       "nuage",
	}
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

// Connectionendpoints retrieves the list of child Connectionendpoints of the InfrastructureAccessProfile
func (o *InfrastructureAccessProfile) Connectionendpoints(info *bambou.FetchingInfo) (ConnectionendpointsList, *bambou.Error) {

	var list ConnectionendpointsList
	err := bambou.CurrentSession().FetchChildren(o, ConnectionendpointIdentity, &list, info)
	return list, err
}

// CreateConnectionendpoint creates a new child Connectionendpoint under the InfrastructureAccessProfile
func (o *InfrastructureAccessProfile) CreateConnectionendpoint(child *Connectionendpoint) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// SSHKeys retrieves the list of child SSHKeys of the InfrastructureAccessProfile
func (o *InfrastructureAccessProfile) SSHKeys(info *bambou.FetchingInfo) (SSHKeysList, *bambou.Error) {

	var list SSHKeysList
	err := bambou.CurrentSession().FetchChildren(o, SSHKeyIdentity, &list, info)
	return list, err
}

// CreateSSHKey creates a new child SSHKey under the InfrastructureAccessProfile
func (o *InfrastructureAccessProfile) CreateSSHKey(child *SSHKey) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
