/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// ApplicationIdentity represents the Identity of the object
var ApplicationIdentity = bambou.Identity{
	Name:     "application",
	Category: "applications",
}

// ApplicationsList represents a list of Applications
type ApplicationsList []*Application

// ApplicationsAncestor is the interface that an ancestor of a Application must implement.
// An Ancestor is defined as an entity that has Application as a descendant.
// An Ancestor can get a list of its child Applications, but not necessarily create one.
type ApplicationsAncestor interface {
	Applications(*bambou.FetchingInfo) (ApplicationsList, *bambou.Error)
}

// ApplicationsParent is the interface that a parent of a Application must implement.
// A Parent is defined as an entity that has Application as a child.
// A Parent is an Ancestor which can create a Application.
type ApplicationsParent interface {
	ApplicationsAncestor
	CreateApplication(*Application) *bambou.Error
}

// Application represents the model of a application
type Application struct {
	ID                                 string  `json:"ID,omitempty"`
	ParentID                           string  `json:"parentID,omitempty"`
	ParentType                         string  `json:"parentType,omitempty"`
	Owner                              string  `json:"owner,omitempty"`
	DSCP                               string  `json:"DSCP,omitempty"`
	Name                               string  `json:"name,omitempty"`
	ReadOnly                           bool    `json:"readOnly"`
	PerformanceMonitorType             string  `json:"performanceMonitorType,omitempty"`
	Description                        string  `json:"description,omitempty"`
	DestinationIP                      string  `json:"destinationIP,omitempty"`
	DestinationPort                    string  `json:"destinationPort,omitempty"`
	EnablePPS                          bool    `json:"enablePPS"`
	OneWayDelay                        int     `json:"oneWayDelay,omitempty"`
	OneWayJitter                       int     `json:"oneWayJitter,omitempty"`
	OneWayLoss                         float64 `json:"oneWayLoss,omitempty"`
	PostClassificationPath             string  `json:"postClassificationPath,omitempty"`
	SourceIP                           string  `json:"sourceIP,omitempty"`
	SourcePort                         string  `json:"sourcePort,omitempty"`
	OptimizePathSelection              string  `json:"optimizePathSelection,omitempty"`
	PreClassificationPath              string  `json:"preClassificationPath,omitempty"`
	Protocol                           string  `json:"protocol,omitempty"`
	AssociatedL7ApplicationSignatureID string  `json:"associatedL7ApplicationSignatureID,omitempty"`
	EtherType                          string  `json:"etherType,omitempty"`
	Symmetry                           bool    `json:"symmetry"`
}

// NewApplication returns a new *Application
func NewApplication() *Application {

	return &Application{
		ReadOnly:               false,
		PerformanceMonitorType: "FIRST_PACKET",
		EnablePPS:              false,
		PostClassificationPath: "ANY",
		PreClassificationPath:  "DEFAULT",
		Protocol:               "NONE",
	}
}

// Identity returns the Identity of the object.
func (o *Application) Identity() bambou.Identity {

	return ApplicationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Application) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Application) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Application from the server
func (o *Application) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Application into the server
func (o *Application) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Application from the server
func (o *Application) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Monitorscopes retrieves the list of child Monitorscopes of the Application
func (o *Application) Monitorscopes(info *bambou.FetchingInfo) (MonitorscopesList, *bambou.Error) {

	var list MonitorscopesList
	err := bambou.CurrentSession().FetchChildren(o, MonitorscopeIdentity, &list, info)
	return list, err
}

// CreateMonitorscope creates a new child Monitorscope under the Application
func (o *Application) CreateMonitorscope(child *Monitorscope) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// ApplicationBindings retrieves the list of child ApplicationBindings of the Application
func (o *Application) ApplicationBindings(info *bambou.FetchingInfo) (ApplicationBindingsList, *bambou.Error) {

	var list ApplicationBindingsList
	err := bambou.CurrentSession().FetchChildren(o, ApplicationBindingIdentity, &list, info)
	return list, err
}
