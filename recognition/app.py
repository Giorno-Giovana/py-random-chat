import atexit
from flask_cors import CORS
from utils import error_response
from flask import Flask, Response, request, jsonify
from internal.recognition import FaceRecognition
from werkzeug.exceptions import BadRequest
from flask_cors import CORS, cross_origin
import os

app = Flask(__name__)
cors = CORS(app)
app.config['CORS_HEADERS'] = 'Content-Type'

os.system("python3 -m spacy download en_core_web_sm")
os.system("python3 -m spacy download en")
os.system("python3 -m nltk.downloader words")
os.system("python3 -m nltk.downloader stopwords")


@app.route('/')
def index():
    return Response('stub')


@app.route('/upload_image', methods=['POST'])
@cross_origin()
def upload_image():
    username = request.form['username']
    base64image = request.form['base64image']
    if base64image and username:
        FaceRecognition().upload(username, base64image)
        return Response('uploaded an image'), 200
    else:
        return error_response('failed to retrieve image', BadRequest.code)


@app.route('/identify_image', methods=['POST'])
@cross_origin()
def image_identify():
    base64image = request.form['base64image']
    if base64image:
        username = FaceRecognition().identify(base64image)
        response = jsonify({
            'username': username,
        })
        return response, 200
    else:
        return error_response('failed to retrieve image', BadRequest.code)

# def parse():
#     if request.method == 'POST':
#         file = request.files['file']
#         tmp_name = str(uuid.uuid4()) + ".pdf"
#         file.save(tmp_name)
#         # with open(tmp_name, 'wb') as f:
#         #     f.write(file.content)
#         parsed_resume = ResumeParser(tmp_name).get_extracted_data()
#         os.remove(tmp_name)
#         print(str(parsed_resume))
#         return str(parsed_resume)
#     return '''
#         <!doctype html>
#         <title>Uplopythoad new File</title>
#         <h1>Upload new File</h1>
#         <form method=post enctype=multipart/form-data>
#           <input type=file name=file>
#           <input type=submit value=Upload>
#         </form>
#         '''


def exit_handler():
    FaceRecognition().sync()


atexit.register(exit_handler)

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', threaded=True, port=1337)
