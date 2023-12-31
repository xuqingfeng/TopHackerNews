import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import * as child from "child_process";

const commitHash = child.execSync("git rev-parse --short HEAD").toString();
// https://vitejs.dev/config/
export default defineConfig({
  define: {
    "process.env.COMMIT_HASH": JSON.stringify(commitHash),
  },
  server: {
    port: 3000,
  },
  plugins: [svelte()],
});
