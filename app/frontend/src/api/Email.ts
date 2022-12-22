import { CONFIG } from "../config";

export interface EmailI {
	content: string;
	id: string;
	from: string;
	to: string;
	subject: string;
}

export class Email {
	static async SearchByText(term: string): Promise<EmailI[]> {
		const response = await fetch(`${CONFIG.API_URL}/emails/` + term);
		if (response.status == 200) {
			let res = await response.json();
			return (res.content.hits.hits as Array<any>).map<EmailI>(
				(result: any) => {
					return {
						content: result._source.content,
						from: result._source.from,
						id: result._source.emailId,
						subject: result._source.subject,
						to: result._source.to,
					};
				}
			);
		}
		throw new Error("Could not get the emails");
	}
}
