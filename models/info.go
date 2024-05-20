package models

type (
	InfoSettings struct {
		CommonSettings
		ConnectionSettings
	}
)

func NewInfoSettings(arguments []string) InfoSettings {
	return InfoSettings{
		CommonSettings: CommonSettings{
			Arguments: arguments,
		},
		ConnectionSettings: ConnectionSettings{
			Driver:   "",
			Url:      "",
			Username: "",
		},
	}
}
