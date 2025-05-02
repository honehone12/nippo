package cognito

import (
	"context"
	"errors"
	"fmt"
	"nippo/schema"
	"nippo/view/viewprops"
	"slices"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func emailAttr() *string {
	attr := "email"
	return &attr
}

func emailCheckedAttr() *string {
	attr := "email_verified"
	return &attr
}

func checked() *string {
	checked := "true"
	return &checked
}

func Exists(props *viewprops.Props, email schema.CleanText) (bool, error) {
	name := string(email)
	getUser := cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: props.UserPool(),
		Username:   &name,
	}
	_, err := props.Cognito().AdminGetUser(props.Ctx(), &getUser)
	if errors.Is(err, &types.ResourceNotFoundException{}) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
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

func MarkEmailChecked(
	ctx context.Context,
	client *cognitoidentityprovider.Client,
	userPool string,
	email schema.CleanText,
) error {
	name := string(email)

	getUser := cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: &userPool,
		Username:   &name,
	}
	out, err := client.AdminGetUser(ctx, &getUser)
	if err != nil {
		return err
	}

	idx := slices.IndexFunc(out.UserAttributes, func(attr types.AttributeType) bool {
		return *attr.Name == *emailCheckedAttr()
	})
	if idx < 0 {
		return errors.New("could not find email checked attribute")
	}
	if *out.UserAttributes[idx].Value == *checked() {
		return nil
	}

	updateAttr := cognitoidentityprovider.AdminUpdateUserAttributesInput{
		UserPoolId: &userPool,
		Username:   &name,
		UserAttributes: []types.AttributeType{
			{
				Name:  emailCheckedAttr(),
				Value: checked(),
			},
		},
	}
	_, err = client.AdminUpdateUserAttributes(ctx, &updateAttr)
	if err != nil {
		return err
	}

	return nil
}
