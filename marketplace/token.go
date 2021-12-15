package marketplace

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// BearerToken represents the bearer token sent by the marketplace, and includes the Unix timestamp of the time when
// it expires.
type BearerToken struct {
	Expiration *int    `json:"expiration"`
	Token      *string `json:"access_token"`
}

// DecodeMarketplaceTokenFromResponse decodes the bearer token and the expiration timestamp from the received
// response.
func DecodeMarketplaceTokenFromResponse(response *http.Response) (*BearerToken, error) {
	token := BearerToken{}

	err := json.NewDecoder(response.Body).Decode(&token)
	if err != nil {
		return nil, err
	}

	err = response.Body.Close()
	if err != nil {
		log.Fatalf("could not close the marketplace JWT token response's body: %s", err)
	}

	if token.Expiration == nil || token.Token == nil {
		return nil, fmt.Errorf("unexpected JSON structure received from the marketplace")
	}

	return &token, nil
}
