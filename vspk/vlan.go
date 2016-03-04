/*
  Copyright (c) 2015, Alcatel-Lucent Inc
  All rights reserved.

  Redistribution and use in source and binary forms, with or without
  modification, are permitted provided that the following conditions are met:
      * Redistributions of source code must retain the above copyright
        notice, this list of conditions and the following disclaimer.
      * Redistributions in binary form must reproduce the above copyright
        notice, this list of conditions and the following disclaimer in the
        documentation and/or other materials provided with the distribution.
      * Neither the name of the copyright holder nor the names of its contributors
        may be used to endorse or promote products derived from this software without
        specific prior written permission.

  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
  ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
  WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY
  DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
  (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
  LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
  ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
  (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
  SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VLANIdentity represents the Identity of the object
var VLANIdentity = bambou.Identity{
	Name:     "vlan",
	Category: "vlans",
}

// VLANsList represents a list of VLANs
type VLANsList []*VLAN

// VLANsAncestor is the interface of an ancestor of a VLAN must implement.
type VLANsAncestor interface {
	VLANs(*bambou.FetchingInfo) (VLANsList, *bambou.Error)
	CreateVLANs(*VLAN) *bambou.Error
}

// VLAN represents the model of a vlan
type VLAN struct {
	ID                          string `json:"ID,omitempty"`
	ParentID                    string `json:"parentID,omitempty"`
	ParentType                  string `json:"parentType,omitempty"`
	Owner                       string `json:"owner,omitempty"`
	AssociatedEgressQOSPolicyID string `json:"associatedEgressQOSPolicyID,omitempty"`
	Description                 string `json:"description,omitempty"`
	EntityScope                 string `json:"entityScope,omitempty"`
	ExternalID                  string `json:"externalID,omitempty"`
	GatewayID                   string `json:"gatewayID,omitempty"`
	LastUpdatedBy               string `json:"lastUpdatedBy,omitempty"`
	PermittedAction             string `json:"permittedAction,omitempty"`
	Readonly                    bool   `json:"readonly"`
	Restricted                  bool   `json:"restricted"`
	Status                      string `json:"status,omitempty"`
	TemplateID                  string `json:"templateID,omitempty"`
	UseUserMnemonic             bool   `json:"useUserMnemonic"`
	UserMnemonic                string `json:"userMnemonic,omitempty"`
	Value                       int    `json:"value,omitempty"`
	VportID                     string `json:"vportID,omitempty"`
}

// NewVLAN returns a new *VLAN
func NewVLAN() *VLAN {

	return &VLAN{}
}

// Identity returns the Identity of the object.
func (o *VLAN) Identity() bambou.Identity {

	return VLANIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VLAN) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VLAN) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VLAN from the server
func (o *VLAN) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VLAN into the server
func (o *VLAN) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VLAN from the server
func (o *VLAN) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Alarms retrieves the list of child Alarms of the VLAN
func (o *VLAN) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

	var list AlarmsList
	err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
	return list, err
}

// CreateAlarm creates a new child Alarm under the VLAN
func (o *VLAN) CreateAlarm(child *Alarm) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EnterprisePermissions retrieves the list of child EnterprisePermissions of the VLAN
func (o *VLAN) EnterprisePermissions(info *bambou.FetchingInfo) (EnterprisePermissionsList, *bambou.Error) {

	var list EnterprisePermissionsList
	err := bambou.CurrentSession().FetchChildren(o, EnterprisePermissionIdentity, &list, info)
	return list, err
}

// CreateEnterprisePermission creates a new child EnterprisePermission under the VLAN
func (o *VLAN) CreateEnterprisePermission(child *EnterprisePermission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the VLAN
func (o *VLAN) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}

// CreateEventLog creates a new child EventLog under the VLAN
func (o *VLAN) CreateEventLog(child *EventLog) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VLAN
func (o *VLAN) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VLAN
func (o *VLAN) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the VLAN
func (o *VLAN) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VLAN
func (o *VLAN) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Permissions retrieves the list of child Permissions of the VLAN
func (o *VLAN) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

	var list PermissionsList
	err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
	return list, err
}

// CreatePermission creates a new child Permission under the VLAN
func (o *VLAN) CreatePermission(child *Permission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
