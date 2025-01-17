/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// CustomPropertyIdentity represents the Identity of the object
var CustomPropertyIdentity = bambou.Identity{
	Name:     "customproperty",
	Category: "customproperties",
}

// CustomPropertiesList represents a list of CustomProperties
type CustomPropertiesList []*CustomProperty

// CustomPropertiesAncestor is the interface that an ancestor of a CustomProperty must implement.
// An Ancestor is defined as an entity that has CustomProperty as a descendant.
// An Ancestor can get a list of its child CustomProperties, but not necessarily create one.
type CustomPropertiesAncestor interface {
	CustomProperties(*bambou.FetchingInfo) (CustomPropertiesList, *bambou.Error)
}

// CustomPropertiesParent is the interface that a parent of a CustomProperty must implement.
// A Parent is defined as an entity that has CustomProperty as a child.
// A Parent is an Ancestor which can create a CustomProperty.
type CustomPropertiesParent interface {
	CustomPropertiesAncestor
	CreateCustomProperty(*CustomProperty) *bambou.Error
}

// CustomProperty represents the model of a customproperty
type CustomProperty struct {
	ID             string `json:"ID,omitempty"`
	ParentID       string `json:"parentID,omitempty"`
	ParentType     string `json:"parentType,omitempty"`
	Owner          string `json:"owner,omitempty"`
	LastUpdatedBy  string `json:"lastUpdatedBy,omitempty"`
	EntityScope    string `json:"entityScope,omitempty"`
	AttributeName  string `json:"attributeName,omitempty"`
	AttributeValue string `json:"attributeValue,omitempty"`
	ExternalID     string `json:"externalID,omitempty"`
}

// NewCustomProperty returns a new *CustomProperty
func NewCustomProperty() *CustomProperty {

	return &CustomProperty{}
}

// Identity returns the Identity of the object.
func (o *CustomProperty) Identity() bambou.Identity {

	return CustomPropertyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CustomProperty) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CustomProperty) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the CustomProperty from the server
func (o *CustomProperty) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the CustomProperty into the server
func (o *CustomProperty) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the CustomProperty from the server
func (o *CustomProperty) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the CustomProperty
func (o *CustomProperty) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the CustomProperty
func (o *CustomProperty) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the CustomProperty
func (o *CustomProperty) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the CustomProperty
func (o *CustomProperty) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
