package mail

import (
	"net/http"
	"os"

	"github.com/matcornic/hermes/v2"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type sendMail struct{}

type Mailer interface {
	SendWelcomeMessage(string, string, string, string, string) (*EmailResponse, error)
}

var (
	SendMail Mailer = &sendMail{}
)

type EmailResponse struct {
	Status   int
	Response string
}

func (s *sendMail) SendWelcomeMessage(ToUser string, FromAdmin string, Email string, SendKey string, AppEnv string) (*EmailResponse, error) {
	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "EventPark",
			Link: "https://eventpark.gr",
		},
	}

	var url string
	if os.Getenv("APP_ENV") == "production" {
		url = "https://eventpark.gr/register" + Email
	} else {
		url = "https://127.0.0.1:3000/register" + Email
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: ToUser,
			Intros: []string{
				"Καλώς ήρθατε στο EventPark! Ολοκληρώστε την εγγραφή σας.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Κάντε κλικ για την επαλήθευση του λογαριασμού σας",
					Button: hermes.Button{
						Color: "#FFFFFF",
						Text:  "Επαλήθευση λογαριασμού",
						Link:  url,
					},
				},
			},
			Outros: []string{
				"Αντιμετωπίζετε κάποιο πρόβλημα ή έχετε κάποια απορία; Θα σας εξυπηρετήσουμε αμέσως!",
			},
		},
	}

	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		return nil, err
	}

	from := mail.NewEmail("eventparkgr", FromAdmin)
	subject := "Επαλήθευση λογαριασμού"
	to := mail.NewEmail("Επαλήθευση λογαριασμού", ToUser)
	message := mail.NewSingleEmail(from, subject, to, emailBody, emailBody)

	client := sendgrid.NewSendClient(SendKey)
	_, err = client.Send(message)
	if err != nil {
		return nil, err
	}

	return &EmailResponse{
		Status:   http.StatusOK,
		Response: "Επιτυχία!",
	}, nil

}
