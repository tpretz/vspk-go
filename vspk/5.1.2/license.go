/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// LicenseIdentity represents the Identity of the object
var LicenseIdentity = bambou.Identity{
	Name:     "license",
	Category: "licenses",
}

// LicensesList represents a list of Licenses
type LicensesList []*License

// LicensesAncestor is the interface that an ancestor of a License must implement.
// An Ancestor is defined as an entity that has License as a descendant.
// An Ancestor can get a list of its child Licenses, but not necessarily create one.
type LicensesAncestor interface {
	Licenses(*bambou.FetchingInfo) (LicensesList, *bambou.Error)
}

// LicensesParent is the interface that a parent of a License must implement.
// A Parent is defined as an entity that has License as a child.
// A Parent is an Ancestor which can create a License.
type LicensesParent interface {
	LicensesAncestor
	CreateLicense(*License) *bambou.Error
}

// License represents the model of a license
type License struct {
	ID                          string  `json:"ID,omitempty"`
	ParentID                    string  `json:"parentID,omitempty"`
	ParentType                  string  `json:"parentType,omitempty"`
	Owner                       string  `json:"owner,omitempty"`
	MajorRelease                int     `json:"majorRelease,omitempty"`
	LastUpdatedBy               string  `json:"lastUpdatedBy,omitempty"`
	AdditionalSupportedVersions string  `json:"additionalSupportedVersions,omitempty"`
	Phone                       string  `json:"phone,omitempty"`
	License                     string  `json:"license,omitempty"`
	LicenseEncryption           string  `json:"licenseEncryption,omitempty"`
	LicenseEntities             string  `json:"licenseEntities,omitempty"`
	LicenseID                   int     `json:"licenseID,omitempty"`
	LicenseType                 string  `json:"licenseType,omitempty"`
	MinorRelease                int     `json:"minorRelease,omitempty"`
	Zip                         string  `json:"zip,omitempty"`
	City                        string  `json:"city,omitempty"`
	AllowedAVRSGsCount          int     `json:"allowedAVRSGsCount,omitempty"`
	AllowedAVRSsCount           int     `json:"allowedAVRSsCount,omitempty"`
	AllowedCPEsCount            int     `json:"allowedCPEsCount,omitempty"`
	AllowedNICsCount            int     `json:"allowedNICsCount,omitempty"`
	AllowedVMsCount             int     `json:"allowedVMsCount,omitempty"`
	AllowedVRSGsCount           int     `json:"allowedVRSGsCount,omitempty"`
	AllowedVRSsCount            int     `json:"allowedVRSsCount,omitempty"`
	Email                       string  `json:"email,omitempty"`
	EncryptionMode              bool    `json:"encryptionMode"`
	UniqueLicenseIdentifier     string  `json:"uniqueLicenseIdentifier,omitempty"`
	EntityScope                 string  `json:"entityScope,omitempty"`
	Company                     string  `json:"company,omitempty"`
	Country                     string  `json:"country,omitempty"`
	ProductVersion              string  `json:"productVersion,omitempty"`
	Provider                    string  `json:"provider,omitempty"`
	IsClusterLicense            bool    `json:"isClusterLicense"`
	UserName                    string  `json:"userName,omitempty"`
	State                       string  `json:"state,omitempty"`
	Street                      string  `json:"street,omitempty"`
	CustomerKey                 string  `json:"customerKey,omitempty"`
	ExpirationDate              float64 `json:"expirationDate,omitempty"`
	ExpiryTimestamp             int     `json:"expiryTimestamp,omitempty"`
	ExternalID                  string  `json:"externalID,omitempty"`
}

// NewLicense returns a new *License
func NewLicense() *License {

	return &License{}
}

// Identity returns the Identity of the object.
func (o *License) Identity() bambou.Identity {

	return LicenseIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *License) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *License) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the License from the server
func (o *License) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the License into the server
func (o *License) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the License from the server
func (o *License) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the License
func (o *License) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the License
func (o *License) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the License
func (o *License) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the License
func (o *License) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the License
func (o *License) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
