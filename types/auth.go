package types

import(
	"fmt"
)

type GrantType string

const (
	AuthorizationCode GrantType = "authorization_code"
	Password          GrantType = "password"
	RefreshToken      GrantType = "refresh_token"
	Saml2Bearer       GrantType = "urn:ietf:params:oauth:grant-type:saml2-bearer"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

type AccessTokenCfg struct{
	Assertion 		string
	ClientId		string
	ClientSecret	string
	CustomerId		string
	Password		string
	Username		string
	Grant			GrantType		
}

func (grant GrantType) IsValid() bool{
	switch grant {
		case AuthorizationCode, Password, RefreshToken, Saml2Bearer:
			return true
	}
	return false
}

func (cfg *AccessTokenCfg) IsValid() error {
	if cfg.Grant == "" {
		return fmt.Errorf("empty grant type")
	}
	
	if cfg.Grant != RefreshToken{
		if cfg.ClientId == ""{
			return fmt.Errorf("client_id required for %s grant type", cfg.Grant)
		}
		if cfg.ClientSecret == ""{
			return fmt.Errorf("client_secret required for %s grant type", cfg.Grant)
		}
	}
	
	if cfg.Grant == Saml2Bearer { 
		if cfg.CustomerId == ""{
			return fmt.Errorf("customer_id required for %s grant type", cfg.Grant)
		}
		if cfg.Assertion == ""{
			return fmt.Errorf("assertion required for %s grant type", cfg.Grant)
		}
	}
	
	if cfg.Grant == Password{
		if cfg.Password == ""{
			return fmt.Errorf("password required for %s grant type", cfg.Grant)	
		}
		if cfg.Username == ""{
			return fmt.Errorf("username required for %s grant type", cfg.Grant)
		}
	}
	
	return nil
}