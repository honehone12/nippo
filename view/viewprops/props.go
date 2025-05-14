package viewprops

import (
	"context"
	"nippo/database"
	"nippo/schema"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type Viewer struct {
	id        []uint8
	createdAt time.Time
	name      schema.CleanText
	email     schema.CleanText
	admin     bool
	subAdmin  bool
}

func NewViewer(
	id []uint8,
	created_at time.Time,
	name schema.CleanText,
	email schema.CleanText,
	admin bool,
	subAdmin bool,
) *Viewer {
	return &Viewer{
		id,
		created_at,
		name,
		email,
		admin,
		subAdmin,
	}
}

func (v *Viewer) Id() []uint8 {
	return v.id
}

func (v *Viewer) CreatedAt() time.Time {
	return v.createdAt
}

func (v *Viewer) Name() schema.CleanText {
	return v.name
}

func (v *Viewer) Email() schema.CleanText {
	return v.email
}

func (v *Viewer) Admin() bool {
	return v.admin
}

func (v *Viewer) SubAdmin() bool {
	return v.subAdmin
}

func (v *Viewer) Privileged() bool {
	return v.admin || v.subAdmin
}

type Props struct {
	ctx           context.Context
	bucket        string
	userPool      string
	orgId         uuid.UUID
	db            database.Database
	s3            *s3.Client
	s3Presign     *s3.PresignClient
	cognitoClient *cognitoidentityprovider.Client
}

func NewProps(
	ctx context.Context,
	bucket string,
	userPool string,
	orgId uuid.UUID,
	db database.Database,
	s3 *s3.Client,
	s3Presign *s3.PresignClient,
	cognitoClient *cognitoidentityprovider.Client,
) *Props {
	return &Props{
		ctx,
		bucket,
		userPool,
		orgId,
		db,
		s3,
		s3Presign,
		cognitoClient,
	}
}

func (p *Props) Ctx() context.Context {
	return p.ctx
}

func (p *Props) Bucket() *string {
	return &p.bucket
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

func (p *Props) S3() *s3.Client {
	return p.s3
}

func (p *Props) S3Presign() *s3.PresignClient {
	return p.s3Presign
}

func (p *Props) Cognito() *cognitoidentityprovider.Client {
	return p.cognitoClient
}
