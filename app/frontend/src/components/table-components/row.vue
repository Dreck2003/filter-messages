<script lang="ts" setup>
interface RowEvents {
	(e: "select"): void;
}

const props = defineProps<{
	pointer?: boolean;
}>();

let isClick = props.pointer ? "cursor-pointer" : "";

defineEmits<RowEvents>();
</script>

<template>
	<tr @click="$emit('select')" :class="isClick">
		<slot />
	</tr>
</template>
