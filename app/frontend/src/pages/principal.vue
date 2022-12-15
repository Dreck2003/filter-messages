<script lang="ts" setup>
import { ref } from "vue";
import ExpandContainer from "../components/expand-container.vue";
import EmailIcon from "../components/icons/email-icon.vue";
import SearchInput from "../components/inputs/search-input.vue";
import Table from "../components/table.vue";

let text = ref("");

function handleChanged(content: string) {
	text.value = content;
}
const HEADERS = [
	{ text: "Subject", type: "subject" },
	{ text: "From", type: "from" },
	{ text: "To", type: "to" },
];

let BodyContent = ref([]);

function handleSelectRow(id: number | string) {
	// Filter bodyContent
}
</script>

<template>
	<header class="header bg-slate-400 flex py-4 px-2 justify-start items-center">
		<EmailIcon height="30" width="30" class="fill-slate-700 mr-2" />
		<span class="text-3xl">MamuroEmail</span>
	</header>
	<ExpandContainer
		justify-content="justify-start"
		flex-direction="flex-col"
		align-items="items-center"
	>
		<SearchInput :value="text" @changed="handleChanged" />
		<section class="flex flex-row mx-4 my-1 gap-4">
			<Table
				:headers="HEADERS"
				:body-content="BodyContent"
				border-color="slate"
				@select="handleSelectRow"
			/>
			<p>
				<!-- Email content -->
			</p>
		</section>
	</ExpandContainer>
</template>
<style>
th,
td {
	padding: 0.2em 0.4em;
}
</style>
