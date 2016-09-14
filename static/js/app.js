
var Ethan = Ethan || {};

var welcomeback;

Ethan.app = {
    initPage: function () {
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