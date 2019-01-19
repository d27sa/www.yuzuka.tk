function getXMLHttpRequest() {
    var obj = new XMLHttpRequest();
    return obj;
}

function prepareForm() {
    var form = document.getElementsByTagName('form')[0];
    var ta = form.getElementsByTagName('textarea')[0];
    var from,to;
    var selects=form.getElementsByTagName('select');
    for (var i=0;i<selects.length;i++){
        if(selects[i].getAttribute("name")=="from"){
            from=selects[i];
        }
        if(selects[i].getAttribute("name")=="to"){
            to=selects[i];
        }
    }
    
    form.onsubmit = function () {
        var xmlHttpRequest = getXMLHttpRequest();
        xmlHttpRequest.open('POST', document.location.origin + '/app/translator/ajax');
        xmlHttpRequest.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
        xmlHttpRequest.send('from='+from.value+'&to='+to.value+'&text=' + ta.value);
        xmlHttpRequest.onreadystatechange = function () {
            if (xmlHttpRequest.readyState == 4) {
                var rp = document.getElementById('result');
                if (rp) {
                    rp.innerText = xmlHttpRequest.responseText;
                } else {
                    var p = document.createElement('p');
                    p.setAttribute('id', 'result');
                    var t = document.createTextNode(xmlHttpRequest.responseText);
                    p.appendChild(t);
                    document.getElementById('content').appendChild(p);
                }
            }
        }
        return false;
    }
    ta.focus();
}

window.onload = prepareForm;