# makedataurl
Make a data-URL from a file, optionally serves it

Install:
<pre>
go get github.com/ungerik/makedataurl
</pre>

Testride:
<pre>
makedataurl test-images/melbourne-cbd.jpg
makedataurl test-images/melbourne-cbd.gif --out=url.txt
makedataurl test-images/melbourne-cbd.png --serve=localhost:8080
</pre>


Usage of makedataurl:
<pre>
  --out string
    	save output to file
  --serve string
    	serve output at this address
</pre>