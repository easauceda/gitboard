package main

import (
	ui "github.com/gizak/termui"
)

func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	// Review Request List
	reviewRequestList := ui.NewList()
	reviewRequestList.Items = getReviewRequests()
	reviewRequestList.Height = ui.TermHeight() / 2
	reviewRequestList.BorderLabel = "Review Requests"

	// Event Handlers
	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		// press q to quit
		ui.StopLoop()
	})

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, reviewRequestList),
		),
	)

	ui.Body.Align()

	// Dashboard Render Loop
	ui.Render(ui.Body)
	ui.Loop()
}

func getReviewRequests() []string {
	//TODO: Implement request to get Review Requests from github
	reviewRequests := []string{
		"[0] platform/this-test-branch ➞ https://github.com/pull/test",
		"[1] editbox.go",
		"[2] interrupt.go",
		"[3] keyboard.go",
		"[4] output.go",
		"[5] random_out.go",
		"[6] dashboard.go",
		"[7] nsf/termbox-go"}

	return reviewRequests
}
