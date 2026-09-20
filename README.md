# TabOverflow

A command-line tool for links you mean to get back to but never do. Save a URL, and when you have spare time, let it pick one for you.

Written in Go with nothing outside the standard library.

## Commands

| Command | What it does |
| --- | --- |
| `tof add '<url>'` | Save a link |
| `tof list` | Show everything unread |
| `tof pick` | Surface one at random |
| `tof done '<url>'` | Mark a link read |
| `tof rm '<url>'` | Delete a link |

`add` tags each link by domain — `video`, `repo`, `book`, `movie`, `article` — and rejects duplicates. Matching ignores tracking parameters, `www.`, and fragments, so the same page saved twice is caught even when the two URLs don't look identical. `done` and `rm` accept whatever form your browser hands you.

## Requirements

- Go 1.26 or later

## Installation

```bash
git clone https://github.com/dillon-goknit/TabOverflow.git
cd TabOverflow
go build -o ~/.local/bin/tof .
```

Any directory on `PATH` works. After that, `tof` runs from anywhere.

## Notes

Quote your URLs. Without quotes the shell discards everything after an `&`, and the truncated link saves with no error:

```bash
tof add 'https://youtube.com/watch?v=abc&t=50s'
```

Links live in `~/.config/taboverflow/links.json`. Set `TABOVERFLOW_DATA` to point somewhere else.

Single machine only, nothing syncs.
