import { RECOGNITION_BASE_URL } from '../constants';
import { RECOGNITION_IDENTIFY_IMAGE, RECOGNITION_UPLOAD_IMAGE } from '../constants';

const recognition = {
    name: 'recognition',
    async uploadBase64Image(username, base64image) {
        var url = RECOGNITION_BASE_URL + RECOGNITION_UPLOAD_IMAGE;

        var data = new FormData();
        data.append('username', username);
        data.append('base64image', base64image);

        try {
            fetch(url, {
                method: "POST",
                body: data,
                headers: {
                    "Access-Control-Allow-Origin": "*",
                }
            });
        } catch (error) {
            console.error(error.message);
        }
    },
    async identifyBase64Image(base64image) {
        var url = RECOGNITION_BASE_URL + RECOGNITION_IDENTIFY_IMAGE;

        var data = new FormData();
        data.append('base64image', base64image);

        var username;
        await fetch(url, {
            method: "POST",
            body: data,
            headers: {
                "Access-Control-Allow-Origin": "*",
            }
        }).then((response) => {
            response.json().then((json) => {
                username = json["username"];
                console.log(username);
            });
        }).catch((error) => {
            console.error(error);
        });

        return username;
    }
}


export default ({ app }, inject) => {
    inject('recognition', recognition)
}

