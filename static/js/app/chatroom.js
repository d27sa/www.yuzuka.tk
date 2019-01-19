var form, msg, content, msgDiv, ws;

function sendMessage() {
    ws.send(msg.value);
    msg.value = '';
    return false;
}

function prepareForm() {
    form = document.createElement('form');
    msg = document.createElement('input');
    msg.setAttribute('type', 'text');
    form.appendChild(msg);
    var submit = document.createElement('input');
    submit.setAttribute('type', 'submit');
    form.appendChild(submit);
    content.appendChild(form);
    form.onsubmit = sendMessage;
}

function prepareMessageDiv() {
    msgDiv = document.createElement('div');
    msgDiv.setAttribute('id', 'messages');
    content.appendChild(msgDiv);
}

function connectWebSocket() {
    if ('WebSocket' in window) {
        prepareMessageDiv();
        prepareForm();
        ws = new WebSocket(document.location.protocol == 'http' ? 'ws://' : 'wss://' + document.location.host + '/app/chatroom/ws');
        ws.onopen = function () {
            var p = document.createElement('p');
            var t = document.createTextNode('Connection succeeded.');
            p.appendChild(t);
            msgDiv.appendChild(p);
            msgDiv.scrollTop = msgDiv.scrollHeight;
        }
        ws.onmessage = function (e) {
            var p = document.createElement('p');
            var t = document.createTextNode(String(e.data));
            p.appendChild(t);
            msgDiv.appendChild(p);
            msgDiv.scrollTop = msgDiv.scrollHeight;
        }
        ws.onclose = function () {
            var p = document.createElement('p');
            var t = document.createTextNode('Connection closed.');
            p.appendChild(t);
            msgDiv.appendChild(p);
            msgDiv.scrollTop = msgDiv.scrollHeight;
        }
    } else {
        var p = document.createElement('p');
        var t = document.createTextNode('Sorry, your browser dosen\'t support websocket.');
        p.appendChild(t);
        content.appendChild(p);
    }
}

window.onload = function () {
    content = document.getElementById('content');
    connectWebSocket();
}
