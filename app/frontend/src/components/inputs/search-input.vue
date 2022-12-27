<script lang="ts" setup>
import SearchIcon from "../icons/search-icon.vue";

defineProps({
	value: String,
	className: String,
	inputClass: String,
	iconClass: String,
	iconHeight: Number,
	iconWidth: Number,
	placeholder: String,
});
const emit = defineEmits<{
	(e: "changed", content: string): void;
}>();
function handleEnter(e: KeyboardEvent) {
	emit("changed", (e.target as HTMLInputElement).value);
}
</script>

<template>
	<div class="flex rounded-md p-1 overflow-hidden" :class="className">
		<SearchIcon
			class="fill-slate-700"
			:class="iconClass"
			:height="iconHeight"
			:width="iconWidth"
		/>
		<input
			:placeholder="placeholder"
			type="search"
			:value.trim="value"
			@keyup.prevent.enter="handleEnter"
			class="ml-1 text-slate-700 w-full"
			:class="inputClass"
		/>
	</div>
</template>
<style scoped>
input {
	outline: none;
}
</style>
