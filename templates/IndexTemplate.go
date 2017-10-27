package templates

import (
	"html/template"
)
var(
	IndexTemplate = template.Must(template.New("index").Parse(`
	<style>
		a {color:blue};
	</style>
	<div>
		<h2>A HTTP test server for clients</h2>
	</div>
	<div>
	<h3>ENDPOINTS</h3>
		<ul>
		<li><a href = "/">/</a>  Returns home page.</li>
		<li><a href = "/ip">/ip</a>  Returns origin ip.</li>
		<li><a href = "/uuid">/uuid</a>  Returns UUID.</li>		
		<li><a href = "/user-agent">/user-agent</a>  Returns user-agent.</li>		
		<li><a href = "/headers">/headers</a>  Return headers map.</li>
		<li><a href = "/get">/get</a> Returns GET data.</li>
		<li><b>/post</b> Returns POST data.</li>
		<li><b>/put</b> Returns PUT data.</li>		
		<li><b>/delete</b> Returns DELETE data.</li>				
		<li><a href = "/anything">/anything</a>  Returns request data, including method used.</li>
		<li><b>/anything/:anything</b>  Returns request data, including the URL.</li>
		<li><a href = "/encoding/utf8">/encoding/utf8</a>  Returns page containing UTF-8 data.</li>	
		<li><a href = "/gzip">/gzip</a> Returns gzip-encoded data.</li>
		<li><a href = "/deflate">/deflate</a> Returns deflate-encoded data.</li>
		<li><a href = "/brolti">/brotli</a> Returns brotli-encoded data.</li>
		<li><a href = "/status/">/status:code</a> Returns given HTTP Status Code.</li>
		<li><a href = "/response-headers">/response-headers?key=value</a> Returns given response headers.</li>
		<li><a href = "/redirect/">/redirect/:n </a> 302 Redirects n times.</li>
		<li><a href = "/redirect-to">/redirect-to?url=foo</a> 302 Redirects to the foo URL.</li>
		<li><a href = "/redirect-to">/redirect-to?url=foo&status_code=307</a> 307 Redirects to the foo URL.</li>
		<li><a href = "/cookies">/cookies</a> Returns cookie data.</li>
		<li><a href = "/cookies/">cookies/set/?name=value</a> Sets one or more simple cookies.</li>
		<li><a href = "/cookies/">/cookies/delete?name</a> Deletes one or more simple cookies.</li>
		
		
		</ul>
	</div>
	`))
)