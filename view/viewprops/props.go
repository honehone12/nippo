package viewprops

import (
	"context"
	"nippo/database"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Props struct {
	ctx      context.Context
	bucket   string
	orgIdStr string
	db       database.Database
	s3       *s3.Client
}

func NewProps(
	ctx context.Context,
	bucket string,
	orgIdStr string,
	db database.Database,
	s3 *s3.Client,
) *Props {
	return &Props{ctx, bucket, orgIdStr, db, s3}
}

func (p *Props) Ctx() context.Context {
	return p.ctx
}

func (p *Props) Bucket() *string {
	return &p.bucket
}

func (p *Props) OrgId() string {
	return p.orgIdStr
}

func (p *Props) Db() database.Database {
	return p.db
}

func (p *Props) S3() *s3.Client {
	return p.s3
}
