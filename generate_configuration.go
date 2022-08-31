package firead

import "github.com/jonnyorman/fireworks"

func GenerateConfiguration() fireworks.Configuration {
	configurationFilePathProvider := fireworks.NewConfigurationFilePathProvider("fireworks-config")

	configurationFileReader := fireworks.NewConfigurationFileReader(configurationFilePathProvider)

	configurationJsonFileReader := fireworks.NewConfigurationJsonFileReader(configurationFileReader)

	configurationJson := configurationJsonFileReader.Read()

	projectIDProvider := fireworks.CreateConfigurationValueProvider("projectID", "PROJECT_ID", configurationJson)

	collectionNameProvider := fireworks.CreateConfigurationValueProvider("collectionName", "COLLECTION_NAME", configurationJson)

	configurationLoader := fireworks.NewApplicationConfigurationLoader(
		projectIDProvider,
		collectionNameProvider,
	)

	configuration := configurationLoader.Load()

	return configuration
}
