// constructing a base layer of entities with ideal states and identifiers for S, G, and K with respect to S. Each entity has a unique identifier and ideal state information.
// this is sample example that the future work will take reference to this work.
package main

import "fmt"

func main() {

	entity := constructor()
	entrytest(entity)

}

func entrytest(entity Entity) {

	So(entity)
	Go(entity)
	Ko(entity)
}

type Others struct {
	Otherinfo string
}
type Entity struct {
	Entityidentifer
	Entitystate
	Others
}

type Entitystate struct {
	Ideal
	Generalstate
	S
	G
	K
	Others
}
type Entityidentifer struct {
	Entityid
	Others
}
type Entityid struct {
	name string
	id   int
	Others
}
type Ideal struct {
	Minidealstate
	Maxidealstate
	Others
}

type Entityinfo struct {
	name  string
	value float64
	Others
}

type Minidealstate struct {
	Entityinfo
	Others
}
type Maxidealstate struct {
	Entityinfo
	Others
}
type Generalstate struct {
	Entityinfo
	Others
}

type S struct {
	Ideal
	Others
}
type G struct {
	Ideal
	Others
}
type K struct {
	Ideal
	Others
}

func constructor() Entity {
	return Entity{
		Entityidentifer: Entityidentifer{
			Entityid: Entityid{
				name: "S",
				id:   0,
			},
		},
		Entitystate: Entitystate{
			Ideal: Ideal{
				Minidealstate: Minidealstate{
					Entityinfo: Entityinfo{
						name:  "SGenesisMinState",
						value: 1.0,
					},
				},
				Maxidealstate: Maxidealstate{
					Entityinfo: Entityinfo{
						name:  "SGenesisMaxState",
						value: 1.0,
					},
				},
			},
			Generalstate: Generalstate{
				Entityinfo: Entityinfo{
					name:  "SGenesisGeneralState",
					value: 1.0,
				},
			},
			S: S{
				Ideal: Ideal{
					Minidealstate: Minidealstate{
						Entityinfo: Entityinfo{
							name:  "SMinState",
							value: 1.0,
						},
					},
					Maxidealstate: Maxidealstate{
						Entityinfo: Entityinfo{
							name:  "SMaxState",
							value: 1.0,
						},
					},
				},
			},
			G: G{
				Ideal: Ideal{
					Minidealstate: Minidealstate{
						Entityinfo: Entityinfo{
							name:  "GMinState",
							value: 1.0,
						},
					},
					Maxidealstate: Maxidealstate{
						Entityinfo: Entityinfo{
							name:  "GMaxState",
							value: 1.0,
						},
					},
				},
			},
			K: K{
				Ideal: Ideal{
					Minidealstate: Minidealstate{
						Entityinfo: Entityinfo{
							name:  "KMinState",
							value: 1.0,
						},
					},
					Maxidealstate: Maxidealstate{
						Entityinfo: Entityinfo{
							name:  "KMaxState",
							value: 1.0,
						},
					},
				},
			},
		},
	}
}

func So(entity Entity) {

	printIdealDetails("S", entity.Entitystate.S.Ideal, entity.Entityidentifer.Entityid)
}

func Go(entity Entity) {
	printIdealDetails("G", entity.Entitystate.G.Ideal, entity.Entityidentifer.Entityid)
}

func Ko(entity Entity) {
	printIdealDetails("K", entity.Entitystate.K.Ideal, entity.Entityidentifer.Entityid)
}

func printIdealDetails(label string, ideal Ideal, entityID Entityid) {
	fmt.Printf(
		"%s: entity=%s id=%d min=%s(%.2f) max=%s(%.2f)\n",
		label,
		entityID.name,
		entityID.id,
		ideal.Minidealstate.Entityinfo.name,
		ideal.Minidealstate.Entityinfo.value,
		ideal.Maxidealstate.Entityinfo.name,
		ideal.Maxidealstate.Entityinfo.value,
	)
}
