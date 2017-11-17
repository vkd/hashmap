package hashmap

import "testing"

func Test_hash(t *testing.T) {
	tests := []struct {
		name    string
		s       Key
		want    uint
		wantErr bool
	}{
	// TODO: Add test cases.
	// {"base string", "hello", 103547413, false},
	// {"base struct", struct {
	// 	I  int
	// 	S  string
	// 	U  uint
	// 	F  float64
	// 	B  bool
	// 	P  *struct{ S []string }
	// 	M  map[string]string
	// 	IF interface{}
	// }{P: &struct{ S []string }{[]string{"name"}}}, 1580146056, false},
	// {"base struct", struct {
	// 	I  int
	// 	S  string
	// 	U  uint
	// 	F  float64
	// 	B  bool
	// 	P  *struct{ S []string }
	// 	PI *int
	// 	M  map[string]interface{}
	// 	IF interface{}
	// }{
	// 	B: true,
	// 	P: &struct{ S []string }{[]string{"name", "age"}},
	// 	M: map[string]interface{}{"name": "igor", "age": 15},
	// }, 3859361036, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hash(64, tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
