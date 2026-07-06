# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`fuzzy-repo-finder` is a Go CLI/TUI tool for navigating locally cloned git repositories. It scans configured root directories for git repos, presents an interactive fuzzy-search terminal UI (termbox), and prints the selected project's path to stdout — designed to be wrapped in a shell function like `cd $(fuzzy-repo-finder --projectRoots ...)`.

## Commands

```sh
make build        # go mod download/tidy/verify + build to bin/fuzzy-repo-finder
make test         # go test -v ./...
make run arg=...  # clean, build, and run with optional args
make clean        # remove bin/, dist/, *.snap
```

Run a single test:

```sh
go test -v ./pkg/proj -run TestName
```

Local run for debugging (bypasses the TUI, prints parsed config and found projects):

```sh
./bin/fuzzy-repo-finder --projectRoots "${HOME}/projects" --debug
```

## Architecture

Entry point is `cmd/fuzzy-repo-finder/main.go`: parse flags → scan disk for projects → run the TUI event loop → print selected path to stdout and exit. Exit code 0 means a project was selected (stdout carries the path); exit code 1 is abnormal exit (Ctrl+C).

Packages under `pkg/`:

- **`config`** — global package-level variables (`Debug`, `ProjectsRoots`, `ProjectNameFilter`, `SelectedProjectPath`) set by `ParseArgsAndFlags()` and read throughout the app. `Version`/`ShortCommit`/`BuildDate` are injected via `-ldflags` at build time (see Makefile and `.goreleaser.yml`). The selected project is communicated back from the TUI to `main` via `config.SelectedProjectPath`, not a return value.
- **`io`** — filesystem scanning. `Disk` (implements the `IDisk` interface) recursively walks a root and stops descending once a directory contains `.git`. `Filesystem` converts those paths into `proj.Project` values, deriving `Group` (path segment(s) between root and project) and `Name`. Tests substitute a fake `IDisk`.
- **`proj`** — the `Project`/`Projects` domain model and fuzzy filtering. `Projects` implements `fuzzy.Source` from `github.com/sahilm/fuzzy`; `FuzzyMatch` filters by project name only.
- **`term`** — termbox-based TUI. `Terminal.Cycle()` is one iteration of the event loop: filter projects by the query, render (`display`), poll a key event, and return an `ExitCode` (`ContinueRunning`/`NormalExit`/`AbnormalExit`). Keys: typing/backspace/Ctrl+W edit the query (`field`), arrows move the selection (`projectselector`), Enter selects, Ctrl+C aborts.

The TUI renders to the terminal via termbox while the final result goes to stdout — keep debug/interactive output separate from the selected-path output, since shell wrappers consume stdout.

## Releases

Tag pushes trigger goreleaser (GitHub release + Homebrew formula in `hoto/homebrew-repo`); see `.github/workflows/release.yml` and `.goreleaser.yml`. CI (`test.yml`) runs `make build`, `make test`, and a goreleaser dry run. Snap packaging is defined in `snapcraft.yaml` with `snap-*` Makefile targets.
