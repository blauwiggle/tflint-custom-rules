package rules

import (
	"github.com/hashicorp/hcl/v2"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// KTMEmptyLineRule checks whether ...
type KTMEmptyLineRule struct {
	tflint.DefaultRule
}

// NewKTMEmptyLineRule returns a new rule
func NewKTMEmptyLineRule() *KTMEmptyLineRule {
	return &KTMEmptyLineRule{}
}

// Name returns the rule name
func (r *KTMEmptyLineRule) Name() string {
	return "ktm_readability_rule"
}

// Enabled returns whether the rule is enabled by default
func (r *KTMEmptyLineRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *KTMEmptyLineRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *KTMEmptyLineRule) Link() string {
	return "https://ktm-ag.atlassian.net/l/cp/YatVa5eK"
}

// Check checks whether ...
func (r *KTMEmptyLineRule) Check(runner tflint.Runner) error {
	files, err := runner.GetFiles()
	if err != nil {
		return err
	}

	for filename, file := range files {
		content := string(file.Bytes)
		lines := strings.Split(content, "\n")

		for i, line := range lines {
			trimmed := strings.TrimSpace(line)
			if trimmed == "" {
				continue
			}

			if strings.HasPrefix(trimmed, "resource") || strings.HasPrefix(trimmed, "module") || strings.HasPrefix(trimmed, "variable") {
				// Check the line before the block start
				if i > 0 && strings.TrimSpace(lines[i-1]) != "" {
					runner.EmitIssue(
						r,
						"Block should have an empty line before the block.",
						hcl.Range{
							Filename: filename,
							Start:    hcl.Pos{Line: i, Column: 1},
							End:      hcl.Pos{Line: i, Column: len(lines[i])},
						},
					)
				}

				// Find the end of the block
				openBraces := 0
				for j := i; j < len(lines); j++ {
					openBraces += strings.Count(lines[j], "{")
					openBraces -= strings.Count(lines[j], "}")

					if openBraces == 0 {
						// Check the line after the block end
						if j+1 < len(lines) && strings.TrimSpace(lines[j+1]) != "" {
							runner.EmitIssue(
								r,
								"Block should have an empty line after the block.",
								hcl.Range{
									Filename: filename,
									Start:    hcl.Pos{Line: j + 2, Column: 1},
									End:      hcl.Pos{Line: j + 2, Column: len(lines[j+1])},
								},
							)
						}
						break
					}
				}
			}
		}
	}

	return nil
}
