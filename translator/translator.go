package translator

import (
	"cdigiuseppe/xctranslate/model"
	"fmt"
	"github.com/bregydoc/gtranslate"
	"github.com/schollz/progressbar/v3"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var mu sync.Mutex

func Transalate(fileContent *model.TranslationContext, targetLang string, wg *sync.WaitGroup) error {
	defer wg.Done()
	bar := progressbar.Default(int64(len(fileContent.SourceStrings)))
	bar.Describe(fmt.Sprintf("translate in %s", targetLang))

	for _, str := range fileContent.SourceStrings {
		translatedText, err := translateText(str.Content, fileContent.SourceLanguage, targetLang)
		if err != nil {
			return err
		}

		fileContent.AddTranslatedString(targetLang, model.StringEntry{
			Identifier: str.Identifier,
			Content:    translatedText,
		})

		bar.Add(1)
	}
	return nil
}

func translateText(text, sourceLang string, targetLang string) (string, error) {
	var protectedText, placeholderMap = protectPlaceholders(text)
	mu.Lock()
	defer mu.Unlock()
	translatedText, err := gtranslate.TranslateWithParams(protectedText, gtranslate.TranslationParams{
		From: sourceLang,
		To:   targetLang,
	})
	if err != nil {
		return "", err
	}
	return restorePlaceholders(translatedText, placeholderMap), nil
}

func protectPlaceholders(text string) (string, map[string]string) {
	pattern := regexp.MustCompile(`(%\d+\$\@|%\w+|\{[^}]+\})`)
	matches := pattern.FindAllString(text, -1)

	placeholderMap := make(map[string]string)
	processed := text

	for i, match := range matches {
		token := "_v" + strconv.Itoa(i)
		placeholderMap[token] = match
		processed = strings.ReplaceAll(processed, match, token)
	}

	return processed, placeholderMap
}

func restorePlaceholders(text string, placeholderMap map[string]string) string {
	for token, original := range placeholderMap {
		text = strings.ReplaceAll(text, token, original)
	}
	return text
}
