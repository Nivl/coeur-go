package secret_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Nivl/coeur/secret"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	t.Parallel()

	rawSecret := "thisIsASecret"
	testSecret := secret.NewString(rawSecret)

	t.Run("should return the secret", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, rawSecret, testSecret.Value())
	})

	t.Run("should not leak secret when using an exported String", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, rawSecret, testSecret.Value())
		assert.Equal(t, secret.Redacted, testSecret.String())
		assert.Equal(t, secret.Redacted, testSecret.GoString())

		b, err := json.Marshal(testSecret)
		require.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("%q", secret.Redacted), string(b))
	})

	t.Run("should not leak secret when using an unexported String", func(t *testing.T) {
		t.Parallel()

		s := struct {
			secret secret.String
		}{
			secret: testSecret,
		}

		assert.NotContains(t, fmt.Sprintf("%v", s), rawSecret)
	})

	t.Run("should free the memory", func(t *testing.T) {
		t.Parallel()

		raw := "thisIsASecret"
		sec := secret.NewString(raw)
		require.Equal(t, raw, sec.Value())
		sec.Free()
		require.Empty(t, sec.Value())
	})

	t.Run("should work with an empty secret", func(t *testing.T) {
		t.Parallel()

		emptySecret := secret.NewString("")
		require.Empty(t, emptySecret.Value())
		emptySecret.Free()
		require.Empty(t, emptySecret.Value())
	})
}
