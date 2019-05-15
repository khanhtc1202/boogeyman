import React, { Component } from "react";
import {Item} from "semantic-ui-react";

export class Results extends Component {
	constructor(props) {
		super(props);
	}

	render() {
		let groupedItems = this.props.results.map((item) => {
			return (
				<Item>
					<Item.Content>
						<Item.Header>{item.title}</Item.Header>
						<Item.Extra as='a' href={item.url} target='_blank'>{item.url}</Item.Extra>
						<Item.Description>{item.description}</Item.Description>
					</Item.Content>
				</Item>
			);
		});

		return (
			<Item.Group divided>
				{groupedItems}
			</Item.Group>
		);
	}
}