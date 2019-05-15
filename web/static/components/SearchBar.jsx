import {Component} from "react";
import PropTypes from "prop-types";
import React from "react";

export class SearchBar extends Component {
	constructor(props) {
		super(props);
		this.state = {
			query: "sample",
			selectedEngine: "",
			selectedStrategy: "",
		}
	}

	render() {
		return (
			<div>

			</div>
		);
	}
}

SearchBar.propTypes = {
	query: PropTypes.func.isRequired
};
