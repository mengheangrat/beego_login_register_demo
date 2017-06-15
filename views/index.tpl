<!DOCTYPE html>
<html lang="en-us">
<head>
  <meta charset="utf-8" />
  <meta http-equiv="X-UA-Compatible" content="IE=edge" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <link rel="icon" href="/static/favicon.ico" />
  <!--jquery库-->
  <script src="/static/js/jquery.min.js"></script>
  <!--bootstrap库-->
  <link href="/static/css/bootstrap.min.css" rel="stylesheet" />
  <script src="/static/js/bootstrap/bootstrap.min.js"></script>
  <!--[if lt IE 9]>
  <script src="/static/js/bootstrap/html5shiv.min.js"></script>
  <script src="/static/js/bootstrap/respond.min.js"></script>
  <![endif]-->
  <!--font-awesome字体库-->
  <link href="/static/css/font-awesome.min.css" rel="stylesheet" />
  <!--页面加载进度条-->
  <link href="/static/css/pace/dataurl.css" rel="stylesheet" />
  <script src="/static/js/pace/pace.min.js"></script>
  <!--jquery.hammer手势插件-->
  <script src="/static/js/jquery.hammer/hammer.min.js"></script>
  <script src="/static/js/jquery.hammer/jquery.hammer.js"></script>
  <!--平滑滚动到顶部库-->
  <script src="/static/js/jquery.scrolltopcontrol/scrolltopcontrol.js" type="text/javascript"></script>
  <!--主要写的jquery拓展方法-->
  <script src="/static/js/jquery.extend.js" type="text/javascript"></script>
  <!--主要写的css代码-->
  <link href="/static/css/default.css" rel="stylesheet" type="text/css" />
  <!--主要写的js代码-->
  <script src="/static/js/default.js" type="text/javascript"></script>
</head>
<body>
<nav class="navbar navbar-inverse navbar-fixed-top">
  <div class="container-fluid">
    <div class="navbar-header">
      <button type="button" class="navbar-toggle show pull-left" data-target="sidebar">
        <span class="sr-only">Toggle navigation</span> <span class="icon-bar"></span><span class="icon-bar"></span><span class="icon-bar"></span>
      </button>
      <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar"
              aria-expanded="false" aria-controls="navbar">
        <span class="sr-only">Toggle navigation</span> <span class="icon-bar"></span><span class="icon-bar"></span><span class="icon-bar"></span>
      </button>
      <a class="navbar-brand" href="index.html">运维平台</a>
    </div>
    <div id="navbar" class="collapse navbar-collapse">

      <ul class="nav navbar-nav navbar-right">
        <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button"
             aria-expanded="false"><i class="fa fa-user fa-fw"></i>&nbsp; {{.username}} &nbsp;<span class="caret"></span></a>
          <ul class="dropdown-menu" role="menu">

            <li><a href="/admin/">&nbsp;后台</a></li>
            <li class="divider"></li>

            <li><a href="/logout"><i class="fa fa-sign-out fa-fw"></i>&nbsp;logout</a></li>
          </ul>
        </li>
      </ul>
    </div>
  </div>
</nav>
<div class="container-fluid all">
  <div class="sidebar">
    <ul class="nav">
      <li><a href="/home/">主页面</a></li>
      <li><a href="/cdn/uploadfile/">CDN资源管理</a></li>
      <li class="has-sub">
        <a href="javascript:void(0);"><span>svn管理</span><i class="fa fa-caret-right fa-fw pull-right"></i></a>
        <ul class="sub-menu">
          <li><a href="/svnmgr/commit/">权限申请</a></li>
          <li><a href="/svnmgr/mylist/">我的权限</a></li>
          <li><a href="/svnmgr/approve/">审批</a></li>
        </ul>
      </li>
      <li><a href="/sss/">上线发布</a></li>
      <li><a href="/aaa/">配置管理</a></li>
    </ul>
  </div>
  <div class="maincontent row">
    <!--我是主要内容 start-->
    <ul class="breadcrumb">
      <!-- <li class="active">首页</li>-->
    </ul>
    <div class="col-sm-12">
<!--主要内容
-->
    </div>
    <!--我是主要内容 end-->
  </div>
</div>
<a href="#top" id="goTop"><i class="fa fa-angle-up fa-3x"></i></a>
</body>
</html>