<template>
  <div class="content" ref="content">
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
  </div>
</template>

<script>
import videoStream from "~/components/video-stream.vue";
import html2canvas from "html2canvas";
import { mapState, mapGetters } from "vuex";

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

  computed: {
    ...mapState({
      authUser: (state) => state.authUser,
    }),
    ...mapGetters({
      isLoggedIn: "isLoggedIn",
    }),
  },

  data() {
    return {
      callId: "",
      remoteStreams: [],
      localStream: undefined,
    };
  },

  methods: {
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
