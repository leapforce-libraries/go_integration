package integration

import (
	"fmt"

	credentials "github.com/leapforce-libraries/go_google/credentials"
)

type Config struct {
	AppName               string
	ProjectID             string
	Bucket                string
	Dataset               string
	SentryDSN             string
	settings              map[string]string
	ServiceAccountJSONKey *credentials.CredentialsJSON
	LogOrganisationID     *int64   // if the integration runs for a single organisation pass it's ID here
	OrganisationIDs       *[]int64 // if nil, app runs for all organisationIDs not specified in any config from OtherConfigs
}

func (config *Config) Set(key interface{}, value string) {
	if config.settings == nil {
		config.settings = make(map[string]string)
	}

	config.settings[fmt.Sprintf("%v", key)] = value
}

func (config *Config) Get(key interface{}) (string, bool) {
	if config.settings == nil {
		return "", false
	}

	value, ok := config.settings[fmt.Sprintf("%v", key)]

	return value, ok
}

func (config *Config) FullTableName(tableName string, quoted bool) string {
	_tableName := fmt.Sprintf("%s.%s.%s", config.ProjectID, config.Dataset, tableName)

	if quoted {
		_tableName = fmt.Sprintf("`%s`", _tableName)
	}

	return _tableName
}
