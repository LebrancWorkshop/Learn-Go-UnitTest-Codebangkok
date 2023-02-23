package services_test

import (
	"lebrancconvas/gounittest/services"
	"testing"
)

func TestCheckGrade(t *testing.T) {
	t.Run("A", func(t *testing.T) {
		grade := services.CheckGrade(80)
		expected := "A"

		if grade != expected {
			t.Errorf("Expected %s but got %s", expected, grade)
		}
	})
	t.Run("B", func(t *testing.T) {
		grade := services.CheckGrade(70)
		expected := "B"

		if grade != expected {
			t.Errorf("Expected %s but got %s", expected, grade)
		}
	})
	t.Run("C", func(t *testing.T) {
		grade := services.CheckGrade(60)
		expected := "C"

		if grade != expected {
			t.Errorf("Expected %s but got %s", expected, grade)
		}
	})
}
func TestHello(t *testing.T) {

}
