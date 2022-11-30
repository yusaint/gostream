package hash

import (
	"testing"
)

func TestMap_Add(t *testing.T) {
	m := NewMap[int, int]()
	type fields struct {
		m Hash[int, int]
	}
	type args struct {
		k int
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "测试添加1",
			fields: fields{
				m: m,
			},
			args: args{
				k: 10,
				v: 1,
			},
			want: false,
		},
		{
			name: "测试添加2",
			fields: fields{
				m: m,
			},
			args: args{
				k: 11,
				v: 3,
			},
			want: false,
		},
		{
			name: "测试添加3",
			fields: fields{
				m: m,
			},
			args: args{
				k: 111,
				v: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.m.Add(tt.args.k, tt.args.v); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
	m.Front()
	var index int
	for {
		v := m.Current()
		t.Log(v)
		index++
		if m.Next() == false {
			break
		}
	}
}
