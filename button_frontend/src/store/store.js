import { createStore } from 'vuex'
import example from './example_data'

export default createStore({
    state () {
        return {
            session_count: 0,
            messages: {}
        }
    },
    mutations: {
        increment (state) {
            state.session_count++
        },
        load_example_messages(state) {
            state.messages = example;
        }
    },
    actions: {
        increment (context) {
            context.commit('increment')
        },
        load_example_messages (context) {
            context.commit('load_example_messages')
        },
        async send_message ({ commit }, payload) {
            const to_send = payload.message_str;
            console.log(to_send);
            commit('increment')
        },
    }
})