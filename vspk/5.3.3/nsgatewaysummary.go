/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// NSGatewaySummaryIdentity represents the Identity of the object
var NSGatewaySummaryIdentity = bambou.Identity{
	Name:     "nsgatewayssummary",
	Category: "nsgatewayssummaries",
}

// NSGatewaySummariesList represents a list of NSGatewaySummaries
type NSGatewaySummariesList []*NSGatewaySummary

// NSGatewaySummariesAncestor is the interface that an ancestor of a NSGatewaySummary must implement.
// An Ancestor is defined as an entity that has NSGatewaySummary as a descendant.
// An Ancestor can get a list of its child NSGatewaySummaries, but not necessarily create one.
type NSGatewaySummariesAncestor interface {
	NSGatewaySummaries(*bambou.FetchingInfo) (NSGatewaySummariesList, *bambou.Error)
}

// NSGatewaySummariesParent is the interface that a parent of a NSGatewaySummary must implement.
// A Parent is defined as an entity that has NSGatewaySummary as a child.
// A Parent is an Ancestor which can create a NSGatewaySummary.
type NSGatewaySummariesParent interface {
	NSGatewaySummariesAncestor
	CreateNSGatewaySummary(*NSGatewaySummary) *bambou.Error
}

// NSGatewaySummary represents the model of a nsgatewayssummary
type NSGatewaySummary struct {
	ID                  string  `json:"ID,omitempty"`
	ParentID            string  `json:"parentID,omitempty"`
	ParentType          string  `json:"parentType,omitempty"`
	Owner               string  `json:"owner,omitempty"`
	NSGVersion          string  `json:"NSGVersion,omitempty"`
	MajorAlarmsCount    int     `json:"majorAlarmsCount,omitempty"`
	LastUpdatedBy       string  `json:"lastUpdatedBy,omitempty"`
	GatewayID           string  `json:"gatewayID,omitempty"`
	GatewayName         string  `json:"gatewayName,omitempty"`
	GatewayType         string  `json:"gatewayType,omitempty"`
	Latitude            float64 `json:"latitude,omitempty"`
	Address             string  `json:"address,omitempty"`
	TimezoneID          string  `json:"timezoneID,omitempty"`
	MinorAlarmsCount    int     `json:"minorAlarmsCount,omitempty"`
	InfoAlarmsCount     int     `json:"infoAlarmsCount,omitempty"`
	EnterpriseID        string  `json:"enterpriseID,omitempty"`
	EntityScope         string  `json:"entityScope,omitempty"`
	Locality            string  `json:"locality,omitempty"`
	Longitude           float64 `json:"longitude,omitempty"`
	BootstrapStatus     string  `json:"bootstrapStatus,omitempty"`
	Country             string  `json:"country,omitempty"`
	CriticalAlarmsCount int     `json:"criticalAlarmsCount,omitempty"`
	State               string  `json:"state,omitempty"`
	ExternalID          string  `json:"externalID,omitempty"`
	SystemID            string  `json:"systemID,omitempty"`
}

// NewNSGatewaySummary returns a new *NSGatewaySummary
func NewNSGatewaySummary() *NSGatewaySummary {

	return &NSGatewaySummary{}
}

// Identity returns the Identity of the object.
func (o *NSGatewaySummary) Identity() bambou.Identity {

	return NSGatewaySummaryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NSGatewaySummary) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NSGatewaySummary) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NSGatewaySummary from the server
func (o *NSGatewaySummary) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NSGatewaySummary into the server
func (o *NSGatewaySummary) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NSGatewaySummary from the server
func (o *NSGatewaySummary) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the NSGatewaySummary
func (o *NSGatewaySummary) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the NSGatewaySummary
func (o *NSGatewaySummary) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the NSGatewaySummary
func (o *NSGatewaySummary) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the NSGatewaySummary
func (o *NSGatewaySummary) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
