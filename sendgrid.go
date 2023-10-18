package sendgrid

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/vottunio/sendgrid/apiwrapper"
)

const (
	BASE_URL                  string = "https://api.sendgrid.com"
	ADD_CONTACT_URL           string = "/v3/marketing/contacts"
	GET_CONTACTS_BY_EMAILS    string = "/v3/marketing/contacts/search/emails"
	REMOVE_CONTACTS_FROM_LIST string = "/v3/marketing/lists/%s/contacts?contact_ids=%s"
	CONTENT_TYPE              string = "Content-Type"
	AUTH_APP_ID               string = "x-application-vkn"
	AUTHORIZATION             string = "Authorization"
	MIME_TYPE_JSON            string = "application/json"
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

type ContactInfo struct {
	Contact struct {
		ID string `json:"id"`
	} `json:"contact"`
}

type GetContactsRequestDTO struct {
	Emails []string `json:"emails"`
}
type AddContactResponseDTO struct {
	JobID string `json:"job_id"`
}

type SendGridClient struct {
	PrivateKey string
}

type GetContactsInfoByEmail struct {
	Result map[string]ContactInfo `json:"result"`
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

func (s SendGridClient) AddContactToLists(c *AddContactRequestDTO, responseData *AddContactResponseDTO) error {
	return s.CreateContact(c, responseData)
}

func (s SendGridClient) RemoveFromList(emails []string, list string, responseData *AddContactResponseDTO) error {

	emailContactByEmailResponseDTO := GetContactsInfoByEmail{}
	err := apiwrapper.RequestApiEndpoint(
		&apiwrapper.RequestApiEndpointInfo{
			EndpointUrl:  BASE_URL + GET_CONTACTS_BY_EMAILS,
			RequestData:  GetContactsRequestDTO{Emails: emails},
			ResponseData: &emailContactByEmailResponseDTO,
			HttpMethod:   http.MethodPost,
			TokenAuth:    s.PrivateKey,
		},
		setReqHeaders,
	)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	ids := strings.Builder{}
	for _, v := range emailContactByEmailResponseDTO.Result {
		ids.WriteString(v.Contact.ID)
		ids.WriteString(",")
	}

	idsStr := ids.String()[:len(ids.String())-1]

	url := fmt.Sprintf(REMOVE_CONTACTS_FROM_LIST, list, idsStr)
	return apiwrapper.RequestApiEndpoint(
		&apiwrapper.RequestApiEndpointInfo{
			EndpointUrl:  BASE_URL + url,
			RequestData:  nil,
			ResponseData: &responseData,
			HttpMethod:   http.MethodDelete,
			TokenAuth:    s.PrivateKey,
		},
		setReqHeaders,
	)

}
