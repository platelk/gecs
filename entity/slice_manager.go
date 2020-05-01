package entity

type SliceManager struct {
	entities []uint32
	garbage  []uint32
}

func NewSliceManager() *SliceManager {
	return &SliceManager{}
}

func (em *SliceManager) New() Entity {
	ne := Entity{}

	if len(em.garbage) > 0 {
		i := em.garbage[0]
		ne.version = em.entities[i-1]
		ne.id = i
		em.garbage = em.garbage[1:]

		return ne
	}

	em.entities = append(em.entities, 1)
	ne.version = 1
	ne.id = uint32(len(em.entities))

	return ne
}

func (em *SliceManager) IsValid(e Entity) bool {
	return em.entities[e.id-1] == e.version
}

func (em *SliceManager) Invalidate(e Entity) {
	if !em.IsValid(e) {
		return
	}
	em.entities[e.id-1]++
	em.garbage = append(em.garbage, e.ID())
}

func (em *SliceManager) Len() int {
	return len(em.entities) - len(em.garbage)
}