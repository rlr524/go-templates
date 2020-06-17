package main

import "fmt"

func main() {
	name := "Rila Fukushima"

	tpl := `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Hello, Rila!</title>
</head>
<body>
<h1>Hello `+ name +`!</h1>
</body>
</html>
`
fmt.Println(tpl)
}
