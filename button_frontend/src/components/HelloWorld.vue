<template>
  <div class="hello">
    <h1>Nyanpasu Message Board</h1>
    
    <div class="card-list overflow-auto" ref="cardList">
      <MessageCard v-bind:msg="msg" v-for="msg in messages" :key="msg.count"/>
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
    'messages'
  ]),
  mounted() {
    this.$store.dispatch('load_backend_messages')
  },
  watch: {
    messages: {
      handler() {
        this.$nextTick(() => this.scrollToEnd());
      },
      deep: true
    }
  },
  methods: {
    sendMessage() {
      this.$store.dispatch('send_message', {message_str: this.messageInput});
      this.messageInput = "";
    },
    scrollToEnd: function () {
      // scroll to the start of the last message
      console.log("hi");
      this.$refs.cardList.scrollTop = this.$refs.cardList.lastElementChild.offsetTop;
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
.card-list {
  height: 50vh;
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
