<script lang="ts" setup>
let className = "bg-blue-100 px-1 text-center py-6";
</script>

<template>
	<div class="my-16 mx-4 flex flex-row gap-1 max-h-24">
		<span :class="className"> </span>
		<span :class="className"> </span>
		<span :class="className"> </span>
		<span :class="className"> </span>
	</div>
</template>

<style scoped>
div > span {
	animation: spin 1.2s infinite;
}
div > span:nth-child(2) {
	animation-delay: -1.1s;
}
div > span:nth-child(3) {
	animation-delay: -1s;
}
div > span:nth-child(4) {
	animation-delay: -0.9s;
}

@keyframes spin {
	0%,
	40%,
	100% {
		transform: scaleY(0.4);
	}
	82% {
		transform: scaleY(1);
	}
}
</style>
