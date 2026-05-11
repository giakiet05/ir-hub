package serial

type IrMessage struct {
	RawCode  string `json:"raw_code"`
	Address  uint16 `json:"address"`
	Command  uint16 `json:"command"`
	IsRepeat bool   `json:"is_repeat"`
}
