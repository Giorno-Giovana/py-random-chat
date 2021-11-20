import firebase from "firebase";

var firebaseConfig = {
  apiKey: "AIzaSyAOzDEIo2XfJKsnxdYLTqGy_DgqyMUVntE",
  authDomain: "juction-1f3a4.firebaseapp.com",
  projectId: "juction-1f3a4",
  storageBucket: "juction-1f3a4.appspot.com",
  messagingSenderId: "1059864498087",
  appId: "1:1059864498087:web:b0381e9160b342d27b3082"
};

firebase.initializeApp(firebaseConfig);

export const FirebaseStorage = firebase.storage();
export const auth = firebase.auth()
export default firebase