/*
 * Go Clean Template API
 *
 * Using a translation service as an example
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type EntityTranslation struct {
	Destination string `json:"destination,omitempty"`

	Original string `json:"original,omitempty"`

	Source string `json:"source,omitempty"`

	Translation string `json:"translation,omitempty"`
}