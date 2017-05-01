var moods = ['happy', 'meh', 'angry'];

function get_counts() {
    $.each(moods, function(i, mood) {
        $.getJSON('/counter/' + mood, function(data) {
            $('#' + mood + '-count').text(data['value']);
        });
    });
}

function on_ready() {
    setInterval(get_counts, 1000);
    $.each(moods, function(i, mood) {
        var img = $('#' + mood);
        img.click(function() {
            $.post('/counter/' + mood);
            var width = img.width()
            img.animate({width: 1.05 * width}, 300)
                .animate({width: width}, 300);
        });
    });
}

$(on_ready);
