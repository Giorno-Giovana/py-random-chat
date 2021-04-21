#!/usr/bin/env sh

python -m uvicorn --host 0.0.0.0 --port 8000 --reload app:app
