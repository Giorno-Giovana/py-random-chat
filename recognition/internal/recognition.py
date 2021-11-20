import enum
import face_recognition
import numpy as np
from PIL import Image
from io import BytesIO
import base64

from werkzeug.datastructures import ContentSecurityPolicy

TEMP_FILEPATH = './upload/temp.png'


def saveBase64image(base64image: str) -> str:
    im = Image.open(BytesIO(base64.b64decode(base64image)))
    im.save(TEMP_FILEPATH)
    return TEMP_FILEPATH


class SingletonMeta(type):
    _instances = {}

    def __call__(cls, *args, **kwargs):
        if cls not in cls._instances:
            instance = super().__call__(*args, **kwargs)
            cls._instances[cls] = instance
        return cls._instances[cls]


class FaceRecognition(metaclass=SingletonMeta):
    UNKOWN = 'unknown'

    def __init__(self):
        self.usernames = []
        self.encodings = []

    def upload(self, username: str, base64image: str) -> None:
        path = saveBase64image(base64image)
        image = face_recognition.load_image_file(path)
        image_encoding = face_recognition.face_encodings(image)[0]
        self.encodings.append(image_encoding)
        self.usernames.append(username)
        print("added encoding for username {}".format(username))
        print(len(self.usernames), len(self.encodings))

    def identify(self, base64image: str) -> str:
        username = self.UNKOWN

        path = saveBase64image(base64image)
        image = face_recognition.load_image_file(path)

        # Find all the faces and face encodings in the image
        face_locations = face_recognition.face_locations(image)
        face_encodings = face_recognition.face_encodings(image, face_locations)

        # determine the closest user to located faces
        min_distance = np.inf
        for face_encoding in face_encodings:
            face_distances = face_recognition.face_distance(self.encodings, face_encoding)
            for i, face_distance in enumerate(face_distances):
                print(face_distances, min_distance)
                if face_distance < min_distance:
                    username = self.usernames[i]
                    min_distance = face_distance

        return username
