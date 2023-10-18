package sendgrid_test

// You must create a constant in your sendgrid_test package to set the Sendgrid API Key. The constant name is sendgridApiKey.
import (
	"testing"

	"github.com/vottun-com/utils/pointer"
	"github.com/vottunio/log"
	"github.com/vottunio/sendgrid"
)

var (
	contacts []sendgrid.ContactDTO = []sendgrid.ContactDTO{
		{Email: "alexlopezt@gmail.com",
			FirstName: pointer.String("Alex"),
			LastName:  pointer.String("Lopez"),
		},
	}

	idsCreate *[]string = &[]string{"c6b4770e-bbe0-462d-8746-1a0983863d80"}
	idsUpdate *[]string = &[]string{"1090d09f-46fb-413f-8b36-d9d243de0e44"}
	idSensei  string    = "1090d09f-46fb-413f-8b36-d9d243de0e44"
)

func TestCreateContact(t *testing.T) {

	client := sendgrid.SendGridClient(sendgrid.SendGridClient{PrivateKey: sendgridApiKey})
	responseData := &sendgrid.AddContactResponseDTO{}

	err := client.CreateContact(
		&sendgrid.AddContactRequestDTO{
			ListIDs:  idsCreate,
			Contacts: contacts,
		},
		responseData,
	)

	if err != nil {
		log.Errorf("An error was raised sending create contacts request to Sendgrid. %+v", err)
	} else {
		log.Infof("A new contacts list has been sent to Sendgrid to be created {%+v}. Job id is: {%s}", contacts, responseData.JobID)
	}

}

func TestUpdateContact(t *testing.T) {

	client := sendgrid.SendGridClient(sendgrid.SendGridClient{PrivateKey: sendgridApiKey})
	responseData := &sendgrid.AddContactResponseDTO{}

	err := client.CreateContact(
		&sendgrid.AddContactRequestDTO{
			ListIDs:  idsUpdate,
			Contacts: contacts,
		},
		responseData,
	)

	if err != nil {
		log.Errorf("An error was raised sending create contacts request to Sendgrid. %+v", err)
	} else {
		log.Infof("A new contacts list has been sent to Sendgrid to be created {%+v}. Job id is: {%s}", contacts, responseData.JobID)
	}

}

func TestRemoveFromList(t *testing.T) {
	client := sendgrid.SendGridClient(sendgrid.SendGridClient{PrivateKey: sendgridApiKey})
	responseData := &sendgrid.AddContactResponseDTO{}

	err := client.RemoveFromList(
		[]string{"alexlopezt@gmail.com"},
		idSensei,
		responseData,
	)

	if err != nil {
		log.Errorf("An error was raised sending create contacts request to Sendgrid. %+v", err)
	} else {
		log.Infof("A new contacts list has been sent to Sendgrid to be created {%+v}. Job id is: {%s}", contacts, responseData.JobID)
	}
}
