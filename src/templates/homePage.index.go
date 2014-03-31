package templates

import (

)

func HomeTemplate() string {
	return homeTemplate
}

const homeTemplate = `
<html>
<head>
<title>Home page</title>
	<div style="text-align: center;margin-top:50px;">
		Welcome {{.}} !
	</div>
</body>
</html>
`