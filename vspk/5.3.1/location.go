/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// LocationIdentity represents the Identity of the object
var LocationIdentity = bambou.Identity{
	Name:     "location",
	Category: "locations",
}

// LocationsList represents a list of Locations
type LocationsList []*Location

// LocationsAncestor is the interface that an ancestor of a Location must implement.
// An Ancestor is defined as an entity that has Location as a descendant.
// An Ancestor can get a list of its child Locations, but not necessarily create one.
type LocationsAncestor interface {
	Locations(*bambou.FetchingInfo) (LocationsList, *bambou.Error)
}

// LocationsParent is the interface that a parent of a Location must implement.
// A Parent is defined as an entity that has Location as a child.
// A Parent is an Ancestor which can create a Location.
type LocationsParent interface {
	LocationsAncestor
	CreateLocation(*Location) *bambou.Error
}

// Location represents the model of a location
type Location struct {
	ID                   string  `json:"ID,omitempty"`
	ParentID             string  `json:"parentID,omitempty"`
	ParentType           string  `json:"parentType,omitempty"`
	Owner                string  `json:"owner,omitempty"`
	LastUpdatedBy        string  `json:"lastUpdatedBy,omitempty"`
	Latitude             float64 `json:"latitude,omitempty"`
	Address              string  `json:"address,omitempty"`
	IgnoreGeocode        bool    `json:"ignoreGeocode"`
	TimeZoneID           string  `json:"timeZoneID,omitempty"`
	EntityScope          string  `json:"entityScope,omitempty"`
	Locality             string  `json:"locality,omitempty"`
	Longitude            float64 `json:"longitude,omitempty"`
	Country              string  `json:"country,omitempty"`
	AssociatedEntityID   string  `json:"associatedEntityID,omitempty"`
	AssociatedEntityName string  `json:"associatedEntityName,omitempty"`
	AssociatedEntityType string  `json:"associatedEntityType,omitempty"`
	State                string  `json:"state,omitempty"`
	ExternalID           string  `json:"externalID,omitempty"`
}

// NewLocation returns a new *Location
func NewLocation() *Location {

	return &Location{}
}

// Identity returns the Identity of the object.
func (o *Location) Identity() bambou.Identity {

	return LocationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Location) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Location) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Location from the server
func (o *Location) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Location into the server
func (o *Location) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Location from the server
func (o *Location) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the Location
func (o *Location) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Location
func (o *Location) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Location
func (o *Location) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Location
func (o *Location) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
