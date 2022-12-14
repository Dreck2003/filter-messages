<script lang="ts" setup>
import { ref } from "vue";
import ExpandContainer from "../components/expand-container.vue";
import EmailIcon from "../components/icons/email-icon.vue";
import SearchInput from "../components/inputs/search-input.vue";

let text = ref("");

function handleChanged(content: string) {
	text.value = content;
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
		<section>
			<!-- Table and email text -->
		</section>
	</ExpandContainer>
</template>
<style scoped></style>
