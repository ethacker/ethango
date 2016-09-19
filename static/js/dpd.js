Ethan = Ethan || {};


Ethan.dpd = {
    initPage:function(){
        $.get({
            url:"http://localhost:8080/dpdinfo/crimes",
            success:function (data) {
                data.forEach(function(crime,index){
                    console.log(crime["beat"]);
                })
            }
        });
    }
};