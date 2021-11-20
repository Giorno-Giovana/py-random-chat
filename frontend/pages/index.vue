<template>
  <div @click="handleClick" class="content" ref="content">
    <button @click="getMedia" class="border-black block">
      Дать доступы к камере
    </button>

    <button @click="createOffer" class="border-black block">
      Сгенерить id комнаты
    </button>

    <button @click="answerCall" :disabled="!callId" class="border border-black">
      Войти в комнату
    </button>

    <input
      v-model="callId"
      class="border border-black"
      placeholder="id комнаты"
    />

    <div class="grid grid-cols-2">
      <video-stream :stream="localStream" />

      <video-stream
        v-for="stream in remoteStreams"
        :stream="stream"
        :key="stream.id"
      />
    </div>

    <div id="box" ref="box">
      <p>Move the mouse around in this box to watch its coordinates change.</p>
      <code>pageX</code>:
      <div id="x" ref="x">n/a</div>
      <code>pageY</code>:
      <div id="y" ref="y">n/a</div>
    </div>
  </div>
</template>

<script>
import videoStream from "~/components/video-stream.vue";
import html2canvas from "html2canvas";
const servers = {
  iceServers: [
    {
      urls: ["stun:stun1.l.google.com:19302", "stun:stun2.l.google.com:19302"],
    },
  ],
  // iceCandidatePoolSize: 10,
};

const pc_constraints = { optional: [{ DtlsSrtpKeyAgreement: true }] };
const pc = new RTCPeerConnection(servers, pc_constraints);

export default {
  components: { videoStream },
  data() {
    return {
      callId: "",
      remoteStreams: [],
      localStream: undefined,
      pageX: null,
      pageY: null,
      rectangleSideLength: 75,
    };
  },
  mounted() {
    this.pageX = this.$refs["x"];
    this.pageY = this.$refs["y"];

    var content = this.$refs["content"];
    content.addEventListener("mousemove", this.updateDisplay, false);
    content.addEventListener("mouseenter", this.updateDisplay, false);
    content.addEventListener("mouseleave", this.updateDisplay, false);
  },

  methods: {
    async updateDisplay(event) {
      this.pageX.innerText = event.pageX;
      this.pageY.innerText = event.pageY;
    },

    async handleClick() {
      console.log("sssssssssssssssss");
      var rectangle = document.createElement("div");
      rectangle.className = "rectangle";

      rectangle.style.visibility = "hidden";
      rectangle.style.width = this.rectangleSideLength + "px";
      rectangle.style.height = this.rectangleSideLength + "px";

      rectangle.style.position = "absolute";
      rectangle.style.left = `${
        parseInt(this.pageX.innerText) - this.rectangleSideLength / 2
      }px`;
      rectangle.style.top = `${
        parseInt(this.pageY.innerText) - this.rectangleSideLength / 2
      }px`;

      document.body.append(rectangle);

      html2canvas(rectangle)
        .then((canvas) => {
          var base64 = canvas.toDataURL("image/jpeg");
          // base64 = base64.replace(/^data:image\/[a-z]+;base64,/, "");
          console.log(base64);
          this.$recognition.identifyBase64Image(base64);
        })
        .catch((error) => {
          console.log(error);
        });

      document.body.removeChild(rectangle);
    },

    async getMedia() {
      console.log("Запрашиваем media у пользователя");
      const localStream = await window.navigator.mediaDevices.getUserMedia({
        video: true,
        audio: true,
      });
      this.addStreamToPeers(localStream);
      this.localStream = new MediaStream();
      localStream.getTracks().forEach((track) => {
        console.log(track);
        if (track.kind === "audio") return;
        this.localStream.addTrack(track);
      });
      this.addRemoteStreamToPeersListener();
    },
    async createOffer() {
      const callDoc = this.$fire.firestore.collection("calls").doc();
      const offerCandidates = callDoc.collection("offerCandidates");
      const answerCandidates = callDoc.collection("answerCandidates");

      this.callId = callDoc.id;

      // Get candidates for caller, save to db
      pc.onicecandidate = (event) => {
        event.candidate && offerCandidates.add(event.candidate.toJSON());
      };

      // Create offer
      const offerDescription = await pc.createOffer();
      await pc.setLocalDescription(offerDescription);

      const offer = {
        sdp: offerDescription.sdp,
        type: offerDescription.type,
      };

      await callDoc.set({ offer });

      // Listen for remote answer
      callDoc.onSnapshot((snapshot) => {
        const data = snapshot.data();
        if (!pc.currentRemoteDescription && data?.answer) {
          const answerDescription = new RTCSessionDescription(data.answer);
          pc.setRemoteDescription(answerDescription);
        }
      });

      // When answered, add candidate to peer connection
      answerCandidates.onSnapshot((snapshot) => {
        snapshot.docChanges().forEach((change) => {
          console.log("Новый answer candidate", change);
          if (change.type === "added") {
            const candidate = new RTCIceCandidate(change.doc.data());
            pc.addIceCandidate(candidate);
          }
        });
      });
    },
    async answerCall() {
      const callDoc = this.$fire.firestore.collection("calls").doc(this.callId);
      const answerCandidates = callDoc.collection("answerCandidates");
      const offerCandidates = callDoc.collection("offerCandidates");

      pc.onicecandidate = (event) => {
        event.candidate && answerCandidates.add(event.candidate.toJSON());
      };

      const callData = (await callDoc.get()).data();

      const offerDescription = callData.offer;
      await pc.setRemoteDescription(
        new RTCSessionDescription(offerDescription)
      );

      const answerDescription = await pc.createAnswer();
      await pc.setLocalDescription(answerDescription);

      const answer = {
        type: answerDescription.type,
        sdp: answerDescription.sdp,
      };

      await callDoc.update({ answer });

      offerCandidates.onSnapshot((snapshot) => {
        snapshot.docChanges().forEach((change) => {
          console.log("Новый offer candidate", change);
          if (change.type === "added") {
            let data = change.doc.data();
            pc.addIceCandidate(new RTCIceCandidate(data));
          }
        });
      });
    },
    addRemoteStreamToPeersListener() {
      pc.ontrack = (event) => {
        const currentStream = event.streams[0];
        if (this.remoteStreams.some((s) => s.id === currentStream.id)) return;
        console.log("Получен новый стрим", currentStream);
        this.remoteStreams.push(currentStream);
      };
    },
    addStreamToPeers(stream) {
      stream.getTracks().forEach((track) => {
        pc.addTrack(track, stream);
      });
    },
  },
};
</script>
