<!DOCTYPE html>
<html lang="zh-cn">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{.page_title}}</title>

    <!-- Bootstrap -->
    <link href="static/css/bootstrap.min.css" rel="stylesheet">

    <link href="static/css/site.css" rel="stylesheet">
  </head>
  <body>
    <h1>你好，世界！</h1>

    {{.LayoutContent}}

    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <script src="static/js/jquery-1.11.1.min.js"></script>
    <!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="static/js/bootstrap.min.js"></script>
    <script src="static/js/site.js"></script>
    {{.Scripts}}
  </body>
</html>
