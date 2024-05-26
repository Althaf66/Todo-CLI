package main

import "github.com/charmbracelet/bubbles/list"

// Provides the mock data to fill the kanban board

func (b *Board) initLists() {
	b.cols = []column{
		newColumn(todo),
		newColumn(inProgress),
		newColumn(done),
	}
	// Init To Do
	b.cols[todo].list.Title = "To Do"
	b.cols[todo].list.SetItems([]list.Item{
		Task{status: todo, title: "Learning technology", description: "information about AI"},
		Task{status: todo, title: "Eating BreakFast", description: "Burger,Apple"},
		Task{status: todo, title: "Bath", description: "Getting Ready"},
	})
	// Init in progress
	b.cols[inProgress].list.Title = "In Progress"
	b.cols[inProgress].list.SetItems([]list.Item{
		Task{status: inProgress, title: "Coding", description: "Completing the Backlog"},
	})
	// Init done
	b.cols[done].list.Title = "Done"
	b.cols[done].list.SetItems([]list.Item{
		Task{status: done, title: "Drink Tea", description: "one cup of Tea"},
	})
}
