type Numberish = number | string;

export interface IconProps {
	class?: string;
	height?: Numberish;
	width?: Numberish;
	"stroke-width"?: Numberish;
}

export const createProps = (props: IconProps): IconProps => ({
	class: props.class ?? "",
	height: props.height ?? 24,
	width: props.width ?? 24,
	"stroke-width": props["stroke-width"] ?? 3,
});
