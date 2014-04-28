function on_ready() {
    /* FIXME: Form validation */
    $('#add').click(function() { $('form').submit(); });
}

$(on_ready);
