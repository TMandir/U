package main

import (
	"errors"
	"math"
	"strings"
	"testing"

	stabilitycalculus "stability-calculus"
)

func TestValidateEntityValuesAcceptsBoundaryValues(t *testing.T) {
	entity := constructor()
	entity.Entitystate.S.Ideal.Minidealstate.Entityinfo.value = 0
	entity.Entitystate.S.Ideal.Maxidealstate.Entityinfo.value = 1

	if err := validateEntityValues(entity); err != nil {
		t.Fatalf("validateEntityValues returned an error for valid boundaries: %v", err)
	}
}

func TestValidateEntityValuesRejectsInvalidValues(t *testing.T) {
	tests := []struct {
		name  string
		value float64
	}{
		{name: "negative", value: -0.01},
		{name: "above one", value: 1.01},
		{name: "NaN", value: math.NaN()},
		{name: "positive infinity", value: math.Inf(1)},
		{name: "negative infinity", value: math.Inf(-1)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			entity := constructor()
			entity.Entitystate.S.Ideal.Minidealstate.Entityinfo.value = test.value

			err := validateEntityValues(entity)
			if !errors.Is(err, stabilitycalculus.ErrValueOutOfRange) {
				t.Fatalf("validateEntityValues error = %v, want ErrValueOutOfRange", err)
			}
			if !strings.Contains(err.Error(), "entity id 0") {
				t.Errorf("error %q does not identify the associated entity ID", err)
			}
			if !strings.Contains(err.Error(), "S minimum") {
				t.Errorf("error %q does not identify the invalid field", err)
			}
		})
	}
}

func TestValidateEntityValuesChecksEveryModeledValue(t *testing.T) {
	tests := []struct {
		name  string
		set   func(*Entity)
		label string
	}{
		{
			name: "ideal minimum",
			set: func(entity *Entity) {
				entity.Entitystate.Ideal.Minidealstate.Entityinfo.value = 1.01
			},
			label: "ideal minimum",
		},
		{
			name: "ideal maximum",
			set: func(entity *Entity) {
				entity.Entitystate.Ideal.Maxidealstate.Entityinfo.value = 1.01
			},
			label: "ideal maximum",
		},
		{
			name: "general state",
			set: func(entity *Entity) {
				entity.Entitystate.Generalstate.Entityinfo.value = 1.01
			},
			label: "general state",
		},
		{
			name: "S minimum",
			set: func(entity *Entity) {
				entity.Entitystate.S.Ideal.Minidealstate.Entityinfo.value = 1.01
			},
			label: "S minimum",
		},
		{
			name: "S maximum",
			set: func(entity *Entity) {
				entity.Entitystate.S.Ideal.Maxidealstate.Entityinfo.value = 1.01
			},
			label: "S maximum",
		},
		{
			name: "G minimum",
			set: func(entity *Entity) {
				entity.Entitystate.G.Ideal.Minidealstate.Entityinfo.value = 1.01
			},
			label: "G minimum",
		},
		{
			name: "G maximum",
			set: func(entity *Entity) {
				entity.Entitystate.G.Ideal.Maxidealstate.Entityinfo.value = 1.01
			},
			label: "G maximum",
		},
		{
			name: "K minimum",
			set: func(entity *Entity) {
				entity.Entitystate.K.Ideal.Minidealstate.Entityinfo.value = 1.01
			},
			label: "K minimum",
		},
		{
			name: "K maximum",
			set: func(entity *Entity) {
				entity.Entitystate.K.Ideal.Maxidealstate.Entityinfo.value = 1.01
			},
			label: "K maximum",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			entity := constructor()
			test.set(&entity)

			err := validateEntityValues(entity)
			if !errors.Is(err, stabilitycalculus.ErrValueOutOfRange) {
				t.Fatalf("validateEntityValues error = %v, want ErrValueOutOfRange", err)
			}
			if !strings.Contains(err.Error(), test.label) {
				t.Errorf("error %q does not identify field %q", err, test.label)
			}
		})
	}
}
