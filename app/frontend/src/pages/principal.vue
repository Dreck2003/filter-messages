<script lang="ts" setup>
import { computed, ref, watchEffect } from "vue";
import { replace_string } from "helpers-wasm";

import ExpandContainer from "../components/expand-container.vue";
import EmailIcon from "../components/icons/email-icon.vue";
import SearchInput from "../components/inputs/search-input.vue";
import Table from "../components/table.vue";
import { Email } from "../api/Email";

let text = ref("");
let emailText = ref<string>("");

function handleChanged(content: string) {
	emailText.value = "";
	text.value = content;
}
const HEADERS = [
	{ text: "Subject", type: "subject" },
	{ text: "From", type: "from" },
	{ text: "To", type: "to" },
];

watchEffect(() => {
	Email.SearchByText(text.value)
		.then((emails) => {
			BodyContent.value = emails.map((email, i) => {
				return {
					id: `table-${i}`,
					selected: false,
					content: email.content,
					cells: [
						{
							id: email.id + `-${i}`,
							text: email.subject.length > 0 ? email.subject : "-",
						},
						{ id: email.id + `-${i}`, text: email.from },
						{ id: email.id + `-${i}`, text: email.to },
					],
				};
			});
		})
		.catch((err) => {
			console.log(err);
		});
});

let BodyContent = ref<
	{
		id: number | string;
		selected: boolean;
		cells: Array<{ text: string; id: string | number }>;
		[key: string]: any;
	}[]
>([]);

function handleSelectRow(id: number | string) {
	BodyContent.value = BodyContent.value.map((rowEmail) => {
		if (rowEmail.id == id) {
			emailText.value = rowEmail["content"];
			return { ...rowEmail, selected: true };
		}
		return { ...rowEmail, selected: false };
	});
}

let stringParsed = computed(() => {
	return replace_string(emailText.value, text.value);
});
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
		<section
			class="SectionEmail flex flex-row mx-4 my-1 gap-4 w-full px-4 py-2 justify-center"
		>
			<Table
				:headers="HEADERS"
				:body-content="BodyContent"
				border-color="slate"
				@select="handleSelectRow"
			/>
			<p class="h-max overflow-auto text-sm" v-html="stringParsed"></p>
		</section>
	</ExpandContainer>
</template>
<style>
.SectionEmail {
	overflow-x: hidden;
}

.SectionEmail > p {
	width: 40%;
}

table tr {
	height: 50px;
}

th,
td {
	padding: 0.2em 0.4em;
	overflow: hidden;
	white-space: nowrap;
	text-overflow: ellipsis;
	width: 100px;
}
table {
	max-width: 30rem;
	table-layout: fixed;
	width: 60%;
	height: max-content;
}
</style>
