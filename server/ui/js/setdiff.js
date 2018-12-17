$(document).ready(function() {
    $('#submit').click(function() {
        Run();
    });
    $("#f3").on('keypress', function(e) {
        if (e.keyCode != 13) return;
        Run();
    });
});

function Run() {
    var f1 = $('#f1').val().trim();
    var f2 = $('#f2').val().trim();
    var f3 = $('#f3').val().trim();
    if (f1 == "" || f2 == "" || f3 == "") {
        alert("File1，File2，File3必须全部输入");
        return;
    }
    $('#editor_holder').html("<h4>loading...</h4>");
    $("#visual").html("<h4>loading...</h4>");
    $.ajax({
        url: "/api/setdiff?f1="+encodeURIComponent(f1)+"&f2="+encodeURIComponent(f2)+"&f3="+encodeURIComponent(f3), cache: false,
        success: function(result) {
            $('#editor_holder').jsonview(result);
            visual(result)
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            alert(XMLHttpRequest.responseText);
        }
    });
}

function one(k, v) {
    return "<fieldset><legend class=\"label label-info left\">"+
        k+"</legend>"+v+"</fieldset>";
}

function visual(doc) {
    var html = "";
    for (k in doc) {
        if (k == "text" || doc[k] == "") continue;
        var v = doc[k];
        if (k == "url" || k == "canonical_url") {
            v = "<a href=\""+v+"\">"+v+"</a>";
        } else if (k == "favicon") {
            v = "<img src=\""+v+"\" />";
        } else if (k == "images") {
            var str = "<ol>";
            for (i in v) str += "<li><img src=\""+v[i]+"\" /></li>";
            v = str + "</ol>";
        }
        html += one(k, v);
    }
    $("#visual").html(html);
}
