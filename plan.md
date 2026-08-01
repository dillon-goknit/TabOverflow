# TabOverflow — Build Guide

A CLI tool to save links (articles, repos, videos, shows, tools) and surface
something to consume later.

## v1 Commands

- `add <url> ["optional title"]` — add a link, unread; guess type from domain
- `pick` — surface one unread link at random (weighted by age)
- `list` — show the pile (filter by type / status)
- `done <id>` — mark consumed, but keep it
- `rm <id>` — delete for good
- `import <file>` — bulk-load URLs from tabs.txt

## The Link

```
ID       identifier
URL      the link
Title    optional (type it manually for now)
Type     article, repo, video, movie, show, tool...
Added    timestamp
Status   unread or consumed
```

Store everything as a JSON file. Load on start, save on any change.

## Build Order

Get it running early, then grow it. Don't start a step until the one before it works.

1. Struct + JSON load/save — hardcode one link, get `list` to print it
2. `add` — real input, guess type, save
3. `done` and `rm`
4. `pick` — start with plain random, then add age-weighting
5. Filters — by type and status
6. `import` — read tabs.txt, loop your `add` logic over the 30 links

## Decisions To Make Yourself

- **IDs** — list position (simple, but shifts on delete) or a stable counter?
- **Type** — free text, or a fixed set?
- **Weighting** — how much should age tilt the random pick?

## Later (after the HTTP course)

- Auto-fetch titles (fixes blank descriptions)
- `stale` — show unread older than X
- GitHub / YouTube metadata
- Episode tracking for shows

## Handy stdlib

- `net/url` — pull the domain out of a URL for type-guessing
- `math/rand` — the random pick
