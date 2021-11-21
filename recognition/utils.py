import io
from flask import jsonify
from flask import Response


def error_response(message: str, code: int) -> Response:
    response = jsonify({
        'code': code,
        'message': message,
    })
    response.status_code = code
    return response
