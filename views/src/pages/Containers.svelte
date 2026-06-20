<script lang="ts">
  import axios from "axios";
  import { onMount } from "svelte";
  import type { DockerContainer } from "../types/docker";
  import * as Table from "$lib/components/ui/table/index.js";
  import { formatUnixTimestamp } from "$lib/formatter";
  import { currentPage, PathContainers } from "../states/page";

  onMount(() => {
    currentPage.set(PathContainers)
  })

  let ContainersDocker: DockerContainer[] = [];

  async function getContainersDocker() {
    await axios
      .post("/api/containers/docker")
      .then((res) => {
        console.log(res.data);
        ContainersDocker = res.data.Items || [];
      })
      .catch((err) => {
        console.error(err);
      });
  }

  onMount(async () => {
    await getContainersDocker();
  });
</script>

<div class="flex flex-col gap-2">
  <Table.Root>
    <Table.Header>
      <Table.Row class="">
        <Table.Head class="w-25">Container ID</Table.Head>
        <Table.Head>Image</Table.Head>
        <Table.Head>Name</Table.Head>
        <Table.Head>Created</Table.Head>
        <Table.Head class="text-end">Status</Table.Head>
      </Table.Row>
    </Table.Header>
    <Table.Body>
      {#each (ContainersDocker || []) as docker}
      {@const createdTime = new Date(docker.Created * 1000)}
      <Table.Row class="group cursor-pointer" onclick={() => window.location.href = `#/containers/docker/${docker.Id}`}>
        <Table.Cell class="group-hover:underline group-hover:text-blue-500 font-medium">{docker.Id.slice(0,12)}</Table.Cell>
        <Table.Cell class="group-hover:underline group-hover:text-blue-500">{docker.Image}</Table.Cell>
        <Table.Cell class="group-hover:underline group-hover:text-blue-500">{docker.Names[0]}</Table.Cell>
        <Table.Cell class="group-hover:underline group-hover:text-blue-500">
          <time datetime={createdTime.toString()}>{formatUnixTimestamp(docker.Created)}</time>
        </Table.Cell>
        <Table.Cell class="text-end">
          <span class="{docker.State} text-xs py-1 px-3 rounded-full">
            {docker.Status}
          </span>
        </Table.Cell>
      </Table.Row>
      {/each}
    </Table.Body>
  </Table.Root>
</div>

<style>
  @reference "../app.css";

  .exited {
    @apply bg-red-300/10 text-red-500;
  }

  .running {
    @apply bg-emerald-300/10 text-emerald-500;
  }
</style>
