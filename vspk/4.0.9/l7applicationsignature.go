/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// L7applicationsignatureIdentity represents the Identity of the object
var L7applicationsignatureIdentity = bambou.Identity{
	Name:     "l7applicationsignature",
	Category: "l7applicationsignatures",
}

// L7applicationsignaturesList represents a list of L7applicationsignatures
type L7applicationsignaturesList []*L7applicationsignature

// L7applicationsignaturesAncestor is the interface that an ancestor of a L7applicationsignature must implement.
// An Ancestor is defined as an entity that has L7applicationsignature as a descendant.
// An Ancestor can get a list of its child L7applicationsignatures, but not necessarily create one.
type L7applicationsignaturesAncestor interface {
	L7applicationsignatures(*bambou.FetchingInfo) (L7applicationsignaturesList, *bambou.Error)
}

// L7applicationsignaturesParent is the interface that a parent of a L7applicationsignature must implement.
// A Parent is defined as an entity that has L7applicationsignature as a child.
// A Parent is an Ancestor which can create a L7applicationsignature.
type L7applicationsignaturesParent interface {
	L7applicationsignaturesAncestor
	CreateL7applicationsignature(*L7applicationsignature) *bambou.Error
}

// L7applicationsignature represents the model of a l7applicationsignature
type L7applicationsignature struct {
	ID                string `json:"ID,omitempty"`
	ParentID          string `json:"parentID,omitempty"`
	ParentType        string `json:"parentType,omitempty"`
	Owner             string `json:"owner,omitempty"`
	Name              string `json:"name,omitempty"`
	Category          string `json:"category,omitempty"`
	Readonly          bool   `json:"readonly"`
	Description       string `json:"description,omitempty"`
	DictionaryVersion int    `json:"dictionaryVersion,omitempty"`
	Guidstring        string `json:"guidstring,omitempty"`
}

// NewL7applicationsignature returns a new *L7applicationsignature
func NewL7applicationsignature() *L7applicationsignature {

	return &L7applicationsignature{
		Readonly: false,
	}
}

// Identity returns the Identity of the object.
func (o *L7applicationsignature) Identity() bambou.Identity {

	return L7applicationsignatureIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *L7applicationsignature) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *L7applicationsignature) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the L7applicationsignature from the server
func (o *L7applicationsignature) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the L7applicationsignature into the server
func (o *L7applicationsignature) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the L7applicationsignature from the server
func (o *L7applicationsignature) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Applications retrieves the list of child Applications of the L7applicationsignature
func (o *L7applicationsignature) Applications(info *bambou.FetchingInfo) (ApplicationsList, *bambou.Error) {

	var list ApplicationsList
	err := bambou.CurrentSession().FetchChildren(o, ApplicationIdentity, &list, info)
	return list, err
}
