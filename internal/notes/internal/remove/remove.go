// WHY: delete_note's slice isn't named delete: that is Go's builtin, which
// the package's name would shadow where it is imported.
package remove

import (
	"context"
	"fmt"

	"github.com/McReaper/rumblestrip-sandbox/internal/notes/internal/domain"
	"github.com/McReaper/rumblestrip-sandbox/internal/platform/op"
)

type Args struct {
	ID string `json:"id" jsonschema:"the id of the note to delete, as add_note returns it and list_notes shows it"`
}

type Out struct {
	ID string `json:"id"`
}

type Deleter interface {
	Delete(ctx context.Context, id domain.NoteID) error
}

type Remove struct {
	store Deleter
}

func New(store Deleter) *Remove {
	return &Remove{store: store}
}

func (r *Remove) Op() op.Op[Args, Out] {
	return op.Op[Args, Out]{
		Name:        "delete_note",
		Description: "Deletes the note with this id for good and returns its id. An id no note has is refused.",
		Effect:      op.Destructive,
		Example:     &Args{ID: "groceries-20261030t220509"},
		Run:         r.Run,
	}
}

func (r *Remove) Run(ctx context.Context, args Args) (Out, error) {
	id, err := domain.ParseNoteID(args.ID)
	if err != nil {
		return Out{}, fmt.Errorf("delete a note: %w", err)
	}
	if err := r.store.Delete(ctx, id); err != nil {
		return Out{}, fmt.Errorf("delete a note: %w", err)
	}
	return Out{ID: string(id)}, nil
}
