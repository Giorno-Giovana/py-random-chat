from flask_cors import CORS
from flask import Flask, Response, request, jsonify
from internal.recognition import FaceRecognition
from utils import error_response, file_to_bytesio
from werkzeug.exceptions import BadRequest
from flask_cors import CORS, cross_origin

app = Flask(__name__)
cors = CORS(app)
app.config['CORS_HEADERS'] = 'Content-Type'


@app.route('/')
def index():
    return Response('stub')


@app.route('/upload_image', methods=['POST'])
@cross_origin()
def user_upload_photo():
    file = request.files['image']
    username = request.form['username']
    if file:
        FaceRecognition().upload(username, file)
        return Response('uploaded an image'), 200
    else:
        return error_response('failed to retrieve image', BadRequest.code)


@app.route('/identify_image', methods=['POST'])
@cross_origin()
def image_identify():
    file = request.files['image']
    if file:
        username = FaceRecognition().identify(file)
        response = jsonify({
            'username': username,
        })
        return response, 200
    else:
        return error_response('failed to retrieve image', BadRequest.code)


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', threaded=True, port=2000)
