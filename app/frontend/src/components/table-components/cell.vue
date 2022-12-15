<script lang="ts" setup>
import { computed } from "vue";

type TypeCell = "body" | "header";

interface CellProps {
	text: string;
	type?: TypeCell;
	className?: string;
}

const props = defineProps<CellProps>();
let typeCell = computed(() => props.type ?? "body");
</script>

<template>
	<td v-if="typeCell == 'body'" :class="className ?? ''">{{ text }}</td>
	<th v-else :class="className ?? ''">{{ text }}</th>
</template>
