package chess_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/chess"
)

func TestCastling_Set(t *testing.T) {
	t.Parallel()
	f := func(c Castling, r CastlingRight) bool {
		c.Set(r)
		return c.HasRight(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCastling_Clear(t *testing.T) {
	t.Parallel()
	f := func(c Castling, r CastlingRight) bool {
		c.Clear(r)
		return !c.HasRight(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCastling_NoRights(t *testing.T) {
	t.Parallel()
	f := func(c Castling) bool {
		c.Clear(FriendOO)
		c.Clear(FriendOOO)
		c.Clear(EnemyOO)
		c.Clear(EnemyOOO)
		return c.NoRights()
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCastling_String(t *testing.T) {
	cases := []struct {
		c    Castling
		want string
	}{
		{Castling(0), "-"},
		{Castling(FriendOO), "K"},
		{Castling(FriendOOO), "Q"},
		{Castling(EnemyOO), "k"},
		{Castling(EnemyOOO), "q"},
		{Castling(FriendOO | FriendOOO), "KQ"},
		{Castling(FriendOO | EnemyOO), "Kk"},
		{Castling(FriendOO | EnemyOOO), "Kq"},
		{Castling(FriendOOO | EnemyOO), "Qk"},
		{Castling(FriendOOO | EnemyOOO), "Qq"},
		{Castling(EnemyOO | EnemyOOO), "kq"},
		{Castling(FriendOO | FriendOOO | EnemyOO), "KQk"},
		{Castling(FriendOO | FriendOOO | EnemyOOO), "KQq"},
		{Castling(FriendOO | EnemyOO | EnemyOOO), "Kkq"},
		{Castling(FriendOOO | EnemyOO | EnemyOOO), "Qkq"},
		{Castling(FriendOO | FriendOOO | EnemyOO | EnemyOOO), "KQkq"},
	}
	for _, c := range cases {
		if got := c.c.String(); got != c.want {
			t.Errorf("0x%x: got %q; want %q", c.c, got, c.want)
		}
	}
}
