<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title></title>
    <meta charset="utf-8" />
    <style>
        html, body { margin: 0; width: 100%; height: 100%; }
        .center-container { position: absolute; top: 0; left: 0; width: 100%; height: 100%; text-align: center; overflow: auto; }
        .center-container:after,
        .center-container:after { content: ''; height: 100%; margin-left: -0.25em; /* To offset spacing. May vary by font */ }
    </style>
    <link rel="stylesheet" type="text/css" href="/static/css/login.css"/>
</head>
<body   data-vide-bg="/static/video/ocean">
<div  class="center-container">

    <div id="login">
        <h1 onclick="hide_log(this)" >Ops</h1>
        <form method="post" action="/login">
            <input type="text"   required="required" placeholder="username" name="username"></input>
            <input type="password" required="required" placeholder="password" name="password"></input>
            <button class="but"  type="submit">登录</button>
        </form>
    </div>
</div>
<script src="/static/js/jquery.min.js"></script>
<script src="/static/js/jquery.vide/jquery.vide.js"></script>
<script>
    function hide_log(obj) {
        $(obj).parent().hide();;
    }
</script>
</body>
</html>