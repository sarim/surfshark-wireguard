<template>
  <header>
    <va-navbar color="dark" shape class="mb-2">
      <template #left>
        <va-navbar-item>
          <img
            alt="Vue logo"
            class="logo"
            src="./assets/logo.svg"
            width="125"
            height="125"
          />
        </va-navbar-item>
      </template>
      <template #center>
        <va-navbar-item>Dark</va-navbar-item>
      </template>
      <template #right>
        <va-navbar-item>Right</va-navbar-item>
      </template>
    </va-navbar>
  </header>
  <div class="content">Content</div>
  <va-sidebar color="dark" minimizedWidth="64px" width="18rem">
    <va-sidebar-item active-color="background">
      <va-sidebar-item-content>
        <va-sidebar-item-title>
          background dark and auto text color
        </va-sidebar-item-title>
      </va-sidebar-item-content>
    </va-sidebar-item>
    <va-sidebar-item
      v-for="item in items"
      :key="item.to"
      :active="item.active"
      active-color="background"
    >
      <va-sidebar-item-content>
        <va-icon :name="item.icon" />
        <va-sidebar-item-title style="height: 24px">
          {{ item.title }}
        </va-sidebar-item-title>
      </va-sidebar-item-content>
    </va-sidebar-item>
  </va-sidebar>
  <main>
    <HelloWorld msg="You did it!" />
  </main>
</template>

<script setup lang="ts">
import HelloWorld from "./components/HelloWorld.vue";

const items = [
  { title: "Dashboard", icon: "dashboard" },
  { title: "Sidebar demo", icon: "room", active: true },
  { title: "Loop", icon: "loop" },
];
</script>

<style lang="scss">
#app {
  font-family: "Source Sans Pro", Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
body {
  margin: 0;
  background: var(--va-background);
}
</style>
