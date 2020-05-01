package storage

import (
	"github.com/stretchr/testify/require"
	"testing"
	"gecs/component"
)

func runManagerTestSuite(t *testing.T, storageCreator func(string) Storage, compCreator func(string) component.Component, m Manager) {
	t.Run("Add", func(t *testing.T) {
		runTestManagerAdd(t, storageCreator, compCreator, m)
	})
	t.Run("Del", func(t *testing.T) {
		runTestManagerDel(t, storageCreator, compCreator, m)
	})
	t.Run("Get", func(t *testing.T) {
		runTestManagerGet(t, storageCreator, compCreator, m)
	})
	t.Run("Exist", func(t *testing.T) {
		runTestManagerExist(t, storageCreator, compCreator, m)
	})
	t.Run("Len", func(t *testing.T) {
		runTestManagerLen(t, storageCreator, compCreator, m)
	})
}

func runTestManagerAdd(t *testing.T, storageCreator func(string) Storage, compCreator func(string) component.Component, m Manager) {
	t.Run("simple add", func(t *testing.T) {
		m.Add(storageCreator("playerTestSimpleAdd"))
	})
	t.Run("multiple add", func(t *testing.T) {
		m.Add(storageCreator("playerTestMultipleAdd1")).
			Add(storageCreator("playerTestMultipleAdd2")).
			Add(storageCreator("playerTestMultipleAdd3")).
			Add(storageCreator("playerTestMultipleAdd4"))
	})
	t.Run("add multiples times same storage", func(t *testing.T) {
		s := storageCreator("playerTestAddSame")
		m.Add(s)
		m.Add(s)
		m.Add(s)
	})
}

func runTestManagerDel(t *testing.T, storageCreator func(string) Storage, compCreator func(string) component.Component, m Manager) {
	t.Run("simple delete", func(t *testing.T) {
		s := storageCreator("playerTestSimpleDelete")
		m.Add(s)
		m.Del(s)
	})
	t.Run("delete multiple times same storage", func(t *testing.T) {
		s := storageCreator("playerTestMultipleDelete")
		m.Add(s)
		m.Del(s)
		m.Del(s)
		m.Del(s)
		m.Del(s)
	})
}

func runTestManagerGet(t *testing.T, storageCreator func(string) Storage, compCreator func(string) component.Component, m Manager) {
	t.Run("simple get", func(t *testing.T) {
		s := storageCreator("playerTestSimpleGet")
		m.Add(s)
		ns := m.Get(compCreator("playerTestSimpleGet"))
		require.Equal(t, s, ns)
		require.Equal(t, s.Type(), ns.Type())
		require.Equal(t, s.Len(), ns.Len())
	})
}

func runTestManagerExist(t *testing.T, storageCreator func(string) Storage, compCreator func(string) component.Component, m Manager) {
	t.Run("simple not exist", func(t *testing.T) {
		require.False(t, m.Exist(compCreator("unknown")))
	})
	t.Run("simple exist", func(t *testing.T) {
		c := compCreator("playerTestSimpleExist")
		require.False(t, m.Exist(c))
		s := storageCreator(c.Type())
		m.Add(s)
		require.True(t, m.Exist(c))
	})
	t.Run("not exist after delete", func(t *testing.T) {
		c := compCreator("playerTestNotExistAfterDelete")
		s := storageCreator(c.Type())
		m.Add(s)
		require.True(t, m.Exist(c))
		m.Del(s)
		require.False(t, m.Exist(c))
	})
}

func runTestManagerLen(t *testing.T, storageCreator func(string) Storage, compCreator func(string) component.Component, m Manager) {
	t.Run("simple len", func(t *testing.T) {
		m.Add(storageCreator("playerTestSimpleAdd"))
		require.NotEqual(t, 0, m.Len())
	})
	t.Run("increase len after add", func(t *testing.T) {
		prev := m.Len()
		m.Add(storageCreator("playerTestIncreaseLenAfterAdd"))
		require.Equal(t, prev+1, m.Len())
	})
	t.Run("decrease len after del", func(t *testing.T) {
		prev := m.Len()
		s := storageCreator("playerTestDecreaseLenAfterDel")
		m.Add(s)
		require.Equal(t, prev+1, m.Len())
		m.Del(s)
		require.Equal(t, prev, m.Len())
	})
}
