package service

var (
	W WhileUrls
)

type WhileUrls struct{}

func (WhileUrls) WhileList(url string) bool {
	flag := false
	wls := []string{
		"/login",
		"/logout",
		"/galogin",
		// test assets
		"/assets/ws",
		"/assets/api",
		"/assets/process/status",
		"/assets/process/update/create",
		"/assets/process/update/list",
		"/assets/list",
		"/assets/upload",
	}

	for i := 0; i < len(wls); i++ {
		if wls[i] == url {
			flag = true
		}
	}

	if !flag {
		return flag
	}

	return true
}

func (WhileUrls) OperateWhileList(url string) bool {
	flag := false
	wls := []string{
		"/perms/list",
		"/role/list",
		"/user/list",
		"/log/list",
		"/role/rolesname",
		"/role/userperms",
		"/user/getinfobyname",
		// test assets
		"/assets/ws",
		"/assets/api",
		"/assets/process/status",
		"/assets/process/update/create",
		"/assets/process/update/list",
		"/assets/list",
		"/assets/upload",
	}

	for i := 0; i < len(wls); i++ {
		if wls[i] == url {
			flag = true
		}
	}

	if !flag {
		return flag
	}

	return true
}
