package document

import (
	doct "doct/internal/document"
	"testing"
)

func TestInsertTransaction_Apply(t *testing.T) {
	type args struct {
		doc *doct.Document
	}
	tests := []struct {
		name string
		tr   *InsertTransaction
		args args
		want string
	}{
		{
			name: "Apply_InsertAtEnd",
			tr:   &InsertTransaction{Position: 3, Data: []byte("000")},
			args: args{doc: getTestDoct("123")},
			want: "123000",
		},

		{
			name: "Apply_InsertAtStart",
			tr:   &InsertTransaction{Position: 0, Data: []byte("000")},
			args: args{doc: getTestDoct("123")},
			want: "000123",
		},

		{
			name: "Apply_InsertAtMiddle",
			tr:   &InsertTransaction{Position: 2, Data: []byte("000")},
			args: args{doc: getTestDoct("1234")},
			want: "1200034",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Apply(tt.args.doc)

			docstring := string(tt.args.doc.Data)
			if docstring != tt.want {
				t.Errorf("InsertTransaction.Apply() = %v, want %v", docstring, tt.want)
			}
		})
	}
}

func TestInsertTransaction_Validate(t *testing.T) {
	type args struct {
		doc *doct.Document
	}
	tests := []struct {
		name string
		tr   *InsertTransaction
		args args
		want bool
	}{
		{
			name: "ValidTransaction_NonEmptyDoc_NonEmptyData_ZeroPosition",
			tr:   &InsertTransaction{Position: 0, Data: []byte("123")},
			args: args{getTestDoct("456")},
			want: true,
		},

		{
			name: "ValidTransaction_EmptyDoc_NonEmptyData_GodPosition",
			tr:   &InsertTransaction{Position: 0, Data: []byte("123")},
			args: args{getTestDoct("")},
			want: true,
		},

		{
			name: "ValidTransaction_NonEmptyDoc_NonEmptyData_GoodPosition",
			tr:   &InsertTransaction{Position: 3, Data: []byte("123")},
			args: args{getTestDoct("456")},
			want: true,
		},

		{
			name: "InvalidTransaction_NonEmptyDoc_NonEmptyData_BadPosition",
			tr:   &InsertTransaction{Position: -1, Data: []byte("123")},
			args: args{getTestDoct("456")},
			want: false,
		},

		{
			name: "InvalidTransaction_NonEmptyDoc_NonEmptyData_BadFarPosition",
			tr:   &InsertTransaction{Position: 100, Data: []byte("123")},
			args: args{getTestDoct("456")},
			want: false,
		},

		{
			name: "InvalidTransaction_NilDoc_NonEmptyData_GoodPosition",
			tr:   &InsertTransaction{Position: 0, Data: []byte("123")},
			args: args{nil},
			want: false,
		},

		{
			name: "InvalidTransaction_NoEmptyDoc_EmptyData_GoodPosition",
			tr:   &InsertTransaction{Position: 0, Data: []byte("")},
			args: args{getTestDoct("123")},
			want: false,
		},

		{
			name: "InvalidTransaction_NoEmptyDoc_NilData_GoodPosition",
			tr:   &InsertTransaction{Position: 1, Data: nil},
			args: args{getTestDoct("123")},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Validate(tt.args.doc); got != tt.want {
				t.Errorf("InsertTransaction.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
