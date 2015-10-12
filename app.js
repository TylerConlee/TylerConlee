var express = require('express');
var exphbs  = require('express-handlebars');

var app = express();

app.engine('handlebars', exphbs({defaultLayout: 'main'}));
app.set('view engine', 'handlebars');
app.use(express.static(__dirname + '/public'));
app.get('/', function (req, res) {
    res.render('home', { title: 'Tyler Conlee - Developer' });
});
app.get('/about', function (req, res) {
    res.render('about', { title: 'Tyler Conlee - Developer' });
});
app.get('/portfolio', function (req, res) {
    res.render('portfolio', { title: 'Tyler Conlee - Developer' });
});
app.get('/contact', function (req, res) {
    res.render('contact', { title: 'Tyler Conlee - Developer' });
});

app.listen(3000);
