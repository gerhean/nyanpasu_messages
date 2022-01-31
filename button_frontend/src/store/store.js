import { createStore } from 'vuex'
import example from './example_data'

const BACKEND_URL = process.env.VUE_APP_BACKEND_URL || "http://localhost:8082/"

export default createStore({
    state () {
        return {
            session_count: 0,
            messages: []
        }
    },
    mutations: {
        append_message(state, message) {
            state.messages.push(message);
        },
        load_example_messages(state) {
            state.messages = example;
        },
        load_messages(state, messages) {
            state.messages = messages;
        }
    },
    actions: {
        load_example_messages (context) {
            context.commit('load_example_messages')
        },

        async send_message ({ commit }, payload) {
            const url = BACKEND_URL + "messages";
            const to_send = {
                "msg": payload.message_str || "にゃんぱすー!",
                "time": (new Date()).toJSON()
            };
            console.log(to_send);
            const rawResponse = await fetch(url, {
                method: 'POST',
                body: JSON.stringify(to_send)
            });
            const data = await rawResponse.json();
            if (data.error) {
                return;
            }
            commit('append_message', data)
        },

        async load_backend_messages({ commit, dispatch }) {
            const url = BACKEND_URL + "messages";
            const response = await fetch(url, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json'
                },
            });
            const data = await response.json();
            if (data.error) {
                dispatch('load_example_messages');
            } else {
                commit('load_messages', data);
            }
        }
    }
})