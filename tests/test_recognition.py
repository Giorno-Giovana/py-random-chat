import os
import pathlib
import unittest
from PIL import Image
from app.utils import file_to_bytesio
from app.internal.recognition import FaceRecognition

DATA_PATH = os.path.join(pathlib.Path(__file__).parent.resolve(), 'data')


class TestFaceRecognition(unittest.TestCase):
    def test(self):
        for root, dirs, _ in os.walk(DATA_PATH):
            for dir in dirs:
                path = os.path.join(root, dir)
                image = Image.open(os.path.join(path, os.listdir(path)[0]))
                FaceRecognition().upload(dir, file_to_bytesio(image))

                for _, _, files in os.walk(path):
                    for file in files:
                        image = Image.open(os.path.join(path, file))
                        username = FaceRecognition().identify(file_to_bytesio(image))
                        assert(username == dir or username == FaceRecognition.UNKOWN)


if __name__ == '__main__':
    unittest.main()
