package storage

import (
	"testing"
	"gecs/component"

	"github.com/stretchr/testify/require"
)

func runManagerTestSuite(t *testing.T, storageCreator func(string) Storage, compCreator func(string) component.Component, m Manager) {
	t.Run("AddStorage", func(t *testing.T) {
		runTestManagerAdd(t, storageCreator, compCreator, m)
	})
	t.Run("DelStorage", func(t *testing.T) {
		runTestManagerDel(t, storageCreator, compCreator, m)
	})
	t.Run("GetStorage", func(t *testing.T) {
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
		m.AddStorage(storageCreator("playerTestSimpleAdd"))
	})
	t.Run("multiple add", func(t *testing.T) {
		m.AddStorage(storageCreator("playerTestMultipleAdd1")).
			AddStorage(storageCreator("playerTestMultipleAdd2")).
			AddStorage(storageCreator("playerTestMultipleAdd3")).
			AddStorage(storageCreator("playerTestMultipleAdd4"))
	})
	t.Run("add multiples times same storage", func(t *testing.T) {
		s := storageCreator("playerTestAddSame")
		m.AddStorage(s)
		m.AddStorage(s)
		m.AddStorage(s)
	})
}

func runTestManagerDel(t *testing.T, storageCreator func(string) Storage, compCreator func(string) component.Component, m Manager) {
	t.Run("simple delete", func(t *testing.T) {
		s := storageCreator("playerTestSimpleDelete")
		m.AddStorage(s)
		m.DelStorage(s)
	})
	t.Run("delete multiple times same storage", func(t *testing.T) {
		s := storageCreator("playerTestMultipleDelete")
		m.AddStorage(s)
		m.DelStorage(s)
		m.DelStorage(s)
		m.DelStorage(s)
		m.DelStorage(s)
	})
}

func runTestManagerGet(t *testing.T, storageCreator func(string) Storage, compCreator func(string) component.Component, m Manager) {
	t.Run("simple get", func(t *testing.T) {
		s := storageCreator("playerTestSimpleGet")
		m.AddStorage(s)
		ns := m.GetStorage(compCreator("playerTestSimpleGet"))
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
		m.AddStorage(s)
		require.True(t, m.Exist(c))
	})
	t.Run("not exist after delete", func(t *testing.T) {
		c := compCreator("playerTestNotExistAfterDelete")
		s := storageCreator(c.Type())
		m.AddStorage(s)
		require.True(t, m.Exist(c))
		m.DelStorage(s)
		require.False(t, m.Exist(c))
	})
}

func runTestManagerLen(t *testing.T, storageCreator func(string) Storage, compCreator func(string) component.Component, m Manager) {
	t.Run("simple len", func(t *testing.T) {
		m.AddStorage(storageCreator("playerTestSimpleAdd"))
		require.NotEqual(t, 0, m.Len())
	})
	t.Run("increase len after add", func(t *testing.T) {
		prev := m.Len()
		m.AddStorage(storageCreator("playerTestIncreaseLenAfterAdd"))
		require.Equal(t, prev+1, m.Len())
	})
	t.Run("decrease len after del", func(t *testing.T) {
		prev := m.Len()
		s := storageCreator("playerTestDecreaseLenAfterDel")
		m.AddStorage(s)
		require.Equal(t, prev+1, m.Len())
		m.DelStorage(s)
		require.Equal(t, prev, m.Len())
	})
}
