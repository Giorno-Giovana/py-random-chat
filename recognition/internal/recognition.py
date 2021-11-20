import face_recognition
import numpy as np
from PIL import Image
from io import BytesIO
import base64


def decode_base64(data, altchars=b'+/'):
    data = re.sub(rb'[^a-zA-Z0-9%s]+' % altchars, b'', data)  # normalize
    missing_padding = len(data) % 4
    if missing_padding:
        data += b'=' * (4 - missing_padding)
    return base64.b64decode(data, altchars)


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
        with open("image.jpeg", "wb") as fh:
            fh.write(decode_base64(base64image))
            image = face_recognition.load_image_file(fh)
            # image_encoding = face_recognition.face_encodings(image)[0]
            # self.encodings.append(image_encoding)
            # self.usernames.append(username)

    def identify(self, base64image: str) -> str:
        username = self.UNKOWN
        print(base64image)

        im = Image.open(BytesIO(base64.b64decode(base64image)))
        im.save('./upload/temp.jpg')
        face_recognition.load_image_file(im)

        # with open("image.jpeg", "wb") as fh:
        #     print(base64image)

        #     image_64_decode = base64.decodebytes(base64image)
        #     image_result = open('deer_decode.jpg', 'wb')  # create a writable image and write the decoding result
        #     image_result.write(image_64_decode)
        #     face_recognition.load_image_file(image_result)

        # fh.write(decode_base64(base64image))
        # image = face_recognition.load_image_file(fh)

        # # Find all the faces and face encodings in the image
        # face_locations = face_recognition.face_locations(image)
        # face_encodings = face_recognition.face_encodings(image, face_locations)

        # # determine the closest user to located faces
        # min_distance = np.inf
        # for face_encoding in face_encodings:
        #     matches = face_recognition.compare_faces(self.encodings, face_encoding)
        #     face_distances = face_recognition.face_distance(self.encodings, face_encoding)
        #     if np.min(face_distances) < min_distance:
        #         best_match_index = np.argmin(face_distances)
        #         if matches[best_match_index]:
        #             username = self.usernames[best_match_index]

        return username
