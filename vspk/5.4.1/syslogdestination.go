/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// SyslogDestinationIdentity represents the Identity of the object
var SyslogDestinationIdentity = bambou.Identity{
	Name:     "syslogdestination",
	Category: "syslogdestinations",
}

// SyslogDestinationsList represents a list of SyslogDestinations
type SyslogDestinationsList []*SyslogDestination

// SyslogDestinationsAncestor is the interface that an ancestor of a SyslogDestination must implement.
// An Ancestor is defined as an entity that has SyslogDestination as a descendant.
// An Ancestor can get a list of its child SyslogDestinations, but not necessarily create one.
type SyslogDestinationsAncestor interface {
	SyslogDestinations(*bambou.FetchingInfo) (SyslogDestinationsList, *bambou.Error)
}

// SyslogDestinationsParent is the interface that a parent of a SyslogDestination must implement.
// A Parent is defined as an entity that has SyslogDestination as a child.
// A Parent is an Ancestor which can create a SyslogDestination.
type SyslogDestinationsParent interface {
	SyslogDestinationsAncestor
	CreateSyslogDestination(*SyslogDestination) *bambou.Error
}

// SyslogDestination represents the model of a syslogdestination
type SyslogDestination struct {
	ID          string `json:"ID,omitempty"`
	ParentID    string `json:"parentID,omitempty"`
	ParentType  string `json:"parentType,omitempty"`
	Owner       string `json:"owner,omitempty"`
	IPAddress   string `json:"IPAddress,omitempty"`
	IPType      string `json:"IPType,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Port        int    `json:"port,omitempty"`
	Type        string `json:"type,omitempty"`
}

// NewSyslogDestination returns a new *SyslogDestination
func NewSyslogDestination() *SyslogDestination {

	return &SyslogDestination{
		IPType: "IPV4",
		Port:   514,
		Type:   "UDP",
	}
}

// Identity returns the Identity of the object.
func (o *SyslogDestination) Identity() bambou.Identity {

	return SyslogDestinationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SyslogDestination) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SyslogDestination) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the SyslogDestination from the server
func (o *SyslogDestination) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the SyslogDestination into the server
func (o *SyslogDestination) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the SyslogDestination from the server
func (o *SyslogDestination) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
