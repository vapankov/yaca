package file_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vapankov/yaca/internal/domains/messaging/infrastructure/storage/file"
)

func TestLineStorage(t *testing.T) {
	t.Parallel()

	t.Run("read empty", func(t *testing.T) {
		t.Parallel()

		storage := file.NewLineStore(t.Name())

		lines, err := storage.Read()
		require.NoError(t, err)
		require.Empty(t, lines)
	})

	t.Run("insert and read", func(t *testing.T) {
		t.Parallel()

		fileName := strings.ReplaceAll(t.Name(), "/", "_")
		storage := file.NewLineStore(fileName)

		t.Cleanup(func() {
			if err := os.Remove(fileName); err != nil {
				t.Logf("failed to close file: %s", err.Error())
			}
		})

		line := "hello world"

		err := storage.Insert(line)
		require.NoError(t, err)

		lines, err := storage.Read()
		require.NoError(t, err)
		require.Len(t, lines, 1)
		require.Equal(t, line, lines[0])
	})

	t.Run("insert and read interleaved", func(t *testing.T) {
		t.Parallel()

		fileName := strings.ReplaceAll(t.Name(), "/", "_")
		storage := file.NewLineStore(fileName)

		t.Cleanup(func() {
			if err := os.Remove(fileName); err != nil {
				t.Logf("failed to close file: %s", err.Error())
			}
		})

		expectedLines := []string{
			"hello",
			"world",
			"!",
		}

		require.NoError(t, storage.Insert(expectedLines[0]))

		actualLines, err := storage.Read()
		require.NoError(t, err)
		require.Len(t, actualLines, 1)
		require.Equal(t, expectedLines[0], actualLines[0])

		require.NoError(t, storage.Insert(expectedLines[1]))

		actualLines, err = storage.Read()
		require.NoError(t, err)
		require.Len(t, actualLines, 2)
		require.Equal(t, expectedLines[0], actualLines[0])
		require.Equal(t, expectedLines[1], actualLines[1])

		require.NoError(t, storage.Insert(expectedLines[2]))

		actualLines, err = storage.Read()
		require.NoError(t, err)
		require.Len(t, actualLines, 3)
		require.Equal(t, expectedLines[0], actualLines[0])
		require.Equal(t, expectedLines[1], actualLines[1])
		require.Equal(t, expectedLines[2], actualLines[2])
	})
}
