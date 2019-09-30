function send() {
  var wrap = $("#responce");
  var displayResources = $("#display-resources");

  var type;
  var name = $("#name").val();
  var time = $("#time").val();
  var mac = $("#mac").val();

  var macRegExp = /[0-9a-fA-F]{12}/;

  if (
    $("#type")
      .children("option:selected")
      .val() == "sw"
  ) {
    type = "sw";
  } else if (
    $("#type")
      .children("option:selected")
      .val() == "dhcp"
  ) {
    type = "dhcp";
    mac = mac.toLowerCase();
    mac = mac.replace(/\.|-|:/g, "");

    if (mac.length == 12) {
      if (!macRegExp.test(mac)) {
        alert("Wrong mac-address!");
        return;
      }
    } else {
      alert("Mac-address too long!");
      return;
    }
  }

  if (type == "dhcp") {
    $.ajax({
      type: "GET",
      url: "/get",
      data: { type: type, mac: mac },
      dataType: "json",
      statusCode: {
        401: function() {
          window.location.href = "/login";
        }
      },
      success: function(data) {
        var output =
          "<table><thead><tr><th>Mac</th><th>IP</th><th>Message</th><th>Time</th></thead><tr>";
        for (var i in data) {
          output +=
            "<td>" +
            data[i].Mac +
            "</td><td>" +
            data[i].IP +
            "</td><td>" +
            data[i].Message +
            "</td><td>" +
            data[i].Time +
            "</td></tr>";
        }
        output += "</tbody></table>";
        displayResources.html(output);
      }
    });
  } else {
    $.ajax({
      type: "GET",
      url: "/get",
      data: { type: type, name: name, time: time, mac: mac },
      dataType: "text",
      statusCode: {
        401: function() {
          window.location.href = "/login";
        }
      },
      success: function(data) {
        wrap.html(data.replace(/\n/g, "<br>"));
      }
    });
  }
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

  $("#show").click(function() {
    return send();
  });

  $("#type option[value=dhcp]").attr("selected", "true");
  $("#time").hide(0);
  $("#name").hide(0);
  $("#mac").show(0);

  $(document).on("input", "#name", function(ev) {
    if ($(ev.target).val().length > 0) {
      $.ajax({
        type: "GET",
        url: "/findsimilar",
        data: { t: $(ev.target).val() },
        dataType: "json",
        statusCode: {
          401: function() {
            window.location.href = "/login";
          }
        },
        success: function(data) {
          $("#similar").remove();
          $("#similars").show(0);
          var output = "";
          for (var i in data) {
            output +=
              "<div id='similar' name='" +
              data[i].SwName +
              "'>" +
              data[i].SwName +
              " - " +
              data[i].SwIP +
              "</div>";
          }
          $("#similars").append(output);
          $("#similar").click(function() {
            $("#name").val($(this).attr("name"));
            $("#similars").hide(0);
          });
        }
      });
    } else {
      $.get("/getavailable", function(data) {
        $("#similar").remove();
        $("#similars").show(0);
        var output = "";
        for (var i in data) {
          output +=
            "<div id='similar' name='" +
            data[i].SwName +
            "'>" +
            data[i].SwName +
            " - " +
            data[i].SwIP +
            "</div>";
        }
        $("#similars").append(output);
        $("#similar").click(function() {
          $("#name").val($(this).attr("name"));
          $("#similars").hide(0);
        });
      });
    }
  });
});
