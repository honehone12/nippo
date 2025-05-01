package viewprops

import (
	"context"
	"nippo/database"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type Props struct {
	ctx       context.Context
	bucket    string
	orgId     uuid.UUID
	db        database.Database
	s3        *s3.Client
	s3Presign *s3.PresignClient
}

func NewProps(
	ctx context.Context,
	bucket string,
	orgId uuid.UUID,
	db database.Database,
	s3 *s3.Client,
	s3Presign *s3.PresignClient,
) *Props {
	return &Props{
		ctx,
		bucket,
		orgId,
		db,
		s3,
		s3Presign,
	}
}

func (p *Props) Ctx() context.Context {
	return p.ctx
}

func (p *Props) Bucket() *string {
	return &p.bucket
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
