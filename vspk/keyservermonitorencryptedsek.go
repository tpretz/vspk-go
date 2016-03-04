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

// KeyServerMonitorEncryptedSEKIdentity represents the Identity of the object
var KeyServerMonitorEncryptedSEKIdentity = bambou.Identity{
	Name:     "keyservermonitorencryptedsek",
	Category: "keyservermonitorencryptedseks",
}

// KeyServerMonitorEncryptedSEKsList represents a list of KeyServerMonitorEncryptedSEKs
type KeyServerMonitorEncryptedSEKsList []*KeyServerMonitorEncryptedSEK

// KeyServerMonitorEncryptedSEKsAncestor is the interface of an ancestor of a KeyServerMonitorEncryptedSEK must implement.
type KeyServerMonitorEncryptedSEKsAncestor interface {
	KeyServerMonitorEncryptedSEKs(*bambou.FetchingInfo) (KeyServerMonitorEncryptedSEKsList, *bambou.Error)
	CreateKeyServerMonitorEncryptedSEKs(*KeyServerMonitorEncryptedSEK) *bambou.Error
}

// KeyServerMonitorEncryptedSEK represents the model of a keyservermonitorencryptedsek
type KeyServerMonitorEncryptedSEK struct {
	ID                                        string  `json:"ID,omitempty"`
	ParentID                                  string  `json:"parentID,omitempty"`
	ParentType                                string  `json:"parentType,omitempty"`
	Owner                                     string  `json:"owner,omitempty"`
	NSGCertificateSerialNumber                float64 `json:"NSGCertificateSerialNumber,omitempty"`
	AssociatedKeyServerMonitorSEKCreationTime float64 `json:"associatedKeyServerMonitorSEKCreationTime,omitempty"`
	AssociatedKeyServerMonitorSEKID           string  `json:"associatedKeyServerMonitorSEKID,omitempty"`
	EntityScope                               string  `json:"entityScope,omitempty"`
	ExternalID                                string  `json:"externalID,omitempty"`
	GatewaySecuredDataID                      string  `json:"gatewaySecuredDataID,omitempty"`
	KeyServerCertificateSerialNumber          float64 `json:"keyServerCertificateSerialNumber,omitempty"`
	LastUpdatedBy                             string  `json:"lastUpdatedBy,omitempty"`
}

// NewKeyServerMonitorEncryptedSEK returns a new *KeyServerMonitorEncryptedSEK
func NewKeyServerMonitorEncryptedSEK() *KeyServerMonitorEncryptedSEK {

	return &KeyServerMonitorEncryptedSEK{}
}

// Identity returns the Identity of the object.
func (o *KeyServerMonitorEncryptedSEK) Identity() bambou.Identity {

	return KeyServerMonitorEncryptedSEKIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *KeyServerMonitorEncryptedSEK) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *KeyServerMonitorEncryptedSEK) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the KeyServerMonitorEncryptedSEK from the server
func (o *KeyServerMonitorEncryptedSEK) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the KeyServerMonitorEncryptedSEK into the server
func (o *KeyServerMonitorEncryptedSEK) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the KeyServerMonitorEncryptedSEK from the server
func (o *KeyServerMonitorEncryptedSEK) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the KeyServerMonitorEncryptedSEK
func (o *KeyServerMonitorEncryptedSEK) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the KeyServerMonitorEncryptedSEK
func (o *KeyServerMonitorEncryptedSEK) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the KeyServerMonitorEncryptedSEK
func (o *KeyServerMonitorEncryptedSEK) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the KeyServerMonitorEncryptedSEK
func (o *KeyServerMonitorEncryptedSEK) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
