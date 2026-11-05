package main

import (
	"fmt"

	stabilitycalculus "stability-calculus"
)

type namedEntityInfo struct {
	label string
	info  Entityinfo
}

// validateEntityValues checks every modeled value with both Stability
// Calculus value types. The entity ID is associated with errors, not validated.
func validateEntityValues(entity Entity) error {
	entityID := entity.Entityidentifer.Entityid.id
	infos := []namedEntityInfo{
		{label: "ideal minimum", info: entity.Entitystate.Ideal.Minidealstate.Entityinfo},
		{label: "ideal maximum", info: entity.Entitystate.Ideal.Maxidealstate.Entityinfo},
		{label: "general state", info: entity.Entitystate.Generalstate.Entityinfo},
		{label: "S minimum", info: entity.Entitystate.S.Ideal.Minidealstate.Entityinfo},
		{label: "S maximum", info: entity.Entitystate.S.Ideal.Maxidealstate.Entityinfo},
		{label: "G minimum", info: entity.Entitystate.G.Ideal.Minidealstate.Entityinfo},
		{label: "G maximum", info: entity.Entitystate.G.Ideal.Maxidealstate.Entityinfo},
		{label: "K minimum", info: entity.Entitystate.K.Ideal.Minidealstate.Entityinfo},
		{label: "K maximum", info: entity.Entitystate.K.Ideal.Maxidealstate.Entityinfo},
	}

	//here we use trust and trustless value both are same just for example purpose.
	//In real-world trust and trustless factors id and value may differ in other cases.

	//Philoshopically, the values are always ideal for S, G, and K.
	//this example is just to show how to use the stability calculus library to validate the values of an entity.
	for _, namedInfo := range infos {
		_, _, err := stabilitycalculus.NewValues(
			namedInfo.info.value,
			namedInfo.info.name,
			namedInfo.info.value,
			namedInfo.info.name,
		)
		if err != nil {
			return fmt.Errorf(
				"entity id %d, %s %q: %w",
				entityID,
				namedInfo.label,
				namedInfo.info.name,
				err,
			)
		}
	}

	return nil
}
