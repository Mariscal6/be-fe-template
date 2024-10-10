<template>
  <div class="container">
    <slot></slot>
    <div class="footer">
      <BramblesLogo />
    </div>
  </div>
</template>

<script lang="ts">
import BramblesLogo from "./BramblesLogo.vue";
import { Options, Vue } from "vue-class-component";
@Options({
  props: {
    msg: String,
  },
  components: {
    BramblesLogo,
  },
})
export default class HelloBrambles extends Vue {
  msg!: string;
}
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
  position: relative;
  /* Add this line */
}

.footer {
  position: absolute;
  bottom: 0;
  right: 0;
  margin: 15px;
  width: 100px;
}
</style>
