/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// UplinkConnectionIdentity represents the Identity of the object
var UplinkConnectionIdentity = bambou.Identity{
	Name:     "uplinkconnection",
	Category: "uplinkconnections",
}

// UplinkConnectionsList represents a list of UplinkConnections
type UplinkConnectionsList []*UplinkConnection

// UplinkConnectionsAncestor is the interface that an ancestor of a UplinkConnection must implement.
// An Ancestor is defined as an entity that has UplinkConnection as a descendant.
// An Ancestor can get a list of its child UplinkConnections, but not necessarily create one.
type UplinkConnectionsAncestor interface {
	UplinkConnections(*bambou.FetchingInfo) (UplinkConnectionsList, *bambou.Error)
}

// UplinkConnectionsParent is the interface that a parent of a UplinkConnection must implement.
// A Parent is defined as an entity that has UplinkConnection as a child.
// A Parent is an Ancestor which can create a UplinkConnection.
type UplinkConnectionsParent interface {
	UplinkConnectionsAncestor
	CreateUplinkConnection(*UplinkConnection) *bambou.Error
}

// UplinkConnection represents the model of a uplinkconnection
type UplinkConnection struct {
	ID                     string `json:"ID,omitempty"`
	ParentID               string `json:"parentID,omitempty"`
	ParentType             string `json:"parentType,omitempty"`
	Owner                  string `json:"owner,omitempty"`
	DNSAddress             string `json:"DNSAddress,omitempty"`
	Password               string `json:"password,omitempty"`
	Address                string `json:"address,omitempty"`
	Netmask                string `json:"netmask,omitempty"`
	Mode                   string `json:"mode,omitempty"`
	Role                   string `json:"role,omitempty"`
	Username               string `json:"username,omitempty"`
	AssociatedVSCProfileID string `json:"associatedVSCProfileID,omitempty"`
}

// NewUplinkConnection returns a new *UplinkConnection
func NewUplinkConnection() *UplinkConnection {

	return &UplinkConnection{
		Address: "IPv4",
		Mode:    "Dynamic",
		Role:    "PRIMARY",
	}
}

// Identity returns the Identity of the object.
func (o *UplinkConnection) Identity() bambou.Identity {

	return UplinkConnectionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *UplinkConnection) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *UplinkConnection) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the UplinkConnection from the server
func (o *UplinkConnection) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the UplinkConnection into the server
func (o *UplinkConnection) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the UplinkConnection from the server
func (o *UplinkConnection) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}