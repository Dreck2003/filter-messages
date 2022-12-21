<script lang="ts" setup>
import { computed } from "vue";
import Cell from "./table-components/cell.vue";
import Row from "./table-components/row.vue";

type BorderColor = "slate" | "blue";

interface RowTable {
	id: number | string;
	selected: boolean;
	cells: Array<{ text: string; id: string | number }>;
}

interface IHeader {
	text: string;
	type: string | number;
}

interface Props {
	headers: IHeader[];
	bodyContent: Array<RowTable>;
	borderColor: BorderColor;
	withClick?: boolean;
}

const props = defineProps<Props>();

defineEmits<{ (e: "select", id: string | number): void }>();

let cleanProps = computed(() => {
	return {
		...props,
		headerColor:
			props.borderColor === "slate" ? "bg-slate-300/80" : "bg-blue-300/80",
		bodyColorPart: (i: number) => {
			return i % 2 != 0 ? "bg-gray-100/90" : "";
		},
		borderColor:
			props.borderColor === "slate"
				? "border-slate-400 border-2"
				: "border-blue-400 border-2",
	};
});

let bodyRows = computed(() => {
	return props.bodyContent.map((row) => {
		return {
			id: row.id,
			cells: row.cells.slice(0, props.headers.length),
			selected: row.selected,
		};
	});
});
</script>

<template>
	<table>
		<thead>
			<Row :class="cleanProps.headerColor">
				<Cell
					v-for="cell in headers"
					:text="cell.text"
					type="header"
					:key="cell.type"
					:class-name="cleanProps.borderColor"
				/>
			</Row>
		</thead>
		<tbody class="text-ellipsis overflow-hidden">
			<Row
				v-for="(row, index) in bodyRows"
				@select="() => $emit('select', row.id)"
				pointer
				:class="row.selected ? 'bg-blue-200' : cleanProps.bodyColorPart(index)"
			>
				<Cell
					v-for="bodyCell in row.cells"
					:text="bodyCell.text"
					:id="bodyCell.id"
					:key="bodyCell.id"
					:class-name="cleanProps.borderColor"
				/>
			</Row>
		</tbody>
	</table>
</template>
