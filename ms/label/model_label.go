/*
 * PinArt Labels MS
 *
 * A labels microservice for PinArt system.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package label

type Label struct {
	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	RelatedLabels []int64 `json:"relatedLabels,omitempty"`
}
type Labels struct {
	Labels []Label `json:"Labels,omitempty"`
}
