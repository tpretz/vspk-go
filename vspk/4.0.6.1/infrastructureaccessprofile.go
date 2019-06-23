/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// InfrastructureaccessprofileIdentity represents the Identity of the object
var InfrastructureaccessprofileIdentity = bambou.Identity{
	Name:     "infrastructureaccessprofile",
	Category: "infrastructureaccessprofiles",
}

// InfrastructureaccessprofilesList represents a list of Infrastructureaccessprofiles
type InfrastructureaccessprofilesList []*Infrastructureaccessprofile

// InfrastructureaccessprofilesAncestor is the interface that an ancestor of a Infrastructureaccessprofile must implement.
// An Ancestor is defined as an entity that has Infrastructureaccessprofile as a descendant.
// An Ancestor can get a list of its child Infrastructureaccessprofiles, but not necessarily create one.
type InfrastructureaccessprofilesAncestor interface {
	Infrastructureaccessprofiles(*bambou.FetchingInfo) (InfrastructureaccessprofilesList, *bambou.Error)
}

// InfrastructureaccessprofilesParent is the interface that a parent of a Infrastructureaccessprofile must implement.
// A Parent is defined as an entity that has Infrastructureaccessprofile as a child.
// A Parent is an Ancestor which can create a Infrastructureaccessprofile.
type InfrastructureaccessprofilesParent interface {
	InfrastructureaccessprofilesAncestor
	CreateInfrastructureaccessprofile(*Infrastructureaccessprofile) *bambou.Error
}

// Infrastructureaccessprofile represents the model of a infrastructureaccessprofile
type Infrastructureaccessprofile struct {
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

// NewInfrastructureaccessprofile returns a new *Infrastructureaccessprofile
func NewInfrastructureaccessprofile() *Infrastructureaccessprofile {

	return &Infrastructureaccessprofile{
		SSHAuthMode:    "PASSWORD_AND_KEY_BASED",
		SourceIPFilter: "DISABLED",
		UserName:       "nuage",
	}
}

// Identity returns the Identity of the object.
func (o *Infrastructureaccessprofile) Identity() bambou.Identity {

	return InfrastructureaccessprofileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Infrastructureaccessprofile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Infrastructureaccessprofile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Infrastructureaccessprofile from the server
func (o *Infrastructureaccessprofile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Infrastructureaccessprofile into the server
func (o *Infrastructureaccessprofile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Infrastructureaccessprofile from the server
func (o *Infrastructureaccessprofile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the Infrastructureaccessprofile
func (o *Infrastructureaccessprofile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Infrastructureaccessprofile
func (o *Infrastructureaccessprofile) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Infrastructureaccessprofile
func (o *Infrastructureaccessprofile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Infrastructureaccessprofile
func (o *Infrastructureaccessprofile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
