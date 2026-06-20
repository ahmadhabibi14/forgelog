<script lang="ts">
  import { onMount } from "svelte";
  import { currentPage, PathDashboard } from "../states/page";

  onMount(() => {
    currentPage.set(PathDashboard)
  })
</script>

<div>
  <h1>Dashboard</h1>
</div>