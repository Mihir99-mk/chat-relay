package auth

import (
	"auth/config"
	"auth/ent/entgen"
	"auth/util"
	"context"
)

type IStorage interface {
	GetAuthProvider(ctx context.Context) (*entgen.AuthProvider, error)
	UpsertSlackUser(ctx context.Context, in *SlackUser) (*entgen.AuthUser, error)
}

type Storage struct {
	db *entgen.Client
}

func NewStorage(config config.IConfig) IStorage {
	return &Storage{
		db: config.DB(),
	}
}

func (s *Storage) UpsertSlackUser(ctx context.Context, in *SlackUser) (*entgen.AuthUser, error) {
	var result *entgen.AuthUser

	txErr := util.ExecuteTransaction(ctx, s.db, func(tx *entgen.Tx) error {
		client := tx.Client()

		existing, err := getAuthUser(client, in).Only(ctx)
		if err == nil {
			result, err = existing.Update().
				SetAccessToken(in.AccessToken).
				SetTokenType(in.TokenType).
				SetRawProfile(in.RawProfile).
				Save(ctx)
			return err
		}

		result, err = createAuthUser(client, in).Save(ctx)
		return err
	})

	return result, txErr
}

// GetAuthProvider implements IStorage.
func (s *Storage) GetAuthProvider(ctx context.Context) (*entgen.AuthProvider, error) {
	return getAuthProvider(s.db).Only(ctx)
}
