import Vuex from 'vuex'
import { auth } from '~/plugins/firebase.js'

const createStore = () => {
  return new Vuex.Store({
    state: {
      user: '',
    },

    getters: {
      user(state) {
        return state.user
      },

      isAuthenticated(state) {
        return !!state.user
      }
    },

    mutations: {
      setUser(state, payload) {
        state.user = payload
      }
    },

    actions: {
      signUp({ commit }, { email, password }) {
        return auth.createUserWithEmailAndPassword(email, password)
      },

      signInWithEmail({ commit }, { email, password }) {
        return auth.signInWithEmailAndPassword(email, password)
      },

      signOut() {
        return auth.signOut()
      },

      async recognitionUploadImage({ commit }, { username, file }) {
        var url = "http://172.17.0.2:2000/api/upload_image"

        var data = new FormData()
        data.append('image', file)
        data.append('username', username)

        fetch(url, {
          method: "POST",
          body: data,
          headers: {
            "Access-Control-Allow-Origin": "*",
          }
        })
          .catch(error => {
            alert(error);
          });
      }
    }
  })
}

export default createStore