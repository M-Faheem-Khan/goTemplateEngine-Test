package engine

import (
	"fmt"
	"regexp"
	"strings"
)

// Parses a given string and renders Template Syntax into HTML

const variable_regex = "{{*.[a-zA-Z0-9_].*}}"

func Parse(htmlFile []byte, passedVariables map[string]string) string {
	// convert []byte to string
	htmlFileStr := string(htmlFile)

	// extract syntax
	// Regular Expression for extracting syntax
	extractedVariables := extractVariables(htmlFileStr)

	// Verify the number of passed variable match the number of extract variables
	if matchExtractedVariableToPassedVariables(extractedVariables, passedVariables) {
		// render syntax with html
		htmlFileStr = renderVariable(htmlFileStr, extractedVariables, passedVariables)

		// return rendered html
		return htmlFileStr
	}

	return "ERROR: Check console for errors."

}

func extractVariables(htmlFile string) map[string]string {
	rexp, _ := regexp.Compile(variable_regex)
	matches := rexp.FindAllString(htmlFile, -1)
	matchesMap := make(map[string]string)

	matches = removeDuplicateVariables(matches) // removing duplicate variables

	// Clean up: remove spaces & syntax from matches
	for i := range matches {
		clean_variable := strings.Clone(matches[i])
		clean_variable = strings.ReplaceAll(clean_variable, "{", "")
		clean_variable = strings.ReplaceAll(clean_variable, "}", "")
		clean_variable = strings.ReplaceAll(clean_variable, " ", "")
		matchesMap[clean_variable] = matches[i]
	}
	return matchesMap
}

func renderVariable(htmlFile string, extractedVariables, passedVariables map[string]string) string {
	// iterate over passedVariable
	// find & replace the identifier aka {{ variable_name }} in the htmlFile
	// return the htmlFile

	for k, v := range passedVariables {
		id := extractedVariables[k]
		htmlFile = strings.ReplaceAll(htmlFile, id, v)
	}

	return htmlFile
}

func matchExtractedVariableToPassedVariables(extracted map[string]string, passed map[string]string) bool {
	// number of variable are equal
	if len(extracted) != len(passed) {
		fmt.Println("Number of variables found in template do not equal the number of variable found in passed variables map.")
		return false
	}

	// variable names are different
	for k := range extracted {
		_, found := passed[k]
		if found {
			continue
		} else {
			fmt.Println("Error: variable in template does not exists in passed variables map.")
			return false
		}
	}

	return true
}

func removeDuplicateVariables(extracted []string) []string {
	var v []string
	vMap := make(map[string]int)
	// convert []string to map[string][string]
	// since map can only have 1 key it will remove duplicates
	// convert map back to []string

	for i, v := range extracted {
		vMap[v] = i
	}

	for k := range vMap {
		v = append(v, k)
	}

	return v
}

// EOF
