import face_recognition
import numpy as np


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

    def upload(self, username: str, file) -> None:
        image = face_recognition.load_image_file(file)
        image_encoding = face_recognition.face_encodings(image)[0]
        self.encodings.append(image_encoding)
        self.usernames.append(username)

    def identify(self, file) -> str:
        image = face_recognition.load_image_file(file)

        # Find all the faces and face encodings in the image
        face_locations = face_recognition.face_locations(image)
        face_encodings = face_recognition.face_encodings(image, face_locations)

        # determine the closest user to located faces
        username = self.UNKOWN
        min_distance = np.inf
        for face_encoding in face_encodings:
            matches = face_recognition.compare_faces(self.encodings, face_encoding)
            face_distances = face_recognition.face_distance(self.encodings, face_encoding)
            if np.min(face_distances) < min_distance:
                best_match_index = np.argmin(face_distances)
                if matches[best_match_index]:
                    username = self.usernames[best_match_index]

        return username
