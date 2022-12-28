<script lang="ts" setup>
import { AlertStore } from "../../store/alert";
import SingleMessage from "./SingleMessage.vue";
</script>

<template>
	<Transition>
		<div
			class="absolute bottom-0 right-1 w-72 h-64 overflow-x-hidden overflow-y-auto normal-scroll"
			v-if="AlertStore.state.length > 0"
		>
			<TransitionGroup name="list" tag="div" class="flex flex-col gap-2">
				<SingleMessage
					v-for="msg in AlertStore.state"
					v-bind:msg="msg"
					:key="msg.id"
				/>
			</TransitionGroup>
		</div>
	</Transition>
</template>

<style scoped>
.v-enter-active,
.v-leave-active {
	transition: opacity 0.5s ease;
}

.v-enter-from,
.v-leave-to {
	opacity: 0;
}

.list-enter-active,
.list-leave-active {
	transition: all 0.5s ease;
}
.list-enter-from,
.list-leave-to {
	opacity: 0;
	transform: translateX(30px);
}
</style>
