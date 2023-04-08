<template>
  <div class="day" 
  @click.native="showEvents">
    <div class="date animate__animated animate__zoomIn animate__delay-{{ dayDelay }}s">{{ day }}</div>

    <div class="day-events animate__animated animate__bounceIn"  v-if="eventsCount > 0" > {{ eventsCount }} </div>

  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import { useStore } from 'vuex';
import { computed, onMounted } from 'vue';

export default {
  props: {
    day: {
      type: Number,
      required: true,
    },
    date: {
      type: String,
      default: '',
    },
 
  },
  setup(props, { emit }) {
    const store = useStore();
    const eventsCount = computed(() => store.getters.getEventsForDate(props.date).length);
    
    onMounted(() => {
        store.dispatch('fetchEvents')
        
        });

    const showEvents = () => {
      if (props.date) {
    const selectedDate = new Date(props.date);
    selectedDate.setDate(selectedDate.getDate() + 1);
    emit('show-events', selectedDate.toISOString().split('T')[0]);
    emit("open-event-dialog");
  }
    };
    return { eventsCount, showEvents };
  },


  computed: {
    ...mapGetters(['getEvents']),
    eventCount() {
      return this.getEvents(this.day).length;
    },
  },
  methods: {
    handleClick() {
      this.$emit('openEventDialog', this.day);
    },
  },
};
</script>

<style>
.day {
  position: relative;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: center;
  background-color: #f8f8f8;
  border-radius: 8px;
  padding: 0.5rem;
  height: 60px;
  cursor: pointer;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
  transition: all 0.3s;
  z-index: 1; 
}

.day:hover {
  background-color: #e6e6e6;
}

.day-date {
  font-size: 1.2rem;
  font-weight: bold;
  color: #333333;
  margin-bottom: 0.5rem;
}

.day-events {
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #ff8c00;
  color: #ffffff;
  border-radius: 50%;
  width: 20px;
  height: 20px;
  font-size: 0.8rem;
  font-weight: bold;
}
</style>
