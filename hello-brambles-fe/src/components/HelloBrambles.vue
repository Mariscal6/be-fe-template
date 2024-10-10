<template>
  <div class="container">
    <slot></slot>
    <div class="footer">
      <BramblesLogo />
    </div>
  </div>
</template>
<script setup lang="ts">
import { onMounted } from "vue";
import BramblesLogo from "./BramblesLogo.vue";
onMounted(() => {
  // change me with your url
  // don't forget to change visibility on the port
  fetch("")
    .then((response) => response.text())
    .then((data) => {
      console.log(data);
    })
    .catch((error) => {
      console.error("Error fetching data:", error);
    });
});
</script>

<style scoped>
.container {
  width: 95%;
  margin: 0 auto;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 90vh;
  border: 2px solid rgb(4, 57, 110);
  border-radius: 0.5rem;
  position: relative;
}

.footer {
  position: absolute;
  bottom: 0;
  right: 0;
  margin: 15px;
  width: 100px;
}
</style>
