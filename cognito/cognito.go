package cognito

import (
	"nippo/schema"
	"nippo/view/viewprops"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func Pooled(props *viewprops.Props, email schema.CleanText) bool {
	name := string(email)
	getUser := cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: props.UserPool(),
		Username:   &name,
	}
	_, err := props.Cognito().AdminGetUser(props.Ctx(), &getUser)
	return err == nil
}
