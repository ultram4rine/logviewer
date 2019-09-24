function send() {
    var wrap = $("#responce");

    var type;
    var name = $("#name").val();
    var time = $("#time").val();
    var mac = $("#mac").val();

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
            if (!macRegExpShort.test(mac)) {
                alert("Wrong mac-address!")
            }
        } else {
            alert("Mac-address too long!")
        }
    }

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