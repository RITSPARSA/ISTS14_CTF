var http = require('http');
var path = require('path'); 
var port = 8080;

var pages = [

    {route: '', output: 'This is not the page you\'re looking for.'},
    {route: 'admin', output: 'Admin page'},
    {route: 'about', output:'There\'s nothing special about you.'},
    {route: 'flag.html', output: 'FLAG: ISTS14-NODEJSMORELIKEBESTJS3245'},
];


//Creating server
http.createServer(function (request, response) {
    var lookup = path.basename(decodeURI(request.url)); //Extracting base name of requested path.
    var UA = request.headers['user-agent'];
    var lookup = path.basename(decodeURI(request.url)) || 'flag.html';

    if (UA != 'SPARSA') //User-agent must be 'SPARSA'.
    {
        response.writeHead(404);
        response.end("You're not what I'm looking for!");
    }
    else
    {
        pages.forEach(function(page) {
            if (page.route === lookup) {
                response.writeHead(200, {'Content-Type':'text/html'});
                response.end(page.output);
            }
        });
        if (!response.finished) {
            response.writeHead(404);
            response.end("Page not found!");
        }
    }
}).listen(port);

