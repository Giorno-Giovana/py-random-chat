import asyncio
import falcon

from typing import Tuple

from falcon import Request, WebSocketDisconnected
from falcon.asgi import WebSocket

from resources import messages


class ChatResource:
    def __init__(self):
        self.waiting_queue = asyncio.Queue()

    async def on_get(self, req, resp):
        """Handles GET requests"""
        resp.status = falcon.HTTP_200  # This is the default status
        resp.content_type = falcon.MEDIA_TEXT  # Default is JSON, so override
        resp.text = ('\nTwo things awe me most, the starry sky '
                     'above me and the moral law within me.\n'
                     '\n'
                     '    ~ Immanuel Kant\n\n')

    async def on_websocket(self, req: Request, ws: WebSocket):
        try:
            await ws.accept()
        except WebSocketDisconnected:
            return

        try:
            read_task, write_task = await self.new_match(ws)
        except WebSocketDisconnected:
            return

        while not read_task.done():
            await write_task

            read_task.cancel()

            try:
                read_task, write_task = await self.new_match(ws)
            except WebSocketDisconnected:
                return

    async def new_match(self, ws: WebSocket) -> Tuple[asyncio.Task, asyncio.Task]:
        """
        finds a match for given user, returns ws reader and ws writer tasks
        :param ws: websocket
        :return: reader task and writer task
        """
        in_feed = asyncio.Queue()
        try:
            out_feed = await self.find_match(ws, in_feed)
        except WebSocketDisconnected as e:
            raise e

        read_task = falcon.create_task(self.read_socket(ws, out_feed))
        write_task = falcon.create_task(self.run_conversation(ws, in_feed))

        return read_task, write_task

    async def find_match(self, ws: WebSocket, mq: asyncio.Queue) -> asyncio.Queue:
        """
        establish connection establishes connection between two users
        :param mq: message queue (clients message queue)
        :return: interlocutors message queue
        """
        if self.waiting_queue.empty():
            await self.waiting_queue.put(mq)

            while ws.ready and not self.waiting_queue.empty():
                await asyncio.sleep(0)

            # left self from waiting queue if disconnected
            if not ws.ready:
                await self.waiting_queue.get()
                raise WebSocketDisconnected

            await mq.put(messages.connected)
            return await mq.get()

        inter_feed = await self.waiting_queue.get()
        await inter_feed.put(mq)

        await mq.put(messages.connected)

        return inter_feed

    @staticmethod
    async def read_socket(ws: WebSocket, out: asyncio.Queue):
        """
        reads data from websocket and forwards it to output queue
        :param ws: websocket to read from
        :param out: output queue
        :return:
        """
        while True:
            try:
                message = await ws.receive_text()

            except WebSocketDisconnected:
                await out.put(messages.disconnect)
                break

            await out.put(messages.Text(message))

    @staticmethod
    async def run_conversation(ws: WebSocket, inp: asyncio.Queue):
        """
        :param ws:
        :param inp:
        :return:
        """
        while True:
            while inp.empty():
                await asyncio.sleep(0)

            msg = await inp.get()

            try:
                await ws.send_text(str(msg))
            except WebSocketDisconnected:
                pass

            if isinstance(msg, messages.Disconnect):
                break
