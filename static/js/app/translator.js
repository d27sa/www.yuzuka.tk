function getXMLHttpRequest() {
    var obj = new XMLHttpRequest();
    return obj;
}
var content;
function prepareForm() {
    content = document.getElementById('content');
    var form = document.getElementsByTagName('form')[0];
    var cbs = document.getElementsByName('engine')
    var ta = form.getElementsByTagName('textarea')[0];
    var from, to;
    var selects = form.getElementsByTagName('select');
    for (var i = 0; i < selects.length; i++) {
        if (selects[i].getAttribute("name") == "from") {
            from = selects[i];
        }
        if (selects[i].getAttribute("name") == "to") {
            to = selects[i];
        }
    }
    form.onsubmit = function () {
        var formData = new Array();
        var ok = false;
        for (var i = 0; i < cbs.length; i++) {
            if (cbs[i].checked) {
                ok = true;
                formData[formData.length] = "engine=" + cbs[i].value;
            }
        }
        if (!ok) {
            return false;
        }
        formData[formData.length] = 'from=' + from.value;
        formData[formData.length] = 'to=' + to.value;
        formData[formData.length] = 'text=' + ta.value;
        var xmlHttpRequest = getXMLHttpRequest();
        xmlHttpRequest.overrideMimeType('application/json');
        xmlHttpRequest.open('POST', document.location.origin + '/app/translator/ajax');
        xmlHttpRequest.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
        xmlHttpRequest.send(formData.join('&'));
        xmlHttpRequest.onreadystatechange = function () {
            if (xmlHttpRequest.readyState == 4) {
                var rp = document.getElementsByClassName('result');
                while (rp.length > 0) {
                    rp[rp.length - 1].parentNode.removeChild(rp[rp.length - 1]);
                }
                var json = JSON.parse(xmlHttpRequest.responseText);
                for (var i = 0; i < json.length; i++) {
                    var p = document.createElement('p');
                    p.setAttribute('class', 'result');
                    var s = document.createElement('span');
                    s.setAttribute('class', 'engine');
                    var e = document.createTextNode(json[i].Engine);
                    s.appendChild(e);
                    p.appendChild(s);
                    var t = document.createTextNode(json[i].Text);
                    p.appendChild(t);
                    content.appendChild(p);
                }
            }
        }
        return false;
    }
    ta.focus();
}

window.onload = prepareForm;