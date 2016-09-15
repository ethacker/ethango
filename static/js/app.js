
var Ethan = Ethan || {};

var welcomeback;

Ethan.app = {
    initPage: function () {
      window.addEventListener("resize",Ethan.app.navbarSizing);
    },
    navbarSizing: function() {
        if($(window).width()<769){
            $(".navigation-list").hide();
        } else {
            $(".navigation-list").show();
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
    }
};