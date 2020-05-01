package storage

import (
	"github.com/stretchr/testify/require"
	"testing"
	"gecs/component"
	"gecs/entity"
)

func runStorageTestSuite(t *testing.T, em entity.Manager, compCreator func() component.Component, s Storage) {
	t.Run("Add", func(t *testing.T) {
		runTestStorageAdd(t, em, compCreator, s)
	})
	t.Run("Del", func(t *testing.T) {
		runTestStorageDel(t, em, compCreator, s)
	})
	t.Run("Get", func(t *testing.T) {
		runTestStorageGet(t, em, compCreator, s)
	})
	t.Run("Len", func(t *testing.T) {
		runTestStorageLen(t, em, compCreator, s)
	})
	t.Run("Type", func(t *testing.T) {
		runTestStorageType(t, em, compCreator, s)
	})
}

func runTestStorageAdd(t *testing.T, em entity.Manager, compCreator func() component.Component, s Storage) {
	t.Run("simple add", func(t *testing.T) {
		require.NoError(t, s.Add(em.New(), compCreator()))
	})
	t.Run("add other type return an error", func(t *testing.T) {
		require.Error(t, s.Add(em.New(), component.NewTag("other_types")))
	})
}

func runTestStorageDel(t *testing.T, em entity.Manager, compCreator func() component.Component, s Storage) {
	t.Run("simple del", func(t *testing.T) {
		e := em.New()
		require.NoError(t, s.Add(e, compCreator()))
		require.NoError(t, s.Del(e))
	})
	t.Run("multiple del on same entity", func(t *testing.T) {
		e := em.New()
		require.NoError(t, s.Add(e, compCreator()))
		require.NoError(t, s.Del(e))
		require.NoError(t, s.Del(e))
		require.NoError(t, s.Del(e))
	})
}

func runTestStorageGet(t *testing.T, em entity.Manager, compCreator func() component.Component, s Storage) {
	t.Run("simple get", func(t *testing.T) {
		e := em.New()
		comp := compCreator()
		require.NoError(t, s.Add(e, comp))
		rComp := s.Get(e)
		require.Equal(t, comp, rComp)
	})
	t.Run("get after delete return nil", func(t *testing.T) {
		e := em.New()
		comp := compCreator()
		require.NoError(t, s.Add(e, comp))
		require.NoError(t, s.Del(e))
		rComp := s.Get(e)
		require.Nil(t, rComp)
	})
}

func runTestStorageLen(t *testing.T, em entity.Manager, compCreator func() component.Component, s Storage) {
	t.Run("retrieve len after add", func(t *testing.T) {
		prev := s.Len()
		require.NoError(t, s.Add(em.New(), compCreator()))
		require.Equal(t, prev+1, s.Len())
	})
	t.Run("correct len after del", func(t *testing.T) {
		prev := s.Len()
		e := em.New()
		require.NoError(t, s.Add(e, compCreator()))
		require.Equal(t, prev+1, s.Len())
		require.NoError(t, s.Del(e))
		require.Equal(t, prev, s.Len())
	})
}

func runTestStorageType(t *testing.T, em entity.Manager, compCreator func() component.Component, s Storage) {
	t.Run("check has type", func(t *testing.T) {
		require.NotEmpty(t, s.Type())
	})
	t.Run("check type match component", func(t *testing.T) {
		require.Equal(t, s.Type(), compCreator().Type())
	})
}
