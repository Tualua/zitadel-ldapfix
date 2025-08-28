package view

import (
	"github.com/jinzhu/gorm"

	"github.com/Tualua/zitadel-ldapfix/internal/crypto"
	"github.com/Tualua/zitadel-ldapfix/internal/database"
	"github.com/Tualua/zitadel-ldapfix/internal/eventstore"
	"github.com/Tualua/zitadel-ldapfix/internal/query"
)

type View struct {
	Db           *gorm.DB
	client       *database.DB
	keyAlgorithm crypto.EncryptionAlgorithm
	query        *query.Queries
	es           *eventstore.Eventstore
}

func StartView(sqlClient *database.DB, keyAlgorithm crypto.EncryptionAlgorithm, queries *query.Queries, es *eventstore.Eventstore) (*View, error) {
	gorm, err := gorm.Open("postgres", sqlClient.DB)
	if err != nil {
		return nil, err
	}
	return &View{
		Db:           gorm,
		client:       sqlClient,
		keyAlgorithm: keyAlgorithm,
		query:        queries,
		es:           es,
	}, nil
}

func (v *View) Health() (err error) {
	return v.Db.DB().Ping()
}
