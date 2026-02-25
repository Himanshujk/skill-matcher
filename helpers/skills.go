package helpers

import (
	"regexp"
	"strings"
)

func NormalizeSkill(skill string) string {
	skill = StripVersion(strings.ToLower(strings.TrimSpace(skill)))

	// skip if skill is just a number or float
	if matched, _ := regexp.MatchString(`^\d+(\.\d+)?$`, skill); matched {
		return ""
	}

	if v, ok := SkillAlias[skill]; ok {
		return v
	}

	return skill
}

func StripVersion(skill string) string {
	// remove spaces inside weird tokens
	skill = strings.ReplaceAll(skill, " ", "")

	if PreserveNumbers[skill] {
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
