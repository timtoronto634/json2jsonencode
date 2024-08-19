package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var jsonString string

	// Check if input is provided as a command-line argument
	if len(os.Args) > 1 {
		jsonString = os.Args[1]
	} else {
		// If no argument is provided, read from stdin
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal("Error reading from stdin: ", err)
		}
		jsonString = string(bytes)
	}

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		log.Fatal("Error parsing JSON: ", err)
	}

	output := convertToJsonEncode(data)
	fmt.Println(output)
}

func convertToJsonEncode(data map[string]interface{}) string {
	return "jsonencode(" + convertMap(data, 0) + ")"
}

func convertMap(m map[string]interface{}, indent int) string {
	var builder strings.Builder
	indentation := strings.Repeat("  ", indent)
	builder.WriteString("{\n")

	for k, v := range m {
		builder.WriteString(fmt.Sprintf("%s  %s = %s\n", indentation, k, convertValue(v, indent+1)))
	}

	builder.WriteString(indentation + "}")
	return builder.String()
}

func convertValue(v interface{}, indent int) string {
	switch v := v.(type) {
	case string:
		if strings.Contains(v, "\n") {
			return fmt.Sprintf("<<-MARKDOWN\n%sMARKDOWN", v)
		}
		return fmt.Sprintf(`"%s"`, v)
	case bool:
		return fmt.Sprintf("%v", v)
	case float64:
		return fmt.Sprintf("%v", v)
	case []interface{}:
		return convertArray(v, indent)
	case map[string]interface{}:
		return convertMap(v, indent)
	default:
		return fmt.Sprintf(`"%v"`, v)
	}
}

func convertArray(a []interface{}, indent int) string {
	var builder strings.Builder
	indentation := strings.Repeat("  ", indent)
	builder.WriteString("[\n")

	for _, v := range a {
		builder.WriteString(fmt.Sprintf("%s  %s\n", indentation, convertValue(v, indent+1)))
	}

	builder.WriteString(indentation[:len(indentation)-2] + "]")
	return builder.String()
}
