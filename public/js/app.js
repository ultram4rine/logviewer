function send() {
    var wrap = document.getElementById('responce');

    var date = document.getElementById('date').value;
    var ip = document.getElementById('name').value;
    var time = document.getElementById('time').value;
    var req = 'date=' + date + '&name=' + ip + '&time=' + time;

    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/get');
    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    xhr.onload = function() {
        if (xhr.status == 200) {
            wrap.innerText = xhr.responseText
        }
    }
    xhr.send(req);
}