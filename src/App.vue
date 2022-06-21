<template>
  <div class="app-layout">
    <div class="app-layout__content">
      <div class="app-layout__sidebar-wrapper" :class="{ minimized: false }">
        <Sidebar :minimized="false" />
      </div>
      <div class="app-layout__page">
        <div class="layout fluid gutter--xl">
          <HelloWorld msg="ggg" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import Sidebar from "./components/Sidebar.vue";
import HelloWorld from "./components/HelloWorld.vue";
</script>

<style lang="scss">
$mobileBreakPointPX: 640px;
$tabletBreakPointPX: 768px;

.app-layout {
  height: 100vh;
  display: flex;
  flex-direction: column;
  &__navbar {
    min-height: 4rem;
  }

  &__content {
    display: flex;
    height: calc(100vh - 4rem);
    flex: 1;

    @media screen and (max-width: $tabletBreakPointPX) {
      height: calc(100vh - 6.5rem);
    }

    .app-layout__sidebar-wrapper {
      position: relative;
      height: 100%;
      background: var(--va-white);

      @media screen and (max-width: $tabletBreakPointPX) {
        &:not(.minimized) {
          width: 100%;
          height: 100%;
          position: fixed;
          top: 0;
          z-index: 999;
        }

        .va-sidebar:not(.va-sidebar--minimized) {
          // Z-index fix for preventing overflow for close button
          z-index: -1;
          .va-sidebar__menu {
            padding: 0;
          }
        }
      }
    }
  }
  &__page {
    flex-grow: 2;
    overflow-y: scroll;
  }
}
</style>
