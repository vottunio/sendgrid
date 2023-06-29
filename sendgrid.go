package sendgrid

import (
	"fmt"
	"net/http"

	"github.com/vottunio/sendgrid/apiwrapper"
)

const (
	BASE_URL        string = "https://api.sendgrid.com"
	ADD_CONTACT_URL string = "/v3/marketing/contacts"

	CONTENT_TYPE   string = "Content-Type"
	AUTH_APP_ID    string = "x-application-vkn"
	AUTHORIZATION  string = "Authorization"
	MIME_TYPE_JSON string = "application/json; charset=UTF-8"
)

type ContactDTO struct {
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name,omitempty"`
	LastName   *string `json:"last_name,omitempty"`
	City       *string `json:"city,omitempty"`
	Country    *string `json:"country,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
}
type AddContactRequestDTO struct {
	ListIDs  *[]string    `json:"list_ids"`
	Contacts []ContactDTO `json:"contacts"`
}

type AddContactResponseDTO struct {
	JobID string `json:"job_id"`
}

type SendGridClient struct {
	PrivateKey string
}

func (s *SendGridClient) CreateContact(contact *AddContactRequestDTO, responseData *AddContactResponseDTO) error {

	return apiwrapper.RequestApiEndpoint(
		&apiwrapper.RequestApiEndpointInfo{
			EndpointUrl:  BASE_URL + ADD_CONTACT_URL,
			RequestData:  contact,
			ResponseData: &responseData,
			HttpMethod:   http.MethodPut,
			TokenAuth:    s.PrivateKey,
		},
		setReqHeaders,
	)
}
func setReqHeaders(req *http.Request, apiKey string) {

	req.Header.Set(CONTENT_TYPE, MIME_TYPE_JSON)

	req.Header.Add(AUTHORIZATION, fmt.Sprintf("Bearer %s", apiKey))

}
