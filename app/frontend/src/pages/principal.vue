<script lang="ts" setup>
import EmailIcon from "../components/icons/email-icon.vue";
import EmailContainer from "../layout/EmailContainer.vue";
</script>

<template>
	<header class="header bg-slate-400 flex py-4 px-2 justify-start items-center">
		<EmailIcon height="30" width="30" class="fill-slate-700 mr-2" />
		<span class="text-3xl">MamuroEmail</span>
	</header>
	<EmailContainer />
</template>
<style>
.Section-Email {
	min-height: 10rem;
}

.Section-Email > :first-child {
	flex-grow: 1;
	min-width: 12rem;
}

.Section-Email > :last-child {
	flex-grow: 2;
}

.SectionEmail > pre {
	width: 40%;
}

.Section-Email-List > :last-child {
	border: 0;
}

.Section-Email-List {
	/*   Add 2.5rem from .Section-Email-Input  */
	height: calc(100% - 3rem);
}

.Section-Email-Input {
	height: 3rem;
}

.Section-Email-Render mark {
	padding: 1px 0.3em;
	border-radius: 2px;
	background-color: rgb(250, 227, 57);
}
</style>
