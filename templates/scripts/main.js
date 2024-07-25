import { initializeSearch } from './search.js';
import { initMap } from './geolocalization.js';

function initializeArtist() {
    document.querySelectorAll(".item").forEach(item => {
        item.querySelector(".header").addEventListener("click", e => {
            item.classList.toggle("active")
        })
    })
}

function loadGoogleMapsScript(locationData) {
    const script = document.createElement('script');
    script.src = 'https://maps.googleapis.com/maps/api/js?key=AIzaSyCY_DORiaW68dDA7JaLmhSzSh0-2QYZQks';
    script.onload = () => initMap(locationData);
    document.body.appendChild(script);
}

document.addEventListener('DOMContentLoaded', () => {
    initializeArtist();
    initializeSearch();
    const locationDataElement = document.getElementById('locationData');
    if (locationDataElement) {
        const locationData = JSON.parse(locationDataElement.textContent);
        loadGoogleMapsScript(locationData);
    } else {
        console.error('Location data not found');
    }
});