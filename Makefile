all:
	python app/app.py

test:
	python -m unittest discover tests

install:
	pip install -r ./requirements.txt
