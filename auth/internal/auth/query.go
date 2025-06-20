package auth

import (
	"auth/ent/entgen"
	"auth/ent/entgen/authprovider"
	"auth/ent/entgen/authuser"
)

func createAuthUser(db *entgen.Client, in *SlackUser) *entgen.AuthUserCreate {
	return db.AuthUser.Create().
		SetProviderID(in.ProviderID).
		SetProviderUserID(in.ProviderUserID).
		SetTeamID(in.TeamID).
		SetName(in.Name).
		SetRealName(in.RealName).
		SetEmail(in.Email).
		SetAccessToken(in.AccessToken).
		SetRefreshToken(in.RefreshToken).
		SetTokenType(in.TokenType).
		SetScope(in.Scope).
		SetRawProfile(in.RawProfile)
}

func getAuthUser(db *entgen.Client, in *SlackUser) *entgen.AuthUserQuery {
	return db.AuthUser.
		Query().
		Where(authuser.ProviderUserID(in.ProviderUserID))
}

func getAuthProvider(db *entgen.Client) *entgen.AuthProviderQuery {
	return db.AuthProvider.Query().Where(authprovider.NameEQ("slack"))
}
