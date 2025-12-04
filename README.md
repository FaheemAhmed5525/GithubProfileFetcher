# GitHub Profile Fetcher

A modular CLI tool built with clean architecture in Go to fetch and display GitHub user profiles.

## Architecture Overview
┌─────────────┐ ┌─────────────┐ ┌─────────────┐
│ CLI Layer │ → │ API Layer │ → │ GitHub API │
└─────────────┘ └─────────────┘ └─────────────┘
↓ ↓ ↓
┌─────────────┐ ┌─────────────┐ ┌─────────────┐
│ Formatters │ ← │ Data Models │ ← │ JSON Parse │
└─────────────┘ └─────────────┘ └─────────────┘


## Key Architectural Decisions

1. **Separation of Concerns**: Each layer has a single responsibility
2. **Interface-Based Design**: Easy to swap implementations
3. **Extensible Structure**: Ready for additional features
4. **Error Handling**: Graceful degradation and informative messages

## Usage

```bash
# Basic usage
go run cmd/ghfetcher/main.go --user torvalds

# With repository details
go run cmd/ghfetcher/main.go --user torvalds --repos --format json

# Using GitHub token for higher limits
export GITHUB_TOKEN=your_token
go run cmd/ghfetcher/main.go --user torvalds