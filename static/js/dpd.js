Ethan = Ethan || {};


Ethan.dpd = {
    geoCoder: variable,
    map: variable,
    initPage:function(){
        $("#crimes-table").hide();
        $.get({
            url:"http://www.ethanthacker.com/dpdinfo/crimes",
            success:function (data) {
                data.forEach(function(crime,index){
                    var table = document.getElementById("crimes-table-body");
                    var row = table.insertRow(index);

                    locRow = row.insertCell(0);
                    locRow.setAttribute("class","locationRow");
                    locRow.innerHTML = crime["location"];

                    Ethan.dpd.geoCode(crime["location"]);
                    natureRow = row.insertCell(1);
                    natureRow.setAttribute("class","natureRow");
                    natureRow.innerHTML = crime["nature_of_call"];

                    priorityRow = row.insertCell(2);
                    priorityRow.setAttribute("class","priorityRow");
                    priorityRow.innerHTML = crime["priority"];

                    timeRow = row.insertCell(3);
                    timeRow.setAttribute("class","timeRow");
                    timeRow.innerHTML = crime["date_time"].substr(11,8);

                    divisionRow = row.insertCell(4);
                    divisionRow.setAttribute("class","divisionRow");
                    divisionRow.innerHTML = crime["division"];

                    statusRow = row.insertCell(5);
                    statusRow.setAttribute("class","statusRow");
                    statusRow.innerHTML = crime["status"];

                    unitRow = row.insertCell(6);
                    unitRow.setAttribute("class","unitRow");
                    unitRow.innerHTML = crime["unit_number"];
                });
                Ethan.dpd.loadMap();
                $("#crimes-table").css("display","block");
            },
            error:function (errorCode) {
                console.log(errorCode);
            }
        });
    },
    geoCode: function(address){
        Ethan.dpd.geocoder.geocode( { 'address': address}, function(results, status) {
            if (status == 'OK') {
                console.log(results);
                var marker = new google.maps.Marker({
                    map: Ethan.dpd.map,
                    position: results[0].geometry.location
                });
            } else {
                alert('Geocode was not successful for the following reason: ' + status);
            }
        });
    },

    loadMap: function(){
        Ethan.dpd.map = new google.maps.Map(document.getElementById('googleMap'), {
                center: {lat: 32.779, lng: -96.798},
                zoom: 10
            });
        Ethan.dpd.geocoder = new google.maps.Geocoder();

    }
};