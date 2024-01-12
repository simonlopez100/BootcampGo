package exercise1_test

import (
	"lesson3-functions_testing/exercise1"
	"testing"
)

func TestExercise1TaxesAndNetSalary(t *testing.T) {
	t.Run("TestExercise1", func(t *testing.T) {
		//Arrange
		salaryCase1 := 151000.0    // Convert salaryCase1 to float64
		expectedTaxesCase1 := 27.0 // Convert expectedTaxesCase1 to float64
		expectedNetSalaryCase1 := 110230.00

		//Act
		netSalaryCase1, taxesCase1 := exercise1.CalculateSalary(float64(salaryCase1)) // Convert salaryCase1 to float64

		//Assert
		if netSalaryCase1 != expectedNetSalaryCase1 {
			t.Errorf("Expected net salary: %.2f, got: %.2f", expectedNetSalaryCase1, netSalaryCase1)
		}

		if taxesCase1 != expectedTaxesCase1 {
			t.Errorf("Expected taxes: %.2f, got: %.2f", expectedTaxesCase1, taxesCase1)
		}
	})
}
