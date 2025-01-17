/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// AlarmIdentity represents the Identity of the object
var AlarmIdentity = bambou.Identity{
	Name:     "allalarm",
	Category: "allalarms",
}

// AlarmsList represents a list of Alarms
type AlarmsList []*Alarm

// AlarmsAncestor is the interface that an ancestor of a Alarm must implement.
// An Ancestor is defined as an entity that has Alarm as a descendant.
// An Ancestor can get a list of its child Alarms, but not necessarily create one.
type AlarmsAncestor interface {
	Alarms(*bambou.FetchingInfo) (AlarmsList, *bambou.Error)
}

// AlarmsParent is the interface that a parent of a Alarm must implement.
// A Parent is defined as an entity that has Alarm as a child.
// A Parent is an Ancestor which can create a Alarm.
type AlarmsParent interface {
	AlarmsAncestor
	CreateAlarm(*Alarm) *bambou.Error
}

// Alarm represents the model of a allalarm
type Alarm struct {
	ID                 string `json:"ID,omitempty"`
	ParentID           string `json:"parentID,omitempty"`
	ParentType         string `json:"parentType,omitempty"`
	Owner              string `json:"owner,omitempty"`
	Name               string `json:"name,omitempty"`
	TargetObject       string `json:"targetObject,omitempty"`
	LastUpdatedBy      string `json:"lastUpdatedBy,omitempty"`
	Acknowledged       bool   `json:"acknowledged"`
	Reason             string `json:"reason,omitempty"`
	Description        string `json:"description,omitempty"`
	Severity           string `json:"severity,omitempty"`
	Timestamp          int    `json:"timestamp,omitempty"`
	EnterpriseID       string `json:"enterpriseID,omitempty"`
	EntityScope        string `json:"entityScope,omitempty"`
	ErrorCondition     int    `json:"errorCondition,omitempty"`
	NumberOfOccurances int    `json:"numberOfOccurances,omitempty"`
	ExternalID         string `json:"externalID,omitempty"`
}

// NewAlarm returns a new *Alarm
func NewAlarm() *Alarm {

	return &Alarm{}
}

// Identity returns the Identity of the object.
func (o *Alarm) Identity() bambou.Identity {

	return AlarmIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Alarm) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Alarm) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Alarm from the server
func (o *Alarm) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Alarm into the server
func (o *Alarm) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Alarm from the server
func (o *Alarm) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the Alarm
func (o *Alarm) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Alarm
func (o *Alarm) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Alarm
func (o *Alarm) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Alarm
func (o *Alarm) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
