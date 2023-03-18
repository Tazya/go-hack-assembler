package utils

import "testing"

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{
			name:  "String is numeric",
			value: "24",
			want:  true,
		},
		{
			name:  "String is not numeric",
			value: "hello2",
			want:  false,
		},
		{
			name:  "String is empty",
			value: "",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumeric(tt.value); got != tt.want {
				t.Errorf("IsNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveComments(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "Remove single line comment",
			value: "// Comment",
			want:  "",
		},
		{
			name:  "Remove instruction line comment",
			value: "D=D+1 // Comment",
			want:  "D=D+1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveComments(tt.value); got != tt.want {
				t.Errorf("RemoveComments() = %v, want %v", got, tt.want)
			}
		})
	}
}
