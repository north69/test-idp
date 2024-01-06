package configuration

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"test-idp/internal/config"
)

const (
	openidIssuerUrl            = "/auth/openid/"
	authorizeUrl               = "/auth/openid/authorize/"
	tokenUrl                   = "/auth/openid/token/"
	checkSessionUrl            = "/auth/openid/check_session/"
	logoutUrl                  = "/auth/openid/logout/"
	publicKeysUrl              = "/auth/openid/.well-known/jwks.json"
	scopeOpenid                = "openid"
	scopeUser                  = "user"
	responseTypeCode           = "code"
	authorizationCodeGrantType = "authorization_code"
	refreshTokenGrantType      = "refresh_token"
	resourceOwnerPassword      = "password"
)

type Handler struct {
	cnf *config.Config
}

func Init(cnf *config.Config) *Handler {
	return &Handler{
		cnf: cnf,
	}
}

func (h Handler) Handle(c echo.Context) error {
	baseHref := h.cnf.App.BaseHref()
	configuration := Configuration{
		Issuer:                 baseHref + openidIssuerUrl,
		AuthorizationEndpoint:  baseHref + authorizeUrl,
		TokenEndpoint:          baseHref + tokenUrl,
		CheckSessionIframe:     baseHref + checkSessionUrl,
		EndSessionEndpoint:     baseHref + logoutUrl,
		JWKSUri:                baseHref + publicKeysUrl,
		ScopesSupported:        []string{scopeOpenid, scopeUser},
		ResponseTypesSupported: []string{responseTypeCode},
		ResponseModesSupported: []string{"form_post"},
		GrantTypesSupported:    []string{authorizationCodeGrantType, refreshTokenGrantType, resourceOwnerPassword},
		SubjectTypesSupported:  []string{"public"},
		TokenEndpointAuthSigningAlgValuesSupported: []string{"ES512"},
		IdTokenSigningAlgValuesSupported:           []string{"ES512"},
		ClaimsSupported:                            []string{"issuer", "id", "subject", "audience", "issued_at", "expiration_time"},
		TokenEndpointAuthMethodsSupported:          []string{"client_secret_post"},
		DisplayValuesSupported:                     []string{"page"},
	}

	return c.JSON(http.StatusOK, configuration)
}
