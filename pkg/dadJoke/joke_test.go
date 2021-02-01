package dadJoke

import (
	"testing"
)

func TestDecodeJoke(t *testing.T) {
  sample := []byte(`{"id": "123abc", "joke": "nothing funny", "status": 200}`)

  t.Run("parses json properly", func(t *testing.T) {
    got_joke, got_ID := DecodeJson(sample)
    want_ID := "123abc"
    want_joke := "nothing funny"

    if got_joke != want_joke {
      t.Errorf("got %q want %q", got_joke, want_joke)
    }

    if got_ID != want_ID {
      t.Errorf("got %q want %q", got_ID, want_ID)
    }
  })

}
