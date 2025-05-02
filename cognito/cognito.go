package cognito

import (
	"fmt"
	"nippo/schema"
	"nippo/view/viewprops"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func emailAttr() *string {
	attr := "email"
	return &attr
}

func Pooled(props *viewprops.Props, email schema.CleanText) bool {
	name := string(email)
	getUser := cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: props.UserPool(),
		Username:   &name,
	}
	_, err := props.Cognito().AdminGetUser(props.Ctx(), &getUser)
	return err == nil
}

func Invite(props *viewprops.Props, email schema.CleanText) error {
	name := string(email)
	userPool := props.UserPool()
	client := props.Cognito()
	ctx := props.Ctx()

	invite := cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId: userPool,
		Username:   &name,
		UserAttributes: []types.AttributeType{
			{
				Name:  emailAttr(),
				Value: &name,
			},
		},
	}
	_, err := client.AdminCreateUser(ctx, &invite)
	if err != nil {
		return err
	}

	group := fmt.Sprintf("nippo-%s", props.OrgUuid())
	groupAdd := cognitoidentityprovider.AdminAddUserToGroupInput{
		UserPoolId: userPool,
		Username:   &name,
		GroupName:  &group,
	}
	_, err = client.AdminAddUserToGroup(ctx, &groupAdd)

	return err
}
