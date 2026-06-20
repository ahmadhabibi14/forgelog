<script lang="ts">
  import Router, { type RouteDefinition } from 'svelte-spa-router';
  import Dashboard from './pages/Dashboard.svelte';
  import Containers from './pages/Containers.svelte';
  import { TextAlignJustify } from '@lucide/svelte';
  import SideMenu from './partials/SideMenu.svelte';
  import ContainerDetailDocker from './pages/containers/Docker.svelte';

  const routes: RouteDefinition = {
    '/': Dashboard,
    '/containers': Containers,
    '/containers/docker/:id': ContainerDetailDocker as any
  };
</script>

<div class="top-0 bottom-0 left-0 right-0 min-h-screen min-w-screen">
  <div class="flex flex-col h-screen w-full bg-neutral-950 text-neutral-300">
    <div class="h-16 w-full flex flex-row justify-between px-4">
      <div class="flex flex-row items-center gap-4">
        <button class="flex justify-center items-center p-2
        hover:bg-neutral-900 active:ring-0 focus:ring-0 focus:border-none active:border-none
        focus:bg-neutral-900 focus:outline-0 rounded-lg cursor-pointer">
          <TextAlignJustify
            size="20"
          />
        </button>
        <div>
          <img
            src="/main-logo.svg"
            alt="Forgelog"
            class="h-6"
          />
        </div>
      </div>
    </div>
    <div class="h-full grow flex flex-row">
      <SideMenu />
      <div class="bg-neutral-900 flex-1 m-3 rounded-lg p-3 overflow-hidden">
  <Router {routes} />
</div>
    </div>
  </div>
</div>
