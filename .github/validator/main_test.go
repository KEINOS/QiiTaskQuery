package main

import (
	"os"
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
//  Main Test
// ----------------------------------------------------------------------------

func Test_main(t *testing.T) {
	// Mock args
	restoreArgs := mockArgs(t, []string{
		"./testdata/schema.json",
		"./testdata/target.json",
	})
	defer restoreArgs()

	// Run main and capture STDOUT/STDERR
	out := capturer.CaptureOutput(func() {
		main()
	})

	assert.Empty(t, out)
}

func Test_main_detect_fixes(t *testing.T) {
	// Mock os.Exit
	var capturedCode int

	restoreOsExit := mockOsExit(t, &capturedCode)
	defer restoreOsExit()

	// Mock args
	restoreArgs := mockArgs(t, []string{
		"./testdata/schema.json",
		"./testdata/target_ng.json",
	})
	defer restoreArgs()

	// Run main and capture STDOUT/STDERR
	out := capturer.CaptureOutput(func() {
		main()
	})

	expectCode := 1
	actualCode := capturedCode
	assert.Equal(t, expectCode, actualCode)

	assert.Contains(t, out, `"firstName" value is required`)
	assert.Contains(t, out, `"lastName" value is required`)
}

func Test_main_missing_args(t *testing.T) {
	// Mock os.Exit
	var capturedCode int

	restoreOsExit := mockOsExit(t, &capturedCode)
	defer restoreOsExit()

	// Mock args
	restoreArgs := mockArgs(t, []string{})
	defer restoreArgs()

	// Run main and capture STDOUT/STDERR
	out := capturer.CaptureOutput(func() {
		main()
	})

	expectCode := 1
	actualCode := capturedCode
	assert.Equal(t, expectCode, actualCode)

	assert.Contains(t, out, "missing arg: it requires path to the schema and the data")
}

func Test_main_fail_to_read_file(t *testing.T) {
	// Mock os.Exit
	var capturedCode int

	restoreOsExit := mockOsExit(t, &capturedCode)
	defer restoreOsExit()

	// Mock args
	restoreArgs := mockArgs(t, []string{
		"./unknownSchema",
		"./unknownTarget",
	})
	defer restoreArgs()

	// Run main and capture STDOUT/STDERR
	out := capturer.CaptureOutput(func() {
		main()
	})

	expectCode := 1
	actualCode := capturedCode
	assert.Equal(t, expectCode, actualCode)

	assert.Contains(t, out, "failed to read JSON schema")
	assert.Contains(t, out, "failed to read target JSON")
}

// ----------------------------------------------------------------------------
//  Function Test
// ----------------------------------------------------------------------------

func TestValidateJSON(t *testing.T) {
	schema := getDummySchema(t)
	data := []byte(`{
		"firstName": "Taro",
		"lastName": "Yamada",
		"friends": [
			{
				"firstName": "Hanako",
				"lastName": "Yamad"
			}
		]
	}`)

	fixes, err := ValidateJSON(schema, data)

	require.NoError(t, err)
	assert.Nil(t, fixes)
}

func TestValidateJSON_invalidFriend(t *testing.T) {
	schema := getDummySchema(t)
	data := []byte(`{
		"firstName": "Alice",
		"lastName": "Yamada",
		"friends": [
			{"firstName": "Bob" }
		]
	}`)

	fixes, err := ValidateJSON(schema, data)

	require.NoError(t, err)
	assert.Contains(t, fixes[0], `{"firstName":"Bob"} "lastName" value is required`)
}

func TestValidateJSON_invalidPerson(t *testing.T) {
	schema := getDummySchema(t)
	data := []byte(`{
		"firstName": "Alice",
		"friends": [
			{"firstName": "Bob", "lastName": "Yamada" }
		]
	}`)

	fixes, err := ValidateJSON(schema, data)

	require.NoError(t, err)
	assert.Contains(t, fixes[0], `{"firstName":"Alice"... "lastName" value is required`)
}

func TestValidateJSON_invalidData(t *testing.T) {
	schema := getDummySchema(t)
	data := []byte(`
	firstName:
		- Alice
	friends:
		firstName: Bob
		lastName: Yamada
`)

	fixes, err := ValidateJSON(schema, data)

	require.Error(t, err)
	assert.Nil(t, fixes, "fixes should be nil on error")
}

// ----------------------------------------------------------------------------
//  Helper Funxtion
// ----------------------------------------------------------------------------

var dummySchema = []byte(`{
    "$id": "https://qri.io/schema/",
    "$comment": "sample comment",
    "title": "Person",
    "type": "object",
    "properties": {
        "firstName": {
            "type": "string"
        },
        "lastName": {
            "type": "string"
        },
        "age": {
            "description": "Age in years",
            "type": "integer",
            "minimum": 0
        },
        "friends": {
            "type": "array",
            "items": {
                "title": "REFERENCE",
                "$ref": "#"
            }
        }
    },
    "required": [
        "firstName",
        "lastName"
    ]
}`)

func getDummySchema(t *testing.T) []byte {
	t.Helper()

	return dummySchema
}

func mockArgs(t *testing.T, args []string) (deferFunc func()) {
	t.Helper()

	// Backup original args
	oldOsArgs := os.Args

	// Mock os.Args
	os.Args = append([]string{t.Name()}, args...)

	return func() {
		os.Args = oldOsArgs // Restrore backup
	}
}

func mockOsExit(t *testing.T, capturedCode *int) (deferFunc func()) {
	t.Helper()

	// Backup original function
	oldOsExit := osExit

	// Mock osExit
	osExit = func(code int) {
		*capturedCode = code
	}

	return func() {
		osExit = oldOsExit // Restrore backup
	}
}
