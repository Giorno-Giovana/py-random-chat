import { RECOGNITION_BASE_URL } from '../constants';
import { RECOGNITION_IDENTIFY_IMAGE, RECOGNITION_UPLOAD_IMAGE } from '../constants';

const recognition = {
    name: 'recognition',
    async uploadBase64Image(username, base64image) {
        let url = RECOGNITION_BASE_URL + RECOGNITION_UPLOAD_IMAGE;

        let data = new FormData();
        data.append('username', username);
        data.append('base64image', base64image);

        await fetch(url, {
            method: "POST",
            body: data,
            headers: {
                "Access-Control-Allow-Origin": "*",
            }
        });
    },
    async identifyBase64Image(base64image) {
        let url = RECOGNITION_BASE_URL + RECOGNITION_IDENTIFY_IMAGE;

        let data = new FormData();
        data.append('base64image', base64image);

        let response = await fetch(url, {
            method: "POST",
            body: data,
            headers: {
                "Access-Control-Allow-Origin": "*",
            }
        });

        if (!response.ok) {
            return "error";
        }

        let json = await response.json();
        return json["username"];
    }
}


export default ({ app }, inject) => {
    inject('recognition', recognition)
}

