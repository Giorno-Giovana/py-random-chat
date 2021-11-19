import io
from flask import jsonify
from flask import Response


def file_to_bytesio(file):
    b = io.BytesIO()
    # to prevent errors with rgba
    file = file.convert('RGB')
    file.save(b, "JPEG")
    b.seek(0)
    return b


def error_response(message: str, code: int) -> Response:
    response = jsonify({
        'code': code,
        'message': message,
    })
    response.status_code = code
    return response
