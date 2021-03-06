package dance

import (
	"github.com/wieku/danser-go/beatmap"
	"github.com/wieku/danser-go/render"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/wieku/danser-go/rulesets/osu"
	"github.com/wieku/danser-go/bmath/difficulty"
	"github.com/wieku/danser-go/bmath"
)

type PlayerController struct {
	bMap    *beatmap.BeatMap
	cursors []*render.Cursor
	window  *glfw.Window
	ruleset *osu.OsuRuleSet
	lastTime int64
	counter int64
}

func NewPlayerController() Controller {
	return &PlayerController{}
}

func (controller *PlayerController) SetBeatMap(beatMap *beatmap.BeatMap) {
	controller.bMap = beatMap
}

func (controller *PlayerController) InitCursors() {
	controller.cursors = []*render.Cursor{render.NewCursor()}
	controller.window = glfw.GetCurrentContext()
	controller.ruleset = osu.NewOsuRuleset(controller.bMap, controller.cursors, []difficulty.Modifier{difficulty.None})
}

func (controller *PlayerController) Update(time int64, delta float64) {

	controller.bMap.Update(time)

	if controller.window != nil {
		glfw.PollEvents()
		x, y := controller.window.GetCursorPos()
		controller.cursors[0].SetScreenPos(bmath.NewVec2d(x, y))
		controller.cursors[0].LeftButton = controller.window.GetKey(glfw.KeyZ) == glfw.Press
		controller.cursors[0].RightButton = controller.window.GetKey(glfw.KeyX) == glfw.Press
	}

	controller.counter += time-controller.lastTime

	if controller.counter >= 12 {
		controller.cursors[0].LastFrameTime = time-12
		controller.cursors[0].CurrentFrameTime = time
		controller.cursors[0].IsReplayFrame = true
		controller.counter-=12
	} else {
		controller.cursors[0].IsReplayFrame = false
	}

	controller.lastTime = time

	controller.ruleset.Update(time)

	controller.cursors[0].Update(delta)
}

func (controller *PlayerController) GetRuleset() *osu.OsuRuleSet {
	return controller.ruleset
}

func (controller *PlayerController) GetCursors() []*render.Cursor {
	return controller.cursors
}
