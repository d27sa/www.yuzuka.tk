function shuffle(a) {
    var j, x, i;
    for (i = a.length - 1; i > 0; i--) {
        j = Math.floor(Math.random() * (i + 1));
        x = a[i];
        a[i] = a[j];
        a[j] = x;
    }
    return a;
}

var colors = ['#578fff', '#8c9ffd', '#ff7ea2', '#ffbf43', '#74dde3', 'red'];

window.onload = function () {
    var appsDiv = document.getElementById('applications');
    var appDiv = document.getElementsByClassName("app");
    shuffle(colors);
    for (var i = 0; i < appDiv.length; i++) {
        appDiv[i].style.backgroundColor = colors[i];
    }
    var m = appDiv.length % 3;
    if (m == 0) { return; }
    for (var i = 0; i < 3 - m; i++) {
        var d = document.createElement('div');
        d.setAttribute('class', 'app');
        d.setAttribute('visibility', 'hidden');
        appsDiv.appendChild(d);
    }
}