/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// L4ServiceIdentity represents the Identity of the object
var L4ServiceIdentity = bambou.Identity{
	Name:     "l4service",
	Category: "l4services",
}

// L4ServicesList represents a list of L4Services
type L4ServicesList []*L4Service

// L4ServicesAncestor is the interface that an ancestor of a L4Service must implement.
// An Ancestor is defined as an entity that has L4Service as a descendant.
// An Ancestor can get a list of its child L4Services, but not necessarily create one.
type L4ServicesAncestor interface {
	L4Services(*bambou.FetchingInfo) (L4ServicesList, *bambou.Error)
}

// L4ServicesParent is the interface that a parent of a L4Service must implement.
// A Parent is defined as an entity that has L4Service as a child.
// A Parent is an Ancestor which can create a L4Service.
type L4ServicesParent interface {
	L4ServicesAncestor
	CreateL4Service(*L4Service) *bambou.Error
}

// L4Service represents the model of a l4service
type L4Service struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewL4Service returns a new *L4Service
func NewL4Service() *L4Service {

	return &L4Service{}
}

// Identity returns the Identity of the object.
func (o *L4Service) Identity() bambou.Identity {

	return L4ServiceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *L4Service) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *L4Service) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the L4Service from the server
func (o *L4Service) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the L4Service into the server
func (o *L4Service) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the L4Service from the server
func (o *L4Service) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// L4ServiceGroups retrieves the list of child L4ServiceGroups of the L4Service
func (o *L4Service) L4ServiceGroups(info *bambou.FetchingInfo) (L4ServiceGroupsList, *bambou.Error) {

	var list L4ServiceGroupsList
	err := bambou.CurrentSession().FetchChildren(o, L4ServiceGroupIdentity, &list, info)
	return list, err
}

// AssignL4ServiceGroups assigns the list of L4ServiceGroups to the L4Service
func (o *L4Service) AssignL4ServiceGroups(children L4ServiceGroupsList) *bambou.Error {

	list := []bambou.Identifiable{}
	for _, c := range children {
		list = append(list, c)
	}

	return bambou.CurrentSession().AssignChildren(o, list, L4ServiceGroupIdentity)
}
