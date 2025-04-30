package parser

import (
	"cdigiuseppe/xctranslate/model"
	"encoding/json"
	"fmt"
	"os"
)

func ReadSourceStrings(filePath string) (*model.TranslationContext, *model.XCStrings, error) {
	var identifier int64

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("errore nella lettura del file: %w", err)
	}

	var xc model.XCStrings
	if err := json.Unmarshal(data, &xc); err != nil {
		return nil, nil, fmt.Errorf("errore nel parsing JSON: %w", err)
	}

	var sourceStrings []model.StringEntry

	for key := range xc.Strings {
		identifier++
		sourceStrings = append(sourceStrings, model.StringEntry{Identifier: identifier, Content: key})
	}

	return &model.TranslationContext{
		SourceLanguage: xc.SourceLanguage,
		SourceStrings:  sourceStrings,
	}, &xc, nil
}
