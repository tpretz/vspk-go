/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// AddressMapIdentity represents the Identity of the object
var AddressMapIdentity = bambou.Identity{
	Name:     "addressmap",
	Category: "addressmaps",
}

// AddressMapsList represents a list of AddressMaps
type AddressMapsList []*AddressMap

// AddressMapsAncestor is the interface that an ancestor of a AddressMap must implement.
// An Ancestor is defined as an entity that has AddressMap as a descendant.
// An Ancestor can get a list of its child AddressMaps, but not necessarily create one.
type AddressMapsAncestor interface {
	AddressMaps(*bambou.FetchingInfo) (AddressMapsList, *bambou.Error)
}

// AddressMapsParent is the interface that a parent of a AddressMap must implement.
// A Parent is defined as an entity that has AddressMap as a child.
// A Parent is an Ancestor which can create a AddressMap.
type AddressMapsParent interface {
	AddressMapsAncestor
	CreateAddressMap(*AddressMap) *bambou.Error
}

// AddressMap represents the model of a addressmap
type AddressMap struct {
	ID                     string `json:"ID,omitempty"`
	ParentID               string `json:"parentID,omitempty"`
	ParentType             string `json:"parentType,omitempty"`
	Owner                  string `json:"owner,omitempty"`
	LastUpdatedBy          string `json:"lastUpdatedBy,omitempty"`
	EntityScope            string `json:"entityScope,omitempty"`
	PrivateIP              string `json:"privateIP,omitempty"`
	PrivatePort            int    `json:"privatePort,omitempty"`
	AssociatedPATNATPoolID string `json:"associatedPATNATPoolID,omitempty"`
	PublicIP               string `json:"publicIP,omitempty"`
	PublicPort             int    `json:"publicPort,omitempty"`
	ExternalID             string `json:"externalID,omitempty"`
	Type                   string `json:"type,omitempty"`
}

// NewAddressMap returns a new *AddressMap
func NewAddressMap() *AddressMap {

	return &AddressMap{}
}

// Identity returns the Identity of the object.
func (o *AddressMap) Identity() bambou.Identity {

	return AddressMapIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AddressMap) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AddressMap) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the AddressMap from the server
func (o *AddressMap) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the AddressMap into the server
func (o *AddressMap) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the AddressMap from the server
func (o *AddressMap) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the AddressMap
func (o *AddressMap) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the AddressMap
func (o *AddressMap) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the AddressMap
func (o *AddressMap) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the AddressMap
func (o *AddressMap) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Statistics retrieves the list of child Statistics of the AddressMap
func (o *AddressMap) Statistics(info *bambou.FetchingInfo) (StatisticsList, *bambou.Error) {

	var list StatisticsList
	err := bambou.CurrentSession().FetchChildren(o, StatisticsIdentity, &list, info)
	return list, err
}

// StatisticsPolicies retrieves the list of child StatisticsPolicies of the AddressMap
func (o *AddressMap) StatisticsPolicies(info *bambou.FetchingInfo) (StatisticsPoliciesList, *bambou.Error) {

	var list StatisticsPoliciesList
	err := bambou.CurrentSession().FetchChildren(o, StatisticsPolicyIdentity, &list, info)
	return list, err
}

// CreateStatisticsPolicy creates a new child StatisticsPolicy under the AddressMap
func (o *AddressMap) CreateStatisticsPolicy(child *StatisticsPolicy) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
