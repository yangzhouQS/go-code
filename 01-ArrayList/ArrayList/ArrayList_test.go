package ArrayList

import (
	"reflect"
	"testing"
)

func TestArrayList_Append(t *testing.T) {
	type fields struct {
		dataStore []interface{}
		size      int
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &ArrayList{
				dataStore: tt.fields.dataStore,
				size:      tt.fields.size,
			}
			if got := list.Append(tt.args.value); got != tt.want {
				t.Errorf("Append() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_Clear(t *testing.T) {
	type fields struct {
		dataStore []interface{}
		size      int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &ArrayList{
				dataStore: tt.fields.dataStore,
				size:      tt.fields.size,
			}
		})
	}
}

func TestArrayList_Delete(t *testing.T) {
	type fields struct {
		dataStore []interface{}
		size      int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &ArrayList{
				dataStore: tt.fields.dataStore,
				size:      tt.fields.size,
			}
			if err := list.Delete(tt.args.index); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArrayList_Get(t *testing.T) {
	type fields struct {
		dataStore []interface{}
		size      int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &ArrayList{
				dataStore: tt.fields.dataStore,
				size:      tt.fields.size,
			}
			got, err := list.Get(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_Insert(t *testing.T) {
	type fields struct {
		dataStore []interface{}
		size      int
	}
	type args struct {
		index int
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &ArrayList{
				dataStore: tt.fields.dataStore,
				size:      tt.fields.size,
			}
			if err := list.Insert(tt.args.index, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArrayList_Set(t *testing.T) {
	type fields struct {
		dataStore []interface{}
		size      int
	}
	type args struct {
		index int
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &ArrayList{
				dataStore: tt.fields.dataStore,
				size:      tt.fields.size,
			}
			if err := list.Set(tt.args.index, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArrayList_Size(t *testing.T) {
	type fields struct {
		dataStore []interface{}
		size      int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &ArrayList{
				dataStore: tt.fields.dataStore,
				size:      tt.fields.size,
			}
			if got := list.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_String(t *testing.T) {
	type fields struct {
		dataStore []interface{}
		size      int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &ArrayList{
				dataStore: tt.fields.dataStore,
				size:      tt.fields.size,
			}
			if got := list.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_checkIsFull(t *testing.T) {
	type fields struct {
		dataStore []interface{}
		size      int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &ArrayList{
				dataStore: tt.fields.dataStore,
				size:      tt.fields.size,
			}
		})
	}
}

func TestNewArrayList(t *testing.T) {
	tests := []struct {
		name string
		want *ArrayList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArrayList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArrayList() = %v, want %v", got, tt.want)
			}
		})
	}
}