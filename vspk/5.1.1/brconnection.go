/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// BRConnectionIdentity represents the Identity of the object
var BRConnectionIdentity = bambou.Identity{
	Name:     "brconnections",
	Category: "brconnections",
}

// BRConnectionsList represents a list of BRConnections
type BRConnectionsList []*BRConnection

// BRConnectionsAncestor is the interface that an ancestor of a BRConnection must implement.
// An Ancestor is defined as an entity that has BRConnection as a descendant.
// An Ancestor can get a list of its child BRConnections, but not necessarily create one.
type BRConnectionsAncestor interface {
	BRConnections(*bambou.FetchingInfo) (BRConnectionsList, *bambou.Error)
}

// BRConnectionsParent is the interface that a parent of a BRConnection must implement.
// A Parent is defined as an entity that has BRConnection as a child.
// A Parent is an Ancestor which can create a BRConnection.
type BRConnectionsParent interface {
	BRConnectionsAncestor
	CreateBRConnection(*BRConnection) *bambou.Error
}

// BRConnection represents the model of a brconnections
type BRConnection struct {
	ID                    string `json:"ID,omitempty"`
	ParentID              string `json:"parentID,omitempty"`
	ParentType            string `json:"parentType,omitempty"`
	Owner                 string `json:"owner,omitempty"`
	DNSAddress            string `json:"DNSAddress,omitempty"`
	Gateway               string `json:"gateway,omitempty"`
	Address               string `json:"address,omitempty"`
	AdvertisementCriteria string `json:"advertisementCriteria,omitempty"`
	Netmask               string `json:"netmask,omitempty"`
	Mode                  string `json:"mode,omitempty"`
	UplinkID              int    `json:"uplinkID,omitempty"`
}

// NewBRConnection returns a new *BRConnection
func NewBRConnection() *BRConnection {

	return &BRConnection{}
}

// Identity returns the Identity of the object.
func (o *BRConnection) Identity() bambou.Identity {

	return BRConnectionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *BRConnection) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *BRConnection) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the BRConnection from the server
func (o *BRConnection) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the BRConnection into the server
func (o *BRConnection) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the BRConnection from the server
func (o *BRConnection) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// BFDSessions retrieves the list of child BFDSessions of the BRConnection
func (o *BRConnection) BFDSessions(info *bambou.FetchingInfo) (BFDSessionsList, *bambou.Error) {

	var list BFDSessionsList
	err := bambou.CurrentSession().FetchChildren(o, BFDSessionIdentity, &list, info)
	return list, err
}

// CreateBFDSession creates a new child BFDSession under the BRConnection
func (o *BRConnection) CreateBFDSession(child *BFDSession) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
