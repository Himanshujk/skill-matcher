# Skill Matcher

A powerful Go application that searches for skills across Excel files using either exact text matching or semantic similarity with word embeddings.

## Features

- **Interactive TUI**: Clean terminal user interface for easy configuration
- **Dual Search Modes**:
  - **Exact Mode**: Direct text matching for precise skill searches  
  - **Exact + Semantic Mode**: Combines exact matching with AI-powered similarity search using pre-trained word embeddings
- **Concurrent Processing**: Fast parallel processing of multiple Excel files
- **Flexible Input**: Search across any column in Excel files
- **Automated Results**: Exports filtered results to timestamped Excel files
- **Skill Extraction & Training**: Built-in tools to extract skills from Excel files and train custom models
- **Smart Skill Normalization**: Comprehensive skill aliases and version handling (e.g., "js" → "javascript", preserves "html5", "css3")

## Prerequisites

- Go 1.25.6 or later
- FastText (for training custom models)

## Installation

1. Clone or download this repository
2. Install FastText (required for custom model training):
   ```bash
   git clone https://github.com/facebookresearch/fastText.git
   cd fastText
   make
   ```
3. Install Go dependencies:
   ```bash
   go mod tidy
   ```
4. Build the application:
   ```bash
   go build -o skill-matcher
   ```

## Usage

### Main Skill Matcher Application

1. Run the application:
   ```bash
   ./skill-matcher
   ```

2. Follow the interactive prompts to configure your search:
   - **Folder**: Directory containing your Excel files
   - **Column**: Name of the column to search within
   - **Skills**: Comma-separated list of skills to search for
   - **Mode**: Choose between "exact" or "exact + semantic" matching
   - **Threshold** (exact + semantic mode only): Similarity threshold (0.0-1.0) for the semantic part

3. The application will:
   - Process all `.xlsx` files in the specified folder
   - Search the specified column for matching skills
   - Export results to a timestamped Excel file (`results_YYYYMMDD_HHMMSS.xlsx`)

### Skill Extractor Tool

For creating custom training data and models from your own Excel files:

1. Navigate to the skill extractor:
   ```bash
   cd skillExtractor
   ```

2. Edit the folder path in `main.go` to point to your Excel files directory

3. Run the skill extractor:
   ```bash
   go run main.go
   ```

4. This will generate `skills_corpus.txt` containing normalized skills from your data

5. To train a custom model (requires FastText):
   ```bash
   cd train
   go run main.go
   ```

   This creates `skill_model.vec` and `skill_model.bin` files for use in the main application.
   **Note**: Model files (*.vec, *.bin, *.txt) are gitignored due to their large size.

## Search Modes

### Exact Mode
Perfect for precise skill matching. Searches for exact text matches within the specified column.

**Example**: Searching for "Python" will match cells containing exactly "Python"

### Exact + Semantic Mode  
Combines the precision of exact matching with the power of semantic similarity. First attempts exact text matching, then falls back to AI-powered similarity search for broader coverage.

**How it works**:
1. **Exact Match**: First tries to find direct text matches (same as Exact Mode)
2. **Semantic Fallback**: If no exact match found, uses custom FastText word embeddings to find semantically similar skills

**Example**: Searching for "Python" will match:
- **Exact matches**: "Python" (direct match)
- **Semantic matches**: "Python Programming", "Python Development", or other related terms (if no exact match found)

**Threshold Guidelines** (for semantic part of exact + semantic mode):
- **Recommended Range**: `0.5-0.8` (best performance for most use cases)
- `0.8-1.0`: Very strict, only near-identical matches
- `0.7-0.8`: High similarity, related concepts  
- `0.6-0.7`: Moderate similarity, broader matches
- `0.5-0.6`: Loose similarity, more experimental
- `Below 0.5`: Very loose, may include many false positives

## Skill Normalization

The application includes intelligent skill normalization to handle common variations and maintain consistency:

### Skill Aliases
Automatically maps common skill variations to standard terms:
- `js`, `node.js`, `node-js` → `nodejs`
- `reactjs`, `react.js` → `react`
- `golang`, `go-lang` → `go`
- `py` → `python`
- `ts` → `typescript`
- `.net`, `c#` → `dotnet`, `csharp`
- And 300+ more mappings...

