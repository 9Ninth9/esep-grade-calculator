package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 48, Assignment)
	gradeCalculator.AddGrade("exam 1", 63, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 37, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
func TestGetGradeC(t *testing.T) {
    gradeCalculator := NewGradeCalculator()

    gradeCalculator.AddGrade("assignment", 70, Assignment)
    gradeCalculator.AddGrade("exam", 70, Exam)
    gradeCalculator.AddGrade("essay", 70, Essay)

    expected_value := "C"
    actual_value := gradeCalculator.GetFinalGrade()

    if expected_value != actual_value {
        t.Errorf("Expected C, got %s", actual_value)
    }
}
func TestGetGradeD(t *testing.T) {
    gradeCalculator := NewGradeCalculator()

    gradeCalculator.AddGrade("assignment", 60, Assignment)
    gradeCalculator.AddGrade("exam", 60, Exam)
    gradeCalculator.AddGrade("essay", 60, Essay)

    expected_value := "D"
    actual_value := gradeCalculator.GetFinalGrade()

    if expected_value != actual_value {
        t.Errorf("Expected D, got %s", actual_value)
    }
}
func TestGradeTypeString(t *testing.T) {
    cases := []struct {
        gradeType GradeType
        expected  string
    }{
        {Assignment, "assignment"},
        {Exam, "exam"},
        {Essay, "essay"},
    }

    for _, tc := range cases {
        actual := tc.gradeType.String()

        if actual != tc.expected {
            t.Errorf("Expected %s, got %s", tc.expected, actual)
        }
    }
}
func TestComputeAverageEmpty(t *testing.T) {
    grades := []Grade{}

    expected_value := 0
    actual_value := computeAverage(grades)

    if expected_value != actual_value {
        t.Errorf("Expected 0, got %d", actual_value)
    }
}