package document

import (
	doct "doct/internal/document"
	"testing"
)

func TestAppendTransaction_Apply(t *testing.T) {

	type args struct {
		doc *doct.Document
	}

	tests := []struct {
		name string
		tr   *AppendTransaction
		args args
		want string
	}{
		{
			name: "Apply_NonEmptyDoc",
			tr:   &AppendTransaction{Data: []byte("456")},
			args: args{doc: getTestDoct("123")},
			want: "123456",
		},

		{
			name: "Apply_EmptyDoc",
			tr:   &AppendTransaction{Data: []byte("456")},
			args: args{doc: getTestDoct("")},
			want: "456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Apply(tt.args.doc)
			docstring := string(tt.args.doc.Data)
			if docstring != tt.want {
				t.Errorf("AppendTransaction.Apply() = %v, want %v", docstring, tt.want)
			}
		})
	}
}

func TestAppendTransaction_Validate(t *testing.T) {
	type args struct {
		doc *doct.Document
	}

	tests := []struct {
		name string
		tr   *AppendTransaction
		args args
		want bool
	}{
		{
			name: "ValidTransaction_NonEmptyData_NonNilDoc",
			tr:   &AppendTransaction{Data: []byte("456")},
			args: args{doc: getTestDoct("123")},
			want: true,
		},

		{
			name: "InvalidTransaction_EmptyData_NonNilDoc",
			tr:   &AppendTransaction{Data: []byte("")},
			args: args{doc: getTestDoct("123")},
			want: false,
		},

		{
			name: "InvalidTransaction_NonEmptyData_NilDoc",
			tr:   &AppendTransaction{Data: []byte("123")},
			args: args{doc: nil},
			want: false,
		},

		{
			name: "InvalidTransaction_EmptyData_NilDoc",
			tr:   &AppendTransaction{Data: []byte("")},
			args: args{doc: nil},
			want: false,
		},

		{
			name: "InvalidTransaction_NilData_NonNilDoc",
			tr:   &AppendTransaction{Data: nil},
			args: args{doc: getTestDoct("123")},
			want: false,
		},

		{
			name: "InvalidTransaction_NilData_NilDoc",
			tr:   &AppendTransaction{Data: nil},
			args: args{doc: nil},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Validate(tt.args.doc); got != tt.want {
				t.Errorf("AppendTransaction.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
