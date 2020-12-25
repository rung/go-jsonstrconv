package jsonstrconv

import (
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		value   string
		want    string
		wantErr bool
	}{
		// Not changed
		{
			value: `{"id":"12345","value1":"1259457","value2":"12345"}`,
			want:  `{"id":"12345","value1":"1259457","value2":"12345"}`,
		},
		// Escaped "
		{
			value: `{"id":"12\"345","value1":"12\\59457","value2":"\"12345"}`,
			want:  `{"id":"12\"345","value1":"12\\59457","value2":"\"12345"}`,
		},
		// One value is changed
		{
			value: `{"id":"12345","value1":1259457,"value2":"12345"}`,
			want:  `{"id":"12345","value1":"1259457","value2":"12345"}`,
		},
		// Two values are changed
		{
			value: `{"id":"12345","value1":1259457,"value2":12345}`,
			want:  `{"id":"12345","value1":"1259457","value2":"12345"}`,
		},
		// Use list
		{
			value: `[{"id":"12345","value1":1259457},{"id":"12345","value1":1259457},{"id":"12345","value1":1259457}]`,
			want:  `[{"id":"12345","value1":"1259457"},{"id":"12345","value1":"1259457"},{"id":"12345","value1":"1259457"}]`,
		},
		// Multi-byte character is used
		{
			value: `{"id":"123こんにちは世界","value1":1259457,"value2":"12345"}`,
			want:  `{"id":"123こんにちは世界","value1":"1259457","value2":"12345"}`,
		},
		// Have whitespace1
		{
			value: `{"id":"123こんにちは世界"   ,   "value1" :  1259457  ,"value2":"12345"}`,
			want:  `{"id":"123こんにちは世界"   ,   "value1" :  "1259457"  ,"value2":"12345"}`,
		},
		// Have whitespace2
		{
			value: `{"id":"123こんにちは世界"   ,   "value1" :

1259457  
,"value2":"12345"}`,
			want: `{"id":"123こんにちは世界"   ,   "value1" :

"1259457"  
,"value2":"12345"}`,
		},
		// Value only
		{
			value: `""`,
			want:  `""`,
		},
		// Value only
		{
			value: `"test"`,
			want:  `"test"`,
		},
		// Value only
		{
			value: `123`,
			want:  `"123"`,
		},
		// Error
		{
			value: `{id:"a","value1":1259457,"value2":"12345"}`,
			wantErr: true,
		},
	}

	for _, test := range tests {
		got, err := ToString([]byte(test.value))
		if (test.wantErr && err == nil) || !test.wantErr && err != nil {
			t.Errorf("err is wrong, value=%v, wantErr=%v, err=%v", test.value, test.wantErr, err)
			continue
		}
		if string(got) != test.want {
			t.Errorf("value is wrong, value=%v, want=%v, got=%v", test.value, test.want, string(got))
		}
	}
}
