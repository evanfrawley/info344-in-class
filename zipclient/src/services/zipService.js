import 'whatwg-fetch';

let HOST = "http://localhost";
let ZIP_PATH = "/zips/";

export function getZipsForCityName(cityName) {
    return fetch(`${HOST}${ZIP_PATH}${cityName}`).then((response) => {
        return response.json();
    })
}
