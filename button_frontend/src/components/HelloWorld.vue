<template>
  <div class="hello">
    <h1>Nyanpasu Message Board</h1>

    <div v-for="msg in messages" :key="msg.count">
      <MessageCard v-bind:msg="msg"/>
    </div>

    <div class="my-footer">
      <div class="mb-3 message-input">
        <input type="text" class="form-control" id="messageForm" v-model="messageInput" placeholder="Nyanpasu!">
      </div>

      <button @click="sendMessage" type="button" class="send-btn btn btn-success">Shout!</button>
    </div>

  </div>
</template>

<script>
import { mapState } from 'vuex'
import MessageCard from './MessageCard'

export default {
  name: 'HelloWorld',
  data() {
    return {
      messageInput: ""
    }
  },
  components: {
    MessageCard
  },
  computed: mapState([
    // map this.count to store.state.count
    'messages'
  ]),
  mounted() {
    this.$store.commit('load_example_messages')
  },
  methods: {
    sendMessage() {
      this.$store.dispatch('send_message', {message_str: this.messageInput});
      this.messageInput = "";
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1 {
  width: 50vw;
  margin: auto;
  color: plum;
  background-color: greenyellow;
}
.my-card {
  margin: 20px auto;
  display: block;
  width: 50vw;
}
.my-footer {
  position: absolute;
  bottom: 20px;
  width: 100%;
}
.message-input {
  margin: auto;
  width: 50vh;
}
.send-btn {
  margin: auto;
}
#messageForm {
  border-radius: 20%;
  font-size: x-large;
}
</style>
