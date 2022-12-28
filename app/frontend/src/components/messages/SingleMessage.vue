<script lang="ts" setup>
import { computed } from "vue";
import { AlertStore, MessageAlert } from "../../store/alert";
import AlertIcon from "../icons/alert-icon.vue";
import CloseIcon from "../icons/close-icon.vue";
import SuccessIcon from "../icons/success.icon.vue";

interface MessageProps {
	msg: MessageAlert;
}

const props = defineProps<MessageProps>();

type TypeIconsI = {
	[key in MessageAlert["type"]]: {
		component: any;
		className: string;
	};
};

const typeIcons: TypeIconsI = {
	Alert: {
		component: AlertIcon,
		className: "bg-orange-400/80",
	},
	Success: {
		component: SuccessIcon,
		className: "bg-green-600/80",
	},
};
</script>

<template>
	<div
		class="flex flex-row items-center justify-between gap-2 w-full py-2 px-2 rounded"
		:class="typeIcons[props.msg.type].className"
	>
		<span class="">
			<component
				:is="typeIcons[props.msg.type].component"
				height="20"
				width="20"
				class="fill-white"
			></component>
		</span>
		<span class="text-sm text-white">
			{{ props.msg.text }}
		</span>
		<span
			@click="AlertStore.removeMessage(props.msg.id)"
			class="cursor-pointer"
		>
			<CloseIcon class="fill-white" height="20" width="20" />
		</span>
	</div>
</template>

<style scoped></style>
