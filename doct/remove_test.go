package doct

import (
	"testing"
)

func TestRemoveTransaction_Apply(t *testing.T) {
	type args struct {
		doc *Document
	}
	tests := []struct {
		name string
		tr   *RemoveTransaction
		args args
		want string
	}{
		{
			name: "Apply_FromStart",
			tr:   &RemoveTransaction{Position: 0, Count: 2},
			args: args{doc: getTestDoct("12345")},
			want: "345",
		},
		{
			name: "Apply_FromEnd",
			tr:   &RemoveTransaction{Position: 3, Count: 2},
			args: args{doc: getTestDoct("12345")},
			want: "123",
		},
		{
			name: "Apply_FromMiddle",
			tr:   &RemoveTransaction{Position: 3, Count: 3},
			args: args{doc: getTestDoct("123000456")},
			want: "123456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Apply(tt.args.doc)

			docstring := string(tt.args.doc.Data)
			if docstring != tt.want {
				t.Errorf("RemoveTransaction.Apply() = %v, want %v", docstring, tt.want)
			}
		})
	}
}

func TestRemoveTransaction_Validate(t *testing.T) {
	type args struct {
		doc *Document
	}
	tests := []struct {
		name string
		tr   *RemoveTransaction
		args args
		want bool
	}{
		{
			name: "ValidTransaction_NoEmptyDoc_GoodRange",
			tr:   &RemoveTransaction{Position: 0, Count: 1},
			args: args{doc: getTestDoct("123")},
			want: true,
		},

		{
			name: "InvalidTransaction_EmptyDoc_GoodRange",
			tr:   &RemoveTransaction{Position: 0, Count: 1},
			args: args{doc: getTestDoct("")},
			want: false,
		},

		{
			name: "InvalidTransaction_NilDoc_GoodRange",
			tr:   &RemoveTransaction{Position: 0, Count: 1},
			args: args{doc: nil},
			want: false,
		},

		{
			name: "InvalidTransaction_NoEmptyDoc_BadRange_TooWide",
			tr:   &RemoveTransaction{Position: 0, Count: 10},
			args: args{doc: getTestDoct("1234")},
			want: false,
		},

		{
			name: "InvalidTransaction_NoEmptyDoc_ZeroRange",
			tr:   &RemoveTransaction{Position: 0, Count: 0},
			args: args{doc: getTestDoct("123")},
			want: false,
		},

		{
			name: "InvalidTransaction_NoEmptyDoc_BadRange_NegativePosition",
			tr:   &RemoveTransaction{Position: -1, Count: 0},
			args: args{doc: getTestDoct("123")},
			want: false,
		},

		{
			name: "InvalidTransaction_NoEmptyDoc_BadRange_NegativeCount",
			tr:   &RemoveTransaction{Position: 0, Count: -5},
			args: args{doc: getTestDoct("123")},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Validate(tt.args.doc); got != tt.want {
				t.Errorf("RemoveTransaction.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
