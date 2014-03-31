package templates

import ()

func IndexTemplate() string {
	return indexTemplate
}

const indexTemplate = `
<html>
<head>
<title>Index page</title>
	<div style="text-align: center;margin-top:50px;">
		<form action="/login" method="POST">
			Username: <input value="" name="username" type="text"> <br>
			Password: <input value="" name="password" type="password"> <br>
			<input value="submit" type="submit"> <br>
		</for>
	<div>
</body>
</html>
`
