# makedataurl
Make a data-URL from a file, optionally serve a HTML preview

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


<pre>
Usage of makedataurl &lt;filename&gt;
  --out string
    	save data-URL to file
  --serve string
    	serve HTML preview at this HTTP address
</pre>