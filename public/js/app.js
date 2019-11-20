function send() {
  let wrap = $("#responce");
  let displayResources = $("#display-resources");

  let type;
  let name = $("#name").val();
  let time = $("#time").val();
  let mac = $("#mac").val();

  let macRegExp = /[0-9a-fA-F]{12}/;

  if (
    $("#type")
      .children("option:selected")
      .val() === "sw"
  ) {
    type = "sw";
  } else if (
    $("#type")
      .children("option:selected")
      .val() === "dhcp"
  ) {
    type = "dhcp";
    mac = mac.toLowerCase();
    mac = mac.replace(/\.|-|:/g, "");

    if (mac.length === 12) {
      if (!macRegExp.test(mac)) {
        alert("Wrong mac-address!");
        return;
      }
    } else {
      alert("Mac-address too long!");
      return;
    }
  }

  if (type === "dhcp") {
    $.ajax({
      type: "GET",
      url: "/get/dhcp",
      data: { type: type, mac: mac, time: time },
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
            data[i].MACStr +
            "</td><td>" +
            data[i].IP +
            "</td><td>" +
            data[i].Message +
            "</td><td>" +
            data[i].TimeStampStr +
            "</td></tr>";
        }
        output += "</tbody></table>";
        displayResources.html(output);
      }
    });
  } else {
    $.ajax({
      type: "GET",
      url: "/get/switch",
      data: { type: type, name: name, time: time },
      dataType: "json",
      statusCode: {
        401: function() {
          window.location.href = "/login";
        }
      },
      success: function(data) {
        var output =
          "<table><thead><tr><th>Time</th><th>Message</th></thead><tr>";
        for (var i in data) {
          output +=
            "<td>" +
            data[i].LogTimeStampStr +
            "</td><td>" +
            data[i].LogMessage +
            "</td></tr>";
        }
        output += "</tbody></table>";
        displayResources.html(output);
      }
    });
  }
}

$(document).ready(function() {
  $("#type").change(function() {
    if (
      $(this)
        .children("option:selected")
        .val() === "sw"
    ) {
      $("#name").show(0);
      $("#mac").hide(0);
    } else if (
      $(this)
        .children("option:selected")
        .val() === "dhcp"
    ) {
      $("#name").hide(0);
      $("#mac").show(0);
    }
  });

  $("#show").click(function() {
    return send();
  });

  $("#type option[value=dhcp]").attr("selected", "true");
  $("#name").hide(0);
  $("#mac").show(0);

  hideOnClickOutside("#similars");

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
          $(".similar").remove();
          $("#similars").show(0);
          var output = "";
          for (var i in data) {
            output +=
              "<div class='similar' name='" +
              data[i].SwName +
              "'>" +
              data[i].SwName +
              " - " +
              data[i].SwIP +
              "</div>";
          }
          $("#similars").html(output);
          $(".similar").click(function() {
            $("#name").val($(this).attr("name"));
            $(".similar").remove();
          });
        }
      });
    } else {
      $.get("/getavailable", function(data) {
        $(".similar").remove();
        $("#similars").show(0);
        var output = "";
        for (var i in data) {
          output +=
            "<div class='similar' name='" +
            data[i].SwName +
            "'>" +
            data[i].SwName +
            " - " +
            data[i].SwIP +
            "</div>";
        }
        $("#similars").html(output);
        $(".similar").click(function() {
          $("#name").val($(this).attr("name"));
          $(".similar").remove();
        });
      });
    }
  });

  $("#name").focus(function() {
    $("#similars").show(0);
  });
});

function hideOnClickOutside(selector) {
  const outsideClickListener = event => {
    $target = $(event.target);

    if (
      !$target.closest(selector).length &&
      $(selector).is(":visible") &&
      JSON.stringify($target) !== JSON.stringify($("#name"))
    ) {
      $(selector).hide();
    }
  };

  document.addEventListener("click", outsideClickListener);
}
