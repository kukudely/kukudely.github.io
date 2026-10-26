package solution

import (
	"leetcodepractice/internal/testutil"
	"reflect"
	"strings"
	"testing"
)

func eq(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got=%v; want=%v", got, want)
	}
}
func TestExamplesAndEdges(t *testing.T) {
	defer testutil.SkipTODO(t)
	t.Run("official_sequence", func(t *testing.T) {
		defer testutil.SkipTODO(t)
		trie := Constructor()
		trie.Insert("apple")
		eq(t, trie.Search("apple"), true)
		eq(t, trie.Search("app"), false)
		eq(t, trie.StartsWith("app"), true)
		trie.Insert("app")
		eq(t, trie.Search("app"), true)
	})
	t.Run("branches_duplicates_missing", func(t *testing.T) {
		defer testutil.SkipTODO(t)
		trie := Constructor()
		eq(t, trie.Search("a"), false)
		trie.Insert("abc")
		trie.Insert("abd")
		trie.Insert("abc")
		trie.Insert("z")
		eq(t, trie.Search("abc"), true)
		eq(t, trie.Search("ab"), false)
		eq(t, trie.StartsWith("ab"), true)
		eq(t, trie.Search("abe"), false)
		eq(t, trie.StartsWith("za"), false)
		eq(t, trie.Search("z"), true)
	})
	t.Run("long_word", func(t *testing.T) {
		defer testutil.SkipTODO(t)
		trie := Constructor()
		word := strings.Repeat("a", 2000)
		trie.Insert(word)
		eq(t, trie.Search(word), true)
		eq(t, trie.Search(word[:1999]), false)
		eq(t, trie.StartsWith(word[:1999]), true)
	})
}
