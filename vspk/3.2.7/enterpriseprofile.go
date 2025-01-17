/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// EnterpriseProfileIdentity represents the Identity of the object
var EnterpriseProfileIdentity = bambou.Identity{
	Name:     "enterpriseprofile",
	Category: "enterpriseprofiles",
}

// EnterpriseProfilesList represents a list of EnterpriseProfiles
type EnterpriseProfilesList []*EnterpriseProfile

// EnterpriseProfilesAncestor is the interface that an ancestor of a EnterpriseProfile must implement.
// An Ancestor is defined as an entity that has EnterpriseProfile as a descendant.
// An Ancestor can get a list of its child EnterpriseProfiles, but not necessarily create one.
type EnterpriseProfilesAncestor interface {
	EnterpriseProfiles(*bambou.FetchingInfo) (EnterpriseProfilesList, *bambou.Error)
}

// EnterpriseProfilesParent is the interface that a parent of a EnterpriseProfile must implement.
// A Parent is defined as an entity that has EnterpriseProfile as a child.
// A Parent is an Ancestor which can create a EnterpriseProfile.
type EnterpriseProfilesParent interface {
	EnterpriseProfilesAncestor
	CreateEnterpriseProfile(*EnterpriseProfile) *bambou.Error
}

// EnterpriseProfile represents the model of a enterpriseprofile
type EnterpriseProfile struct {
	ID                            string        `json:"ID,omitempty"`
	ParentID                      string        `json:"parentID,omitempty"`
	ParentType                    string        `json:"parentType,omitempty"`
	Owner                         string        `json:"owner,omitempty"`
	DHCPLeaseInterval             int           `json:"DHCPLeaseInterval,omitempty"`
	Name                          string        `json:"name,omitempty"`
	LastUpdatedBy                 string        `json:"lastUpdatedBy,omitempty"`
	ReceiveMultiCastListID        string        `json:"receiveMultiCastListID,omitempty"`
	SendMultiCastListID           string        `json:"sendMultiCastListID,omitempty"`
	Description                   string        `json:"description,omitempty"`
	AllowAdvancedQOSConfiguration bool          `json:"allowAdvancedQOSConfiguration"`
	AllowGatewayManagement        bool          `json:"allowGatewayManagement"`
	AllowTrustedForwardingClass   bool          `json:"allowTrustedForwardingClass"`
	AllowedForwardingClasses      []interface{} `json:"allowedForwardingClasses,omitempty"`
	FloatingIPsQuota              int           `json:"floatingIPsQuota,omitempty"`
	EncryptionManagementMode      string        `json:"encryptionManagementMode,omitempty"`
	EntityScope                   string        `json:"entityScope,omitempty"`
	ExternalID                    string        `json:"externalID,omitempty"`
}

// NewEnterpriseProfile returns a new *EnterpriseProfile
func NewEnterpriseProfile() *EnterpriseProfile {

	return &EnterpriseProfile{}
}

// Identity returns the Identity of the object.
func (o *EnterpriseProfile) Identity() bambou.Identity {

	return EnterpriseProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EnterpriseProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EnterpriseProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the EnterpriseProfile from the server
func (o *EnterpriseProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the EnterpriseProfile into the server
func (o *EnterpriseProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the EnterpriseProfile from the server
func (o *EnterpriseProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the EnterpriseProfile
func (o *EnterpriseProfile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the EnterpriseProfile
func (o *EnterpriseProfile) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the EnterpriseProfile
func (o *EnterpriseProfile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the EnterpriseProfile
func (o *EnterpriseProfile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Enterprises retrieves the list of child Enterprises of the EnterpriseProfile
func (o *EnterpriseProfile) Enterprises(info *bambou.FetchingInfo) (EnterprisesList, *bambou.Error) {

	var list EnterprisesList
	err := bambou.CurrentSession().FetchChildren(o, EnterpriseIdentity, &list, info)
	return list, err
}

// AssignEnterprises assigns the list of Enterprises to the EnterpriseProfile
func (o *EnterpriseProfile) AssignEnterprises(children EnterprisesList) *bambou.Error {

	list := []bambou.Identifiable{}
	for _, c := range children {
		list = append(list, c)
	}

	return bambou.CurrentSession().AssignChildren(o, list, EnterpriseIdentity)
}

// MultiCastLists retrieves the list of child MultiCastLists of the EnterpriseProfile
func (o *EnterpriseProfile) MultiCastLists(info *bambou.FetchingInfo) (MultiCastListsList, *bambou.Error) {

	var list MultiCastListsList
	err := bambou.CurrentSession().FetchChildren(o, MultiCastListIdentity, &list, info)
	return list, err
}

// EventLogs retrieves the list of child EventLogs of the EnterpriseProfile
func (o *EnterpriseProfile) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}

// ExternalServices retrieves the list of child ExternalServices of the EnterpriseProfile
func (o *EnterpriseProfile) ExternalServices(info *bambou.FetchingInfo) (ExternalServicesList, *bambou.Error) {

	var list ExternalServicesList
	err := bambou.CurrentSession().FetchChildren(o, ExternalServiceIdentity, &list, info)
	return list, err
}

// AssignExternalServices assigns the list of ExternalServices to the EnterpriseProfile
func (o *EnterpriseProfile) AssignExternalServices(children ExternalServicesList) *bambou.Error {

	list := []bambou.Identifiable{}
	for _, c := range children {
		list = append(list, c)
	}

	return bambou.CurrentSession().AssignChildren(o, list, ExternalServiceIdentity)
}
