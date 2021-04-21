from falcon.asgi import App
from resources import ChatResource


app = App()
app.add_route("/ws", ChatResource())
