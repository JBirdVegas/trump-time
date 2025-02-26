package main

import (
	"encoding/json"
	"runtime/debug"
	"syscall/js"
	"time"

	"github.com/jbirdvegas/tilltrump.com/internal/timex"
)

const layout = "Mon Jan 02 2006 03:04PM"

type git struct {
	Revision string `json:"revision"`
}

var (
	NoBuildInfoGit = func() string {
		g := git{Revision: "no_build_info"}
		js, _ := json.Marshal(g)
		return string(js)
	}()
	NotFoundGit = func() string {
		g := git{Revision: "not_found"}
		js, _ := json.Marshal(g)
		return string(js)
	}()
	FoundGit = func(revision string) string {
		g := git{Revision: revision}
		js, _ := json.Marshal(g)
		return string(js)
	}
)

var (
	sinceLocalizedDate = js.FuncOf(func(this js.Value, args []js.Value) any {
		return makeResponse(args, timex.DateOfNextPres())
	})
	tillLocalizedDate = js.FuncOf(func(this js.Value, args []js.Value) any {
		return makeResponse(args, timex.DateOfPreviousPres())
	})
	githash = js.FuncOf(func(this js.Value, args []js.Value) any {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			return NoBuildInfoGit
		}

		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return FoundGit(setting.Value)
			}
		}
		return NotFoundGit
	})
)

type jsResponse struct {
	Date string `json:"date"`
}

func main() {
	js.Global().Set("nextPrez", js.FuncOf(func(_ js.Value, _ []js.Value) any {
		return formatter(timex.TillNextPresident())
	}))
	js.Global().Set("previousPrez", js.FuncOf(func(_ js.Value, _ []js.Value) any {
		return formatter(timex.SinceLastPresident())
	}))
	js.Global().Set("dateNextPrez", sinceLocalizedDate)
	js.Global().Set("datePrevPrez", tillLocalizedDate)
	js.Global().Set("gitRevision", githash)
	<-make(chan struct{})
}

func makeResponse(args []js.Value, t time.Time) any {
	response := jsResponse{
		Date: t.In(time.FixedZone(args[0].String(), -1*args[1].Int()*60)).Format(layout),
	}
	switch marshal, err := json.Marshal(response); {
	case err != nil:
		return err.Error()
	default:
		return string(marshal)
	}
}

func formatter(pt timex.PresTracker) string {
	switch marshal, err := json.Marshal(pt); {
	case err != nil:
		return err.Error()
	default:
		return string(marshal)
	}
}
