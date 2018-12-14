$(document).ready(function() {
    $('#submit').click(function() {
        Run(query, file);
    });
    $("#query").on('keypress', function(e) {
        if (e.keyCode != 13) return;
        Run(query, file);
    });
});

function Run() {
    var query = $('#query').val().trim();
    var file = $('#file').val().trim();
    if (query == "" || file == "") {
        alert("请输入Query和FileName");
        return;
    }
    $('#editor_holder').html("<h4>loading...</h4>");
    $("#visual").html("<h4>loading...</h4>");
    $.ajax({
        url: "/api/?query="+encodeURIComponent(query)+"&file="+encodeURIComponent(file), cache: false,
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
