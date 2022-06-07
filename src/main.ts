import { createApp } from "vue";
import App from "./App.vue";
import { VuesticPluginsWithoutComponents } from 'vuestic-ui'
import 'vuestic-ui/dist/styles/essential.css'
import 'vuestic-ui/dist/styles/grid/grid.scss'
import 'vuestic-ui/dist/styles/global/normalize.scss'
import 'vuestic-ui/dist/styles/global/typography.scss'

createApp(App).use(VuesticPluginsWithoutComponents).mount("#app");
