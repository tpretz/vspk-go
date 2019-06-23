/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// LicensestatusIdentity represents the Identity of the object
var LicensestatusIdentity = bambou.Identity{
	Name:     "licensestatus",
	Category: "licensestatus",
}

// LicensestatusList represents a list of Licensestatus
type LicensestatusList []*Licensestatus

// LicensestatusAncestor is the interface that an ancestor of a Licensestatus must implement.
// An Ancestor is defined as an entity that has Licensestatus as a descendant.
// An Ancestor can get a list of its child Licensestatus, but not necessarily create one.
type LicensestatusAncestor interface {
	Licensestatus(*bambou.FetchingInfo) (LicensestatusList, *bambou.Error)
}

// LicensestatusParent is the interface that a parent of a Licensestatus must implement.
// A Parent is defined as an entity that has Licensestatus as a child.
// A Parent is an Ancestor which can create a Licensestatus.
type LicensestatusParent interface {
	LicensestatusAncestor
	CreateLicensestatus(*Licensestatus) *bambou.Error
}

// Licensestatus represents the model of a licensestatus
type Licensestatus struct {
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

// NewLicensestatus returns a new *Licensestatus
func NewLicensestatus() *Licensestatus {

	return &Licensestatus{}
}

// Identity returns the Identity of the object.
func (o *Licensestatus) Identity() bambou.Identity {

	return LicensestatusIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Licensestatus) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Licensestatus) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Licensestatus from the server
func (o *Licensestatus) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Licensestatus into the server
func (o *Licensestatus) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Licensestatus from the server
func (o *Licensestatus) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
