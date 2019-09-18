function send() {
  var wrap = document.getElementById("responce");

  var name = document.getElementById("name").value;
  var time = document.getElementById("time").value;
  var req = "name=" + name + "&time=" + time;

  var xhr = new XMLHttpRequest();
  xhr.open("POST", "/get");
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.onload = function() {
    if (xhr.status == 200) {
      wrap.innerText = xhr.responseText;
    }
  };
  xhr.send(req);
}
