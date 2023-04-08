<template>
  <div 
  v-if="visible"
  class="event-dialog  animate__animated"  
  :class="dialogClass">

    <div class="event-dialog-header">
    <h3>Events on {{ formattedDate }}</h3>
    </div>

    <ul>
      <li v-for="event in events" :key="event.id" class="event-item">
        <template v-if="editingEventId === event.id">
          <input v-model="editingEventContent" class="editing-event-input" />
        </template>
        <template v-else>
          {{ event.content }}
        </template>
        <div class="event-actions">
        <button class="event-button edit" @click.stop="editEvent(event)">Edit</button>
        <button class="event-button delete" @click.stop="deleteEvent(event)">Delete</button>
        </div>
      </li>
    </ul>
    <div class="add-event">
      <input class="add-event-input" v-model="newEventTitle" placeholder="New event title" />
      <button class="add-event-button" @click="addNewEvent">Add Event</button>
      <button class="close-button" @click="$emit('close')">Close</button>

    </div>
  </div>
</template>

<script>
import { ref, computed } from 'vue';
import { useStore } from 'vuex';

export default {
  props: {
    date: {
      type: String,
      default: null,
    },
     visible: {
      type: Boolean,
      default: false,
    },
  },
  setup(props) {
    const store = useStore();
    const newEventTitle = ref('');

    const events = computed(() => store.getters.getEventsForDate(props.date));

    const formattedDate = computed(() => {
      const dateObj = new Date(props.date);
  const dateFormatter = new Intl.DateTimeFormat('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    timeZone: 'Pacific/Auckland',
  });
  return dateFormatter.format(dateObj);
    });

    const addNewEvent = () => {
      if (newEventTitle.value.trim()) {
        const event = {
        //   id: Date.now(),
          date: props.date,
          content: newEventTitle.value.trim(),
        };
        store.dispatch('addEvent', event);
        newEventTitle.value = '';
      }
    };

    const editEvent = (event) => {

  
      const updatedTitle = prompt('Edit event content:', event.content);
      if (updatedTitle && updatedTitle.trim()) {
        store.dispatch('updateEvent', { ...event, content: updatedTitle.trim() });
      }
    };

    const deleteEvent = (event) => {
      if (confirm('Are you sure you want to delete this event?')) {
        store.dispatch('deleteEvent', event);
      }
    };

    return { events, formattedDate, newEventTitle, addNewEvent, editEvent, deleteEvent };
  },
  computed: {
    dialogClass() {
      return this.visible ? 'animate__slideInDown' : 'animate__slideOutUp';
    },
  },
};
</script>

<style scoped>
.event-actions {
  display: inline-block;
}
.event-dialog {
  background-color: #ffffff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  border-radius: 4px;
  padding: 1rem;
  width: 300px;
  font-family: 'Arial', sans-serif;
  z-index: 2;
}

.event-dialog-header {
  margin-bottom: 1rem;
  text-align: center;
}

.event-list {
  list-style-type: none;
  padding: 0;
  margin-bottom: 1rem;
}

.event-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0.5rem;
  padding: 0.25rem;
  border: 1px solid transparent;
  border-radius: 4px;
  transition: all 0.3s;
}
.event-item:hover {
  background-color: #f3f3f3;
  border-color: #e0e0e0;
}

.event-button {
  background-color: #ffffff;
  color: #333333;
  border: 1px solid #333333;
  border-radius: 4px;
  padding: 0.25rem 0.5rem;
  font-size: 0.9rem;
  cursor: pointer;
  transition: all 0.3s;
  text-align: center;
}

.event-button.edit {
  margin-right: 0.5rem;
}

.event-button.delete {
  color: #ff3333;
  border-color: #ff3333;
}

.event-button:hover {
  background-color: #f3f3f3;
}

.add-event {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.add-event-input {
  flex-grow: 1;
  padding: 0.25rem;
  border: 1px solid #cccccc;
  border-radius: 4px;
  margin-right: 0.5rem;
}

.add-event-button {
  background-color: #ffffff;
  color: #333333;
  border: 1px solid #333333;
  border-radius: 4px;
  padding: 0.25rem 0.5rem;
  font-size: 0.9rem;
  cursor: pointer;
  transition: all 0.3s;
}

.add-event-button:hover {
  background-color: #f3f3f3;
}

.close-button {
  background-color: #ffffff;
  color: #333333;
  border: 1px solid #333333;
  border-radius: 4px;
  padding: 0.25rem 0.5rem;
  font-size: 0.9rem;
  cursor: pointer;
  transition: all 0.3s;
}
.close-button:hover {
  background-color: #f3f3f3;
}

.event-dialog h3 {
  font-size: 1.2rem;
  font-weight: normal;
  color: #333333;
  margin: 0;
}
</style>
