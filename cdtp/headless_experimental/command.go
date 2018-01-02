package headlessExperimental

import "github.com/mkenney/go-chrome/cdtp/runtime"

/*
BeginFrameParams represents HeadlessExperimental.beginFrame parameters.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-beginFrame
*/
type BeginFrameParams struct {
	// Optional. Timestamp of this BeginFrame (milliseconds since epoch). If not
	// set, the current time will be used.
	FrameTime runtime.Timestamp `json:"frameTime,omitempty"`

	// Optional. Deadline of this BeginFrame (milliseconds since epoch). If not
	// set, the deadline will be calculated from the frameTime and interval.
	Deadline runtime.Timestamp `json:"deadline,omitempty"`

	// Optional. The interval between BeginFrames that is reported to the
	// compositor, in milliseconds. Defaults to a 60 frames/second interval,
	// i.e. about 16.666 milliseconds.
	Interval float64 `json:"interval,omitempty"`

	// Optional. If set, a screenshot of the frame will be captured and returned
	// in the response. Otherwise, no screenshot will be captured.
	Screenshot *ScreenshotParams `json:"screenshot,omitempty"`
}

/*
BeginFrameResult represents the result of calls to HeadlessExperimental.beginFrame.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-beginFrame
*/
type BeginFrameResult struct {
	// Whether the BeginFrame resulted in damage and, thus, a new frame was
	// committed to the display.
	HasDamage bool `json:"hasDamage"`

	// Whether the main frame submitted a new display frame in response to this
	// BeginFrame.
	MainFrameContentUpdated bool `json:"mainFrameContentUpdated"`

	// Base64-encoded image data of the screenshot, if one was requested and
	// successfully taken.
	ScreenshotData string `json:"screenshotData"`
}