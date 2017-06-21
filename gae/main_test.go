package askmyip

import (
	"testing"
	"net/http/httptest"
)

func TestGetRemoteAddr(t *testing.T) {
	cases := []struct{
		Input string
		Want  string
	}{
		{
			Input: "",
			Want: "0.0.0.0",
		},
		{
			Input: "1.2.3.4",
			Want: "1.2.3.4",
		},
		{
			Input: "1.2.3.4, 5.6.7.8",
			Want: "1.2.3.4",
		},
		{
			Input: " 1.2.3.4 , 5.6.7.8 ",
			Want: "1.2.3.4",
		},
	}

	for _, tc := range cases {
		r := httptest.NewRequest("GET", "/", nil)
		r.RemoteAddr = "0.0.0.0"
		r.Header.Set("X-Forwarded-For", tc.Input)

		if got := GetRemoteAddr(r); got != tc.Want {
			t.Errorf("Expect GetRemoteAddr() to equal %#v, got %#v", tc.Want, got)
		}
	}
}
