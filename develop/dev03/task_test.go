package main

import (
	"reflect"
	"testing"
)

func TestGetValue(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		col      int
		sep      string
		excepted string
		err      error
	}{
		{name: "test1", str: "Яблоко Собака 1", col: 0, sep: " ", excepted: "Яблоко", err: nil},
		{name: "test2", str: "Яблоко; СОбака; 1", col: 0, sep: ";", excepted: "Яблоко", err: nil},
		{name: "test2", str: "Яблоко Собака 1", col: 1, sep: " ", excepted: "Собака", err: nil},
		{name: "test1", str: "Яблоко Собака 1", col: 2, sep: " ", excepted: "1", err: nil},
		{name: "test1", str: "Яблоко Собака 1", col: 3, sep: " ", excepted: "", err: errOutOfRange},
		{name: "test1", str: "Яблоко Собака 1", col: -1, sep: " ", excepted: "", err: errOutOfRange},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := getValue(test.str, test.col, test.sep)
			if actual != test.excepted {
				t.Errorf("actual: %v, excepted: %v", actual, test.excepted)
			}
			if err != test.err {
				t.Errorf("actual: %v, excepted: %v", err, test.err)
			}
		})
	}
}

func TestGetUnique(t *testing.T) {
	tests := []struct {
		name     string
		rows     []string
		excepted []string
	}{
		{name: "test1", rows: []string{"test", "test1", "test2"}, excepted: []string{"test", "test1", "test2"}},
		{name: "test2", rows: []string{"test", "test", "test2"}, excepted: []string{"test", "test2"}},
		{name: "test3", rows: []string{}, excepted: []string{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := getUnique(test.rows)
			if !reflect.DeepEqual(actual, test.excepted) {
				t.Errorf("actual: %v, excepted: %v", actual, test.excepted)
			}
		})
	}
}

func TestCompareAsString(t *testing.T) {
	tests := []struct {
		name     string
		a, b     string
		excepted bool
		err      error
	}{
		{name: "test1", a: "a", b: "b", excepted: true, err: nil},
		{name: "test2", a: "b", b: "a", excepted: false, err: nil},
		{name: "test3", a: "1", b: "2", excepted: true, err: nil},
		{name: "test4", a: "2", b: "1", excepted: false, err: nil},
		{name: "test5", a: "20", b: "3", excepted: true, err: nil},
		{name: "test6", a: "3", b: "20", excepted: false, err: nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := compareAsStrings(test.a, test.b)
			if actual != test.excepted {
				t.Errorf("actual: %v, excepted: %v", actual, test.excepted)
			}
			if err != test.err {
				t.Errorf("actual: %v, excepted: %v", err, test.err)
			}
		})
	}
}

func TestCompareAsNums(t *testing.T) {
	tests := []struct {
		name     string
		a, b     string
		excepted bool
		errNil   bool
	}{
		{name: "test1", a: "a", b: "b", excepted: false, errNil: false},
		{name: "test2", a: "b", b: "a", excepted: false, errNil: false},
		{name: "test3", a: "1", b: "2", excepted: true, errNil: true},
		{name: "test4", a: "2", b: "1", excepted: false, errNil: true},
		{name: "test5", a: "20", b: "3", excepted: false, errNil: true},
		{name: "test6", a: "3", b: "20", excepted: true, errNil: true},
		{name: "test7", a: "3", b: "b", excepted: false, errNil: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := compareAsNums(test.a, test.b)
			if actual != test.excepted {
				t.Errorf("actual: %v, excepted: %v", actual, test.excepted)
			}
			if !((err == nil) == test.errNil) {
				t.Errorf("actual: %v, excepted errNil: %v", err == nil, test.errNil)
			}
		})
	}
}
