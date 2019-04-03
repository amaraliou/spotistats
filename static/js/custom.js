$(document).ready(function(){
    $(".top-track-large").hide();
    $(".top-artist-large").hide();
    $(".top-track").hide();
    $(".top-artist").hide();
    $("div.medium-term").hide();
    $("div.long-term").hide();
    $("a#short-term").click()
});

$(window).load(function(){
    $("div.top-track-large.recent-track, div.top-track, div.top-artist").each(function(index) {
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
    $("div.top-track-large.short, div.top-artist-large.short").each(function(index) {
        $(this).delay(100*index).fadeIn(200);
    });
}

function mediumTerm(){
    $('#long-term').removeClass("active");
    $('#short-term').removeClass("active");
    $('#medium-term').addClass("active");

    $("div.short-term").hide();
    $("div.long-term").hide();
    $("div.medium-term").show();
    $("div.top-track-large.medium, div.top-artist-large.medium").each(function(index) {
        $(this).delay(100*index).fadeIn(200);
    });
}

function longTerm(){
    $('#medium-term').removeClass("active");
    $('#short-term').removeClass("active");
    $('#long-term').addClass("active");

    $("div.short-term").hide();
    $("div.medium-term").hide();
    $("div.long-term").show();
    $("div.top-track-large.long, div.top-artist-large.long").each(function(index) {
        $(this).delay(100*index).fadeIn(200);
    });
}

function collapseBoi(){
    //$(".collapsed-navbar").toggle();
    if($(".collapsed-navbar").css('display') == 'none'){
        $(".float").css("transform","rotate(45deg)");
        $(".collapsed-navbar").fadeIn(200);
    } else {
        $(".float").css("transform","");
        $(".collapsed-navbar").fadeOut(200);
    }
}

