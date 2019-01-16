var form=document.getElementById('form');
var msg=document.getElementById('msg');
var content=document.getElementById('content')
form.onsubmit=connectWebSocket;


function connectWebSocket() {
    if ('WebSocket' in window) {
        var ws = new WebSocket('ws://'+document.location.host+'/app/chatroom/ws');
        ws.onopen = function () {
            var p = document.createElement('p');
            var t = document.createTextNode('Connection succeeded.');
            p.appendChild(t);
            content.appendChild(p);
            ws.send(msg.value);
        }
        ws.onmessage = function (e) {
            var p = document.createElement('p');
            var t = document.createTextNode(String(e.data));
            p.appendChild(t);
            content.appendChild(p);
        }
        ws.onclose = function () {
            var p = document.createElement('p');
            var t = document.createTextNode('Connection closed.');
            p.appendChild(t);
            content.appendChild(p);
        }

    } else {
        alert("WebSocket not supported by browser.")
    }
    return false;
}