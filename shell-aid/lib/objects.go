package objects

type SystemInfo struct {
	OS    string
	Shell string
}

type GenerateCommandRequest struct {
	Command string
	SysInfo SystemInfo
}

type GenerateCommandResponse struct {
	Workflow    string `json:"Workflow"`
	SideEffects string `json:"SideEffects"`
	Command     string `json:"Command"`
}
