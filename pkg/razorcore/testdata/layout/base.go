// This file is generated by gorazor 1.2.2
// DON'T modified manually
// Should edit source file and re-generate: cases/layout/base.gohtml

package layout

import (
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
	"tpl/admin/helper"
)

// Base generates cases/layout/base.gohtml
func Base(body string, title string, js string) string {
	var _b strings.Builder

	_body := func(_buffer io.StringWriter) {
		_buffer.WriteString(body)
	}

	_title := func(_buffer io.StringWriter) {
		_buffer.WriteString(title)
	}

	_js := func(_buffer io.StringWriter) {
		_buffer.WriteString(js)
	}

	RenderBase(&_b, _body, _title, _js)
	return _b.String()
}

// RenderBase render cases/layout/base.gohtml
func RenderBase(_buffer io.StringWriter, body func(_buffer io.StringWriter), title func(_buffer io.StringWriter), js func(_buffer io.StringWriter)) {

	companyName := "深圳思品科技有限公司"

	// Line: 11
	_buffer.WriteString("\n<!DOCTYPE html>\n<html>\n<head>\n\t<meta charset=\"utf-8\" />\n    <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n\t<link rel=\"stylesheet\" href=\"/css/bootstrap.min.css\">\n\t<link rel=\"stylesheet\" href=\"/css/dashboard.css\">\n    <!-- HTML5 shim and Respond.js IE8 support of HTML5 elements and media queries -->\n    <!--[if lt IE 9]>\n      <script src=\"https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js\"></script>\n      <script src=\"https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js\"></script>\n    <![endif]-->\n\t<title>")
	// Line: 25
	title(_buffer)
	// Line: 25
	_buffer.WriteString("</title>\n</head>\n<body>\n    <div class=\"navbar navbar-inverse navbar-fixed-top\" role=\"navigation\">\n      <div class=\"container-fluid\">\n        <div class=\"navbar-header\">\n          <button type=\"button\" class=\"navbar-toggle\" data-toggle=\"collapse\" data-target=\".navbar-collapse\">\n            <span class=\"sr-only\">Toggle navigation</span>\n            <span class=\"icon-bar\"></span>\n            <span class=\"icon-bar\"></span>\n            <span class=\"icon-bar\"></span>\n          </button>\n          <a class=\"navbar-brand\" href=\"http://wethinkwith.com\">")
	// Line: 37
	_buffer.WriteString(gorazor.HTMLEscape(companyName))
	// Line: 37
	_buffer.WriteString("</a>我们在<a href=\"http://www.v2ex.com/t/109162\">招聘</a>\n        </div>\n        <div class=\"navbar-collapse collapse\">\n          <ul class=\"nav navbar-nav navbar-right\">\n            <li><a href=\"/admin/setting\">设置</a></li>\n            <li><a href=\"/admin/help\">帮助</a></li>\n            <li><a href=\"/admin/logout\">退出</a></li>\n          </ul>\n          <form class=\"navbar-form navbar-right\">\n            <input type=\"text\" class=\"form-control\" placeholder=\"搜索...\">\n          </form>\n        </div>\n      </div>\n    </div>\n\n    <div class=\"container-fluid\">\n      <div class=\"row\">\n        <div class=\"col-sm-3 col-md-2 sidebar\">\n\t\t\t")
	// Line: 55
	_buffer.WriteString((helper.Menu()))
	// Line: 55
	_buffer.WriteString("\n        </div>\n        <div class=\"col-sm-9 col-sm-offset-3 col-md-10 col-md-offset-2 main\">\n          ")
	// Line: 58
	body(_buffer)
	// Line: 58
	_buffer.WriteString("\n        </div>\n      </div>\n    </div>\n    ")
	if js == nil {
		// Line: 63
		_buffer.WriteString("<script src=\"/js/jquery.min.js\"></script>")

		// Line: 64
		_buffer.WriteString("<div id=\"footer\">@Copyright 2019</div>")

	} else {

		// Line: 66
		_buffer.WriteString("<div id=\"footer\">")
		// Line: 66
		js(_buffer)
		// Line: 66
		_buffer.WriteString("</div>")

	}
	// Line: 67
	_buffer.WriteString("\n\t<script src=\"/js/bootstrap.min.js\"></script>\n\t")
	// Line: 69
	js(_buffer)
	// Line: 69
	_buffer.WriteString("\n  </body>\n</html>")

}
