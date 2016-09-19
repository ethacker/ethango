Ethan = Ethan || {};


Ethan.dpd = {
    initPage:function(){
        $("#crimes-table").hide();
        $.get({
            url:"http://www.ethanthacker.com/dpdinfo/crimes",
            success:function (data) {
                data.forEach(function(crime,index){
                    var table = document.getElementById("crimes-table-body");
                    var row = table.insertRow(index);
                    row.insertCell(0).innerHTML = crime["location"];
                    row.insertCell(1).innerHTML = crime["nature_of_call"];
                    row.insertCell(2).innerHTML = crime["priority"];
                    row.insertCell(3).innerHTML = crime["date_time"].substr(11,8);
                    row.insertCell(4).innerHTML = crime["division"];
                    row.insertCell(5).innerHTML = crime["status"];
                    row.insertCell(6).innerHTML = crime["unit_number"];
                });
                $("#crimes-table").css("display","block");
            },
            error:function () {

            }
        });
    }
};
