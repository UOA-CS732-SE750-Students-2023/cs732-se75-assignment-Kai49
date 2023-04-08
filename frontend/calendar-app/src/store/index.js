import Vuex from 'vuex';
import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:3000', // 将此值替换为您的后端 API 地址
});

export default new Vuex.Store({
  state: {
    currentDate: new Date(),
    events: []
  },
  mutations: {
    setCurrentDate(state, payload) {
      state.currentDate = payload;
    },

    addEvent(state, payload) {
      state.events.push(payload);
    },

    setEvents(state, events) {
      state.events = events;
    },

    updateEvent(state, payload) {
      const index = state.events.findIndex((event) => event.id === payload.id);
      if (index !== -1) {
        state.events[index] = payload;
      }
    },
    deleteEvent(state, payload) {
      const index = state.events.findIndex((event) => event.id === payload.id);
      if (index !== -1) {
        state.events.splice(index, 1);
      }
    },
  },
  actions: {

    async fetchEvents({ commit }) {
      try {
        const response = await api.get('/events');
        commit('setEvents', response.data);
      } catch (error) {
        console.error('Error fetching events:', error);
      }
    },
    async addEvent({ commit }, event) {
      try {
        const response = await api.post('/events', event);
        commit('addEvent', response.data);
      } catch (error) {
        console.error('Error adding event:', error);
      }
    },


    updateCurrentDate({ commit }, date) {
      commit('setCurrentDate', date);
    },
    // addEvent({ commit }, event) {
    //   commit('addEvent', event);
    // },

    async updateEvent({ commit }, event) {
      try {
        const response = await api.put(`/events/${event.id}`, event);
        commit('updateEvent', response.data);
      } catch (error) {
        console.error('Error updating event:', error);
      }
    },
    // updateEvent({ commit }, event) {
    //   commit('updateEvent', event);
    // },
    async deleteEvent({ commit }, event) {
      try {
        await api.delete(`/events/${event.id}`);
        commit('deleteEvent', event);
      } catch (error) {
        console.error('Error deleting event:', error);
      }
    },
    // deleteEvent({ commit }, event) {
    //   commit('deleteEvent', event);
    // },
  },
  getters: {
    getCurrentDate: state => state.currentDate,
    getEvents(state) {
      return (date) => {
        const dateString = date.toISOString().split('T')[0];
        return state.events.filter((event) => event.date === dateString);
      };
      
    },

    getEventsForDate: (state) => (date) => {
      return state.events.filter((event) => event.date === date);
    },
  },
});
