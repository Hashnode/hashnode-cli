package posts

import "testing"

func Test_getResponses(t *testing.T) {
	type args struct {
		postID    string
		perPage   int
		page      int
		sortOrder string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				postID:    "5c34ffe75aa5738b323c386b",
				page:      2,
				perPage:   5,
				sortOrder: "totalReactions",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getResponses(tt.args.postID, tt.args.perPage, tt.args.page, tt.args.sortOrder)
		})
	}
}
