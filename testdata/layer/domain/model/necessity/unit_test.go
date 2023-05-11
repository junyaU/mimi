package necessity

import "testing"

func TestNewUnit(t *testing.T) {
	tests := []struct {
		value   string
		wantErr bool
	}{
		{value: "kg", wantErr: false},
		{value: "g", wantErr: false},
		{value: "l", wantErr: false},
		{value: "ml", wantErr: false},
		{value: "kgg", wantErr: true},
		{value: "mml", wantErr: true},
		{value: "ll", wantErr: true},
		{value: "", wantErr: true},
	}

	for i, test := range tests {
		if !test.wantErr {
			if unit, err := newUnit(test.value); err != nil {
				t.Fatalf("#%d: bad return value: %#v", i, unit)
			}
		} else {
			if unit, err := newUnit(test.value); err == nil {
				t.Fatalf("#%d: bad return value: %#v", i, unit)
			}
		}
	}
}

func TestUnit_String(t *testing.T) {
	tests := []struct {
		unit    unit
		textVal string
		wantErr bool
	}{
		{unit: kg, textVal: "kg", wantErr: false},
		{unit: g, textVal: "g", wantErr: false},
		{unit: l, textVal: "l", wantErr: false},
		{unit: ml, textVal: "ml", wantErr: false},
		{unit: kg, textVal: "g", wantErr: true},
		{unit: g, textVal: "kg", wantErr: true},
		{unit: l, textVal: "ml", wantErr: true},
		{unit: ml, textVal: "l", wantErr: true},
	}

	for i, test := range tests {
		if !test.wantErr {
			if test.unit.String() != test.textVal {
				t.Fatalf("#%d: bad return value: %#v", i, test.unit.String())
			}
		} else {
			if test.unit.String() == test.textVal {
				t.Fatalf("#%d: want error but no error: %#v", i, test.unit.String())
			}
		}

	}
}
