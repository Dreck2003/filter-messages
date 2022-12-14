<script lang="ts" setup>
import { computed } from "vue";

type AlignItems = "items-start" | "items-end" | "items-center";
type JustifyContent = "justify-start" | "justify-end" | "justify-center";
type FlexDirection = "flex-row" | "flex-col";

interface Container {
	justifyContent?: JustifyContent;
	alignItems?: AlignItems;
	flexDirection?: FlexDirection;
}

const props = defineProps<Container>();

const data = computed(() =>
	Object.values({
		alignItems: props.alignItems ?? "items-center",
		flexDirection: props.flexDirection ?? "flex-col",
		justifyContent: props.justifyContent ?? "justify-center",
	})
);
</script>
<template>
	<main class="flex w-full h-full" :class="data">
		<slot />
	</main>
</template>
<style scoped></style>
