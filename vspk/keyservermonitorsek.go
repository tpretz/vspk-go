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

// KeyServerMonitorSEKIdentity represents the Identity of the object
var KeyServerMonitorSEKIdentity = bambou.Identity{
	Name:     "keyservermonitorsek",
	Category: "keyservermonitorseks",
}

// KeyServerMonitorSEKsList represents a list of KeyServerMonitorSEKs
type KeyServerMonitorSEKsList []*KeyServerMonitorSEK

// KeyServerMonitorSEKsAncestor is the interface of an ancestor of a KeyServerMonitorSEK must implement.
type KeyServerMonitorSEKsAncestor interface {
	KeyServerMonitorSEKs(*bambou.FetchingInfo) (KeyServerMonitorSEKsList, *bambou.Error)
	CreateKeyServerMonitorSEKs(*KeyServerMonitorSEK) *bambou.Error
}

// KeyServerMonitorSEK represents the model of a keyservermonitorsek
type KeyServerMonitorSEK struct {
	ID                                 string `json:"ID,omitempty"`
	ParentID                           string `json:"parentID,omitempty"`
	ParentType                         string `json:"parentType,omitempty"`
	Owner                              string `json:"owner,omitempty"`
	CreationTime                       int    `json:"creationTime,omitempty"`
	EntityScope                        string `json:"entityScope,omitempty"`
	ExternalID                         string `json:"externalID,omitempty"`
	LastUpdatedBy                      string `json:"lastUpdatedBy,omitempty"`
	Lifetime                           int    `json:"lifetime,omitempty"`
	SeedPayloadAuthenticationAlgorithm string `json:"seedPayloadAuthenticationAlgorithm,omitempty"`
	SeedPayloadEncryptionAlgorithm     string `json:"seedPayloadEncryptionAlgorithm,omitempty"`
	StartTime                          int    `json:"startTime,omitempty"`
}

// NewKeyServerMonitorSEK returns a new *KeyServerMonitorSEK
func NewKeyServerMonitorSEK() *KeyServerMonitorSEK {

	return &KeyServerMonitorSEK{}
}

// Identity returns the Identity of the object.
func (o *KeyServerMonitorSEK) Identity() bambou.Identity {

	return KeyServerMonitorSEKIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *KeyServerMonitorSEK) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *KeyServerMonitorSEK) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the KeyServerMonitorSEK from the server
func (o *KeyServerMonitorSEK) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the KeyServerMonitorSEK into the server
func (o *KeyServerMonitorSEK) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the KeyServerMonitorSEK from the server
func (o *KeyServerMonitorSEK) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the KeyServerMonitorSEK
func (o *KeyServerMonitorSEK) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the KeyServerMonitorSEK
func (o *KeyServerMonitorSEK) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// KeyServerMonitorEncryptedSeeds retrieves the list of child KeyServerMonitorEncryptedSeeds of the KeyServerMonitorSEK
func (o *KeyServerMonitorSEK) KeyServerMonitorEncryptedSeeds(info *bambou.FetchingInfo) (KeyServerMonitorEncryptedSeedsList, *bambou.Error) {

	var list KeyServerMonitorEncryptedSeedsList
	err := bambou.CurrentSession().FetchChildren(o, KeyServerMonitorEncryptedSeedIdentity, &list, info)
	return list, err
}

// CreateKeyServerMonitorEncryptedSeed creates a new child KeyServerMonitorEncryptedSeed under the KeyServerMonitorSEK
func (o *KeyServerMonitorSEK) CreateKeyServerMonitorEncryptedSeed(child *KeyServerMonitorEncryptedSeed) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// KeyServerMonitorEncryptedSEKs retrieves the list of child KeyServerMonitorEncryptedSEKs of the KeyServerMonitorSEK
func (o *KeyServerMonitorSEK) KeyServerMonitorEncryptedSEKs(info *bambou.FetchingInfo) (KeyServerMonitorEncryptedSEKsList, *bambou.Error) {

	var list KeyServerMonitorEncryptedSEKsList
	err := bambou.CurrentSession().FetchChildren(o, KeyServerMonitorEncryptedSEKIdentity, &list, info)
	return list, err
}

// CreateKeyServerMonitorEncryptedSEK creates a new child KeyServerMonitorEncryptedSEK under the KeyServerMonitorSEK
func (o *KeyServerMonitorSEK) CreateKeyServerMonitorEncryptedSEK(child *KeyServerMonitorEncryptedSEK) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the KeyServerMonitorSEK
func (o *KeyServerMonitorSEK) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the KeyServerMonitorSEK
func (o *KeyServerMonitorSEK) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}