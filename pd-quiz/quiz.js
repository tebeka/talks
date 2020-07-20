var duration = 30;

function on_hashchange() {
    // No code section
    if ($('article.current div.playground').length == 0) {
	return;
    }

    // Already have a timer running
    if ($('article.current .clock').length != 0) {
	return;
    }

    var v = $('<div class="clock">' + duration + '</div>');
    $('article.current').append(v);
    setTimeout(function() {
	countdown(v, duration);
    }, 1000);
}

function countdown(v, n) {
    n -= 1;
    v.text(n);
    if (n == 0) {
	return;
    }

    setTimeout(function() {
	countdown(v, n);
    }, 1000);
}

function highlight() {
    document.querySelectorAll('pre span').forEach((block) => {
	hljs.highlightBlock(block);
    });
}

$(function() {
    //highlight();
    $(window).on('hashchange', on_hashchange);
});
