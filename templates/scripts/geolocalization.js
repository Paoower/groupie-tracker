let map;
let markers = [];

export function initMap(locationData) {
    map = new google.maps.Map(document.getElementById('map'), {
        center: { lat: 0, lng: 0 },
        zoom: 2,
        mapTypeId: google.maps.MapTypeId.HYBRID,
        mapTypeControl: true,
        mapTypeControlOptions: {
            style: google.maps.MapTypeControlStyle.DROPDOWN_MENU,
            mapTypeIds: [
                google.maps.MapTypeId.ROADMAP,
                google.maps.MapTypeId.SATELLITE,
                google.maps.MapTypeId.HYBRID,
                google.maps.MapTypeId.TERRAIN
            ]
        }
    });

    loadLocations(locationData);
}

function loadLocations(locationData) {
    if (Object.keys(locationData).length > 0) {
        const bounds = new google.maps.LatLngBounds();
        const geocoder = new google.maps.Geocoder();
        let markerPositions = [];

        for (const [location, dates] of Object.entries(locationData)) {
            geocoder.geocode({ address: location }, (results, status) => {
                if (status === 'OK') {
                    const marker = new google.maps.Marker({
                        map: map,
                        position: results[0].geometry.location,
                        title: location
                    });

                    markers.push(marker);
                    markerPositions.push({
                        position: marker.getPosition(),
                        dates: dates
                    });
                    bounds.extend(marker.getPosition());

                    const infoWindowContent = `
                        <div style="background-color: #1e293b; color: white; font-family: Arial, sans-serif; padding: 10px; border-radius: 8px;">
                            <h3 style="font-size: 1.2em; font-weight: bold;">${location}</h3>
                            <br>
                            <h4 style="font-size: 1.2em; font-weight: bold;">Dates:</h4>
                            <br>
                            <ul>
                                ${dates.map(date => `<li>${date}</li>`).join('')}
                            </ul>
                        </div>
                    `;

                    const infoWindow = new google.maps.InfoWindow({
                        content: infoWindowContent
                    });

                    marker.addListener('click', () => {
                        infoWindow.open(map, marker);
                    });

                    if (markers.length === Object.keys(locationData).length) {
                        map.fitBounds(bounds);
                        drawLines(markerPositions);
                    }
                } else {
                    console.error('Geocode was not successful for the following reason: ' + status);
                }
            });
        }
    } else {
        console.error('No location data found for this artist');
    }
}

function drawLines(markerPositions) {
    markerPositions.sort((a, b) => {
        const parseDate = (dateStr) => {
            const [day, month, year] = dateStr.split('-');
            return new Date(`${year}-${month}-${day}`);
        };
        const aDate = parseDate(a.dates[0]);
        const bDate = parseDate(b.dates[0]);
        return aDate - bDate;
    });

    const pathCoordinates = markerPositions.map(pos => pos.position);
    const concertPath = new google.maps.Polyline({
        path: pathCoordinates,
        geodesic: true,
        strokeColor: '#FF0000',
        strokeOpacity: 1.0,
        strokeWeight: 2
    });

    concertPath.setMap(map);
}