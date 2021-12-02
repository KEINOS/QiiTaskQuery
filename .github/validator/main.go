/*
This is a tool to validate a JSON data according to the schema.

    go run ./myschema.json ./target_data.json

*/
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/qri-io/jsonschema"
)

var osExit = os.Exit // osExit is a copy of os.Exit to ease testing

func main() {
	if len(os.Args) < 2 {
		exitOnError(errors.New("it requires path to the schema and the data"), "missing arg")
	} else {
		schemaData, err := os.ReadFile(os.Args[1])
		exitOnError(err, "failed to read JSON schema")

		valid, err := os.ReadFile(os.Args[2])
		exitOnError(err, "failed to read target JSON")

		fixes, err := ValidateJSON(schemaData, valid)
		exitOnError(err, "failed to parse during validation")

		if len(fixes) != 0 {
			fmt.Fprintln(os.Stderr, "Err# : Tree : Error msg")

			for i, msg := range fixes {
				fmt.Fprintln(os.Stderr, "#", i, ":", msg)
			}
			osExit(1)
		}
	}
}

// ValidateJSON validates the target with the schema and returns the found
// fixes.
func ValidateJSON(schema, target []byte) (fixes []string, err error) {
	ctx := context.Background()
	rs := &jsonschema.Schema{}

	if err = json.Unmarshal(schema, rs); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal schema")
	}

	errs, err := rs.ValidateBytes(ctx, target)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal target JSON")
	}

	if len(errs) > 0 {
		for _, fix := range errs {
			fixes = append(fixes, fix.Error())
		}
	}

	return fixes, nil
}

func exitOnError(err error, msgErr string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, errors.Wrap(err, msgErr))
		osExit(1)
	}
}
