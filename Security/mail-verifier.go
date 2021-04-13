package Security

import (
	"log"

	. "github.com/msrexe/portfolio-tracker/Core/EnvVariables"
	"github.com/quickemailverification/quickemailverification-go"
)

func VerifyMail(email string) bool {
	qev := quickemailverification.CreateClient(GoDotEnvVariable("MAIL_VERIFIER_API_KEY"))

	response, err := qev.Verify(email)

	if err != nil {
		log.Println(err)
	}
	return response.Result == "valid"
}
