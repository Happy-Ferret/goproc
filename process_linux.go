package process

func nameOf(pid int) string {
	procName := ""

	statusFile, err := openPid(pid, "status")
	if err == nil {
		defer statusFile.Close()
		procName = parseProcName(statusFile)
	}
	
	return procName
}
