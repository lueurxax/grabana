package row

import (
	"github.com/K-Phoen/grabana/graph"
	"github.com/K-Phoen/grabana/table"
	"github.com/K-Phoen/grabana/text"
	"github.com/grafana-tools/sdk"
)

type Option func(row *Row)

type Row struct {
	builder *sdk.Row
}

func New(board *sdk.Board, title string, options ...Option) *Row {
	panel := &Row{builder: board.AddRow(title)}

	for _, opt := range append(defaults(), options...) {
		opt(panel)
	}

	return panel
}

func defaults() []Option {
	return []Option{
		ShowTitle(),
	}
}

// WithGraph adds a "graph" panel in the row.
func WithGraph(title string, options ...graph.Option) Option {
	return func(row *Row) {
		graphPanel := graph.New(title, options...)

		row.builder.Add(graphPanel.Builder)
	}
}

// WithTable adds a "table" panel in the row.
func WithTable(title string, options ...table.Option) Option {
	return func(row *Row) {
		tablePanel := table.New(title, options...)

		row.builder.Add(tablePanel.Builder)
	}
}

// WithText adds a "text" panel in the row.
func WithText(title string, options ...text.Option) Option {
	return func(row *Row) {
		textPanel := text.New(title, options...)

		row.builder.Add(textPanel.Builder)
	}
}

// ShowTitle ensures that the title of the row will be displayed.
func ShowTitle() Option {
	return func(row *Row) {
		row.builder.ShowTitle = true
	}
}

// HideTitle ensures that the title of the row will NOT be displayed.
func HideTitle() Option {
	return func(row *Row) {
		row.builder.ShowTitle = false
	}
}
