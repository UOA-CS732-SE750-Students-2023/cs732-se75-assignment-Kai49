<template>
  <div class="navbar">
    <button class="nav-button" @click="prevMonth">Previous Month</button>
    <button class="nav-button" @click="nextMonth">Next Month</button>
  </div>
</template>

<script>
export default {
  name: 'Navbar',
  methods: {
    prevMonth() {
      this.$emit('prevMonth');
    },
    nextMonth() {
      this.$emit('nextMonth');
    },
  },
};
</script>

<style scoped>
.navbar {
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #f8f8f8;
  padding: 1rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.nav-button {
  background-color: #ffffff;
  color: #333333;
  border: none;
  border-radius: 4px;
  padding: 0.5rem 1rem;
  font-size: 1rem;
  cursor: pointer;
  margin: 0 0.5rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: all 0.3s;
}

.nav-button:hover {
  background-color: #eeeeee;
}

.nav-button:focus {
  outline: none;
}
</style>
