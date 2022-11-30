package set

import (
	"fmt"
	"testing"
)

type A struct {
	age int
}

func TestHashSet_Add(t *testing.T) {

	m := NewHashSet[*A]()
	type fields struct {
		m Set[*A]
	}
	type args struct {
		v *A
	}
	sameA := &A{age: 5}
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
				v: sameA,
			},
			want: false,
		},
		{
			name: "测试添加2",
			fields: fields{
				m: m,
			},
			args: args{
				v: sameA,
			},
			want: false,
		},
		{
			name: "测试添加3",
			fields: fields{
				m: m,
			},
			args: args{
				v: &A{
					age: 2,
				},
			},
			want: false,
		},
		{
			name: "测试添加3",
			fields: fields{
				m: m,
			},
			args: args{
				v: &A{
					age: 2,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.m.Add(tt.args.v); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
	fmt.Println(m.ToArray())
}
