package entity

import "github.com/willf/bitset"

type BitSetManager struct {
	max       uint32
	entities  bitset.BitSet
	nbGarbage uint
	garbages  bitset.BitSet
}

func NewBitSetManager() *BitSetManager {
	return &BitSetManager{}
}

func (b *BitSetManager) New() Entity {
	if b.nbGarbage > 0 {
		i, _ := b.garbages.NextSet(0)
		b.entities.Set(i)
		b.garbages.Flip(i)
		b.nbGarbage--
		return Entity{
			id: uint32(i),
		}
	}
	e := Entity{id: b.max}
	b.entities.Set(uint(e.id))
	b.max++
	return e
}

func (b *BitSetManager) IsValid(e Entity) bool {
	return b.entities.Test(uint(e.id))
}

func (b *BitSetManager) Invalidate(e Entity) {
	if b.IsValid(e) {
		b.nbGarbage++
		b.entities.Flip(uint(e.id))
	}
}

func (b *BitSetManager) Len() uint {
	return b.entities.Len() - b.garbages.Len()
}
