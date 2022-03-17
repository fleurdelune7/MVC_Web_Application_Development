{{define "base"}}

<!DOCTYPE html>

<html lang="en">

<head>

<meta charset="UTF-8" />

<meta http-equiv="X-UA-Compatible" content="IE=edge" />

<meta name="viewport" content="width=device-width, initial-scale=1.0" />

<title>Document</title>

</head>

<body>

<h1>{{block "title" .}} {{end}}</h1>

<p>

<em>Some very interesting information</em> Lorem ipsum dolor sit amet

consectetur adipisicing elit. Praesentium voluptatibus et deleniti non

perferendis, perspiciatis, consectetur facere rem totam corrupti dicta.

Quo facilis repudiandae fuga officiis nisi quis corrupti dolor.

</p>

<h2>Summary</h2>

{{block "summary" .}} {{end}}

<h2>Content</h2>

{{block "content" .}} {{end}}

<h2>Suggestions</h2>

{{block "suggestions" .}} {{end}}


<h2>Footer</h2>
{{block "footer" .}} {{end}}

</body>

</html>

{{end}}