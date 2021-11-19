FROM python:3.8.10

RUN apt update -y && apt install -y cmake

WORKDIR /usr/src/app
COPY ./requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

CMD ["python app.py"]
