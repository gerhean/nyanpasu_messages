<template>
    <div class="card my-card">
      <div class="card-header">
        Nyanpasu {{msg.count}}, sent: {{time}}
      </div>
      <div class="card-body">
        <blockquote class="blockquote mb-0">
          <p>{{msg.msg}}</p>
        </blockquote>
      </div>
    </div>
</template>

<script>
export default {
    name: 'MessageCard',
    props: {
        msg: Object
    },
    computed: {
        time() {
            return (new Date(this.msg.time)).toLocaleTimeString();
        }
    }
}
</script>

<style scoped>
p {
    color: black;
}
.card-header {
    color: gray;
}
.my-card {
    margin: 20px auto;
    display: block;
    width: 50vw;
    opacity: .8;
}
</style>
