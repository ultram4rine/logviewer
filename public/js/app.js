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

$(document).ready(function() {
  $("#type").change(function() {
    if (
      $(this)
        .children("option:selected")
        .val() == "sw"
    ) {
      $("#time").show(0);
      $("#name").show(0);
      $("#mac").hide(0);
    } else if (
      $(this)
        .children("option:selected")
        .val() == "dhcp"
    ) {
      $("#time").hide(0);
      $("#name").hide(0);
      $("#mac").show(0);
    }
  });

  $("#type option[value=sw]").attr("selected", "true");
  $("#time").show(0);
  $("#name").show(0);
  $("#mac").hide(0);
});