### Number Preservation
Preserves important version numbers and technical terms:
- `html5`, `css3` (preserved as-is)
- `es6`, `es2015`, `es2020` (preserved as-is)
- `gpt2`, `gpt3`, `gpt4` (AI models)
- `vue2`, `vue3`, `angular2` (framework versions)
- `python3`, `java8`, `java11` (language versions)

### Version Stripping
For skills not in the preserve list, automatically strips version numbers:
- `python3.9` → `python`
- `node16.14` → `nodejs`
- `react18.2` → `react`

## Dependencies

- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)**: Terminal UI framework
- **[Lipgloss](https://github.com/charmbracelet/lipgloss)**: Style definitions for TUI
- **[Excelize](https://github.com/xuri/excelize)**: Excel file processing
- **[Bubbles](https://github.com/charmbracelet/bubbles)**: TUI components

## Model Information

For exact + semantic search, the application uses custom FastText models trained on your specific data:

- **FastText**: Skip-gram model with character n-grams
- **Dimensions**: 150 (configurable)
- **Context Window**: 8 words
- **Character N-grams**: 2-6 characters
- **Training**: 15 epochs with negative sampling
- **Negative Sampling**: 15 samples per positive example
- **Learning Rate**: 0.05

Custom models are stored in the `skillExtractor/train/` directory as `skill_model.vec` and `skill_model.bin` files.

**Note**: Model and corpus files (*.vec, *.bin, *.txt) are not tracked in git due to their large size. You'll need to generate them locally using the skillExtractor tools.

## File Structure

```
skill-matcher/
├── main.go              # Main application entry point
├── tui.go               # Terminal user interface
├── search.go            # Search algorithms and similarity functions  
├── model.go             # Word embedding model handling
├── downloader.go        # Model download functionality
├── downloader_tui.go    # Download progress UI
├── index.go             # Data indexing utilities
├── go.mod               # Go module dependencies
├── helpers/             # Skill normalization utilities
│   ├── skills.go        # Core normalization functions
│   ├── skillsAlias.go   # Skill alias mappings (300+ entries)
│   └── preserveNumbers.go # Version preservation rules
├── skillExtractor/      # Training data extraction tools
│   ├── main.go          # Extract skills from Excel files
│   └── train/           # Model training utilities
│       └── main.go      # FastText training script
└── README.md           # This file
```

## Example

### Basic Search
```bash
$ ./skill-matcher

# Follow prompts:
Folder: ./data/resumes
Column: Skills  
Skills: Python, Machine Learning, Data Science
Mode: exact + semantic
Threshold: 0.8

# Output:
✅ Search Complete! Output written to results_20240224_143052.xlsx (23 matches found)
```

### Skill Extraction & Training
```bash
# Extract skills from your Excel files
$ cd skillExtractor
$ go run main.go
Found 150 Excel files to process
Processing complete! Generated skills_corpus.txt

# Train custom model (optional)
$ cd train
$ go run main.go
# Creates skill_model.vec and skill_model.bin
```

### Skill Normalization Examples
The application automatically normalizes common skill variations:
- Input: "js, reactjs, node.js, typescript"
- Normalized: "javascript, react, nodejs, typescript"
- Input: "python3.9, html5, vue2"  
- Normalized: "python, html5, vue2" (html5 and vue2 preserved)

## Performance

- **Concurrent Processing**: Processes multiple Excel files simultaneously
- **Memory Efficient**: Streams data without loading entire files into memory
- **Progress Tracking**: Real-time progress updates during processing

## Troubleshooting

**Model Not Found**: Ensure you have trained a custom model using the skillExtractor tools. The application requires `skill_model.vec` in the `skillExtractor/train/` directory for semantic search.

**Column Not Found**: Verify the exact column name in your Excel files. Column names are case-sensitive. The application searches for exact matches.

**No Results**: Try adjusting the semantic threshold or using exact mode for debugging. Check that skill names match expected patterns.

**Memory Usage**: For very large datasets, the application loads word vectors into memory for exact + semantic searches. Ensure adequate RAM for semantic searches.

**Custom Model Training**: 
- Requires FastText binary installed and accessible in PATH
- Update the FastText path in `skillExtractor/train/main.go` if needed
- Ensure sufficient training data in `skills_corpus.txt` (minimum 1000 skill instances recommended)

**Skill Normalization Issues**: 
- Check `helpers/skillsAlias.go` for unexpected mappings
- Add new aliases or modify existing ones as needed
- Verify number preservation rules in `helpers/preserveNumbers.go`

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