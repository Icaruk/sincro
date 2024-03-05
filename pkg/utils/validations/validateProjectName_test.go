package validation

import "testing"

func TestValidateProjectName(t *testing.T) {
	tests := []struct {
		testName    string
		projectName string
		wantIsValid bool
	}{
		{
			testName:    "valid name",
			projectName: "project",
			wantIsValid: true,
		},
		{
			testName:    "valid name",
			projectName: "prj1",
			wantIsValid: true,
		},
		{
			testName:    "valid name",
			projectName: "123prj",
			wantIsValid: true,
		},
		{
			testName:    "valid name",
			projectName: "PROJECT",
			wantIsValid: true,
		},
		{
			testName:    "too long",
			projectName: "projectaaasssddd",
			wantIsValid: false,
		},
		{
			testName:    "invalid chars",
			projectName: "prj_?)/",
			wantIsValid: false,
		},
		{
			testName:    "empty",
			projectName: "",
			wantIsValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			gotIsValid, _ := ValidateProjectName(tt.projectName)
			if gotIsValid != tt.wantIsValid {
				t.Errorf("ValidateProjectName() = %v, want %v", gotIsValid, tt.wantIsValid)
			}
		})
	}
}
