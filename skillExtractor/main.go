package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"skillsearch/helpers"
	"strings"
	"sync"

	"github.com/xuri/excelize/v2"
)

func main() {

	fmt.Println("Processing Excel files...")
	files, _ := filepath.Glob(filepath.Join("/Users/himanshu/Downloads/For Greg Marketing", "*.xlsx"))
	fmt.Printf("Found %d Excel files to process\n", len(files))

	// Setup concurrent processing for multiple Excel files
	var wg sync.WaitGroup // Wait group for goroutine synchronization
	var mu sync.Mutex     // Mutex for thread-safe access to shared data
	processedFiles := 0   // Counter for progress tracking

	out, err := os.Create("skills_corpus.txt")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// Process each Excel file in a separate goroutine for better performance
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			defer func() {
				mu.Lock()
				processedFiles++
				fmt.Printf("\rProcessed %d/%d files", processedFiles, len(files))
				mu.Unlock()
			}()

			f, err := excelize.OpenFile(file)
			if err != nil {
				return
			}

			sheet := f.GetSheetName(0)
			rows, _ := f.GetRows(sheet)

			if len(rows) == 0 {
				return
			}

			colIndex := -1
			for i, h := range rows[0] {
				if h == "skills" {
					colIndex = i
					break
				}
			}

			if colIndex == -1 {
				return
			}

			for i, row := range rows {
				if i == 0 {
					continue // skip header
				}
				if colIndex >= len(row) {
					continue
				}
				skills := row[colIndex] // adjust column index
				tokens := []string{}

				for _, s := range strings.Split(skills, ",") {
					n := normalizeSkill(s)
					if n != "" {
						tokens = append(tokens, n)
					}
				}

				if len(tokens) > 0 {
					mu.Lock()
					_, err = out.WriteString(strings.Join(tokens, " ") + "\n")
					if err != nil {
						mu.Unlock()
						return
					}
					mu.Unlock()
				}
			}

		}(file)
	}

	wg.Wait()
	fmt.Printf("\nProcessing complete!\n")
}

func normalizeSkill(skill string) string {
	skill = stripVersion(strings.ToLower(strings.TrimSpace(skill)))

	// skip if skill is just a number or float
	if matched, _ := regexp.MatchString(`^\d+(\.\d+)?$`, skill); matched {
		return ""
	}

	if v, ok := helpers.SkillAlias[skill]; ok {
		return v
	}

	return skill
}

func stripVersion(skill string) string {
	// remove spaces inside weird tokens
	skill = strings.ReplaceAll(skill, " ", "")

	if helpers.PreserveNumbers[skill] {
		// Preserve the skill as-is if it is in the PreserveNumbers map
		return skill
	}

	// extract leading alphabetic part
	re := regexp.MustCompile(`^[a-z+#]+`)
	match := re.FindString(skill)

	if match != "" {
		return match
	}
	// remove special characters
	re = regexp.MustCompile(`[^a-z0-9+#.]`)
	return re.ReplaceAllString(skill, "")
}
