/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// NSGInfoIdentity represents the Identity of the object
var NSGInfoIdentity = bambou.Identity{
	Name:     "nsginfo",
	Category: "nsginfos",
}

// NSGInfosList represents a list of NSGInfos
type NSGInfosList []*NSGInfo

// NSGInfosAncestor is the interface that an ancestor of a NSGInfo must implement.
// An Ancestor is defined as an entity that has NSGInfo as a descendant.
// An Ancestor can get a list of its child NSGInfos, but not necessarily create one.
type NSGInfosAncestor interface {
	NSGInfos(*bambou.FetchingInfo) (NSGInfosList, *bambou.Error)
}

// NSGInfosParent is the interface that a parent of a NSGInfo must implement.
// A Parent is defined as an entity that has NSGInfo as a child.
// A Parent is an Ancestor which can create a NSGInfo.
type NSGInfosParent interface {
	NSGInfosAncestor
	CreateNSGInfo(*NSGInfo) *bambou.Error
}

// NSGInfo represents the model of a nsginfo
type NSGInfo struct {
	ID                        string      `json:"ID,omitempty"`
	ParentID                  string      `json:"parentID,omitempty"`
	ParentType                string      `json:"parentType,omitempty"`
	Owner                     string      `json:"owner,omitempty"`
	MACAddress                string      `json:"MACAddress,omitempty"`
	AARApplicationReleaseDate string      `json:"AARApplicationReleaseDate,omitempty"`
	AARApplicationVersion     string      `json:"AARApplicationVersion,omitempty"`
	BIOSReleaseDate           string      `json:"BIOSReleaseDate,omitempty"`
	BIOSVersion               string      `json:"BIOSVersion,omitempty"`
	SKU                       string      `json:"SKU,omitempty"`
	TPMStatus                 int         `json:"TPMStatus,omitempty"`
	TPMVersion                string      `json:"TPMVersion,omitempty"`
	CPUType                   string      `json:"CPUType,omitempty"`
	NSGVersion                string      `json:"NSGVersion,omitempty"`
	UUID                      string      `json:"UUID,omitempty"`
	Name                      string      `json:"name,omitempty"`
	Family                    string      `json:"family,omitempty"`
	PatchesDetail             string      `json:"patchesDetail,omitempty"`
	SerialNumber              string      `json:"serialNumber,omitempty"`
	Personality               string      `json:"personality,omitempty"`
	Libraries                 string      `json:"libraries,omitempty"`
	CmdDetailedStatus         string      `json:"cmdDetailedStatus,omitempty"`
	CmdDetailedStatusCode     int         `json:"cmdDetailedStatusCode,omitempty"`
	CmdDownloadProgress       interface{} `json:"cmdDownloadProgress,omitempty"`
	CmdID                     string      `json:"cmdID,omitempty"`
	CmdLastUpdatedDate        string      `json:"cmdLastUpdatedDate,omitempty"`
	CmdStatus                 string      `json:"cmdStatus,omitempty"`
	CmdType                   string      `json:"cmdType,omitempty"`
	EnterpriseID              string      `json:"enterpriseID,omitempty"`
	EnterpriseName            string      `json:"enterpriseName,omitempty"`
	EntityScope               string      `json:"entityScope,omitempty"`
	BootstrapStatus           string      `json:"bootstrapStatus,omitempty"`
	ProductName               string      `json:"productName,omitempty"`
	AssociatedEntityType      string      `json:"associatedEntityType,omitempty"`
	AssociatedNSGatewayID     string      `json:"associatedNSGatewayID,omitempty"`
	ExternalID                string      `json:"externalID,omitempty"`
	SystemID                  string      `json:"systemID,omitempty"`
}

// NewNSGInfo returns a new *NSGInfo
func NewNSGInfo() *NSGInfo {

	return &NSGInfo{}
}

// Identity returns the Identity of the object.
func (o *NSGInfo) Identity() bambou.Identity {

	return NSGInfoIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NSGInfo) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NSGInfo) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NSGInfo from the server
func (o *NSGInfo) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NSGInfo into the server
func (o *NSGInfo) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NSGInfo from the server
func (o *NSGInfo) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
