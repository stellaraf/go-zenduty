package zenduty

import "github.com/deepmap/oapi-codegen/pkg/securityprovider"

var APIEndpoint string = "https://www.zenduty.com"

func New(key string) (*Zenduty, error) {
	auth, err := securityprovider.NewSecurityProviderBearerToken(key)
	if err != nil {
		return nil, err
	}
	client, err := NewClient(APIEndpoint, WithRequestEditorFn(auth.Intercept))
	if err != nil {
		return nil, err
	}
	return client, nil
}
