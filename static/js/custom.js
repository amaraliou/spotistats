$(document).ready(function(){
    $(".top-track-large").hide();
    $(".medium-term.top-tracks").hide();
    $(".long-term.top-tracks").hide();
});

$(window).load(function(){
    $("div.top-track-large").each(function(index) {
        $(this).delay(100*index).fadeIn(200);
    });
})

function shortTerm(){
    $('#long-term').removeClass("active");
    $('#medium-term').removeClass("active");
    $('#short-term').addClass("active");

    $("div.medium-term.top-tracks").hide();
    $("div.long-term.top-tracks").hide();
    $("div.short-term.top-tracks").show();
}

function mediumTerm(){
    $('#long-term').removeClass("active");
    $('#short-term').removeClass("active");
    $('#medium-term').addClass("active");

    $("div.short-term.top-tracks").hide();
    $("div.long-term.top-tracks").hide();
    $("div.medium-term.top-tracks").show();
}

function longTerm(){
    $('#medium-term').removeClass("active");
    $('#short-term').removeClass("active");
    $('#long-term').addClass("active");

    $("div.short-term.top-tracks").hide();
    $("div.medium-term.top-tracks").hide();
    $("div.long-term.top-tracks").show();
}