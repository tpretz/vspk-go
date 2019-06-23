/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// LicenseStatusIdentity represents the Identity of the object
var LicenseStatusIdentity = bambou.Identity{
	Name:     "licensestatus",
	Category: "licensestatus",
}

// LicenseStatusList represents a list of LicenseStatus
type LicenseStatusList []*LicenseStatus

// LicenseStatusAncestor is the interface that an ancestor of a LicenseStatus must implement.
// An Ancestor is defined as an entity that has LicenseStatus as a descendant.
// An Ancestor can get a list of its child LicenseStatus, but not necessarily create one.
type LicenseStatusAncestor interface {
	LicenseStatus(*bambou.FetchingInfo) (LicenseStatusList, *bambou.Error)
}

// LicenseStatusParent is the interface that a parent of a LicenseStatus must implement.
// A Parent is defined as an entity that has LicenseStatus as a child.
// A Parent is an Ancestor which can create a LicenseStatus.
type LicenseStatusParent interface {
	LicenseStatusAncestor
	CreateLicenseStatus(*LicenseStatus) *bambou.Error
}

// LicenseStatus represents the model of a licensestatus
type LicenseStatus struct {
	ID                          string `json:"ID,omitempty"`
	ParentID                    string `json:"parentID,omitempty"`
	ParentType                  string `json:"parentType,omitempty"`
	Owner                       string `json:"owner,omitempty"`
	TotalLicensedNICsCount      string `json:"totalLicensedNICsCount,omitempty"`
	TotalLicensedNSGsCount      string `json:"totalLicensedNSGsCount,omitempty"`
	TotalLicensedUsedNICsCount  string `json:"totalLicensedUsedNICsCount,omitempty"`
	TotalLicensedUsedNSGsCount  string `json:"totalLicensedUsedNSGsCount,omitempty"`
	TotalLicensedUsedVMsCount   string `json:"totalLicensedUsedVMsCount,omitempty"`
	TotalLicensedUsedVRSGsCount string `json:"totalLicensedUsedVRSGsCount,omitempty"`
	TotalLicensedUsedVRSsCount  string `json:"totalLicensedUsedVRSsCount,omitempty"`
	TotalLicensedVMsCount       string `json:"totalLicensedVMsCount,omitempty"`
	TotalLicensedVRSGsCount     string `json:"totalLicensedVRSGsCount,omitempty"`
	TotalLicensedVRSsCount      string `json:"totalLicensedVRSsCount,omitempty"`
}

// NewLicenseStatus returns a new *LicenseStatus
func NewLicenseStatus() *LicenseStatus {

	return &LicenseStatus{}
}

// Identity returns the Identity of the object.
func (o *LicenseStatus) Identity() bambou.Identity {

	return LicenseStatusIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *LicenseStatus) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *LicenseStatus) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the LicenseStatus from the server
func (o *LicenseStatus) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the LicenseStatus into the server
func (o *LicenseStatus) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the LicenseStatus from the server
func (o *LicenseStatus) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
