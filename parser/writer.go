package parser

import (
	"cdigiuseppe/xctranslate/model"
	"encoding/json"
	"os"
)

func WriteFile(filePath string, xcStrings model.XCStrings) error {
	data, err := json.MarshalIndent(xcStrings, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func BuildFinalXCStringFile(sourceXCString *model.XCStrings, fileContent *model.TranslationContext) (model.XCStrings, error) {
	var output model.XCStrings

	output.SourceLanguage = sourceXCString.SourceLanguage
	output.Version = sourceXCString.Version
	output.Strings = make(map[string]model.XCStringEntry)

	for _, entry := range fileContent.SourceStrings {
		localizations := make(map[string]model.LocalizedString)
		for lang, translations := range fileContent.GetTranslatedStrings() {
			for _, t := range translations {
				if t.Identifier == entry.Identifier {
					localizations[lang] = model.LocalizedString{
						StringUnit: model.StringUnit{
							State: "translated",
							Value: t.Content,
						},
					}
				}
			}
		}
		output.Strings[entry.Content] = model.XCStringEntry{
			Localizations: localizations,
		}
	}
	return output, nil
}
