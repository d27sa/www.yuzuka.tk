var form=document.getElementById('form');
var msg=document.getElementById('msg');
form.onsubmit=connectWebSocket;


function connectWebSocket() {
    if ('WebSocket' in window) {
        var ws = new WebSocket('ws://www.yuzuka.tk/app/chatroom/ws');
        ws.onopen = function () {
            var p = document.createElement('p');
            var t = document.createTextNode('Connection succeeded.');
            p.appendChild(t);
            document.appendChild(p);
            ws.send(msg.getAttribute('value'));
        }
        ws.onmessage = function (e) {
            var p = document.createElement('p');
            var t = document.createTextNode(String(e.data));
            p.appendChild(t);
            document.appendChild(p);
        }
        ws.onclose = function () {
            var p = document.createElement('p');
            var t = document.createTextNode('Connection closed.')
            p.appendChild(t);
            document.appendChild(p);
        }

    } else {
        alert("WebSocket not supported by browser.")
    }
}