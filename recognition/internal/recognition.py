from os import path
import pickle
import base64
import numpy as np
from PIL import Image
from io import BytesIO
import face_recognition

STORAGE_PATH = './storage/'
TEMP_PATH = STORAGE_PATH + 'temp.png'
USERNAMES_PATH = STORAGE_PATH + 'usernames'
ENCODINGS_PATH = STORAGE_PATH + 'encodings'


def saveBase64image(base64image: str) -> str:
    im = Image.open(BytesIO(base64.b64decode(base64image)))
    im.save(TEMP_PATH)
    return TEMP_PATH


class SingletonMeta(type):
    _instances = {}

    def __call__(cls, *args, **kwargs):
        if cls not in cls._instances:
            instance = super().__call__(*args, **kwargs)
            cls._instances[cls] = instance
        return cls._instances[cls]


class FaceRecognition(metaclass=SingletonMeta):
    UNKOWN = 'unknownn'

    def __init__(self):
        self.usernames = []
        self.encodings = []

        if path.isfile(USERNAMES_PATH):
            with open(USERNAMES_PATH, "rb") as fp:
                self.usernames = pickle.load(fp)

        if path.isfile(ENCODINGS_PATH):
            with open(ENCODINGS_PATH, "rb") as fp:
                self.encodings = pickle.load(fp)

    def upload(self, username: str, base64image: str) -> None:
        path = saveBase64image(base64image)
        image = face_recognition.load_image_file(path)
        image_encoding = face_recognition.face_encodings(image)[0]
        self.encodings.append(image_encoding)
        self.usernames.append(username)

    def identify(self, base64image: str) -> str:
        username = self.UNKOWN

        path = saveBase64image(base64image)
        image = face_recognition.load_image_file(path)

        # find all the faces and face encodings in the image
        face_locations = face_recognition.face_locations(image)
        face_encodings = face_recognition.face_encodings(image, face_locations)

        # determine the closest user to located faces
        min_distance = np.inf
        for face_encoding in face_encodings:
            face_distances = face_recognition.face_distance(self.encodings, face_encoding)
            for i, face_distance in enumerate(face_distances):
                if face_distance < min_distance:
                    username = self.usernames[i]
                    min_distance = face_distance

        return username

    def sync(self) -> None:
        with open(USERNAMES_PATH, "wb") as fp:
            pickle.dump(self.usernames, fp)

        with open(ENCODINGS_PATH, "wb") as fp:
            pickle.dump(self.encodings, fp)
