package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"tflint-ktm-rules/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "KTM Rules",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewKTMEmptyLineRule(),
			},
		},
	})
}
