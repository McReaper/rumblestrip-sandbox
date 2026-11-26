package remove_test

import (
	"context"
	"errors"
	"testing"

	"github.com/McReaper/rumblestrip-sandbox/internal/notes/internal/domain"
	"github.com/McReaper/rumblestrip-sandbox/internal/notes/internal/remove"
)

type memStore struct {
	deleted []domain.NoteID
	err     error
}

func (s *memStore) Delete(_ context.Context, id domain.NoteID) error {
	if s.err != nil {
		return s.err
	}
	s.deleted = append(s.deleted, id)
	return nil
}

func TestDeleteDeletesTheNoteAndReturnsItsID(t *testing.T) {
	store := &memStore{}
	out, err := remove.New(store).Run(context.Background(), remove.Args{ID: "buy-milk-20261030t220509"})
	if err != nil {
		t.Fatal(err)
	}
	if out.ID != "buy-milk-20261030t220509" {
		t.Errorf("ID = %q", out.ID)
	}
	if len(store.deleted) != 1 || store.deleted[0] != "buy-milk-20261030t220509" {
		t.Errorf("deleted %q, want the note's id alone", store.deleted)
	}
}

func TestDeleteRefusesABadIDWithoutDeleting(t *testing.T) {
	store := &memStore{}
	_, err := remove.New(store).Run(context.Background(), remove.Args{ID: "../journal"})
	if !errors.Is(err, domain.ErrBadID) {
		t.Fatalf("err = %v, want ErrBadID", err)
	}
	if len(store.deleted) != 0 {
		t.Errorf("deleted %q", store.deleted)
	}
}

func TestDeleteReportsTheStoreFailure(t *testing.T) {
	boom := errors.New("no note has the id groceries-20261030t220509")
	_, err := remove.New(&memStore{err: boom}).Run(context.Background(), remove.Args{ID: "groceries-20261030t220509"})
	if !errors.Is(err, boom) {
		t.Fatalf("err = %v, want it to wrap %v", err, boom)
	}
}
