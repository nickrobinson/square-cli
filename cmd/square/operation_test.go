package main

import "testing"

func Test_formatURL(t *testing.T) {
	url := formatURL("/v2/customers/{customer_id}", []string{"39030VYCW8WTF92871ZEEWATX8"})
	if url != "/v2/customers/39030VYCW8WTF92871ZEEWATX8" {
		t.Errorf("Url format was not correct, got: %s, expected: %s", url, "/v2/customers/39030VYCW8WTF92871ZEEWATX8")
	}
}
