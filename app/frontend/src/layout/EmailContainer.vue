<script lang="ts" setup>
import { computed, ref, watchEffect } from "vue";
import { Email, EmailI } from "../api/Email";
import { parse_email_to_html } from "helpers-wasm";
import SearchInput from "../components/inputs/search-input.vue";

let text = ref("");
let emailText = ref<string>("");
let emailSelected = ref<EmailI | null>(null);

let listOfEmails = ref<EmailI[]>([]);

function handleChanged(content: string) {
	emailText.value = "";
	emailSelected.value = null;
	text.value = content;
}

function selectedEmail(id: EmailI["id"]) {
	listOfEmails.value = listOfEmails.value.map((email) => {
		if (email.id == id) {
			emailSelected.value = { ...email };
		}
		return email;
	});
}

watchEffect(() => {
	if (text.value.length > 0) {
		Email.SearchByText(text.value)
			.then((emails) => {
				listOfEmails.value = emails;
			})
			.catch((err) => {
				console.log(err);
			});
	}
});

let stringParsed = computed(() => {
	if (emailSelected.value && text.value.length) {
		return parse_email_to_html(emailSelected.value.content, text.value);
	}
	return emailSelected.value?.content || "";
});

// TODO: Change the image to user icon
// TODO: Change the Date to email Date
</script>
<template>
	<div
		class="w-full h-full flex items-center justify-center bg-gray-100 overflow-hidden"
	>
		<section
			class="Section-Email bg-white flex flex-col md:flex-row h-3/5 border-2 h-4/5 border-gray-200 overflow-hidden w-4/5 max-w-screen-lg"
		>
			<div class="md:border-r-2 border-gray-200 basis-1/4 md:h-full h-72">
				<div class="bg-white border-b-2 border-gray-200 pl-1">
					<SearchInput
						:value="text"
						placeholder="Search..."
						@changed="handleChanged"
						:icon-height="18"
						:icon-width="18"
						icon-class="fill-gray-400"
						class-name="Section-Email-Input min-w-0 items-center"
					/>
				</div>
				<div class="Section-Email-List normal-scroll overflow-auto">
					<div
						class="flex flex-row border-b border-gray-200 cursor-pointer overflow-hidden text-sm h-24 hover:bg-blue-100/50"
						v-for="(email, i) in listOfEmails"
						:key="email.id"
						@click="() => selectedEmail(email.id)"
						:style="{
							borderLeft:
								email.id == emailSelected?.id ? '2px solid #8e8ec9' : '0',
						}"
					>
						<div
							class="rounded-full bg-slate-200 h-8 w-8 aspect-square ml-2 mt-2.5 overflow-hidden shrink-0"
						>
							<img
								class="h-full w-full"
								src="https://miro.medium.com/max/1024/1*jpHGc8WuZ2BHqgaThJrFoA.jpeg"
								alt=""
							/>
						</div>
						<div class="flex flex-col px-2 py-1 overflow-hidden shrink">
							<span class="text-base text-gray-700 font-semibold truncate">{{
								email.name ?? email.from
							}}</span>
							<span class="text-xs text-gray-700 truncate mb-2">{{
								email.subject ?? "---"
							}}</span>
							<span class="text-xs text-gray-500 truncate">{{
								email.content.slice(0, 30)
							}}</span>
						</div>
					</div>
				</div>
			</div>

			<div
				v-if="emailSelected != null"
				class="normal-scroll p-2 basis-3/4 overflow-auto px-2 md:px-14 py-8 h-96 md:h-full border-t-2 border-slate-300 md:border-0"
			>
				<header class="flex flex-row items-center mb-6 gap-3">
					<div
						class="rounded-full bg-slate-200 w-14 h-14 aspect-square overflow-hidden shrink-0"
					>
						<img
							src="https://miro.medium.com/max/1024/1*jpHGc8WuZ2BHqgaThJrFoA.jpeg"
							alt=""
						/>
					</div>
					<div class="flex flex-col flex-1 shrink">
						<div class="flex flex-row justify-between flex-wrap">
							<span class="font-semibold">{{ emailSelected.from }}</span>
							<span class="text-sm text-gray-600">{{
								new Date(emailSelected.date).toDateString()
							}}</span>
						</div>
						<div class="text-sm mt-2">
							<span class="text-gray-400"> to: </span>
							<span class="text-gray-600">
								{{ emailSelected.to }}
							</span>
						</div>
					</div>
				</header>
				<hr class="h-px bg-slate-200 my-6" />
				<section class="mx-4 md:mx-1">
					<div class="mb-10 font-semibold">
						{{ emailSelected.subject ?? "" }}
					</div>
					<pre
						class="Section-Email-Render h-max text-sm whitespace-pre-line"
						v-html="stringParsed"
					></pre>
				</section>
			</div>
			<div
				v-else
				class="basis-3/4 overflow-auto border-t-2 border-slate-300 md:border-0 flex items-center justify-center"
			>
				Has no content
			</div>
		</section>
	</div>
</template>
