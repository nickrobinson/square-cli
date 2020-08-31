package main

import "testing"

func Test_formatURL(t *testing.T) {
	url := formatURL("/v2/customers/{customer_id}", []string{"39030VYCW8WTF92871ZEEWATX8"})
	if url != "/v2/customers/39030VYCW8WTF92871ZEEWATX8" {
		t.Errorf("Url format was not correct, got: %s, expected: %s", url, "/v2/customers/39030VYCW8WTF92871ZEEWATX8")
	}
}

func Test_getUseString(t *testing.T) {
	name := "test"
	args := []string{"{customer-id}"}
	useString := buildUseString(name, args)
	expectedString := "test <customer-id>"
	if useString != expectedString {
		t.Errorf("Use string was not correct, got: %s, expected: %s", useString, expectedString)
	}
}
