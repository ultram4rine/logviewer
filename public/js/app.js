function send() {
    var wrap = $("#responce");
    var displayResources = $("#display-resources");

    var type;
    var name = $("#name").val();
    var time = $("#time").val();
    var mac = $("#mac").val();

    var macRegExp = /([0-9a-fA-F][0-9a-fA-F]){5}([0-9a-fA-F][0-9a-fA-F])/;

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
        mac = mac.replace(/\.|-|:/g, '');

        if (mac.length == 12) {
            if (!macRegExp.test(mac)) {
                alert("Wrong mac-address!")
            }
        } else {
            alert("Mac-address too long!")
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
                console.log(data);
                var output =
                    "<table><thead><tr><th>Mac</th><th>Message</th><th>Time</th></thead><tbody>";
                for (var i in data) {
                    output +=
                        "<tr><td>" +
                        data[i].Mac +
                        "</td><td>" +
                        data[i].Message +
                        "</td><td>" +
                        data[i].Time +
                        "</td></tr>";
                }
                output += "</tbody></table>";
                displayResources.html(output);
                $("table").addClass("table");
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

    $("#show").click(send);

    $("#type option[value=sw]").attr("selected", "true");
    $("#time").show(0);
    $("#name").show(0);
    $("#mac").hide(0);
});