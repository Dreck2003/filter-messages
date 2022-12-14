<script lang="ts" setup>
import SearchIcon from "../icons/search-icon.vue";

defineProps({
	value: String,
});
const emit = defineEmits<{
	(e: "changed", content: string): void;
}>();
function handleEnter(e: KeyboardEvent) {
	emit("changed", (e.target as HTMLInputElement).value);
}
</script>

<template>
	<div class="flex rounded-md p-1 border-2 w-10/12 mx-4 my-2 border-slate-400">
		<SearchIcon height="24" width="24" class="fill-slate-700" />
		<input
			type="search"
			:value.trim="value"
			@keyup.prevent.enter="handleEnter"
			class="grow h-full ml-1 text-slate-700"
		/>
	</div>
</template>
<style scoped>
input {
	outline: none;
}
</style>
