/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// LDAPConfigurationIdentity represents the Identity of the object
var LDAPConfigurationIdentity = bambou.Identity{
	Name:     "ldapconfiguration",
	Category: "ldapconfigurations",
}

// LDAPConfigurationsList represents a list of LDAPConfigurations
type LDAPConfigurationsList []*LDAPConfiguration

// LDAPConfigurationsAncestor is the interface that an ancestor of a LDAPConfiguration must implement.
// An Ancestor is defined as an entity that has LDAPConfiguration as a descendant.
// An Ancestor can get a list of its child LDAPConfigurations, but not necessarily create one.
type LDAPConfigurationsAncestor interface {
	LDAPConfigurations(*bambou.FetchingInfo) (LDAPConfigurationsList, *bambou.Error)
}

// LDAPConfigurationsParent is the interface that a parent of a LDAPConfiguration must implement.
// A Parent is defined as an entity that has LDAPConfiguration as a child.
// A Parent is an Ancestor which can create a LDAPConfiguration.
type LDAPConfigurationsParent interface {
	LDAPConfigurationsAncestor
	CreateLDAPConfiguration(*LDAPConfiguration) *bambou.Error
}

// LDAPConfiguration represents the model of a ldapconfiguration
type LDAPConfiguration struct {
	ID                    string `json:"ID,omitempty"`
	ParentID              string `json:"parentID,omitempty"`
	ParentType            string `json:"parentType,omitempty"`
	Owner                 string `json:"owner,omitempty"`
	SSLEnabled            bool   `json:"SSLEnabled"`
	Password              string `json:"password,omitempty"`
	LastUpdatedBy         string `json:"lastUpdatedBy,omitempty"`
	AcceptAllCertificates bool   `json:"acceptAllCertificates"`
	Certificate           string `json:"certificate,omitempty"`
	Server                string `json:"server,omitempty"`
	Enabled               bool   `json:"enabled"`
	EntityScope           string `json:"entityScope,omitempty"`
	Port                  string `json:"port,omitempty"`
	GroupDN               string `json:"groupDN,omitempty"`
	GroupNamePrefix       string `json:"groupNamePrefix,omitempty"`
	GroupNameSuffix       string `json:"groupNameSuffix,omitempty"`
	UserDNTemplate        string `json:"userDNTemplate,omitempty"`
	UserNameAttribute     string `json:"userNameAttribute,omitempty"`
	AuthorizationEnabled  bool   `json:"authorizationEnabled"`
	AuthorizingUserDN     string `json:"authorizingUserDN,omitempty"`
	ExternalID            string `json:"externalID,omitempty"`
}

// NewLDAPConfiguration returns a new *LDAPConfiguration
func NewLDAPConfiguration() *LDAPConfiguration {

	return &LDAPConfiguration{}
}

// Identity returns the Identity of the object.
func (o *LDAPConfiguration) Identity() bambou.Identity {

	return LDAPConfigurationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *LDAPConfiguration) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *LDAPConfiguration) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the LDAPConfiguration from the server
func (o *LDAPConfiguration) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the LDAPConfiguration into the server
func (o *LDAPConfiguration) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the LDAPConfiguration from the server
func (o *LDAPConfiguration) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the LDAPConfiguration
func (o *LDAPConfiguration) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the LDAPConfiguration
func (o *LDAPConfiguration) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the LDAPConfiguration
func (o *LDAPConfiguration) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the LDAPConfiguration
func (o *LDAPConfiguration) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
