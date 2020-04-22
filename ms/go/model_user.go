/*
 * PinArt Labels MS
 *
 * A labels microservice for PinArt system.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type User struct {
	Id int64 `json:"id,omitempty"`

	RelatedLabels []Label `json:"relatedLabels,omitempty"`
}
