$(document).ready(function(){
    $(".top-track-large").hide();
    $(".top-artist-large").hide();
    $("div.medium-term").hide();
    $("div.long-term").hide();
});

$(window).load(function(){
    $("div.top-track-large, div.top-artist-large").each(function(index) {
        $(this).delay(100*index).fadeIn(200);
    });
})

function shortTerm(){
    $('#long-term').removeClass("active");
    $('#medium-term').removeClass("active");
    $('#short-term').addClass("active");

    $("div.medium-term").hide();
    $("div.long-term").hide();
    $("div.short-term").show();
}

function mediumTerm(){
    $('#long-term').removeClass("active");
    $('#short-term').removeClass("active");
    $('#medium-term').addClass("active");

    $("div.short-term").hide();
    $("div.long-term").hide();
    $("div.medium-term").show();
}

function longTerm(){
    $('#medium-term').removeClass("active");
    $('#short-term').removeClass("active");
    $('#long-term').addClass("active");

    $("div.short-term").hide();
    $("div.medium-term").hide();
    $("div.long-term").show();
}