package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"unique"
)

var z = unique.Make("0")

type stone struct {
	inscription unique.Handle[string]
	produces    []*stone
}

type stoneInstance struct {
	of    *stone
	count int
}

func Day11(in io.Reader, totalBlinks int) error {
	stonesRaw, err := io.ReadAll(in)
	if err != nil {
		return err
	}

	stoneInstances := map[unique.Handle[string]]*stone{}
	currentState := map[unique.Handle[string]]*stoneInstance{}

	for _, inscription := range strings.Fields(string(stonesRaw)) {
		sh := unique.Make(inscription)
		newStone := &stone{
			inscription: sh,
		}
		currentState[newStone.inscription] = &stoneInstance{
			of:    newStone,
			count: 1,
		}
	}

	blink := 1

	for ; blink <= totalBlinks; blink++ {
		newState := map[unique.Handle[string]]*stoneInstance{}

		for _, instance := range currentState {
			if instance.of.produces != nil {
				for _, producedStone := range instance.of.produces {
					if newInstance, exists := newState[producedStone.inscription]; exists {
						newInstance.count += instance.count
					} else {
						newState[producedStone.inscription] = &stoneInstance{
							of:    producedStone,
							count: instance.count,
						}
					}
				}
				continue
			}

			next := nextStoneInscriptions(instance.of.inscription)
			for _, inscription := range next {
				var newStone *stone
				if existingInstance, exists := stoneInstances[inscription]; exists {
					newStone = existingInstance
				} else {
					newStone = &stone{
						inscription: inscription,
					}
					stoneInstances[inscription] = newStone
				}
				instance.of.produces = append(instance.of.produces, newStone)

				if ns, exists := newState[newStone.inscription]; exists {
					ns.count += instance.count
				} else {
					newState[newStone.inscription] = &stoneInstance{
						of:    newStone,
						count: instance.count,
					}
				}
			}
		}

		currentState = newState
	}

	// remainingBlinks := totalBlinks - blink

	// var total int
	// for _, stone := range stoneInstances {
	// 	if !stone.looped {
	// 		continue
	// 	}
	// 	produces := len(nextStoneInscriptions(stone.inscription))
	// 	for _, i := range stone.occurrences {
	// 		if (totalBlinks)%(len(stone.eachBlink)+i) == 0 {
	// 			total += produces
	// 		}
	// 		// total += stone.eachBlink[(totalBlinks-i)%len(stone.eachBlink)]
	// 	}
	// }

	total := 0

	for _, s := range currentState {
		total += s.count
	}

	fmt.Println(total)
	return nil
}

func nextStoneInscriptions(inscription unique.Handle[string]) []unique.Handle[string] {
	if inscription == z {
		return []unique.Handle[string]{unique.Make("1")}
	}
	inscriptionValue := inscription.Value()
	inscriptionWidth := len(inscriptionValue)
	if inscriptionWidth%2 == 0 {
		leftPart := unique.Make(inscriptionValue[:inscriptionWidth/2])
		rightPart := strings.TrimLeft(inscriptionValue[inscriptionWidth/2:], "0")
		if rightPart == "" {
			return []unique.Handle[string]{leftPart, unique.Make("0")}
		}

		return []unique.Handle[string]{leftPart, unique.Make(rightPart)}
	}

	newInscriptionValue, err := strconv.ParseInt(inscriptionValue, 10, 64)
	if err != nil {
		panic(err)
	}

	return []unique.Handle[string]{
		unique.Make(strconv.Itoa(int(newInscriptionValue * 2024))),
	}
}
