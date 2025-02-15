package keys

import (
	"github.com/algorandfoundation/nodekit/ui/style"
)

func (m ViewModel) View() string {
	table := style.ApplyBorder(m.Width, m.Height, m.BorderColor).Render(m.table.View())
	return style.WithNavigation(
		m.Navigation,
		style.WithControls(
			m.Controls,
			style.WithTitle(
				m.Title,
				table,
			),
		),
	)
}
