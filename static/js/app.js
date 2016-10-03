
var Ethan = Ethan || {};

var welcomeback;

Ethan.app = {
    menuToggleVar:variable = false,

    initPage: function () {
        Ethan.app.navbarSizing();
        window.addEventListener("resize",Ethan.app.navbarSizing);
        $(".nav-menu-responsive-icon").click(Ethan.app.menuToggle);
        if($(window).width()<769){
            $(".main-content").click(Ethan.app.menuClose);
        }
    },
    navbarSizing: function() {
        if($(window).width()<769){
            $(".navigation-list").hide();
            Ethan.app.menuToggleVar = false;
            $(".nav-menu-responsive-icon").show();
            $(".social-media-list-item").hide();
        } else {
            $(".navigation-list").show();
            $(".nav-menu-responsive-icon").hide();
        }
    },
    menuClose: function() {
        $(".navigation-list").hide();
        Ethan.app.menuToggleVar = false;
    },
    menuToggle: function() {
        if(Ethan.app.menuToggleVar == false){
            $(".navigation-list").show();
            Ethan.app.menuToggleVar = true;
        } else {
            $(".navigation-list").hide();
            Ethan.app.menuToggleVar = false;
        }
    },
    setCookie: function () {
        if(Cookies.get("returning")=="true"){
            welcomeback = document.createElement("p");
            welcomeback.setAttribute("id","welcome-back-banner");
            welcomeback.innerHTML = "Welcome Back to EthanThacker.com!";
            document.getElementById("welcome-back").appendChild(welcomeback);
        }
        else {
            Cookies.set("returning","true",{ expires: Infinity });
        }
    },

    checkStrings: function () {

        var strings = [$("#firstString").val(),$("#secondString").val()];
        console.log(strings);
        $.post({
            url:"http://localhost:8080/api/strings",
            data: JSON.stringify(strings)
        }).done(function (data) {
            var results = JSON.parse(data);
            document.getElementById("same-holder").innerHTML = results["Same"];
            document.getElementById("execution-time-holder").innerHTML = results["Time"] + " nanoseconds";
            console.log(data);
        });

    }
};