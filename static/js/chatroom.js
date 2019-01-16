var msgDiv = document.getElementById('msg');
function connectWebSocket() {
    if ('WebSocket' in window) {
        var ws = new WebSocket('ws://www.yuzuka.tk:8080/app/chatroom');
        ws.onopen = function () {
            var p = new HTMLParagraphElement();
            p.textContent = "Connection succeeded.";
            msgDiv.appendChild(p);
            ws.send('Hello!');
        }
        ws.onmessage = function (e) {
            var msg = e.data;
            var p = new HTMLParagraphElement();
            p.textContent = msg;
            msgDiv.appendChild(p);
        }
        ws.onclose = function () {
            var p = new HTMLParagraphElement();
            p.textContent = "Connection closed.";
            msgDiv.appendChild(p);
        }

    } else {
        alert("WebSocket not supported by browser.")
    }
}