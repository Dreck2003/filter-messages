import { reactive } from "vue";

type TypeMessage = "Alert" | "Success";

export interface MessageAlert {
	text: string;
	type: TypeMessage;
	id: number;
}

interface AlertStoreI {
	state: MessageAlert[];
	addMessage(msg: { text: string; type: TypeMessage }): void;
	removeMessage(id: MessageAlert["id"]): void;
}

export const AlertStore = reactive<AlertStoreI>({
	state: [],
	addMessage(msg) {
		this.state.push({ ...msg, id: this.state.length + 1 });
	},
	removeMessage(msgId) {
		console.log("remove: ", msgId);
		this.state = this.state.filter(({ id }) => id != msgId);
	},
});
