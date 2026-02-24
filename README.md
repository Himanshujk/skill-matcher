# Skill Matcher

A powerful Go application that searches for skills across Excel files using either exact text matching or semantic similarity with word embeddings.

## Features

- **Interactive TUI**: Clean terminal user interface for easy configuration
- **Dual Search Modes**:
  - **Exact Mode**: Direct text matching for precise skill searches  
  - **Semantic Mode**: AI-powered similarity search using pre-trained word embeddings
- **Concurrent Processing**: Fast parallel processing of multiple Excel files
- **Flexible Input**: Search across any column in Excel files
- **Automated Results**: Exports filtered results to timestamped Excel files

## Prerequisites

- Go 1.25.6 or later
- Internet connection (for initial model download)

## Installation

1. Clone or download this repository
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Build the application:
   ```bash
   go build -o skill-matcher
   ```

## Usage

1. Run the application:
   ```bash
   ./skill-matcher
   ```

2. Follow the interactive prompts to configure your search:
   - **Folder**: Directory containing your Excel files
   - **Column**: Name of the column to search within
   - **Skills**: Comma-separated list of skills to search for
   - **Mode**: Choose between "exact" or "semantic" matching
   - **Threshold** (semantic mode only): Similarity threshold (0.0-1.0)

3. The application will:
   - Process all `.xlsx` files in the specified folder
   - Search the specified column for matching skills
   - Export results to a timestamped Excel file (`results_YYYYMMDD_HHMMSS.xlsx`)

## Search Modes

### Exact Mode
Perfect for precise skill matching. Searches for exact text matches within the specified column.

**Example**: Searching for "Python" will match cells containing exactly "Python"

### Semantic Mode  
Uses pre-trained GloVe word embeddings to find semantically similar skills. Great for finding related or alternative skill descriptions.

**Example**: Searching for "Python" might also match "Python Programming", "Python Development", or other related terms

**Threshold Guidelines**:
- `0.9-1.0`: Very strict, only near-identical matches
- `0.8-0.9`: High similarity, related concepts  
- `0.7-0.8`: Moderate similarity, broader matches
- `0.6-0.7`: Loose similarity, more experimental

## Dependencies

- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)**: Terminal UI framework
- **[Lipgloss](https://github.com/charmbracelet/lipgloss)**: Style definitions for TUI
- **[Excelize](https://github.com/xuri/excelize)**: Excel file processing
- **[Bubbles](https://github.com/charmbracelet/bubbles)**: TUI components

## Model Information

For semantic search, the application automatically downloads and uses:
- **GloVe**: Global Vectors for Word Representation
- **Version**: 2024 WikiGiga 100d
- **Source**: Stanford NLP
- **Size**: ~100 dimensional word vectors

The model is downloaded once and cached locally as `model.vec`.

## File Structure

```
skill-matcher/
├── main.go           # Main application entry point
├── tui.go            # Terminal user interface
├── search.go         # Search algorithms and similarity functions  
├── model.go          # Word embedding model handling
├── downloader.go     # Model download functionality
├── downloader_tui.go # Download progress UI
├── index.go          # Data indexing utilities
├── go.mod           # Go module dependencies
└── README.md        # This file
```

## Example

```bash
$ ./skill-matcher

# Follow prompts:
Folder: ./data/resumes
Column: Skills  
Skills: Python, Machine Learning, Data Science
Mode: semantic
Threshold: 0.8

# Output:
✅ Search Complete! Output written to results_20240224_143052.xlsx (23 matches found)
```

## Performance

- **Concurrent Processing**: Processes multiple Excel files simultaneously
- **Memory Efficient**: Streams data without loading entire files into memory
- **Progress Tracking**: Real-time progress updates during processing

## Troubleshooting

**Model Download Issues**: Ensure you have a stable internet connection. The GloVe model (~400MB) is downloaded once on first semantic search.

**Column Not Found**: Verify the exact column name in your Excel files. Column names are case-sensitive.

**No Results**: Try adjusting the semantic threshold or using exact mode for debugging.

**Memory Usage**: For very large datasets, the application loads word vectors into memory. Ensure adequate RAM for semantic searches.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes  
4. Add tests if applicable
5. Submit a pull request

## License

This project is open source.

---

*Built with ❤️ using Go and modern TUI libraries*