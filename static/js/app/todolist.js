function newRequest() {
    return new XMLHttpRequest();
}

window.onload = function () {
    var deletes = document.getElementsByClassName('delete');
    for (var i = 0; i < deletes.length; i++) {
        deletes[i].onclick = function () {
            var li = this.parentElement;
            var req = newRequest();
            req.open('POST', document.location.origin + '/app/todolist/' + this.getAttribute('href'));
            req.onload = function () {
                if (req.status == 200) {
                    li.parentElement.removeChild(li);
                }
            };
            var todo = li.getElementsByTagName('span')[0];
            req.send(todo.innerText);
            return false;
        }
    }
    var moveups = document.getElementsByClassName('moveup');
    for (var i = 0; i < moveups.length; i++) {
        moveups[i].onclick = function () {
            var li = this.parentElement;
            var req = newRequest();
            req.open('POST', document.location.origin + '/app/todolist/' + this.getAttribute('href'));
            req.onload = function () {
                if (req.status == 200) {
                    if (li.previousElementSibling) {
                        li.parentElement.insertBefore(li, li.previousElementSibling);
                    }
                }
            };
            var todo = li.getElementsByTagName('span')[0];
            req.send(todo.innerText);
            return false;
        }
    }
    var movedowns = document.getElementsByClassName('movedown');
    for (var i = 0; i < movedowns.length; i++) {
        movedowns[i].onclick = function () {
            var li = this.parentElement;
            var req = newRequest();
            req.open('POST', document.location.origin + '/app/todolist/' + this.getAttribute('href'));
            req.onload = function () {
                if (req.status == 200) {
                    if (li.nextElementSibling) {
                        li.parentElement.insertBefore(li, li.nextElementSibling.nextElementSibling);
                    }else{
                        li.parentElement.appendChild(li);
                    }
                }
            };
            var todo = li.getElementsByTagName('span')[0];
            req.send(todo.innerText);
            return false;
        }
    }
};