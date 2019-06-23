/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// GatewaySecurityProfileIdentity represents the Identity of the object
var GatewaySecurityProfileIdentity = bambou.Identity{
	Name:     "gatewaysecurityprofile",
	Category: "gatewaysecurityprofiles",
}

// GatewaySecurityProfilesList represents a list of GatewaySecurityProfiles
type GatewaySecurityProfilesList []*GatewaySecurityProfile

// GatewaySecurityProfilesAncestor is the interface that an ancestor of a GatewaySecurityProfile must implement.
// An Ancestor is defined as an entity that has GatewaySecurityProfile as a descendant.
// An Ancestor can get a list of its child GatewaySecurityProfiles, but not necessarily create one.
type GatewaySecurityProfilesAncestor interface {
	GatewaySecurityProfiles(*bambou.FetchingInfo) (GatewaySecurityProfilesList, *bambou.Error)
}

// GatewaySecurityProfilesParent is the interface that a parent of a GatewaySecurityProfile must implement.
// A Parent is defined as an entity that has GatewaySecurityProfile as a child.
// A Parent is an Ancestor which can create a GatewaySecurityProfile.
type GatewaySecurityProfilesParent interface {
	GatewaySecurityProfilesAncestor
	CreateGatewaySecurityProfile(*GatewaySecurityProfile) *bambou.Error
}

// GatewaySecurityProfile represents the model of a gatewaysecurityprofile
type GatewaySecurityProfile struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	GatewayID     string `json:"gatewayID,omitempty"`
	Revision      int    `json:"revision,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewGatewaySecurityProfile returns a new *GatewaySecurityProfile
func NewGatewaySecurityProfile() *GatewaySecurityProfile {

	return &GatewaySecurityProfile{}
}

// Identity returns the Identity of the object.
func (o *GatewaySecurityProfile) Identity() bambou.Identity {

	return GatewaySecurityProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *GatewaySecurityProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *GatewaySecurityProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the GatewaySecurityProfile from the server
func (o *GatewaySecurityProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the GatewaySecurityProfile into the server
func (o *GatewaySecurityProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the GatewaySecurityProfile from the server
func (o *GatewaySecurityProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the GatewaySecurityProfile
func (o *GatewaySecurityProfile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the GatewaySecurityProfile
func (o *GatewaySecurityProfile) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the GatewaySecurityProfile
func (o *GatewaySecurityProfile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the GatewaySecurityProfile
func (o *GatewaySecurityProfile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
