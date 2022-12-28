<script lang="ts" setup>
import { createProps, IconProps } from "./interfaces";
interface Props extends IconProps {}

const propsPassed = defineProps<Props>();

let props = createProps(propsPassed);
</script>

<template>
	<svg v-bind="props" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
		<path
			d="m16.192 6.344-4.243 4.242-4.242-4.242-1.414 1.414L10.535 12l-4.242 4.242 1.414 1.414 4.242-4.242 4.243 4.242 1.414-1.414L13.364 12l4.242-4.242z"
		></path>
	</svg>
</template>
