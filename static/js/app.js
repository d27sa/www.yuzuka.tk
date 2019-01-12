function randomColor(t, c) {
    var used = [];
    for (var i = 0; i < t.length; i++) {
        var j;
        do {
            j = Math.floor(Math.random() * colors.length);
        } while (used.includes(j, 0));
        used.push(j)
        t[i].style.backgroundColor = c[j];
    }
}

var ds = document.getElementsByClassName("app")
var colors = ['#578fff', '#8c9ffd', '#ff7ea2', '#ffbf43', '#74dde3', 'red']
randomColor(ds, colors);