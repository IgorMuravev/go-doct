package doct

import (
	"testing"
)

func TestEraseTransaction_Apply(t *testing.T) {
	type args struct {
		doc *Document
	}
	tests := []struct {
		name string
		tr   *EraseTransaction
		args args
		want string
	}{
		{
			name: "Apply",
			tr:   &EraseTransaction{},
			args: args{doc: getTestDoct("123")},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Apply(tt.args.doc)

			docstring := string(tt.args.doc.Data)
			if docstring != tt.want {
				t.Errorf("EraseTransaction.Apply() = %v, want %v", docstring, tt.want)
			}
		})
	}
}

func TestEraseTransaction_Validate(t *testing.T) {
	type args struct {
		doc *Document
	}
	tests := []struct {
		name string
		tr   *EraseTransaction
		args args
		want bool
	}{
		{
			name: "InvalidTransaction_NilDoc",
			tr:   &EraseTransaction{},
			args: args{doc: nil},
			want: false,
		},

		{
			name: "InvalidTransaction_EmptyDoc",
			tr:   &EraseTransaction{},
			args: args{doc: getTestDoct("")},
			want: false,
		},

		{
			name: "ValidTransaction_NoNilDoc",
			tr:   &EraseTransaction{},
			args: args{doc: getTestDoct("123")},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Validate(tt.args.doc); got != tt.want {
				t.Errorf("EraseTransaction.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
