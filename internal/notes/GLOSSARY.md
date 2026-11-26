# The notes module's words

Each type of the domain, what it means here, and what makes one.

- **Note**: a short text a person or an agent keeps: its title, its body (the text it holds: empty or at most 1000 characters, and `ParseText` refuses a longer one), when it was made, and its id. Made by `New`.
- **Title**: what a note is called: trimmed, 1 to 80 characters, with at least one ASCII letter or digit once compatibility forms are unfolded and accents dropped (`Été` gives `ete`, `①` gives `1`, `日本` gives nothing and is refused), to name the note by. Made by `ParseTitle`.
- **NoteID**: the name a note is stored and found under: its title's slug and the UTC second it was made (`buy-milk-20261030t220509`). A second note with the same title in the same second gets the same id, and adding it is refused. Made by `New`; `ParseNoteID` checks one read back: lowercase letters and digits joined by single hyphens.
