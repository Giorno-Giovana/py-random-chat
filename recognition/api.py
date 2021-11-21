from flask import Blueprint, Flask, request, Response, jsonify
from internal.recognition import FaceRecognition
from utils import error_response, file_to_bytesio
from werkzeug.exceptions import BadRequest

api = Blueprint("api", __name__)


@api.route("/user/<username>/upload_image", methods=["POST"])
def user_upload_photo(username):
    file = request.files["image"]
    if file:
        FaceRecognition().upload(username, file_to_bytesio(file))
    else:
        return error_response("failed to retrieve image", BadRequest.code)


@api.route("/image/identify", methods=["POST"])
def image_identify():
    file = request.files["image"]
    if file:
        username = FaceRecognition().identify(file_to_bytesio(file))
        response = jsonify(
            {
                "username": username,
            }
        )
        return response, 200
    else:
        return error_response("failed to retrieve image", BadRequest.code)
