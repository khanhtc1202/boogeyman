import {Component} from "react";
import PropTypes from "prop-types";
import React from "react";

export class Result extends Component {
	render() {
		return (
			<div>
				<h3>{this.props.title}</h3>
				<a href={this.props.link}>{this.props.link}</a>
				<p>{this.props.description}</p>
			</div>
		);
	}
}

Result.propTypes = {
	link: PropTypes.string.isRequired,
	title: PropTypes.string.isRequired,
	description: PropTypes.string
};
