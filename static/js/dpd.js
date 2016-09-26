Ethan = Ethan || {};


Ethan.dpd = {
    geoCoder: variable,
    map: variable,
    locations:[],
    index:variable = 0,
    latlongSW: variable,
    latlongNE: variable,
    initPage:function(){
        $("#crimes-table").hide();
        $.get({
            url:"http://localhost:8080/dpdinfo/crimes",
            success:function (data) {
                data.forEach(function(crime,index){
                    var table = document.getElementById("crimes-table-body");
                    var row = table.insertRow(index);

                    locRow = row.insertCell(0);
                    locRow.setAttribute("class","locationRow crimes-table-cell");
                    locRow.innerHTML = crime["location"];
                    Ethan.dpd.locations.push(crime["location"]);

                    natureRow = row.insertCell(1);
                    natureRow.setAttribute("class","natureRow crimes-table-cell");
                    natureRow.innerHTML = crime["nature_of_call"];

                    priorityRow = row.insertCell(2);
                    priorityRow.setAttribute("class","priorityRow crimes-table-cell");
                    priorityRow.innerHTML = crime["priority"];

                    timeRow = row.insertCell(3);
                    timeRow.setAttribute("class","timeRow crimes-table-cell");
                    timeRow.innerHTML = crime["date_time"].substr(11,8);

                    divisionRow = row.insertCell(4);
                    divisionRow.setAttribute("class","divisionRow crimes-table-cell");
                    divisionRow.innerHTML = crime["division"];

                    statusRow = row.insertCell(5);
                    statusRow.setAttribute("class","statusRow crimes-table-cell");
                    statusRow.innerHTML = crime["status"];

                    unitRow = row.insertCell(6);
                    unitRow.setAttribute("class","unitRow crimes-table-cell");
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
                var marker = new google.maps.Marker({
                    map: Ethan.dpd.map,
                    position: results[0].geometry.location
                });
                console.log(marker);
                if(Ethan.dpd.index<Ethan.dpd.locations.length){
                    Ethan.dpd.delayedGeoCode(Ethan.dpd.locations[++Ethan.dpd.index])
                }
            } else {
                console.log('Geocode was not successful for the following reason: ' + status);
                if(Ethan.dpd.index<Ethan.dpd.locations.length){
                    Ethan.dpd.delayedGeoCode(Ethan.dpd.locations[++Ethan.dpd.index])
                }
            }
        });
    },
    activateGeoCoder:function(){
        $(".geocode-button").hide();
        Ethan.dpd.geoCode(Ethan.dpd.locations[Ethan.dpd.index])
        $(".progress-meter").fadeIn();
    },
    delayedGeoCode: function(address){
        setTimeout(function(){Ethan.dpd.geoCode(address)},750);
        var progress = (Ethan.dpd.index/Ethan.dpd.locations.length)*100;
        console.log(progress);
        if(progress>6){
            $(".progress-filler").width(progress+"%");
        }
        if(progress==100){
            $(".progress-meter").css('background-color','rgb(255, 254, 254)');
            $(".progress-filler").css('background-color','rgb(255, 254, 254)');
            $(".progress-filler").text("GeoCoding Complete");
        }
    },

    loadMap: function(){
        Ethan.dpd.map = new google.maps.Map(document.getElementById('googleMap'), {
                center: {lat: 32.779, lng: -96.798},
                zoom: 10
            });
        Ethan.dpd.geocoder = new google.maps.Geocoder();

        //TODO add in geocoding bounds
        //Ethan.dpd.latlongSW = new google.maps.LatLng(32.584681, -97.005844);
        //Ethan.dpd.latlongNE = new google.maps.LatLng(32.943825, -96.562271);
    }
};