/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// AllAlarmIdentity represents the Identity of the object
var AllAlarmIdentity = bambou.Identity{
	Name:     "allalarm",
	Category: "allalarms",
}

// AllAlarmsList represents a list of AllAlarms
type AllAlarmsList []*AllAlarm

// AllAlarmsAncestor is the interface that an ancestor of a AllAlarm must implement.
// An Ancestor is defined as an entity that has AllAlarm as a descendant.
// An Ancestor can get a list of its child AllAlarms, but not necessarily create one.
type AllAlarmsAncestor interface {
	AllAlarms(*bambou.FetchingInfo) (AllAlarmsList, *bambou.Error)
}

// AllAlarmsParent is the interface that a parent of a AllAlarm must implement.
// A Parent is defined as an entity that has AllAlarm as a child.
// A Parent is an Ancestor which can create a AllAlarm.
type AllAlarmsParent interface {
	AllAlarmsAncestor
	CreateAllAlarm(*AllAlarm) *bambou.Error
}

// AllAlarm represents the model of a allalarm
type AllAlarm struct {
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

// NewAllAlarm returns a new *AllAlarm
func NewAllAlarm() *AllAlarm {

	return &AllAlarm{}
}

// Identity returns the Identity of the object.
func (o *AllAlarm) Identity() bambou.Identity {

	return AllAlarmIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AllAlarm) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AllAlarm) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the AllAlarm from the server
func (o *AllAlarm) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the AllAlarm into the server
func (o *AllAlarm) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the AllAlarm from the server
func (o *AllAlarm) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the AllAlarm
func (o *AllAlarm) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the AllAlarm
func (o *AllAlarm) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the AllAlarm
func (o *AllAlarm) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the AllAlarm
func (o *AllAlarm) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
