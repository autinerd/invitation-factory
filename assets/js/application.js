require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
$(() => {
    if (function(){if (self.innerWidth) {
        return self.innerWidth;
     }
     else if (document.documentElement && document.documentElement.clientHeight){
         return document.documentElement.clientWidth;
     }
     else if (document.body) {
         return document.body.clientWidth;
     }
     return 0;}() < 770) {
         var hammertime = new Hammer(document.getElementsByTagName("body")[0]);
    hammertime.on("swiperight", function (ev) {
        $('#sidebar').addClass('active');
    });
    hammertime.on("swipeleft", function (ev) {
        $('#sidebar').removeClass('active');
    });
    $('#sidebarCollapse').on('click', function () {
        $('#sidebar').toggleClass('active');
    });
     }
    
});
