package handler

func RegisterDefaultActions(handler *IrEventHandler) {
	handler.RegisterAction(IrActionVolume, Volume)
	handler.RegisterAction(IrActionMute, Mute)
	handler.RegisterAction(IrActionMouseClick, MouseClick)
	handler.RegisterAction(IrActionMouseMove, MouseMove)
	handler.RegisterAction(IrActionMouseScroll, MouseScroll)

	handler.RegisterAction(IrActionRunCommand, RunCommand)

	handler.RegisterAction(IrActionPlayPause, PlayPause)
	handler.RegisterAction(IrActionNextTrack, NextTrack)
	handler.RegisterAction(IrActionPreviousTrack, PreviousTrack)

	handler.RegisterAction(IrActionKeyboardShortcut, KeyboardShortcut)
}
