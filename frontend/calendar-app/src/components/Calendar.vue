<template>
  <div class="calendar animate__animated animate__fadeIn">
     <Navbar @prevMonth="goToPrevMonth" @nextMonth="goToNextMonth" />
    <h1>{{ formattedCurrentDate }}</h1>
     <div class="calendar-header">
      <div class="calendar-header-day" v-for="day in ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']" :key="day">
        {{ day }}
      </div>
    </div>
    <div class="calendar-grid">
      <Day
        v-for="dateObj in dates"
        :key="dateObj.date"
        :day="dateObj.day"
        :date="dateObj.date"
        :currentDate="currentDate.value"
        @show-events="showEvents"
        @update-selected-date="updateSelectedDate"
        @open-event-dialog="() => {
        eventDialogVisible = true;
        selectedDate = dateObj.date;
        }"    

      ></Day>
    </div>
    <EventDialog v-if="selectedDate" :visible="eventDialogVisible" :date="selectedDate" @close="eventDialogVisible = false"></EventDialog>
  </div>
</template>

<script>
import { ref, computed, watch } from 'vue';
import Day from './Day.vue';
import EventDialog from './EventDialog.vue';
import { getMonthDates } from '../utils/dateUtils';
import Navbar from './Navbar.vue';


export default {
  components: {
    Day,
    EventDialog,
    Navbar
  },
   data() {
    return {
      eventDialogVisible: false,
      vantaEffect: null,
    };
  },
  setup() {

    const currentDate = ref(new Date());
    
    const selectedDate = ref(null);
    const dates = ref(getMonthDates(currentDate.value));
    
   
    
    const updateSelectedDate = (newDate) => {
      selectedDate.value = newDate;
    };



    const formattedCurrentDate = computed(() => {
      return currentDate.value.toLocaleString('en-NZ', { month: 'long', year: 'numeric' });
    });


    watch(currentDate, () => {
    dates.value = getMonthDates(currentDate.value);
  });
    
 

    const showEvents = (date) => {
      selectedDate.value = new Date(date).toISOString().split('T')[0];
    };

     const goToPrevMonth = () => {
          

      const newDate = new Date(currentDate.value.getFullYear(), currentDate.value.getMonth() - 1);
      currentDate.value = newDate;

    };

    const goToNextMonth = () => {
      const newDate = new Date(currentDate.value.getFullYear(), currentDate.value.getMonth() + 1);
      currentDate.value = newDate;
    };

    return { formattedCurrentDate, dates, showEvents, selectedDate,currentDate, goToPrevMonth, goToNextMonth, updateSelectedDate };
  },

};
</script>

<style scoped>
#snow-particles {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1; /* Add this line to display snowflakes behind other elements */
}

.calendar {
  display: flex;
  flex-direction: column;
  align-items: center;
  font-family: 'Arial', sans-serif;
}

.calendar-header {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 1rem;
  margin: 1rem;
  font-size: 1.1rem;
  font-weight: bold;
}

.calendar-header-day {
  display: flex;
  justify-content: center;
  align-items: center;
}

.calendar-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 1rem;
  margin: 1rem;
}
</style>
