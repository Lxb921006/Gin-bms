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
		"/assets/ws",
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
		// "/login",
		// "/logout",
		// "/galogin",
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
