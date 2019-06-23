/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// OverlayMirrorDestinationTemplateIdentity represents the Identity of the object
var OverlayMirrorDestinationTemplateIdentity = bambou.Identity{
	Name:     "overlaymirrordestinationtemplate",
	Category: "overlaymirrordestinationtemplates",
}

// OverlayMirrorDestinationTemplatesList represents a list of OverlayMirrorDestinationTemplates
type OverlayMirrorDestinationTemplatesList []*OverlayMirrorDestinationTemplate

// OverlayMirrorDestinationTemplatesAncestor is the interface that an ancestor of a OverlayMirrorDestinationTemplate must implement.
// An Ancestor is defined as an entity that has OverlayMirrorDestinationTemplate as a descendant.
// An Ancestor can get a list of its child OverlayMirrorDestinationTemplates, but not necessarily create one.
type OverlayMirrorDestinationTemplatesAncestor interface {
	OverlayMirrorDestinationTemplates(*bambou.FetchingInfo) (OverlayMirrorDestinationTemplatesList, *bambou.Error)
}

// OverlayMirrorDestinationTemplatesParent is the interface that a parent of a OverlayMirrorDestinationTemplate must implement.
// A Parent is defined as an entity that has OverlayMirrorDestinationTemplate as a child.
// A Parent is an Ancestor which can create a OverlayMirrorDestinationTemplate.
type OverlayMirrorDestinationTemplatesParent interface {
	OverlayMirrorDestinationTemplatesAncestor
	CreateOverlayMirrorDestinationTemplate(*OverlayMirrorDestinationTemplate) *bambou.Error
}

// OverlayMirrorDestinationTemplate represents the model of a overlaymirrordestinationtemplate
type OverlayMirrorDestinationTemplate struct {
	ID                string `json:"ID,omitempty"`
	ParentID          string `json:"parentID,omitempty"`
	ParentType        string `json:"parentType,omitempty"`
	Owner             string `json:"owner,omitempty"`
	Name              string `json:"name,omitempty"`
	RedundancyEnabled bool   `json:"redundancyEnabled"`
	Description       string `json:"description,omitempty"`
	EndPointType      string `json:"endPointType,omitempty"`
	TriggerType       string `json:"triggerType,omitempty"`
}

// NewOverlayMirrorDestinationTemplate returns a new *OverlayMirrorDestinationTemplate
func NewOverlayMirrorDestinationTemplate() *OverlayMirrorDestinationTemplate {

	return &OverlayMirrorDestinationTemplate{}
}

// Identity returns the Identity of the object.
func (o *OverlayMirrorDestinationTemplate) Identity() bambou.Identity {

	return OverlayMirrorDestinationTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *OverlayMirrorDestinationTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *OverlayMirrorDestinationTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the OverlayMirrorDestinationTemplate from the server
func (o *OverlayMirrorDestinationTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the OverlayMirrorDestinationTemplate into the server
func (o *OverlayMirrorDestinationTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the OverlayMirrorDestinationTemplate from the server
func (o *OverlayMirrorDestinationTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
