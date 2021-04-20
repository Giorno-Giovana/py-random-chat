from falcon.asgi import App
from backend.resources import ChatResource


app = App()
app.add_route("/ws", ChatResource())
