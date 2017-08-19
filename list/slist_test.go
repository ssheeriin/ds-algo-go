package list

import "testing"

func TestSLinkedList_AddToHead(t *testing.T) {
	type fields struct {
		head *slnode
	}
	type args struct {
		data int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := &SLinkedList{
				head: tt.fields.head,
			}
			ll.AddToHead(tt.args.data)
		})
	}
}
