package configuration

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Handle(c echo.Context) error {
	configuration := Configuration{
		Issuer: "",
	}

	return c.JSON(http.StatusOK, configuration)
}

//[
//                'issuer'                                           => Dictionary::getIssuerUrl(),
//                'authorization_endpoint'                           => $base_href . Authorize::URL,
//                'token_endpoint'                                   => $base_href . Token::URL,
//                'check_session_iframe'                             => $base_href . CheckSession::URL,
//                'end_session_endpoint'                             => $base_href . Logout::URL,
//                'jwks_uri'                                         => $base_href . PublicKeys::URL,
//                'scopes_supported'                                 => [
//                    Dictionary::SCOPE_OPENID,
//                    Dictionary::SCOPE_PROFILE,
//                    Dictionary::SCOPE_FIRM,
//                ],
//                'response_types_supported'                         => [
//                    Code::RESPONSE_TYPE,
//                ],
//                'response_modes_supported'                         => [
//                    'form_post',
//                ],
//                'grant_types_supported'                            => [
//                    AuthorizationCode::GRANT_TYPE,
//                    RefreshToken::GRANT_TYPE,
//                    ResourceOwnerPassword::GRANT_TYPE,
//                ],
//                'subject_types_supported'                          => [
//                    'public',
//                ],
//                'token_endpoint_auth_signing_alg_values_supported' => [
//                    'ES512',
//                    'ES384',
//                    'ES256',
//                    'RS512',
//                    'RS384',
//                    'RS256',
//                ],
//                'id_token_signing_alg_values_supported'            => [
//                    'ES512',
//                    'ES384',
//                    'ES256',
//                    'RS512',
//                    'RS384',
//                    'RS256',
//                ],
//                'claims_supported'                                 => [
//                    JwtToken\RegisteredClaims::ISSUER,
//                    JwtToken\RegisteredClaims::ID,
//                    JwtToken\RegisteredClaims::SUBJECT,
//                    JwtToken\RegisteredClaims::AUDIENCE,
//                    JwtToken\RegisteredClaims::ISSUED_AT,
//                    JwtToken\RegisteredClaims::EXPIRATION_TIME,
//                    Dictionary::JWT_CLAIM_TYPE,
//                    Dictionary::JWT_CLAIM_SCOPE,
//                    Dictionary::JWT_CLAIM_AUTH_TIME,
//                    Dictionary::JWT_CLAIM_NONCE,
//                ],
//                'token_endpoint_auth_methods_supported'            => [
//                    'client_secret_post',
//                ],
//                'display_values_supported'                         => [
//                    'page',
//                ],
//
//                // Reserved for the future
//                // 'userinfo_endpoint'     => null,
//                // 'registration_endpoint' => null,
//            ],
