package generate

import (
	"testing"
)

func TestConvertModel(t *testing.T) {
	s := "account"
	if r := ConvertModel(s); r == "Account" {
		t.Log(r)
	} else {
		t.Error("the test is failed")
	}

	s1 := "cloud-account"
	if r := ConvertModel(s1); r == "CloudAccount" {
		t.Log(r)
	} else {
		t.Error("the test is failed")
	}
}

func TestConvert2Camel(t *testing.T) {
	s := "account"
	if r := Convert2Camel(s); r == "account" {
		t.Log(r)
	} else {
		t.Error("the test is failed")
	}

	s1 := "cloud-account"
	if r := Convert2Camel(s1); r == "cloudAccount" {
		t.Log(r)
	} else {
		t.Error("the test is failed")
	}
}

func TestGenerateSimplyName(t *testing.T) {
	s := "account"
	if r := GenerateSimplyName(s); r == "as" {
		t.Log(r)
	} else {
		t.Error("the test is failed")
	}

	s1 := "cloud-account"
	if r := GenerateSimplyName(s1); r == "cas" {
		t.Log(r)
	} else {
		t.Error("the test is failed")
	}
}
