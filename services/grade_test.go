package services_test

import (
	"lebrancconvas/gounittest/services"
	"testing"
)

func TestCheckGradeA(t *testing.T) {
	grade := services.CheckGrade(80)
	expected := "A"

	if grade != expected {
		t.Errorf("Expected %s but got %s", expected, grade)
	}
}

func TestCheckGradeB(t *testing.T) {
	grade := services.CheckGrade(70)
	expected := "B"

	if grade != expected {
		t.Errorf("Expected %s but got %s", expected, grade)
	}
}

func TestHello(t *testing.T) {

}
