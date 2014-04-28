function on_top10(data) {
    var ul = $('#top10');
    $.each(data.rows, function(i, row) {
        var li = $('<li>').addClass('list-group-item');
        li.append($('<a>').attr('href', '/u/' + row['short']).text(row['short']));
        li.append($('<span>').addClass('badge').text(row['hits']));
        ul.append(li);
    });
    $('#loading').hide();
}

function on_ready() {
    /* FIXME: Form validation */
    $('#add').click(function() { $('form').submit(); });

    $.getJSON('/api/v1/top/10', on_top10);
}

$(on_ready);
