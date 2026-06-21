<script lang="ts">
	import Router, { type RouteDefinition } from 'svelte-spa-router';
	import Dashboard from './pages/Dashboard.svelte';
	import Containers from './pages/Containers.svelte';
	import SideMenu from './partials/SideMenu.svelte';
	import ContainerDetailDocker from './pages/containers/Docker.svelte';
	import Navbar from './partials/Navbar.svelte';


	const routes: RouteDefinition = {
		'/': Dashboard,
		'/containers': Containers,
		'/containers/docker/:id': ContainerDetailDocker as any
	};
</script>

<div class="top-0 bottom-0 left-0 right-0 min-h-screen min-w-screen">
	<div class="h-screen w-full bg-neutral-950 text-neutral-300">
		<!-- NAVBAR -->
		 <Navbar />
		<div class="grow flex flex-row fixed top-16 right-0 left-0 bottom-0">
			<SideMenu />
			<div class="grow p-3 h-full w-full flex min-h-0">
			<div class="bg-neutral-900 rounded-lg p-3 overflow-hidden w-full">
				<Router {routes} />
			</div>
			</div>
		</div>
	</div>
</div>
