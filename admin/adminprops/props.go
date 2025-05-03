package adminprops

import (
	"context"
	"nippo/database"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/google/uuid"
)

type Props struct {
	ctx           context.Context
	userPool      string
	orgId         uuid.UUID
	db            database.Database
	cognitoClient *cognitoidentityprovider.Client
}

func NewProps(
	ctx context.Context,
	userPool string,
	orgId uuid.UUID,
	db database.Database,
	cognitoClient *cognitoidentityprovider.Client,
) *Props {
	return &Props{
		ctx,
		userPool,
		orgId,
		db,
		cognitoClient,
	}
}

func (p *Props) Ctx() context.Context {
	return p.ctx
}

func (p *Props) UserPool() *string {
	return &p.userPool
}

func (p *Props) OrgId() []byte {
	return p.orgId[:]
}

func (p *Props) OrgUuid() uuid.UUID {
	return p.orgId
}

func (p *Props) Db() database.Database {
	return p.db
}

func (p *Props) Cognito() *cognitoidentityprovider.Client {
	return p.cognitoClient
}
