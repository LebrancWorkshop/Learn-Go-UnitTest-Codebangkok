package services_test

import (
	"github.com/LebrancWorkshop/Learn-Go-UnitTest-Codebangkok/services"
	"testing"
)

func TestCheckGrade(t *testing.T) {
	grade := services.CheckGrade();
	expected := "A";

	if grade != expected {
		t.Errorf("got %v expected %v", expected, grade);
	}
}