package main

import (
	"cdigiuseppe/xctranslate/parser"
	"cdigiuseppe/xctranslate/translator"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

var inputFile *string
var outputFile *string
var targetLangs []string

func main() {
	fmt.Println("XCTranslater - A tool for translating XCStrings files")

	inputFile = flag.String("input", "Localizable.xcstrings", "Path al file xcstrings da tradurre")
	outputFile = flag.String("output", "Translated.xcstrings", "Path file xcstrings tradotto")
	languages := flag.String("langs", "", "Codici lingua di destinazione separati da virgola (es: it,fr,de)")
	flag.Parse()

	if *languages == "" {
		fmt.Println("You must specify destination languages with -langs es: -langs=it,fr,de")
		os.Exit(1)
	}

	targetLangs = strings.Split(*languages, ",")

	fmt.Printf("Input file: %s\n", *inputFile)
	fmt.Printf("Output file: %s\n", *outputFile)
	fmt.Printf("Target languages: %s\n", targetLangs)

	// Read the source strings from the input file
	sourceStrings, xc, err := parser.ReadSourceStrings(*inputFile)
	if err != nil {
		fmt.Printf("Error reading source strings: %v\n", err)
		os.Exit(1)
	}
	fmt.Println()
	fmt.Printf("Readed source file, this is what will i do:\n")
	fmt.Printf(" - source language: %s\n", sourceStrings.SourceLanguage)
	fmt.Printf(" - strings to be translated : %d\n", len(sourceStrings.SourceStrings))
	fmt.Printf(" - will be translated in : %s\n", targetLangs)

	var wg sync.WaitGroup
	wg.Add(len(targetLangs))

	for _, lang := range targetLangs {
		//go translator.Transalate(sourceStrings, lang, &wg)
		translator.Transalate(sourceStrings, lang, &wg)
	}

	wg.Wait()
	fmt.Printf("tradotto : %d stringhe in %d lingue\n", len(sourceStrings.SourceStrings), len(targetLangs))
	xcfileOutout, err := parser.BuildFinalXCStringFile(xc, sourceStrings)
	if err != nil {
		fmt.Printf("Error building final xcstrings file: %v\n", err)
		os.Exit(1)
	}
	err = parser.WriteFile(*outputFile, xcfileOutout)
	if err != nil {
		fmt.Printf("Error writing final xcstrings file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("File saved in %s\n", *outputFile)
}
