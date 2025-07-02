package fetch

import (
  "testing"
)

func TestFetchPageStub(t *testing.T) {
  resp, err := FetchPage("bypass-token", "irrelevant")
  if err != nil {
    t.Fatalf("unexpected error: %v", err)
  }
  if len(resp.Items) != 2 {
    t.Errorf("got %d items; want 2", len(resp.Items))
  }
  if resp.NextPage != "" {
    t.Errorf("got nextPage %q; want empty", resp.NextPage)
  }
}
