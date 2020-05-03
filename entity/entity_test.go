package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func runEntityTestSuite(t *testing.T, m Manager) {
	t.Run("New", func(t *testing.T) {
		runTestNewManager(t, m)
	})
	t.Run("Invalidate", func(t *testing.T) {
		runTestInvalidateManager(t, m)
	})
	t.Run("IsValid", func(t *testing.T) {
		runTestIsValidManager(t, m)
	})
	t.Run("Len", func(t *testing.T) {
		runTestLenManager(t, m)
	})
}

func runTestLenManager(t *testing.T, m Manager) {
	t.Run("Len after one creation", func(t *testing.T) {
		prev := m.Len()
		m.New()
		require.Equal(t, prev+1, m.Len())
	})
	t.Run("Len after multiple creation", func(t *testing.T) {
		prev := m.Len()
		m.New()
		m.New()
		m.New()
		m.New()
		require.Equal(t, prev+4, m.Len())
	})
	t.Run("Len after invalidate", func(t *testing.T) {
		prev := m.Len()
		e := m.New()
		m.Invalidate(e)
		require.Equal(t, prev, m.Len())
		m.New()
		require.Equal(t, prev+1, m.Len())
	})
}

func runTestInvalidateManager(t *testing.T, m Manager) {
	t.Run("invalidate after creation", func(t *testing.T) {
		e := m.New()
		require.True(t, m.IsValid(e))
		m.Invalidate(e)
		require.False(t, m.IsValid(e))
	})
	t.Run("invalidate won't impact if already invalidated", func(t *testing.T) {
		e := m.New()
		m.Invalidate(e)
		prev := m.Len()
		m.Invalidate(e)
		require.Equal(t, prev, m.Len())
	})
	t.Run("invalidated isn't valid after other creation", func(t *testing.T) {
		e := m.New()
		m.Invalidate(e)
		ne := m.New()
		require.False(t, m.IsValid(e))
		require.True(t, m.IsValid(ne))
	})
}

func runTestIsValidManager(t *testing.T, m Manager) {
	t.Run("isValid", func(t *testing.T) {
		e := m.New()
		require.True(t, m.IsValid(e))
		m.Invalidate(e)
		require.False(t, m.IsValid(e))
	})
}

func runTestNewManager(t *testing.T, m Manager) {
	t.Run("simple new", func(t *testing.T) {
		e := m.New()
		require.NotEqual(t, e.ID(), 0)
	})
	t.Run("multiple new generate different ID", func(t *testing.T) {
		e1 := m.New()
		e2 := m.New()
		require.NotEqual(t, e1.ID(), e2.ID())
	})
}
