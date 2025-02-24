package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/sourcegraph/sourcegraph/lib/output"

	"github.com/sourcegraph/src-cli/internal/api"
	"github.com/sourcegraph/src-cli/internal/batches/graphql"
	"github.com/sourcegraph/src-cli/internal/batches/service"
	"github.com/sourcegraph/src-cli/internal/batches/ui"
)

func init() {
	usage := `
'src batch repositories' works out the repositories that a batch spec would
apply to.

Usage:

    src batch repositories -f FILE

Examples:

    $ src batch repositories -f batch.spec.yaml

`

	flagSet := flag.NewFlagSet("repositories", flag.ExitOnError)

	var (
		fileFlag = flagSet.String("f", "", "The batch spec file to read.")
		apiFlags = api.NewFlags(flagSet)
	)

	var (
		allowUnsupported bool
		allowIgnored     bool
	)
	flagSet.BoolVar(
		&allowUnsupported, "allow-unsupported", false,
		"Allow unsupported code hosts.",
	)
	flagSet.BoolVar(
		&allowIgnored, "force-override-ignore", false,
		"Do not ignore repositories that have a .batchignore file.",
	)

	handler := func(args []string) error {
		if err := flagSet.Parse(args); err != nil {
			return err
		}

		ctx := context.Background()
		client := cfg.apiClient(apiFlags, flagSet.Output())

		svc := service.New(&service.Opts{
			Client:           client,
			AllowUnsupported: allowUnsupported,
			AllowIgnored:     allowIgnored,
			AllowFiles:       true,
		})

		if err := svc.DetermineFeatureFlags(ctx); err != nil {
			return err
		}

		out := output.NewOutput(flagSet.Output(), output.OutputOpts{Verbose: *verbose})
		spec, _, err := parseBatchSpec(fileFlag, svc)
		if err != nil {
			ui := &ui.TUI{Out: out}
			ui.ParsingBatchSpecFailure(err)
			return err
		}

		queryTmpl, err := parseTemplate(batchRepositoriesTemplate)
		if err != nil {
			return err
		}

		totalTmpl, err := parseTemplate(batchRepositoriesTotalTemplate)
		if err != nil {
			return err
		}

		final := []*graphql.Repository{}
		finalMax := 0
		for _, on := range spec.On {
			repos, _, err := svc.ResolveRepositoriesOn(ctx, &on)
			if err != nil {
				return errors.Wrapf(err, "Resolving %q", on.String())
			}

			max := 0
			for _, repo := range repos {
				if len(repo.Name) > max {
					max = len(repo.Name)
				}

				final = append(final, repo)
			}

			if max > finalMax {
				finalMax = max
			}

			if err := execTemplate(queryTmpl, batchRepositoryTemplateInput{
				Max:                 max,
				Query:               on.String(),
				RepoCount:           len(repos),
				Repos:               repos,
				SourcegraphEndpoint: cfg.Endpoint,
			}); err != nil {
				return err
			}
		}

		return execTemplate(totalTmpl, batchRepositoryTemplateInput{
			RepoCount: len(final),
		})
	}

	batchCommands = append(batchCommands, &command{
		flagSet: flagSet,
		aliases: []string{"repos"},
		handler: handler,
		usageFunc: func() {
			fmt.Fprintf(flag.CommandLine.Output(), "Usage of 'src batch %s':\n", flagSet.Name())
			flagSet.PrintDefaults()
			fmt.Println(usage)
		},
	})
}

const batchRepositoriesTemplate = `
{{- color "logo" -}}✱{{- color "nc" -}}
{{- " " -}}
{{- if eq .RepoCount 0 -}}
    {{- color "warning" -}}
{{- else -}}
    {{- color "success" -}}
{{- end -}}
{{- .RepoCount }} workspace{{ if ne .RepoCount 1 }}s{{ end }}{{- color "nc" -}}
{{- if ne (len .Query) 0 -}}
    {{- " for " -}}{{- color "search-query"}}"{{.Query}}"{{ color "nc" -}}
{{- end -}}
{{- "\n" -}}

{{- range .Repos -}}
    {{- "  "}}{{ color "success" }}{{ padRight .Name $.Max " " }}{{ color "nc" -}}
    {{- if ne (len .Branch.Name) 0 -}}{{ " " }}{{- color "search-branch" -}}{{- .Branch.Name -}}{{ color "nc" -}}{{- end -}}
    {{- color "search-border"}}{{" ("}}{{color "nc" -}}
    {{- color "search-repository"}}{{$.SourcegraphEndpoint}}{{.URL}}{{color "nc" -}}
    {{- color "search-border"}}{{")\n"}}{{color "nc" -}}
{{- end -}}
`

const batchRepositoriesTotalTemplate = `
{{- color "logo" -}}✱{{- color "nc" -}}
{{- " " -}}
{{- if eq .RepoCount 0 -}}
    {{- color "warning" -}}
{{- else -}}
    {{- color "success" -}}
{{- end -}}
{{- .RepoCount }} workspace{{ if ne .RepoCount 1 }}s{{ end }} total
{{- color "nc" -}}
`

type batchRepositoryTemplateInput struct {
	Max                 int
	Query               string
	RepoCount           int
	Repos               []*graphql.Repository
	SourcegraphEndpoint string
}
