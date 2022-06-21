<template>
  <va-sidebar
    :width="width"
    :minimized="minimized"
    :minimizedWidth="minimizedWidth"
  >
    <va-sidebar-item>
      <va-sidebar-item-content>
        <va-icon name="dashboard" />
        <!-- User can hide item with css if he wants -->
        <va-sidebar-item-title> Dashboard </va-sidebar-item-title>
      </va-sidebar-item-content>
    </va-sidebar-item>

    <va-sidebar-item active>
      <va-sidebar-item-content>
        <va-icon name="room" />
        <va-sidebar-item-title> Sidebar demo </va-sidebar-item-title>
      </va-sidebar-item-content>
    </va-sidebar-item>
  </va-sidebar>
</template>

<script>
export default {
  name: "app-sidebar",
  components: {},
  props: {
    width: { type: String, default: "16rem" },
    color: { type: String, default: "secondary" },
    minimized: { type: Boolean, required: true },
    minimizedWidth: {
      type: String,
      required: false,
      default: undefined,
    },
  },
  data() {
    return {};
  },
  computed: {
    computedClass() {
      return {
        "app-sidebar--minimized": this.minimized,
      };
    },
  },
};
</script>

<style lang="scss">
.va-sidebar {
  .va-collapse__body {
    margin-top: 0 !important;
  }

  &__menu {
    padding: 2rem 0;
    &__inner {
      padding-bottom: 8rem;
    }
  }

  &-item {
    &-content {
      padding: 0.75rem 1rem;
    }

    &__icon {
      width: 1.5rem;
      height: 1.5rem;
      display: flex;
      justify-content: center;
      align-items: center;
    }
  }
}
</style>

<style lang="scss" scoped>
.va-sidebar {
  flex-shrink: 0;
}

// .va-sidebar--minimized {
//   width: auto !important;
// }
</style>
