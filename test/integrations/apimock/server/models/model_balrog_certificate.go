/*
 * Firefox Guardian API
 *
 * API to manage Guardian accounts, devices and servers
 *
 * API version: 0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import "time"

type BalrogCertificate struct {
	AuthorityKeyID                           []byte    `json:"oid,omitempty"`
	NotBefore                                time.Time `json:"notbefore,string,omitempty"`
	Subject                                  string    `json:"subject,omitempty"`
	AdditionalIntermediate                   bool      `json:"additionalInter,omitempty"`
	AdditionalRoot                           bool      `json:"additionalRoot,omitempty"`
	AdditionalRootTopOrBot                   bool      `json:"additionRootTopOrBot,omitempty"`
	AdditionalIrrelevantIntermediate         bool      `json:"additionalIrrelevantIntermediate,omitempty"`
	AdditionalIrrelevantIntermediateTopOrBot bool      `json:"additionalIrrelevantIntermediateTopOrBot,omitempty"`
}