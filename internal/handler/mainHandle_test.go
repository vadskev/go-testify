package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getResponse(url string) *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", url, nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MainHandle)
	handler.ServeHTTP(responseRecorder, req)
	return responseRecorder
}

func TestMainHandlerWhenOk(t *testing.T) {
	type paramArgs struct {
		city  string
		count string
	}
	type expectedArgs struct {
		expectedCode int

		expectedArray    bool
		expectedCount    int
		expectedResponse string
	}

	tests := []struct {
		name     string
		param    paramArgs
		expected expectedArgs
	}{
		{
			name: "Test When Ok",
			param: paramArgs{
				city:  "moscow",
				count: "2",
			},

			expected: expectedArgs{
				expectedCode: 200,

				expectedArray: true,
				expectedCount: 2,
			},
		},
		{
			name: "Test When Wrong City",
			param: paramArgs{
				city:  "UnExistsCity",
				count: "2",
			},

			expected: expectedArgs{
				expectedCode: 400,

				expectedArray:    false,
				expectedResponse: "wrong city value",
			},
		},
		{
			name: "Test When Count More Than Total",
			param: paramArgs{
				city:  "moscow",
				count: "12",
			},

			expected: expectedArgs{
				expectedCode: 200,

				expectedArray: true,
				expectedCount: 4,
			},
		},
		{
			name: "Test Count missing",
			param: paramArgs{
				city: "moscow",
			},

			expected: expectedArgs{
				expectedCode: 400,

				expectedArray:    false,
				expectedResponse: "count missing",
			},
		},
		{
			name: "Test Wrong Count Value",
			param: paramArgs{
				city:  "moscow",
				count: "sdf",
			},

			expected: expectedArgs{
				expectedCode: 400,

				expectedArray:    false,
				expectedResponse: "wrong count value",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			target := fmt.Sprintf("/cafe?city=%s&count=%s", tt.param.city, tt.param.count)
			responseRecorder := getResponse(target)

			require.Equal(t, tt.expected.expectedCode, responseRecorder.Code)

			if tt.expected.expectedArray {
				require.NotEmpty(t, responseRecorder.Body)
				body := strings.Split(responseRecorder.Body.String(), ",")
				assert.Len(t, body, tt.expected.expectedCount)
			}
			if !tt.expected.expectedArray {
				require.NotEmpty(t, responseRecorder.Body)
				body := responseRecorder.Body.String()
				assert.Equal(t, tt.expected.expectedResponse, body)
			}
		})
	}
}
